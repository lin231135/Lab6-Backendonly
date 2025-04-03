package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/lin231135/Backend/database"
)

func main() {
    database.ConnectDb()
    app := fiber.New()

    // Configura CORS para permitir solicitudes de cualquier origen
    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowMethods: "GET, POST, HEAD, PUT, DELETE, PATCH",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))

    setupRoutes(app)

    log.Fatal(app.Listen(":8080"))
}
