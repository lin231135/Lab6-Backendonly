package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/lin231135/Backend/database"
    "github.com/lin231135/Backend/models"
    "strconv"
    "errors"
    "gorm.io/gorm"
)

func GetSeries(c *fiber.Ctx) error {
    series, err := models.GetAllSeries(database.DB.Db)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.JSON(series)
}

func GetSeriesByID(c *fiber.Ctx) error {
    idParam := c.Params("id")
    if idParam == "" {
        return c.Status(400).SendString("ID parameter is missing")
    }

    id, err := strconv.Atoi(idParam)
    if err != nil {
        return c.Status(400).SendString("ID must be an integer")
    }

    serie, err := models.GetSeriesByID(database.DB.Db, id)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return c.Status(404).SendString("Series not found")
        }
        return c.Status(500).SendString(err.Error())
    }
    return c.JSON(serie)
}

func CreateSeries(c *fiber.Ctx) error {
    serie := new(models.Serie)
    if err := c.BodyParser(serie); err != nil {
        return c.Status(400).SendString(err.Error())
    }
    if err := models.CreateSerie(database.DB.Db, serie); err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.Status(201).JSON(serie)
}

func UpdateSeries(c *fiber.Ctx) error {
    idParam := c.Params("id")
    serie := new(models.Serie)
    if err := c.BodyParser(serie); err != nil {
        return c.Status(400).SendString(err.Error())
    }

    id, err := strconv.Atoi(idParam)
    if err != nil {
        return c.Status(400).SendString("ID must be an integer")
    }

    if err := models.UpdateSerie(database.DB.Db, id, serie); err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.JSON(serie)
}

func DeleteSeries(c *fiber.Ctx) error {
    idParam := c.Params("id")

    id, err := strconv.Atoi(idParam)
    if err != nil {
        return c.Status(400).SendString("ID must be an integer")
    }

    if err := models.DeleteSerie(database.DB.Db, id); err != nil {
        return c.Status(500).SendString(err.Error())
    }
    return c.SendStatus(204)
}

func IncrementEpisode(c *fiber.Ctx) error {
    id := c.Params("id")
    var serie models.Serie
    if err := database.DB.Db.Model(&serie).Where("id = ?", id).First(&serie).Error; err != nil {
        return c.Status(404).SendString("Series not found")
    }
    serie.LastEpisodeWatched += 1 
    if err := database.DB.Db.Save(&serie).Error; err != nil {
        return c.Status(500).SendString("Failed to update series")
    }
    return c.JSON(serie)
}

func UpdateSeriesStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	status := c.Query("status")
	var serie models.Serie
	if err := database.DB.Db.Model(&serie).Where("id = ?", id).First(&serie).Error; err != nil {
		return c.Status(404).SendString("Series not found")
	}
	serie.Status = status
	if err := database.DB.Db.Save(&serie).Error; err != nil {
		return c.Status(500).SendString("Failed to update series status")
	}
	return c.JSON(serie)
}

func UpvoteSeries(c *fiber.Ctx) error {
	id := c.Params("id")
	var serie models.Serie
	if err := database.DB.Db.Model(&serie).Where("id = ?", id).First(&serie).Error; err != nil {
		return c.Status(404).SendString("Series not found")
	}
	serie.Ranking += 1 
	if err := database.DB.Db.Save(&serie).Error; err != nil {
		return c.Status(500).SendString("Failed to upvote series")
	}
	return c.JSON(serie)
}

func DownvoteSeries(c *fiber.Ctx) error {
	id := c.Params("id")
	var serie models.Serie
	if err := database.DB.Db.Model(&serie).Where("id = ?", id).First(&serie).Error; err != nil {
		return c.Status(404).SendString("Series not found")
	}
	serie.Ranking -= 1
	if err := database.DB.Db.Save(&serie).Error; err != nil {
		return c.Status(500).SendString("Failed to downvote series")
	}
	return c.JSON(serie)
}