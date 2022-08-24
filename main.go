package main

import (
	"io"
	"net/http"

	"github.com/laenur/simple-analytic/pkg/config"
	"github.com/laenur/simple-analytic/pkg/logger"
	"github.com/laenur/simple-analytic/pkg/server"
)

func main() {
	logger.Setup()
	err := config.Load(".env")
	if err != nil {
		logger.Error("config.Load", "Err", err)
	}

	port, _ := config.GetInt("PORT")
	server := server.NewServer(port)
	server.AddHandler("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")

	})
	err = server.StartServer()
	if err != nil {
		logger.Error("server.StartServer", "Err", err)
	}
}
