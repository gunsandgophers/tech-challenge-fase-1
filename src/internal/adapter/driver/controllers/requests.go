package controllers

import "fmt"

type RegiterCustomerRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   string  `json:"cpf"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (r *RegiterCustomerRequest) Validate() error {
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.CPF == "" {
		return errParamIsRequired("cpf", "string")
	}

	return nil
}
