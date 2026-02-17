package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ravirraj/rest-api/internal/config"
	"github.com/ravirraj/rest-api/internal/http/handlers/student"
	"github.com/ravirraj/rest-api/internal/storage/sqlite"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//databse setup
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	//setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	router.HandleFunc("GET /api/students/", student.GetList(storage))
	router.HandleFunc("DELETE /api/students/{id}",student.DeleteStudentById(storage))

	//setup server

	server := http.Server{
		Addr:    cfg.Adr,
		Handler: router,
	}

	fmt.Println("server starting on", cfg.Adr)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("faild to start server", err)
		}
	}()

	<-done

	slog.Info("shutting down the server ")

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown server", slog.String("err", err.Error()))
	}
	slog.Info("server shutdown successfully")

}
