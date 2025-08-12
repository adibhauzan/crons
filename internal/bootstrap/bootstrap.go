package bootstrap

import (
	"github.com/adibhauzan/crons/internal/broker"
	"github.com/adibhauzan/crons/internal/crons"
	"github.com/adibhauzan/crons/internal/handlers"
	"github.com/adibhauzan/crons/internal/repositories"
	"github.com/adibhauzan/crons/internal/services"
	"github.com/minio/minio-go/v7"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB           *gorm.DB
	Log          *logrus.Logger
	MinioClient  *minio.Client
	RabbitmqConn *amqp.Connection
	Cron         *cron.Cron
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	applicationSetting := repositories.NewApplicationSettingRepository(config.DB)
	claimBlastingRepo := repositories.NewClaimBlastingRepository(config.DB)
	documentRepo := repositories.NewDocumentsRepository(config.DB)
	storageRepo := repositories.NewStorageRepository(config.MinioClient)

	// setup producer
	producer, err := broker.NewRabbitMQProducer(config.RabbitmqConn)
	if err != nil {
		config.Log.Fatalf("Failed to create RabbitMQ producer: %v", err)
	}

	storageService := services.NewStorageService(storageRepo)
	claimService := services.NewClaimBastingService(applicationSetting, claimBlastingRepo, documentRepo, storageService, producer, config.DB, config.Log)

	// setup handlers
	claimBlastingHandler := handlers.NewClaimBlastingHandler(claimService, config.Cron, config.Log)
	routeConfig := crons.CronConfig{
		ClaimBlasting: claimBlastingHandler,
	}
	routeConfig.Setup()
}
