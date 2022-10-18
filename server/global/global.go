package global

import (
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"server/config"
	"sync"
)

var (
	CONFIG              config.Config
	VP                  *viper.Viper
	DB                  *gorm.DB
	REDIS               *redis.Client
	LOG                 *zap.SugaredLogger
	Concurrency_Control = &singleflight.Group{}

	//BlackCache local_cache.Cache
	lock sync.RWMutex

	// Timer               timer.Timer = timer.NewTimerTask()

)
