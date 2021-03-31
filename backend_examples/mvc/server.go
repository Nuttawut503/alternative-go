package mvc

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	indexCtl "app/main/backend_examples/mvc/controllers"
	userCtl "app/main/backend_examples/mvc/controllers/user"
)

func SetUp() *fiber.App {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		Views:         html.New("./backend_examples/mvc/views", ".html"),
	})

	app.Mount("/", indexCtl.IndexHandler())
	app.Mount("/user", userCtl.UserHandler())

	return app
}

func handleShutdown(app *fiber.App, done chan<- bool) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	s := <-sig
	fmt.Println()
	fmt.Println("Signal Name:", s)
	app.Shutdown()
	done <- true
}

func CreateAPI() {
	app := SetUp()

	done := make(chan bool)
	go handleShutdown(app, done)

	app.Listen(":3000")

	<-done
	fmt.Println("Closed API")
}
