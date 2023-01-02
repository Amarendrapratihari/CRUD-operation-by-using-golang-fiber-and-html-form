package database

import (
	"database/sql"
	"fmt"
	"strconv"
)


func AddStudent(db *sql.DB, student_name string,father_name string,class int, age int) bool {

	insertuser := `INSERT INTO user(student_name, father_name, class, age) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insertuser)
	if err != nil {
		fmt.Println(err.Error()+"1" )
	}
	_, err = statement.Exec(student_name, father_name , class, age)
	if err != nil {
		fmt.Println(err.Error() + "2")
	}else{
		return true

	}
	return false
}


func GetStudent(db *sql.DB)[][] string{
	
	rows, err := db.Query("SELECT * FROM user")
	
	
	if err != nil{
		panic(err)
	}
	defer rows.Close()


	myCourse := [][]string{}
	for rows.Next(){
		var temp = []string{}
		var id int
		var student_name string
		var father_name string
		var class int	
		var age int
		err = rows.Scan(&id,&student_name, &father_name , &class, &age)
		if err != nil{
			panic(err)
		}
		temp = append(temp,strconv.Itoa(id),student_name,father_name,strconv.Itoa(class),strconv.Itoa(age))
		myCourse = append(myCourse,temp)



	}

	return myCourse
}

func GetStudentById(db *sql.DB, id string)[][] string{

	rows, err := db.Query("SELECT id, student_name, father_name, class, age FROM user where id='"+id+"';")
	
	
	if err != nil{
		panic(err)
	}
	defer rows.Close()


	myCourse := [][]string{}
	for rows.Next(){
		var temp = []string{}
		var id int
		var student_name string
		var father_name string
		var class int	
		var age int
		err = rows.Scan(&id,&student_name, &father_name , &class, &age)
		if err != nil{
			panic(err)
		}
		temp = append(temp,strconv.Itoa(id),student_name,father_name,strconv.Itoa(class),strconv.Itoa(age))
		myCourse = append(myCourse,temp)



	}

	return myCourse
}

func UpdateStudent(db *sql.DB, student_name string,father_name string,class int, age int, id int) bool {

	statement, err := db.Prepare("update user SET student_name = ?, father_name = ?, class = ?, age = ? where id=?;")
	if err != nil {
		fmt.Println(err.Error()+"1" )
	}
	_, err = statement.Exec(student_name, father_name , class, age,id)
	if err != nil {
		fmt.Println(err.Error() + "2")
	}else{
		return true

	}
	return false
	
}

func DeleteStudentById(db *sql.DB,id string) bool{
	query := "delete from user where id = ?"
	statement,err := db.Prepare(query)
	if err != nil{
		panic(err)
	}
	_,err = statement.Exec(id)
	if err != nil{
		panic(err)
	}
	return true

}