package config

func GetDBUrl() string {
	return "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable"
}
