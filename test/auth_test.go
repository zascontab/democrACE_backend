package test

import (
	"github.com/sonderkevin/governance/application"
	"github.com/sonderkevin/governance/config"
	"github.com/sonderkevin/governance/graph/model"
)

func (s *IntTestSuite) TestAuthService_Register() {
	tests := []struct {
		name    string
		i       *model.RegistrarUsuarioInput
		wantErr bool
	}{
		{
			name: "usuario should register succesfully",
			i: &model.RegistrarUsuarioInput{
				Nombres:       "nombres3",
				Apellidos:     "apellidos3",
				NombreUsuario: "nombreusuario3",
				Email:         "email3",
				Password:      "password3",
			},
			wantErr: false,
		},
		{
			name: "usuario with duplicate nombres and apellidos should register succesfully",
			i: &model.RegistrarUsuarioInput{
				Nombres:       "nombres3",
				Apellidos:     "apellidos3",
				NombreUsuario: "nombreusuario4",
				Email:         "email4",
				Password:      "password4",
			},
			wantErr: false,
		},
		{
			name: "duplicate email error",
			i: &model.RegistrarUsuarioInput{
				Nombres:       "nombres5",
				Apellidos:     "apellidos5",
				NombreUsuario: "nombreusuario5",
				Email:         "email4",
				Password:      "password5",
			},
			wantErr: true,
		},
		{
			name: "duplicate nombreusuario error",
			i: &model.RegistrarUsuarioInput{
				Nombres:       "nombres6",
				Apellidos:     "apellidos6",
				NombreUsuario: "nombreusuario4",
				Email:         "email6",
				Password:      "password6",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		u, err := s.authService.UserRegister(tt.i, s.usuarioService)
		if tt.wantErr {
			s.NotNil(err)
		} else {
			s.Nil(err)
			s.NotNil(u)
			s.Equal(u.Nombres, tt.i.Nombres)
			s.Equal(u.Apellidos, tt.i.Apellidos)
			s.Equal(u.NombreUsuario, tt.i.NombreUsuario)
			s.Equal(u.Email, tt.i.Email)
			s.Equal(u.Verificado, false)
			s.Equal(u.Status, true)

			// Grupo Invitado
			grupoID, err := application.ConvertAndEncrypt(3)
			s.Nil(err)
			s.NotNil(grupoID)
			s.Equal(u.GrupoID, grupoID)
		}
	}
}

func (s *IntTestSuite) TestAuthService_Login() {
	tests := []struct {
		name    string
		i       *model.LoginUsuarioInput
		wantErr bool
	}{
		{
			name: "admin should login succesfully",
			i: &model.LoginUsuarioInput{
				Email:    s.config.ADMIN_EMAIL,
				Password: s.config.ADMIN_PASSWORD,
			},
			wantErr: false,
		},
		{
			name: "admin wrong password",
			i: &model.LoginUsuarioInput{
				Email:    s.config.ADMIN_EMAIL,
				Password: "",
			},
			wantErr: true,
		},
		{
			name: "invitado should login succesfully",
			i: &model.LoginUsuarioInput{
				Email:    "email1",
				Password: "password1",
			},
			wantErr: false,
		},
		{
			name: "invitado wrong password",
			i: &model.LoginUsuarioInput{
				Email:    "email1",
				Password: "wrong",
			},
			wantErr: true,
		},
		{
			name: "invitado no password",
			i: &model.LoginUsuarioInput{
				Email: "email1",
			},
			wantErr: true,
		},
		{
			name:    "empty",
			i:       &model.LoginUsuarioInput{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		refreshToken, err := s.authService.UserLogin(tt.i)
		if tt.wantErr {
			s.Assert().NotNil(err)
		} else {
			s.Assert().Nil(err)
			s.Require().NotNil(refreshToken)
		}
	}
}

func (s *IntTestSuite) TestAuthService_VerificarUsuario() {
	tests := []struct {
		name    string
		id      uint
		wantErr bool
	}{
		{
			name:    "verify admin should error",
			id:      1,
			wantErr: true,
		},
		{
			name:    "usuario already verified",
			id:      2,
			wantErr: true,
		},
		{
			name:    "should verify succesfully",
			id:      3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		u, err := s.authService.UR.Get(tt.id)
		if !tt.wantErr {
			s.Require().Nil(err)
		}

		codigoVerificacion := generateCodigoVerificacion()

		err = s.authService.SaveCodigoVerificacion(u, codigoVerificacion)
		if !tt.wantErr {
			s.Require().Nil(err)
		}

		err = s.authService.VerificarUsuario(u.Email, codigoVerificacion)
		if !tt.wantErr {
			s.Require().Nil(err)
		}

		u, err = s.authService.UR.Get(tt.id)
		if !tt.wantErr {
			s.Require().Nil(err)
			s.Require().NotNil(u)
			s.Equal(u.Verificado, true)
			s.Equal(u.CodigoVerificacion, "")
		}
	}
}

func (s *IntTestSuite) TestAuthService_ChangePassword() {
	tests := []struct {
		name    string
		i       *model.ChangePasswordInput
		wantErr bool
	}{
		{
			name: "can't change admin password",
			i: &model.ChangePasswordInput{
				Email:       config.ADMIN_EMAIL,
				NewPassword: config.ADMIN_PASSWORD,
			},
			wantErr: true,
		},
		{
			name: "should change usuario password",
			i: &model.ChangePasswordInput{
				Email:       "email1",
				NewPassword: "newPassword1",
			},
			wantErr: false,
		},
		{
			name: "should not change unverified usuario password",
			i: &model.ChangePasswordInput{
				Email:       "email2",
				NewPassword: "newPassword2",
			},
			wantErr: true,
		},
		{
			name: "unexistent usuario",
			i: &model.ChangePasswordInput{
				Email:       "wrong",
				NewPassword: "wrong",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		codigoVerificacion, _ := s.authService.SendCodigoVerificacion(tt.i.Email, SendEmailMock)

		tt.i.Code = codigoVerificacion
		err := s.authService.ChangePassword(tt.i)
		if tt.wantErr {
			s.NotNil(err)
		} else {
			s.Require().Nil(err)
			u, err := s.authService.UR.GetByEmail(tt.i.Email)
			s.Require().Nil(err)
			s.Require().NotNil(u)
			s.Equal(u.Verificado, true)
			s.Equal(u.CodigoVerificacion, "")
			err = application.ComparePassword(u.Password, tt.i.NewPassword)
			s.Nil(err)
		}
	}
}

func (s *IntTestSuite) TestAuthService_GetTempToken() {
	tests := []struct {
		name    string
		i       *model.LoginUsuarioInput
		wantErr bool
	}{
		{
			name: "should get admin tempToken succesfully",
			i: &model.LoginUsuarioInput{
				Email:    config.ADMIN_EMAIL,
				Password: config.ADMIN_PASSWORD,
			},
			wantErr: false,
		},
		{
			name: "should get usuario tempToken succesfully",
			i: &model.LoginUsuarioInput{
				Email:    "email1",
				Password: "password1",
			},
			wantErr: false,
		},
		{
			name: "should not get tempToken for unexistent usuario",
			i: &model.LoginUsuarioInput{
				Email:    "wrong",
				Password: "wrong",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		refreshToken, err := s.authService.UserLogin(tt.i)
		if tt.wantErr {
			s.NotNil(err)
			s.Nil(refreshToken)
		} else {
			s.Nil(err)
			s.Require().NotNil(refreshToken)
			s.Require().NotNil(refreshToken.Token)

			tempToken, err := s.authService.GetTempToken(refreshToken.Token)
			s.Require().Nil(err)
			s.Require().NotNil(tempToken)
			s.Assert().NotNil(tempToken.Token)
		}
	}
}
