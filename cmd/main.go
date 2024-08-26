package main

import (
	server "Carbon/internal/app"
	"fmt"
	"log/slog"
	"time"
)

func main() {
	startTime := time.Now()
	server.NewApp()
	timeTaken := time.Since(startTime)
	slog.Info(
		fmt.Sprintf("Server is running on http://localhost:8080/"),
		slog.String("time_taken", timeTaken.String()),
	)
	err := server.Run()
	if err != nil {
		slog.Error("Failed to start server:", "error", err)
	}
}
