package registry

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/khafidprayoga/psychic-octo-disco/config"
	"gorm.io/gorm"
	"log"
)

// StartBackend  Used to init new fiber app with initialized domain services
func StartBackend(app *fiber.App, dbMysql *gorm.DB) {
	// TODO: mount all domain interface implementation instance

	socketListener := fmt.Sprintf(":%v", config.ServerPort)

	if !fiber.IsChild() {
		// Log starting server only on main process when pre-forked config are true
		log.Printf("starting the server... at %v", socketListener)
	}

	if err := app.Listen(socketListener); err != nil {
		log.Fatalf("failed to start backend services: %v", err)
	}
}
