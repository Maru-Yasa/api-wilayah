package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	allowedDirs := []string{
		"district", "districts", "province", "regencies", "regency", "village", "villages",
	}

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("API Wilayah")
	})

	app.Get("/api/provinces", func(c fiber.Ctx) error {
		return c.SendFile("./api/provinces.json")
	})

	app.Get("*", func(c fiber.Ctx) error {
		requestedPath := strings.TrimPrefix(c.OriginalURL(), "/api/")
		requestedPath = strings.TrimSuffix(requestedPath, ".json")

		for _, dir := range allowedDirs {
			if strings.HasPrefix(requestedPath, dir) {
				filePath := filepath.Join("./api", requestedPath+".json")

				if _, err := os.Stat(filePath); os.IsNotExist(err) {
					return c.Status(fiber.StatusNotFound).SendStatus(404)
				}

				return c.SendFile(filePath)
			}
		}

		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":3000"))
}
