package web

import "fmt"

type ErrHttp struct {
	Description string `json:"description,omitempty"`
	Metadata    string `json:"metadata,omitempty"`
	StatusCode  int    `json:"statusCode"`
}

func (e ErrHttp) Error() string {
	return fmt.Sprintf("description: %s, metadata: %s", e.Description, e.Metadata)
}

func NewHttpError(description, metadata string, statusCode int) ErrHttp {
	return ErrHttp{
		Description: description,
		Metadata:    metadata,
		StatusCode:  statusCode,
	}
}
