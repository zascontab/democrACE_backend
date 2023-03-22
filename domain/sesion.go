package domain

import "time"

type Sesion struct {
	UsuarioID   uint
	Dispositivo string
	Exp         time.Time
}
