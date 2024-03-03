package main

import (
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/FernandoJVideira/velox"
)

type application struct {
	App        *velox.Velox
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
	wg         sync.WaitGroup
}

func main() {
	v := initApplication()
	go v.listenForShutdwn()
	err := v.App.ListenAndServe()
	if err != nil {
		v.App.ErrorLog.Println(err)
	}
}

func (a *application) shutdown() {
	// Put any cleanup tasks here
	a.App.InfoLog.Println("Shutting down")

	// Block until the waitgroup is empty
	a.wg.Wait()
}

func (a *application) listenForShutdwn() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit

	a.App.InfoLog.Println("Received signal:", s.String())
	a.shutdown()

	os.Exit(0)
}
