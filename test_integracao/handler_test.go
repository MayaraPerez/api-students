package test_integracao

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MayaraPerez/api-students/internal/api"
	"github.com/MayaraPerez/api-students/internal/schema"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateStudentIntegration(t *testing.T){
	server := api.NewServer()
	server.Routes()
	
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&schema.Student{})
	assert.NoError(t, err)

	student := schema.Student{
		Name: "test" ,
		Email: "test@test",
		CPF: 123456789,
		Age: 12,
		Active: true,
	}

	body, _ := json.Marshal(student)

	req := httptest.NewRequest(http.MethodPost, "/students", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	server.Echo.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusCreated, rec.Code)

}