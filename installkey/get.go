package installkey

import (
	"github.com/ChatFalcon/ChatFalcon/config"
	"github.com/ChatFalcon/ChatFalcon/mongo"
	"github.com/ChatFalcon/ChatFalcon/redis"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"os"
)

// InstallationKey defines the installation key.
var InstallationKey = ""

// GetInstallKey is used to get the install key if needed.
func GetInstallKey() {
	cfg, err := config.GetConfig()
	if err != nil {
		if err != mongo.ErrNoDocuments {
			panic(err)
		}
		InstallationKey = os.Getenv("INSTALL_KEY")
		if InstallationKey == "" {
			InstallationKey = uuid.Must(uuid.NewUUID()).String()
		} else {
			logrus.Info("Installation key got from INSTALL_KEY variable.")
		}
		logrus.Info("Your installation key is ", InstallationKey)
	} else if cfg.RedisConfig != nil {
		err = redis.CreateRedisClient(cfg.RedisConfig.Host, cfg.RedisConfig.Password)
		if err != nil {
			logrus.Error("Failed to connect to Redis at boot.")
		}
	}
}
