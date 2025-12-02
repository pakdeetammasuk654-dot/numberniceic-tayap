package handlers

import "github.com/gofiber/fiber/v2"

func APIList(c *fiber.Ctx) error {
	// Pass the layout to the Render function
	return c.Render("api_list", fiber.Map{}, "layouts/main")
}

// New handler for the search page
func SearchPage(c *fiber.Ctx) error {
	return c.Render("search", fiber.Map{}, "layouts/main")
}
