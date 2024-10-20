package api

import (
	"errors"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/MayaraPerez/api-students/schema"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// Handlers
func (api *API) GetStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Error to List students")
	}

	active := c.QueryParam("active")
	if active != "" {
		act , err := strconv.ParseBool(active)
		if err != nil {
			log.Error().Err(err).Msgf("[api] error to parse boolean")
			return c.String(http.StatusInternalServerError, "Failed to parse boolean")
		}
		students, _ = api.DB.GetFilterStudents(act)
	}

	listOfStudents := map[string][]schema.StudentResponse{"students": schema.NewStudentResponse(students)}

	return c.JSON(http.StatusOK, listOfStudents)
}

func (api *API) createStudent(c echo.Context) error {
	studentRequest := StudentRequest{}
	if err := c.Bind(&studentRequest); err != nil {
		return err
	}

	if err := studentRequest.Validate(); err != nil {
		log.Error().Err(err).Msgf("[api] error validation struct")
		return c.String(http.StatusBadRequest, "Error to validation student")
	}

	student := schema.Student{
		Name:   studentRequest.Name,
		Email:  studentRequest.Email,
		CPF:    studentRequest.CPF,
		Age:    studentRequest.Age,
		Active: *studentRequest.Active,
	}

	if err := api.DB.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to created student")
	}
	return c.JSON(http.StatusOK, student)
}

func (api *API) getStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to get student id ")
	}

	student, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}
	return c.JSON(http.StatusOK, student)
}

func (api *API) updateStudent(c echo.Context) error {
	//converto string para int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to get student id ")
	}
	//faco bind do que recebeu com a minha struct
	receivedStudent := schema.Student{}
	if err := c.Bind(&receivedStudent); err != nil {
		return err
	}

	//recebo o regitro que eu tenho no BD
	updateStudent, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}

	//chamo a func que verifica os que existe com o que veio
	student := updateVerify(receivedStudent, updateStudent)

	//passo o student oara fun que vai fazer o update
	if err := api.DB.UpdateStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to save student")
	}
	return c.JSON(http.StatusOK, updateStudent)
}

func (api *API) deleteStudent(c echo.Context) error {
	//converto string para int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to get student id ")
	}

	//recebo o regitro que eu tenho no BD
	deleteStudent, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete student")
	}

	//passo o student oara fun que vai fazer o update
	if err := api.DB.DeleteStudent(deleteStudent); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete student")
	}

	return c.JSON(http.StatusOK, deleteStudent)
}

func updateVerify(receivedStudent, student schema.Student) schema.Student {
	if receivedStudent.Name != "" {
		student.Name = receivedStudent.Name
	}
	if receivedStudent.Email != "" {
		student.Email = receivedStudent.Email
	}
	if receivedStudent.CPF > 0 {
		student.CPF = receivedStudent.CPF
	}
	if receivedStudent.Age > 0 {
		student.Age = receivedStudent.Age
	}
	if receivedStudent.Active != student.Active {
		student.Active = receivedStudent.Active
	}
	return student
}
