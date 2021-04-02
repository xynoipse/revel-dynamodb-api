package controllers

import (
	"fmt"
	"net/http"
	"revel-dynamodb-api/app/entities/user"
	"revel-dynamodb-api/app/utils/response"
	validate "revel-dynamodb-api/app/utils/validation"

	"github.com/mitchellh/mapstructure"
	"github.com/revel/revel"
)

type Auth struct {
	*revel.Controller
}

func (c Auth) Register() revel.Result {
	var reqBody struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	c.Params.BindJSON(&reqBody)

	var User user.Model
	mapstructure.Decode(reqBody, &User)

	// Validate fields
	User.Validate(c.Validation, user.CreateService)
	if c.Validation.HasErrors() {
		c.Response.Status = http.StatusUnprocessableEntity
		res := response.Failed(c.Validation.ErrorMap(), response.Invalid)
		return c.RenderJSON(res)
	}

	// Check if email doesn't already exist
	if ok := User.IsEmailNew(); !ok {
		err := make(map[string]*revel.ValidationError)
		err["email"] = &revel.ValidationError{
			Message: user.EmailAlreadyTaken,
			Key:     "email",
		}

		c.Response.Status = http.StatusUnprocessableEntity
		res := response.Failed(err, response.AlreadyExist)
		return c.RenderJSON(res)
	}

	// Create the User
	if err := User.Create(); err != nil {
		c.Response.Status = http.StatusInternalServerError
		res := response.Failed(err.Error(), response.ServerError)
		return c.RenderJSON(res)
	}

	accessToken, err := User.GenerateAuthToken()
	if err != nil {
		c.Response.Status = http.StatusInternalServerError
		res := response.Failed(err.Error(), response.ServerError)
		return c.RenderJSON(res)
	}

	data := make(map[string]interface{})
	data["access_token"] = accessToken
	data["user"] = User
	res := response.Success(data, user.RegisterSuccess)
	return c.RenderJSON(res)
}

func (c Auth) Login() revel.Result {
	var reqBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	c.Params.BindJSON(&reqBody)

	validate.Email(c.Validation, reqBody.Email).Key("email")
	validate.Password(c.Validation, reqBody.Password).Key("password")
	if c.Validation.HasErrors() {
		c.Response.Status = http.StatusUnprocessableEntity
		res := response.Failed(c.Validation.ErrorMap(), response.Invalid)
		fmt.Printf("%+v", c.Validation.ErrorMap())
		return c.RenderJSON(res)
	}

	User, err := user.FindByEmail(reqBody.Email)
	if err != nil {
		err := make(map[string]*revel.ValidationError)
		err["email"] = &revel.ValidationError{
			Message: user.InvalidCredentials,
			Key:     "email",
		}

		c.Response.Status = http.StatusUnprocessableEntity
		res := response.Failed(err, response.Invalid)
		return c.RenderJSON(res)
	}

	if valid, err := User.ValidatePassword(reqBody.Password); err != nil {
		c.Response.Status = http.StatusInternalServerError
		res := response.Failed(err.Error(), response.ServerError)
		return c.RenderJSON(res)
	} else if !valid {
		err := make(map[string]*revel.ValidationError)
		err["email"] = &revel.ValidationError{
			Message: user.InvalidCredentials,
			Key:     "email",
		}

		c.Response.Status = http.StatusUnprocessableEntity
		res := response.Failed(err, response.Invalid)
		return c.RenderJSON(res)
	}

	accessToken, err := User.GenerateAuthToken()
	if err != nil {
		c.Response.Status = http.StatusInternalServerError
		res := response.Failed(err.Error(), response.ServerError)
		return c.RenderJSON(res)
	}

	data := make(map[string]interface{})
	data["access_token"] = accessToken
	res := response.Success(data, user.LogInSuccess)
	return c.RenderJSON(res)
}
