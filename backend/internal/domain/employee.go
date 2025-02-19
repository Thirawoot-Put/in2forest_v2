package domain

type Employee struct {
	ID uint

	Email    string
	Password string
	Name     string
	Mobile   string

	Role EmployeeRole
}
