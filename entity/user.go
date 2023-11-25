package entity

import (
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// req untuk create user
func NewUser(name, email, password string) *User {
	return &User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}
}

// req untuk update user
func UpdateUser(id int64, name, email, password string) *User {
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		UpdatedAt: time.Now(),
	}
}

// req untuk login
// func Login(email, password string) *User {
// 	return &User{
// 		Email:    email,
// 		Password: password,
// 	}
// }

//note : ketika type data untuk ID hanya int, maka akan error ketika dijalankan. karena ID tidak bisa di tambahkan otmatis oleh database
// namun ketika type data untuk ID diubah menjadi int64, maka tidak akan error ketika dijalankan. karena ID bisa di tambahkan otmatis oleh database melalui postman.
