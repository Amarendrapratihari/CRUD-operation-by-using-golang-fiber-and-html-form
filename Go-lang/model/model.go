package model

import (
	"database/sql"
	"fmt"
	"student/database"
)





func InsertStudent(db *sql.DB, student_name string,father_name string,class int, age int) bool{
	data := database.AddStudent(db, student_name, father_name,class,age)	
	return data
}


func GetStudent(db *sql.DB )[][]string {
	data:= database.GetStudent(db)	
	return data
}

func GetStudentById(db *sql.DB, id string )[][]string {
	data:= database.GetStudentById(db, id)	
	return data
}

func UpdateStudent(db *sql.DB , student_name string,father_name string,class int, age int,id int) bool{
	data := database.UpdateStudent(db, student_name, father_name,class,age, id)	
	fmt.Print(data)
	return data
}

func DeleteStudentById(db *sql.DB, id string ) bool {
	data:= database.DeleteStudentById(db, id)	
	return data
}
