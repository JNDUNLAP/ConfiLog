# Confilog

Log here - log there (MongoDB) - log in stealth - log with Color - log location info on where you be slippin

## Installation
```go
go get github.com/your-username/logger
```

## Usage
```go
package main

import (
    "fmt"
    "os"
    "github.com/joho/godotenv"
    "github.com/your-username/logger"
)

func main() {
    if err := godotenv.Load(); err != nil {
        fmt.Println("Error loading .env file:", err)
        return
    }

    // Configure 
    loggerConfig := logger.LoggerConfig{
        PrintLogs:      true,
        MongoURI:       os.Getenv("YOUR_MONGO_URI"),                 // MongoDB URI from environment variable
        DatabaseName:   os.Getenv("YOUR_MONGO_DATABASE_NAME"),       // Name of the database
        CollectionName: os.Getenv("YOUR_MONGO_COLLECTION_NAME"),     // Name of the collection for logs
    }

    logger.InitializeLogger(loggerConfig)
}
```