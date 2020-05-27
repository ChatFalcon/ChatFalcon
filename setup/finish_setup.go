package setup

import (
	"context"
	"github.com/ChatFalcon/ChatFalcon/config"
	"github.com/ChatFalcon/ChatFalcon/mongo"
	"github.com/ChatFalcon/ChatFalcon/redis"
	"github.com/ChatFalcon/ChatFalcon/user"
	"github.com/labstack/echo"
	"net/http"
)

func finishSetup(c echo.Context, s *setup) (bool, error) {
	// Create the config.
	var r *config.RedisConfig
	if s.RedisHostname != "" {
		r = &config.RedisConfig{
			Host:     s.RedisHostname,
			Password: s.RedisPassword,
		}
		err := redis.CreateRedisClient(r.Host, r.Password)
		if err != nil {
			return true, c.String(http.StatusInternalServerError, err.Error())
		}
	}
	cfg := &config.ServerConfig{
		S3Config:       &config.S3Config{
			Endpoint:        s.S3Endpoint,
			Bucket:          s.S3Bucket,
			AccessKeyId:     s.S3AccessKeyID,
			SecretAccessKey: s.S3SecretAccessKey,
			Region:          s.S3Region,
			BucketURL:       s.S3BucketURL,
		},
		RedisConfig:    r,
		MailgunConfig:  &config.MailgunConfig{
			Domain:     s.MailgunDomain,
			PrivateKey: s.MailgunPrivateKey,
			From:       s.MailgunFrom,
		},
		Name:           "ChatFalcon",
		Description:    "Your ChatFalcon forum has been configured!",
		Warnings:       []*config.Warning{},
		SignupsEnabled: true,
		CurrentTheme: "default",
	}

	// Prep the bucket.
	err := cfg.S3Config.BucketPrep()
	if err != nil {
		return true, c.String(http.StatusInternalServerError, err.Error())
	}

	// Write the config.
	_, err = config.GetConfig()
	if err == nil {
		return false, nil
	}
	_, err = mongo.DB.Collection("serverConfig").InsertOne(context.TODO(), cfg)
	if err != nil {
		return true, c.String(http.StatusInternalServerError, err.Error())
	}

	// Write the first user.
	url, _ := cfg.S3Config.GenerateURL("default_pfp.png")
	u := &user.User{
		Perms: &user.Permissions{
			Admin: true,
		},
		Username: s.FirstUserUsername,
		Email: s.FirstUserEmail,
		PFPUrl: url,
		Confirmed: true,
	}
	token, err := u.Create(s.FirstUserPassword)
	if err != nil {
		return true, c.String(http.StatusInternalServerError, err.Error())
	}
	c.SetCookie(&http.Cookie{
		Name:       "token",
		Value:      token,
		MaxAge:		0,
	})

	// Return no errors.
	return false, nil
}
