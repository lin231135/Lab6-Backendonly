package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lin231135/Backend/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/series", handlers.GetSeries)
	app.Get("/api/series/:id", handlers.GetSeriesByID)
	app.Post("/api/series", handlers.CreateSeries)
	app.Put("/api/series/:id", handlers.UpdateSeries)
	app.Delete("/api/series/:id", handlers.DeleteSeries)

	app.Patch("/api/series/:id/status", handlers.UpdateSeriesStatus)
    app.Patch("/api/series/:id/episode", handlers.IncrementEpisode)
    app.Patch("/api/series/:id/upvote", handlers.UpvoteSeries)
    app.Patch("/api/series/:id/downvote", handlers.DownvoteSeries)
}