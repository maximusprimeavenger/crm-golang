package repository

import (
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ProxyToCRM(c *fiber.Ctx, path string) error {
	resp, err := http.Get("http://crm-analytics:9092" + path)
	if err != nil {
		return c.
			Status(fiber.StatusBadGateway).
			JSON(fiber.Map{"error": "crm-analytics is unavailable"})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "failed to read crm-analytics response"})
	}

	c.Status(resp.StatusCode)

	if ct := resp.Header.Get("Content-Type"); ct != "" {
		c.Set("Content-Type", ct)
	}

	return c.Send(body)
}
