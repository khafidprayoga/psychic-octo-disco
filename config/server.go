package config

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"strconv"
)

var Server *fiber.App
var ServerPort int

func init() {

	portEnv := os.Getenv("APP_PORT")
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		log.Fatalf("error parsing app port: %v", err)
	}

	// Init server port from env
	ServerPort = port

	// Init server if App instance does not exist
	if Server == nil {
		fiberConfig := fiber.Config{
			DisableStartupMessage: os.Getenv("PRODUCTION") == "1",
			JSONEncoder:           sonic.Marshal,
			JSONDecoder:           sonic.Unmarshal,
			Prefork:               os.Getenv("PRODUCTION") == "1",
		}

		// Init exported variable `Server` for new fiber app instance
		Server = fiber.New(fiberConfig)

		newCompressor := compress.New(compress.Config{
			Level: compress.LevelBestSpeed,
		})

		newLogger := logger.New(logger.Config{
			TimeZone:   "Asia/Jakarta",
			Format:     "${pid} [${ip}]:${port} ${locals:requestid} ${status} - ${method} ${path}  ${latency}\n",
			TimeFormat: "02-Jan-2006",
		})

		// Mount middleware
		Server.
			Use(requestid.New()).
			Use(etag.New()).
			Use(newCompressor).
			Use(newLogger)

		Server.All("/", func(ctx *fiber.Ctx) error {
			return ctx.SendString("Todo list API challenge")
		})
	}
}
