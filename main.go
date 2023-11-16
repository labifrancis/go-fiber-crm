package main

import (
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get(GetLeads)
	app.Post(NewLeads)
	app.Delete(DeleteLead)
	app.Get(Lead)
}

func initDatabase(){
	
}
func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")

}
