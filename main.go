package main

import (
	"log"
	"student/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", routes.Index)
	app.Post("/studentDataUpdated", routes.StudentUpdated)
	app.Get("/getStudent", routes.GetStudent)
	app.Get("/addStudent", routes.StudentAdd)
	app.Get("/updateStudent", routes.GetUpdateStudent)
	app.Get("/deleteStudent", routes.DeleteStudent)
	app.Post("/studentDataAdded", routes.AddStudent)
	app.Post("/deleteStudentData", routes.DeleteStudentData)

}

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	SetupRoutes(app)

	log.Fatal(app.Listen(":9090")) //Server will be start in 9090 port.
}
