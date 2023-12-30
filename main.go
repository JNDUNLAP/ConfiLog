package main

import (
	"dunlap/internal/logger"

	"github.com/joho/godotenv"
)

func ExampleFunction() {
	logger.Info("Application started")
	logger.Warning("Application Warning")
}

func main() {
	if err := godotenv.Load(); err != nil {
		return
	}

	loggerConfig := logger.LoggerConfig{
		PrintLogs: true,
	}

	logger.InitializeLogger(loggerConfig)
	ExampleFunction()

}

// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/jackc/pgx/v4"
// )

// func main() {
// 	dsn := "postgresql://jack:idYr12sxM0wC5TeHxKU-4g@lowest-rabbit-6710.g8z.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
// 	ctx := context.Background()
// 	conn, err := pgx.Connect(ctx, dsn)
// 	defer conn.Close(context.Background())
// 	if err != nil {
// 		log.Fatal("failed to connect database", err)
// 	}

// 	var now time.Time
// 	err = conn.QueryRow(ctx, "SELECT NOW()").Scan(&now)
// 	if err != nil {
// 		log.Fatal("failed to execute query", err)
// 	}

// 	fmt.Println(now)
// }
// idYr12sxM0wC5TeHxKU-4g
