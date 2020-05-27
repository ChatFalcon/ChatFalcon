package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/url"
	"strings"
)

// Warning is used to define the warning.
type Warning struct {
	// 0 = info, 1 = warning, 2 = error
	Type int `json:"type" bson:"type"`

	// Defines the message.
	Message string `json:"message" bson:"message"`
}

// S3Config is used to define the S3 compatible host configuration.
type S3Config struct {
	Endpoint        string `json:"endpoint" bson:"endpoint"`
	Bucket          string `json:"bucket" bson:"bucket"`
	AccessKeyId     string `json:"accessKeyId" bson:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey" bson:"secretAccessKey"`
	Region          string `json:"region" bson:"region"`
	BucketURL       string `json:"bucketUrl" bson:"bucketUrl"`
}

// GenerateURL is used to generate a S3 URL with a path.
func (s *S3Config) GenerateURL(Path string) (string, error) {
	u, err := url.Parse(s.BucketURL)
	if err != nil {
		return "", err
	}
	if !strings.HasPrefix(Path, "/") {
		Path = "/" + Path
	}
	u.Path = Path
	return u.String(), nil
}

// CreateS3Client is used to create the AWS S3 client.
func (s *S3Config) CreateS3Client() *s3.S3 {
	StaticCredential := credentials.NewStaticCredentials(s.AccessKeyId, s.SecretAccessKey, "")
	var rptr *string
	if s.Region != "" {
		rptr = &s.Region
	}
	var eptr *string
	if s.Endpoint != "" {
		eptr = &s.Endpoint
	}
	s3sess := session.Must(session.NewSession(&aws.Config{
		Endpoint:    eptr,
		Credentials: StaticCredential,
		Region:      rptr,
	}))
	return s3.New(s3sess)
}

// RedisConfig is used to define the Redis configuration.
type RedisConfig struct {
	Host     string `json:"host" bson:"host"`
	Password string `json:"password" bson:"password"`
}

// MailgunConfig is used to define the Mailgun configuration.
type MailgunConfig struct {
	Domain string `json:"domain"`
	PrivateKey string `json:"privateKey"`
	From string `json:"from"`
}

// ServerConfig is used to define the servers configuration.
type ServerConfig struct {
	S3Config       *S3Config    `json:"s3Config" bson:"s3Config"`
	RedisConfig    *RedisConfig `json:"redisConfig" bson:"redisConfig"`
	MailgunConfig  *MailgunConfig `json:"mailgunConfig" bson:"mailgunConfig"`
	Name           string       `json:"name" bson:"name"`
	Description    string       `json:"description" bson:"description"`
	Warnings       []*Warning   `json:"warnings" bson:"warnings"`
	CustomScript   string       `json:"customScript" bson:"customScript"`
	CustomHead     string       `json:"customHead" bson:"customHead"`
	CustomBody     string       `json:"customBody" bson:"customBody"`
	CurrentTheme string `json:"currentTheme" bson:"currentTheme"`
	SignupsEnabled bool         `json:"signupsEnabled" bson:"signupsEnabled"`
}
