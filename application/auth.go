package application

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sonderkevin/governance/config"
	"github.com/sonderkevin/governance/domain"
	"github.com/sonderkevin/governance/domain/repository"
	"github.com/sonderkevin/governance/graph/model"
	"github.com/sonderkevin/governance/infrastructure/log"
	"github.com/vektah/gqlparser/v2/gqlerror"
	mail "github.com/xhit/go-simple-mail/v2"
	"gorm.io/gorm"
)

type AuthService struct {
	UR repository.UsuarioRepository
	SR repository.SesionRepository
}

func (s *AuthService) UserRegister(i *model.RegistrarUsuarioInput, us *UsuarioService) (*model.Usuario, error) {
	user, err := s.UR.GetByEmail(i.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("email already exists")
	}

	return us.Save(context.Background(), i)
}

func (s *AuthService) UserLogin(ctx context.Context, input *model.LoginUsuarioInput) (*model.RefreshToken, error) {
	// TODO: Check user number of sessions & user devices.

	user, err := s.UR.GetByEmail(input.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &gqlerror.Error{
				Message: "Email not found",
			}
		}

		return nil, err
	}

	if !user.Verificado {
		return nil, errors.New("usuario no verificado")
	}

	if err := ComparePassword(user.Password, input.Password); err != nil {
		return nil, err
	}

	token, err := GenerateToken(strconv.Itoa(int(user.ID)), "REFRESH_TOKEN_EXP")
	if err != nil {
		return nil, err
	}
	sessionExp, err := GetSessionExp(token)

	if err := s.HandleSession(ctx, user.ID, sessionExp); err != nil {
		return nil, err
	}

	return &model.RefreshToken{Token: token}, nil
}

func (s *AuthService) Logout(ctx context.Context) error {
	u, ok := ctx.Value("usuario").(*domain.Usuario)
	if !ok {
		return errors.New("access denied")
	}
	device, ok := ctx.Value("device").(string)
	if !ok {
		return errors.New("access denied: no hay dispositivo para esta sesion")
	}

	s.SR.Delete(u.ID, device)
	return nil
}

// TODO: Delete Session after expiration time
func (s *AuthService) HandleSession(ctx context.Context, userID uint, sessionExp *time.Time) error {
	device, ok := ctx.Value("device").(string)
	if !ok {
		return errors.New("access denied: no hay dispositivo para esta sesion")
	}

	sesiones, err := s.SR.GetByUsuario(userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	containsDevice := false
	for _, s := range sesiones {
		if s.Dispositivo == device {
			containsDevice = true
		}
	}

	session := domain.Sesion{UsuarioID: userID, Dispositivo: device, Exp: *sessionExp}
	if !containsDevice {
		if len(sesiones) >= config.MAX_SESIONES {
			return errors.New("acces denied: maximo numero de sesiones activas para el usuario")
		}

		if err := s.SR.Save(&session); err != nil {
			return err
		}
	} else {
		if err := s.SR.Update(&session); err != nil {
			return err
		}
	}

	return nil
}

func GetSessionExp(auth string) (*time.Time, error) {
	token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.SECRET), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		floatExp, ok := claims["exp"].(float64)
		if !ok {
			return nil, errors.New("internal")
		}
		exp := time.Unix(int64(floatExp), 0)
		return &exp, nil
	} else {
		return nil, err
	}
}

func (s *AuthService) GetTempToken(refreshToken string) (*model.TempToken, error) {
	tempToken, err := ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: "invalid refresh token",
		}
	}

	return &model.TempToken{Token: tempToken}, nil
}

func (s *AuthService) VerificarUsuario(email, code string) error {
	u, err := s.UR.GetByEmail(email)
	if err != nil {
		return err
	}

	// Verify code
	if u.CodigoVerificacion != code {
		return errors.New("wrong code")
	}

	u.Verificado = true
	u.CodigoVerificacion = ""
	return s.UR.Update(u)
}

func (s *AuthService) SendCodigoVerificacion(email string, sendEmailFunc func(from, to, subject, body string) error) (string, error) {
	// Get user
	u, err := s.UR.GetByEmail(email)
	if err != nil {
		return "", err
	}

	// Generate verification code
	codigoVerificacion, err := generateCodigoVerificacion()
	if err != nil {
		return "", err
	}

	// Send email with verification code
	if err = sendEmailFunc(config.EMAIL_SENDER, u.Email, "Codigo de Verificacion", "Su codigo de verificacion es: "+codigoVerificacion); err != nil {
		log.Info("Failed to send email: " + err.Error())
		return "", err
	}

	// Save verification code to database or cache
	s.SaveCodigoVerificacion(u, codigoVerificacion)
	log.Info("Verification code sent to email: " + u.Email)
	return codigoVerificacion, nil
}

func (s *AuthService) ChangePassword(input *model.ChangePasswordInput) error {
	// Get user
	u, err := s.UR.GetByEmail(input.Email)
	if err != nil {
		return err
	}

	if !u.Verificado {
		return errors.New("usuario no verificado")
	}

	if u.CodigoVerificacion == "" {
		return errors.New("can't change password")
	}

	// Verify code
	if u.CodigoVerificacion != input.Code {
		return errors.New("wrong code")
	}

	// Update user
	u.Verificado = true
	u.CodigoVerificacion = ""
	u.Password = HashPassword(input.NewPassword)
	return s.UR.Update(u)
}

func (s *AuthService) SaveCodigoVerificacion(u *domain.Usuario, code string) error {
	if u == nil {
		return errors.New("nil usuario")
	}
	u.CodigoVerificacion = code
	err := s.UR.Update(u)
	return err
}

func generateCodigoVerificacion() (string, error) {
	// Generate a random 6 digit number
	n := rand.Intn(1000000)
	code := fmt.Sprintf("%06d", n)

	return code, nil
}

func SendEmail(from, to, subject, body string) error {
	server := mail.NewSMTPClient()
	server.Host = config.EMAIL_HOST
	server.Port = 587
	server.Username = config.EMAIL_SENDER
	server.Password = config.EMAIL_SENDER_PASSWORD
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {

		return err
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom("From <" + from + ">")
	email.AddTo(to)
	email.SetSubject(subject)

	email.SetBody(mail.TextPlain, body)

	// Send email
	err = email.Send(smtpClient)
	if err != nil {

		return err
	}

	return nil
}
