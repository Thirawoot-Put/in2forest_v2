package constants

type UserRole struct {
	Admin string
}

var Role = UserRole{
	Admin: "ADMIN",
}
