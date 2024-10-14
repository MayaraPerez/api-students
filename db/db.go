package db

import (
	"github.com/rs/zerolog/log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StudentHandler struct {
	DB *gorm.DB
}

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json:"registration"`
}

func NewStudentHandle( db *gorm.DB) *StudentHandler  {
	return &StudentHandler{DB: db}
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to inicialize SQlite: %s", err.Error())
	}

	db.AutoMigrate(&Student{})

	return db
}

func (s *StudentHandler) AddStudent(student Student) error {

	if result := s.DB.Create(&student); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *StudentHandler) GetStudents() ([]Student, error) {
	students := []Student{}

	err := s.DB.Find(&students).Error 
		return students, err
}

func (s *StudentHandler) GetStudent(id int) (Student, error) {
	var student Student

	err := s.DB.First(&student, id)
		return student, err.Error
}

func (s *StudentHandler) UpdateStudent(updateStudent Student) error {
	return s.DB.Save(&updateStudent).Error
}