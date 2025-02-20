package dto

type Employee struct {
	ID       uint         `json:"id"`
	Email    string       `json:"email"`
	Password string       `json:"password"`
	Name     string       `json:"name"`
	Mobile   string       `json:"mobile"`
	Role     EmployeeRole `json:"role"`
}

type EmployeeResponse struct {
	ID     uint         `json:"id"`
	Email  string       `json:"email"`
	Name   string       `json:"name"`
	Mobile string       `json:"mobile"`
	Role   EmployeeRole `json:"role"`
}
