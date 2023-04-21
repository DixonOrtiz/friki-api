package gormdb

var (
	testDBUser     = "user"
	testDBPassword = "password"
	testDBHost     = "localhost"
	testDBName     = "database"
	testDBPort     = "3006"
)

var expectedConfig = Config{
	User:     "user",
	Password: "password",
	Host:     "localhost",
	Port:     3006,
	DBName:   "database",
}
