package models

type UserResponse struct {
	Status   int   `json:"status"`
	Response User  `json:"response,omitempty"`
	Error    error `json:"error,omitempty"`
}

type UsersResponse struct {
	Status   int    `json:"status"`
	Response []User `json:"response,omitempty"`
	Error    error  `json:"error,omitempty"`
}
