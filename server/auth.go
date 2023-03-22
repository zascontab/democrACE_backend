package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gamebtc/devicedetector"
	"github.com/sonderkevin/governance/application"
	"github.com/sonderkevin/governance/db"
	"github.com/sonderkevin/governance/infrastructure/persistence"
	"gorm.io/gorm"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		device, err := GetDevice(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		ctx := context.WithValue(r.Context(), "device", device)
		r = r.WithContext(ctx)

		auth := r.Header.Get("Authorization")
		if auth == "" {
			next.ServeHTTP(w, r)
			return
		}

		bearer := "Bearer "
		auth = auth[len(bearer):]

		stringID, err := application.ValidateToken(auth)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		id, err := strconv.Atoi(stringID)
		if err != nil {
			http.Error(w, "Acces denied", http.StatusForbidden)
			return
		}

		ur := persistence.NewUsuarioRepository(db.New())
		usuario, err := ur.Get(uint(id))
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				http.Error(w, "Acces denied", http.StatusBadRequest)
				return
			}
			http.Error(w, "Acces denied", http.StatusForbidden)
			return
		}

		/* 		if err := HandleSession(w, r, uint(id), device); err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		} */

		ctx = context.WithValue(ctx, "usuario", usuario)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func HandleSession(w http.ResponseWriter, r *http.Request, id uint, device string) error {
	sr := persistence.NewSesionRepository(db.New())
	sesiones, err := sr.GetByUsuario(id)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	containsDispositivo := false
	for _, s := range sesiones {
		if s.Dispositivo == device {
			containsDispositivo = true
		}
	}

	if !containsDispositivo {
		return errors.New("access denied: la sesion no existe")
	}

	return nil
}

func GetDevice(r *http.Request) (string, error) {
	dd, err := devicedetector.NewDeviceDetector("regexes")
	if err != nil {
		return "", err
	}
	userAgent := r.Header.Get("User-Agent")
	info := dd.Parse(userAgent)
	deviceName := GetDeviceName(info)
	os := info.GetOs()
	return fmt.Sprintf("%s %s %s", deviceName, os.ShortName, os.Version), nil
}

func GetDeviceName(info *devicedetector.DeviceInfo) string {
	d := info.GetDevice()
	if d.Model != "" {
		return d.Model
	}
	if d.Brand != "" {
		return d.Brand
	}
	if d.Type != "" {
		return d.Type
	}
	return "unknown"
}
