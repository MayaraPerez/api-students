package tests_unit

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MayaraPerez/api-students/internal/api"
	"github.com/magiconair/properties/assert"
)

func TestGetStudents(t *testing.T) {
    // Configurar servidor API
    server := api.NewServer()
    server.Routes()

    // Criar nova requisição
    req := httptest.NewRequest(http.MethodGet, "/students", nil)
    rec := httptest.NewRecorder()

    // Executar a requisição
    server.Echo.ServeHTTP(rec, req)

    // Validar resposta
    assert.Equal(t, http.StatusOK, rec.Code)
}

func TestCreateStudent(t *testing.T) {
    //Preciso inicializar o servidor
    server := api.NewServer()
    server.Routes()

    //Depois eu crio um payload para enviar na minha requisicao
    trueValue := true
    payloadReq := api.StudentRequest{
        Name: "test",
        Email: "test@test",
        CPF: 123456789,
        Age: 20,
        Active: &trueValue,
    }
    //transformo os dados do meu payload em json para enviar na requisicao
    payload, err := json.Marshal(payloadReq)
    if err != nil {
        t.Fatalf("Error marshaling payload: %v", err)
    }
    //crio a requisicao e passo payload como parametro
    req := httptest.NewRequest(http.MethodPost, "/students", bytes.NewReader(payload))
    req.Header.Set("Content-Type", "application/json") // definicao do cabecalho da requisicao
    //gravo a resposta da requisicao
    rec := httptest.NewRecorder()

    //eu executo a requisicao
    server.Echo.ServeHTTP(rec, req)

    //valido o codigo de resposta
    assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestGetStudent(t *testing.T){
    server := api.NewServer()
    server.Routes()

    req := httptest.NewRequest(http.MethodGet, "/students/1", nil)
    rec := httptest.NewRecorder()

    server.Echo.ServeHTTP(rec, req)
}

func TestUpdateStudent(t *testing.T){  
    server := api.NewServer()
    server.Routes()

    updatePayload := api.StudentRequest{
        Name: "updateTest",
    }

    payload, err := json.Marshal(updatePayload)
    if err != nil {
        t.Fatalf("Error marshaling payload: %v", err)
    }
    req := httptest.NewRequest(http.MethodPut, "/students/1", bytes.NewReader(payload))
    req.Header.Set("Content-Type", "application/json") 
    rec := httptest.NewRecorder()

    server.Echo.ServeHTTP(rec, req)
   
}

// func TestDeleteStudent(t *testing.T){
//     server := api.NewServer()
//     server.Routes()

//     var trueValue = true
//     payloadDelete := api.StudentRequest {
//         Name: "deleteTest",
//         Email: "delete@test",
//         CPF: 123456789,
//         Age: 20,
//         Active: &trueValue,
//     }

//     payload, err := json.Marshal(payloadDelete)
//     if err != nil {
//         t.Fatalf("Error marshaling payload: %v", err)
//     }

//     req := httptest.NewRequest(http.MethodPost, "/students", bytes.NewReader(payload))
//     req.Header.Set("Content-Type", "application/json") 
//     rec := httptest.NewRecorder()

//     server.Echo.ServeHTTP(rec, req)
//     assert.Equal(t, http.StatusCreated, rec.Code)

//     //iniciar o Delete
//     deleteReq := httptest.NewRequest(http.MethodDelete, "/students/1", nil)
//     deleteRec := httptest.NewRecorder()

//     server.Echo.ServeHTTP(deleteRec, deleteReq)
//     assert.Equal(t, http.StatusNoContent, deleteRec.Code)

//     getReq := httptest.NewRequest(http.MethodGet, "/students/1", nil)
//     getRec := httptest.NewRecorder()

//     server.Echo.ServeHTTP(getRec, getReq)
//     assert.Equal(t, http.StatusNotFound, getRec.Code)
//}



