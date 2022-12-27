package config

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"time"
)

var Server *fiber.App
var ServerProd bool
var ServerDebug bool

func init() {
	debugServer := os.Getenv("APP_DEBUG")
	if debugServer == "1" {
		ServerDebug = true
	}

	serverProd := os.Getenv("PRODUCTION")
	if serverProd == "1" {
		ServerProd = true
	}

	// Init server if App instance does not exist
	if Server == nil {
		fiberConfig := fiber.Config{
			AppName:               "Todo App Challenge By Khafid",
			DisableStartupMessage: ServerProd,
			JSONEncoder:           sonic.Marshal,
			JSONDecoder:           sonic.Unmarshal,
			Prefork:               true, // for test
			//Prefork:               ServerProd,
			BodyLimit: 1 * 1024, // 1MB
			//EnablePrintRoutes:     ServerDebug,
			ReadTimeout:  1 * time.Second,
			WriteTimeout: 1 * time.Second,
		}

		// Init exported variable `Server` for new fiber app instance
		Server = fiber.New(fiberConfig)

		newCompressor := compress.New(compress.Config{
			Level: compress.LevelBestSpeed,
		})

		//newLogger := logger.New(logger.Config{
		//	TimeZone:   "Asia/Jakarta",
		//	Format:     "${pid} [${ip}]:${port} ${locals:requestid} ${status} - ${method} ${path}  ${latency}\n",
		//	TimeFormat: "02-Jan-2006",
		//})

		// Mount middleware
		Server.
			Use(newCompressor)
		//Use(requestid.New()).
		//Use(etag.New()).
		//Use(newLogger)
		// Initialize default config
		Server.Use(cache.New())

		// Or extend your config for customization
		Server.Use(cache.New(cache.Config{
			Next: func(c *fiber.Ctx) bool {
				return c.Query("refresh") == "true"
			},
			Expiration:   30 * time.Minute,
			CacheControl: true,
		}))
		//Server.Get("/", func(ctx *fiber.Ctx) error {
		//	return ctx.SendString("Todo list API challenge")
		//})
	}
}
