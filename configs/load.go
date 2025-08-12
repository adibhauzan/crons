package configs

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	MODE string

	AppOrigin string

	// // MSSQL variables
	MSSQLDBHost     string
	MSSQLDBUser     string
	MSSQLSAPassword string
	MSSQLDBName     string
	MSSQLDBPort     string

	// Minio variables
	MinioBucketName string
	MinioURL        string
	MinioAccessKey  string
	MinioSecretKey  string

	// RabbitMQ variables
	rabbitmqUser string
	rabbitmqPass string
	rabbitmqHost string
)

func LoadConfig() {
	_ = godotenv.Load()

	MODE = os.Getenv("MODE")

	AppOrigin = os.Getenv("APP_ORIGIN")

	MSSQLDBHost = os.Getenv("MSSQL_DB_HOST")
	MSSQLDBUser = os.Getenv("MSSQL_DB_USER")
	MSSQLSAPassword = os.Getenv("MSSQL_SA_PASSWORD")
	MSSQLDBName = os.Getenv("MSSQL_DB_NAME")
	MSSQLDBPort = os.Getenv("MSSQL_DB_PORT")

	MinioBucketName = os.Getenv("MINIO_BUCKET_NAME")
	MinioURL = os.Getenv("MINIO_URL")
	MinioAccessKey = os.Getenv("MINIO_ACCESS_KEY")
	MinioSecretKey = os.Getenv("MINIO_SECRET_KEY")

	rabbitmqUser = os.Getenv("RABBITMQ_USER")
	rabbitmqPass = os.Getenv("RABBITMQ_PASS")
	rabbitmqHost = os.Getenv("RABBITMQ_HOST")

}
