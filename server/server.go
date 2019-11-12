package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/NYTimes/gziphandler"
	"github.com/romainmenke/css/runtime"
)

// Server owns all server methods and state
type Server struct {
	runtime           *runtime.Runtime
	publicHTTPHandler http.Handler
}

// New returns a new server
func New() *Server {
	s := &Server{
		publicHTTPHandler: http.StripPrefix("/public/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, r.URL.Path)
		})),
		runtime: runtime.New(),
	}

	return s
}

// Run starts the server
func (s *Server) Run() {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	server := &http.Server{
		Addr:              ":" + os.Getenv("PORT"),
		Handler:           gziphandler.GzipHandler(s),
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 16,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	go func() {
		<-signalChannel

		ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			log.Println(err)
		}
	}()

	err := server.ListenAndServe()
	if err == http.ErrServerClosed {
		return
	}
	if err != nil {
		log.Println(err)
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/public/") {
		s.publicHTTPHandler.ServeHTTP(w, r)
		return
	}

	s.runtime.ServeHTTP(w, r)
}
