package models

import "time"

// Base -> base struct
type Base struct {
	ID        BINARY16  `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UintBase -> base struct
type UintBase struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UintBase -> base struct
type UserBase struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// EmailParams -> parameters to send email from gmail
type EmailParams struct {
	To              string      //recivername
	SubjectData     interface{} //subject data for templates
	SubjectTemplate string      //template file name
	BodyData        interface{} //body data for templates
	BodyTemplate    string      //body template filename
	Lang            string      //languages
	From            string
}

type EmailSubject struct {
	Name    string
	Company string
	Title   string
}

type EmailBody struct {
	ToName   string
	Company  string
	Name     string
	Title    string
	URL      string
	Password string
}
