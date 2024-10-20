package db

import (
	"github.com/MayaraPerez/api-students/schema"
	"github.com/rs/zerolog/log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StudentHandler struct {
	DB *gorm.DB
}

func NewStudentHandle( db *gorm.DB) *StudentHandler  {
	return &StudentHandler{DB: db}
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to inicialize SQlite: %s", err.Error())
	}

	db.AutoMigrate(&schema.Student{})

	return db
}

func (s *StudentHandler) AddStudent(student schema.Student) error {

	if result := s.DB.Create(&student); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *StudentHandler) GetStudents() ([]schema.Student, error) {
	students := []schema.Student{}

	err := s.DB.Find(&students).Error 
		return students, err
}

func (s *StudentHandler) GetStudent(id int) (schema.Student, error) {
	var student schema.Student

	err := s.DB.First(&student, id)
		return student, err.Error
}

func (s *StudentHandler) UpdateStudent(updateStudent schema.Student) error {
	return s.DB.Save(&updateStudent).Error
}

func (s *StudentHandler) DeleteStudent(deleteStudent schema.Student) error {
	return s.DB.Delete(&deleteStudent).Error
}