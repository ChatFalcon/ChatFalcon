package setup

import (
	"context"
	"github.com/ChatFalcon/ChatFalcon/config"
	"github.com/ChatFalcon/ChatFalcon/installkey"
	"github.com/ChatFalcon/ChatFalcon/redis"
	"github.com/ChatFalcon/ChatFalcon/router"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jakemakesstuff/structuredhttp"
	"github.com/labstack/echo"
	"github.com/mailgun/mailgun-go/v4"
	"net/http"
	"strings"
	"time"
)

type setup struct {
	InstallKey    string `json:"installKey"`
	Stage         uint   `json:"stage"`
	RedisHostname string `json:"redisHostname"`
	RedisPassword string `json:"redisPassword"`
	S3Region string `json:"s3Region"`
	S3AccessKeyID string `json:"s3AccessKeyId"`
	S3SecretAccessKey string `json:"s3SecretAccessKey"`
	S3Endpoint string `json:"s3Endpoint"`
	S3Bucket string `json:"s3Bucket"`
	S3BucketURL string `json:"s3BucketUrl"`
	MailgunDomain string `json:"mailgunDomain"`
	MailgunPrivateKey string `json:"mailgunPrivateKey"`
	MailgunFrom string `json:"mailgunFrom"`
	FirstUserEmail string `json:"firstUserEmail"`
	FirstUserUsername string `json:"firstUserUsername"`
	FirstUserPassword string `json:"firstUserPassword"`
}

var setupHandlers = map[uint]func(c echo.Context, s *setup) (bool, error){
	0: func(c echo.Context, s *setup) (bool, error) {
		return false, nil
	},
	1: setupRedis,
	2: setupS3,
	3: setupMailgun,
	4: setupUser,
	5: finishSetup,
}

func setupUser(c echo.Context, s *setup) (bool, error) {
	if s.FirstUserEmail == "" || s.FirstUserPassword == "" || s.FirstUserUsername == "" {
		return true, c.String(http.StatusBadRequest, "No fields in this page can be blank.")
	}
	return false, nil
}

func setupMailgun(c echo.Context, s *setup) (bool, error) {
	mg := mailgun.NewMailgun(s.MailgunDomain, s.MailgunPrivateKey)
	mg.SetAPIBase(mailgun.APIBaseEU)
	_, err := mg.GetDomain(context.TODO(), s.MailgunDomain)
	if err != nil {
		return true, c.String(http.StatusBadRequest, err.Error())
	}
	return false, nil
}

func setupS3(c echo.Context, s *setup) (bool, error) {
	// Create the S3 session.
	x := config.S3Config{
		Endpoint:        s.S3Endpoint,
		Bucket:          s.S3Bucket,
		AccessKeyId:     s.S3AccessKeyID,
		SecretAccessKey: s.S3SecretAccessKey,
		Region:          s.S3Bucket,
		BucketURL: s.S3BucketURL,
	}
	ses := x.CreateS3Client()

	// Put a object with public read.
	key := "testfile"
	MimeType := "text/plain"
	FileReader := strings.NewReader("test")
	Len := int64(4)
	UploadParams := &s3.PutObjectInput{
		Bucket:             &s.S3Bucket,
		Key:                &key,
		ContentType:        &MimeType,
		Body:               FileReader,
		ACL:                aws.String("public-read"),
		ContentLength:      &Len,
		ContentDisposition: aws.String("attachment"),
	}
	_, err := ses.PutObject(UploadParams)
	if err != nil {
		return true, c.String(http.StatusBadRequest, err.Error())
	}

	// Get the test URL.
	url, err := x.GenerateURL("testfile")
	if err != nil {
		return true, c.String(http.StatusBadRequest, err.Error())
	}

	// Check the URL works.
	res, err := structuredhttp.GET(url).Timeout(2 * time.Minute).Run()
	if err != nil {
		return true, c.String(http.StatusBadRequest, err.Error())
	}
	err = res.RaiseForStatus()
	if err != nil {
		return true, c.String(http.StatusBadRequest, err.Error())
	}
	t, err := res.Text()
	if err != nil {
		return true, c.String(http.StatusBadRequest, err.Error())
	}
	if t != "test" {
		return true, c.String(http.StatusBadRequest, "string doesn't match")
	}

	// Delete the item from the bucket.
	_, err = ses.DeleteObject(&s3.DeleteObjectInput{Bucket: &s.S3Bucket, Key: &key})
	if err != nil {
		return true, c.String(http.StatusBadRequest, err.Error())
	}

	// Return no errors.
	return false, nil
}

func setupRedis(c echo.Context, s *setup) (bool, error) {
	// Check if the hostname isn't blank. If so, run the check.
	if s.RedisHostname != "" {
		err := redis.CreateRedisClient(s.RedisHostname, s.RedisPassword)
		if err != nil {
			return true, c.String(http.StatusBadRequest, err.Error())
		}
	}

	// Return no errors.
	return false, nil
}

func postSetup(c echo.Context) (err error) {
	// Small helper function to return 204.
	Return204 := func() {
		err = c.String(http.StatusNoContent, "")
	}

	// Decode the body.
	var s setup
	if err = c.Bind(&s); err != nil {
		err = c.String(http.StatusBadRequest, "Could not bind the setup JSON.")
		return
	}

	// Check the install key.
	if s.InstallKey != installkey.InstallationKey || installkey.InstallationKey == "" {
		err = c.String(http.StatusUnauthorized, "Installation key is incorrect.")
		return
	}

	// Get the handler.
	m, ok := setupHandlers[s.Stage]
	if !ok {
		err = c.String(http.StatusBadRequest, "Unknown stage.")
		return
	}
	var handled bool
	handled, err = m(c, &s)
	if handled {
		return
	}

	// Return a 204.
	Return204()
	return
}

func init() {
	router.Router.POST("/setup/", postSetup)
}
