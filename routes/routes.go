package routes

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"student/model"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{"Title": "Student Information"})
}

func StudentAdd(c *fiber.Ctx) error {
	tenantData := 0
	//fmt.Print(tenantData)
	return c.Render("addStudent", fiber.Map{
		"tenantData": tenantData,
	})
}

func AddStudent(c *fiber.Ctx) error {
	filename := "./data.db"
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		os.Exit(0)
	}
	defer db.Close()
	//var adminUser model.AdminUser
	studName := c.FormValue("studentName")
	fatherName := c.FormValue("fatherName")
	class, _ := strconv.Atoi(c.FormValue("class"))
	age, _ := strconv.Atoi(c.FormValue("age"))
	fmt.Println("ffffffff", studName, fatherName, class, age)
	model.InsertStudent(db, studName, fatherName, class, age)
	return c.Render("insert", fiber.Map{"Title": " Data inserted sucessfully!!!"})

}
func GetStudent(c *fiber.Ctx) error {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	res := model.GetStudent(db)
	return c.Render("getStudent", fiber.Map{"Title": "All Student Information", "data": res})
}

func GetUpdateStudent(c *fiber.Ctx) error {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	id := c.Query("uid")
	res := model.GetStudentById(db, id)
	return c.Render("updateStudent", fiber.Map{"Title": " Update Student Information", "data": res})

}

func StudentUpdated(c *fiber.Ctx) error {
	filename := "./data.db"
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		os.Exit(0)
	}
	defer db.Close()
	//var adminUser model.AdminUser
	id, _ := strconv.Atoi(c.FormValue("studentId"))
	studName := c.FormValue("studentName")
	fatherName := c.FormValue("fatherName")
	class, _ := strconv.Atoi(c.FormValue("class"))
	age, _ := strconv.Atoi(c.FormValue("age"))
	model.UpdateStudent(db, studName, fatherName, class, age, id)
	return c.Render("insert", fiber.Map{"Title": " Data Updated sucessfully!!!"})

}

func DeleteStudent(c *fiber.Ctx) error {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	id := c.Query("uid")
	res := model.GetStudentById(db, id)
	return c.Render("deleteStudent", fiber.Map{"Title": "Delete Student Information", "data": res})

}

func DeleteStudentData(c *fiber.Ctx) error {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	id := c.FormValue("studentId")
	model.DeleteStudentById(db, id)
	return c.Render("insert", fiber.Map{"Title": " Data Deleted sucessfully!!!"})

}
