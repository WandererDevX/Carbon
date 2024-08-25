package main

import (
	server "Carbon/internal/app"
	"fmt"
	"log/slog"
	"time"
)

func main() {
	startTime := time.Now()
	app := server.NewApp()
	timeTaken := time.Since(startTime)
	slog.Info(
		fmt.Sprintf("Server is running on http://localhost:8080/"),
		slog.String("time_taken", timeTaken.String()),
	)
	app.Run()
}
