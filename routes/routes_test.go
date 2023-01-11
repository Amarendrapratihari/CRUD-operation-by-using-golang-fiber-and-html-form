package routes

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestIndex(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Index(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Index() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddStudent(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddStudent(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AddStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetStudent(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetStudent(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUpdateStudent(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetUpdateStudent(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetUpdateStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStudentUpdated(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StudentUpdated(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("StudentUpdated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteStudent(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteStudent(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("DeleteStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteStudentData(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteStudentData(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("DeleteStudentData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
