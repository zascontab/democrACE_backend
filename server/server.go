package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/sonderkevin/governance/application"
	"github.com/sonderkevin/governance/graph/generated"
	"github.com/sonderkevin/governance/graph/resolver"
	"github.com/sonderkevin/governance/infrastructure/persistence"
	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func Run(srv *handler.Server, port string) {
	//Setup the router
	r := mux.NewRouter()
	r.Use(AuthMiddleware)
	r.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	r.Handle("/graphql", srv)
	r.HandleFunc("/archivo/{id}", ServeArchivo).Methods("GET")
	r.HandleFunc("/archivo", UploadArchivo).Methods("POST")
	r.HandleFunc("/archivo/{id}", DeleteArchivo).Methods("DELETE")

	// Setup the HTTP Server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Set up a signal handler to listen for SIGINT or SIGTERM signals
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fmt.Println("Server listening on port :" + port + "...")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Wait for a signal to shutdown
	sig := <-interrupt
	log.Printf("Received signal %s. Shutting down...\n", sig)

	// Set a timeout to wait for pending requests to complete
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Attempt to gracefully shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down server: %v\n", err)
	} else {
		log.Println("Server shut down gracefully")
	}
}

func GenerateExecutableSchema(db *gorm.DB) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &resolver.Resolver{
				AuthService: &application.AuthService{
					UR: persistence.NewUsuarioRepository(db),
					SR: persistence.NewSesionRepository(db),
				},

				ArchivoService: &application.ArchivoService{
					R: persistence.NewArchivoRepository(db),
				},
				CategoriaService: &application.CategoriaService{
					R: persistence.NewCategoriaRepository(db),
				},

				GrupoService: &application.GrupoService{
					R:  persistence.NewGrupoRepository(db),
					PR: persistence.NewPermisoRepository(db),
				},
				PermisoService: &application.PermisoService{
					R: persistence.NewPermisoRepository(db),
				},

				UsuarioService: &application.UsuarioService{
					R:  persistence.NewUsuarioRepository(db),
					PR: persistence.NewPermisoRepository(db),
				},
			},
		},
	)
}

func Setup(es graphql.ExecutableSchema) *handler.Server {
	srv := handler.NewDefaultServer(es)
	srv.AroundRootFields(RootPermisoMiddleware)
	return srv
}
