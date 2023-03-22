package test

import (
	"fmt"
	"math/rand"

	"github.com/vektah/gqlparser/v2/ast"
)

func generateCodigoVerificacion() string {
	// Generate a random 6 digit number
	n := rand.Intn(1000000)
	code := fmt.Sprintf("%06d", n)

	return code
}

func SendEmailMock(from, to, subject, body string) error {
	return nil
}

func makePermisos(s *ast.Schema) []string {
	pp := []string{"all", "uploadArchivo", "serveArchivo", "deleteArchivo"}

	for _, f := range s.Query.Fields {
		pp = append(pp, f.Name)
	}

	for _, f := range s.Mutation.Fields {
		pp = append(pp, f.Name)
	}

	return pp
}

func parseColumns(cc []string) string {
	out := ""
	for i, c := range cc {
		if i > 0 {
			out += ", "
		}
		out += c
	}
	return out
}

func strPtr(s string) *string {
	return &s
}
