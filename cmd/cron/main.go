package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/adibhauzan/crons/configs"
	"github.com/adibhauzan/crons/databases"
	"github.com/adibhauzan/crons/internal/bootstrap"
)

func main() {
	configs.LoadConfig()
	loggerConfig := configs.NewLogger()
	cronConfig := configs.NewCronConfig()
	minioConnection, err := configs.NewMinioConnection()
	if err != nil {
		loggerConfig.Fatal("Failed to connect to RabbitMQ:", err)
	}

	conn, ch, err := configs.NewRabbitMQConnection()
	if err != nil {
		loggerConfig.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer ch.Close()

	db, err := databases.InitSqlServerConnection()
	if err != nil {
		loggerConfig.Fatal("Failed to connect to RabbitMQ:", err)
	}

	bootstrap.Bootstrap(&bootstrap.BootstrapConfig{
		DB:           db,
		Log:          loggerConfig,
		MinioClient:  minioConnection,
		RabbitmqConn: conn,
		Cron:         cronConfig,
	})

	cronConfig.Start()
	loggerConfig.Info("Cron service started successfully")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	cronConfig.Stop()
	loggerConfig.Info("Cron stopped")
}
