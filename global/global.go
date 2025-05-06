package global

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"

	"Go/pkg/logger"
	"Go/pkg/setting"

)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb *redis.Client
	Mdb *gorm.DB
	Mdbc *sql.DB
	KafkaProducer *kafka.Writer
)

