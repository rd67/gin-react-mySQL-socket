package helpers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rd67/gin-react-mySQL-socket/configs"
	"gorm.io/gorm"
)

///////////////////////
//	Validation Response
///////////////////////

// Handle validator error formatting
type ValidationErrorDataStruct struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func FormatValidationError(verr validator.ValidationErrors) []ValidationErrorDataStruct {
	errs := []ValidationErrorDataStruct{}

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs = append(errs, ValidationErrorDataStruct{Field: f.Field(), Reason: err})
	}

	return errs
}

//	Validation Error Response
type ValidationResponseStruct struct {
	StatusCode int                         `json:"statusCode"`
	Message    string                      `json:"message"`
	Data       []ValidationErrorDataStruct `json:"data"`
	Error      string                      `json:"error"`
}

func ValidationResponse(c *gin.Context, err error) {

	var ve validator.ValidationErrors

	var errorData = make([]ValidationErrorDataStruct, 0)
	if errors.As(err, &ve) {
		errorData = FormatValidationError(ve)
	}

	response := ValidationResponseStruct{
		StatusCode: http.StatusBadRequest,
		Message:    "Validation failed, kindly check your parameters",
		Data:       errorData,
		Error:      err.Error(),
	}

	c.JSON(response.StatusCode, response)
	return
}

///////////////////////
//	Error Response
///////////////////////
type ErrorResponseStruct struct {
	configs.CommonResponseStruct
	Error string `json:"error"`
}

func ErrorResponse(c *gin.Context, err error) {

	var statusCode int
	var message string

	if errors.Is(err, gorm.ErrRecordNotFound) {
		statusCode = http.StatusNotFound
		message = "Not found"
	} else {
		statusCode = http.StatusInternalServerError
		message = "Something went wrong, please try again"
	}

	fmt.Println(err)

	response := ErrorResponseStruct{
		CommonResponseStruct: configs.CommonResponseStruct{
			StatusCode: statusCode,
			Message:    message,
		},
		Error: err.Error(),
	}

	c.JSON(response.StatusCode, response)
	return
}

///////////////////////
//	Success Response
///////////////////////
type ActionFailedResponseStruct struct {
	configs.CommonResponseStruct
}

func ActionFailedResponse(c *gin.Context, message string) {

	response := ActionFailedResponseStruct{
		CommonResponseStruct: configs.CommonResponseStruct{
			StatusCode: http.StatusBadRequest,
			Message:    message,
		},
	}

	c.JSON(response.StatusCode, response)
	return
}
