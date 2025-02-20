package handlers

import portin "thirawoot/in2forest_shop_backend/internal/ports/port_in"

type EmployeeHandler struct {
	app portin.EmployeeApp
}

func NewEmployeeHandler(app portin.EmployeeApp) *EmployeeHandler {
	return &EmployeeHandler{app: app}
}
