package api

import "fmt"

type StudentRequest struct {
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active *bool   `json:"registration"` //using bool as a pointer to forcer true/false
}

func msgValidationError(param, typ string) error {
	return fmt.Errorf("params '%s' of type '%s' is required", param, typ)
}

func (s *StudentRequest) Validate() error {
	if s.Name == "" {
		return msgValidationError("name", "string")
	}
	if s.Email == "" {
		return msgValidationError("email", "string")
	}
	if s.CPF == 0 {
		return msgValidationError("cpf", "int")
	}
	if s.Age == 0 {
		return msgValidationError("age", "int")
	}
	if s.Active == nil{
		return msgValidationError("active", "bool")
	}
	return nil
}