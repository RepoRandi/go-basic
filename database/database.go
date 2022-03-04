package database

var connection string

func init() {
	connection = "MySql"
}

func GetConnection() string {
	return connection
}
