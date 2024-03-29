package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labifrancis.com/go-fiber-crm/database"
	"github.com/labifrancis.com/go-fiber-crm/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/id", lead.Getlead)
	app.Post("/api/v1/lead", lead.NewLeads)
	app.Delete("/api/v1/lead/id", lead.DeleteLead)

}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.Automigrate(&lead.Lead{})
	fmt.Println("Database migrated")

}
func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(":3000")
	defer database.DBConn.Close()

}

```

"github.com/gofiber/fiber/v2"
```

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labifrancis.com/go-fiber-crm/database"
	"github.com/labifrancis.com/go-fiber-crm/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/id", lead.Getlead)
	app.Post("/api/v1/lead", lead.NewLeads)
	app.Delete("/api/v1/lead/id", lead.DeleteLead)

}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.Automigrate(&lead.Lead{})
	fmt.Println("Database migrated")

}
func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(":3000")
	defer database.DBConn.Close()

}
