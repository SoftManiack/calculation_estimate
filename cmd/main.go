package main

import (
	handler "calculation-estimate/pkg/handlers"
	service "calculation-estimate/pkg/services"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	port, _ := os.LookupEnv("PORT")

	services := service.NewService()
	handlers := handler.NewHandler(services)

	srv := new(Server)

	fmt.Println("app start")

	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		fmt.Printf("error occured while running http server: %s", err.Error())
	}

}

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 / time.Second,
	}

	return s.httpServer.ListenAndServe()
}
