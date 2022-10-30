package registry

import (
	"github.com/khafidprayoga/psychic-octo-disco/interface/activity"
	"github.com/khafidprayoga/psychic-octo-disco/interface/todo"
	"github.com/khafidprayoga/psychic-octo-disco/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/khafidprayoga/psychic-octo-disco/config"
	"gorm.io/gorm"
	"log"
)

// StartBackend  Used to init new fiber app with initialized domain interface
func StartBackend(app *fiber.App, dbMysql gorm.DB) {
	// TODO: mount all domain interface implementation instance
	todo.ServicesImpl(
		todo.InitServicesTodo{
			App: app,
			Db:  dbMysql,
		},
	)

	activity.ServiceImpl(
		activity.WithDependency{
			App: app,
			DB:  dbMysql,
		},
	)

	socketListener := ":3030" // sesuai kriteria

	if !fiber.IsChild() {
		if config.ServerProd {
			routeList := utils.GetAppRouteList(app)
			log.Printf("\n%v\n", routeList)
		}
		// Log starting server only on main process when pre-forked config are true
		log.Printf("starting the server... at %v", socketListener)
	}

	if err := app.Listen(socketListener); err != nil {
		log.Fatalf("failed to start backend service: %v", err)
	}
}
