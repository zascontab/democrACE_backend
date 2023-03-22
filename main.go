package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/sonderkevin/governance/config"
	"github.com/sonderkevin/governance/db"
	"github.com/sonderkevin/governance/infrastructure/persistence"
	"github.com/sonderkevin/governance/server"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	port := config.PORT
	if port == "" {
		port = defaultPort
	}

	// New db
	d := db.New()
	defer db.Close()

	// Setup the server
	es := server.GenerateExecutableSchema(d)
	srv := server.Setup(es)

	// Optional Automigration
	automigrate := flag.Bool("auto-migrate", false, "Auto-migrate the database")
	flag.Parse()
	if *automigrate {
		pp := makePermisos(es.Schema())
		fmt.Println("Automigrating the database...")
		db.Migrate(pp)
	}

	// Remove expired sessions
	go func() {
		r := persistence.NewSesionRepository(d)
		interval, err := time.ParseDuration(config.REMOVE_EXPIRED_SESSIONS_TIME_INTERVAL)
		if err != nil {
			fmt.Printf("failed to get remove expired sessions time interval error: %s", err.Error())
			return
		}
		for {
			fmt.Print("removing expired sessions...")
			r.RemoveExpiredSessions()
			time.Sleep(interval)
		}
	}()

	server.Run(srv, port)
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
