package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// untuk membuat validator
type FormValidator struct {
	validator *validator.Validate
}

// untuk memvalidasi struct yang diinputkan
func (fv *FormValidator) Validate(i interface{}) error {
	// untuk melakukan validasi
	return fv.validator.Struct(i)
}

// func ini digunakan ketika ingin mereturn data json yang diinputkan
func NewFormValidator() *FormValidator {
	// diguanakn ketik ingin melakuka validate ke struct akan di enable
	validate := validator.New(validator.WithRequiredStructEnabled())

	//untk memvalidate type data json yang diinputkan. misal ketika ingin menginputkan nama harus huruf kecil semua. gabisa satu kecil satu gede.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &FormValidator{validate}
}

// untuk menampilkan error yang terjadi / maping error validator pada field atau json yang required.
func ValidatorErrors(err error) map[string]string {
	// untuk membuat map
	fields := map[string]string{}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				fields[err.Field()] = fmt.Sprintf("%s harus diisi", err.Field())
			case "password":
				// Validasi khusus untuk field "password"
				fields[err.Field()] = "Password harus mengandung setidaknya satu huruf besar dan satu digit"
			default:
				fields[err.Field()] = fmt.Sprintf("Kesalahan pada %s dengan tag %s seharusnya %s", err.Field(), err.Tag(), err.Param())
			}
		}
	}

	return fields
}
