package database

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestAddStudent(t *testing.T) {
	type args struct {
		db           *sql.DB
		student_name string
		father_name  string
		class        int
		age          int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddStudent(tt.args.db, tt.args.student_name, tt.args.father_name, tt.args.class, tt.args.age); got != tt.want {
				t.Errorf("AddStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStudent(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStudent(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStudentById(t *testing.T) {
	type args struct {
		db *sql.DB
		id string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStudentById(tt.args.db, tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStudentById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateStudent(t *testing.T) {
	type args struct {
		db           *sql.DB
		student_name string
		father_name  string
		class        int
		age          int
		id           int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateStudent(tt.args.db, tt.args.student_name, tt.args.father_name, tt.args.class, tt.args.age, tt.args.id); got != tt.want {
				t.Errorf("UpdateStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteStudentById(t *testing.T) {
	type args struct {
		db *sql.DB
		id string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteStudentById(tt.args.db, tt.args.id); got != tt.want {
				t.Errorf("DeleteStudentById() = %v, want %v", got, tt.want)
			}
		})
	}
}
