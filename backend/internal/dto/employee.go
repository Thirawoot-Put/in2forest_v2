package dto

type Employee struct {
	ID       uint
	Email    string       `json:"email"`
	Password string       `json:"password"`
	Name     string       `json:"name"`
	Mobile   string       `json:"mobile"`
	Role     EmployeeRole `json:"role"`
}
