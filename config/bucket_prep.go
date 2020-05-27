package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func fileS3Copy(ses *s3.S3, bucket, acl, fp, bucketPath, mimeType string) error {
	FileReader, err := os.Open(fp)
	if err != nil {
		return err
	}
	stat, err := FileReader.Stat()
	if err != nil {
		return err
	}
	Len := stat.Size()
	UploadParams := &s3.PutObjectInput{
		Bucket:             &bucket,
		Key:                &bucketPath,
		ContentType:        &mimeType,
		Body:               FileReader,
		ACL:                aws.String(acl),
		ContentLength:      &Len,
		ContentDisposition: aws.String("attachment"),
	}
	_, err = ses.PutObject(UploadParams)
	if err != nil {
		return err
	}
	return nil
}

// BucketPrep is used to prepare the S3 bucket.
func (s *S3Config) BucketPrep() error {
	// Create the session.
	ses := s.CreateS3Client()

	// Set default_pfp.png.
	err := fileS3Copy(ses, s.Bucket, "public-read", "./default_pfp.png", "default_pfp.png", "image/png")
	if err != nil {
		return err
	}

	// Set the default theme.
	err = fileS3Copy(ses, s.Bucket, "public-read", "./frontend/dist/base.html", "themes/default/base.html", "text/html; charset=utf-8")
	if err != nil {
		return err
	}
	err = fileS3Copy(ses, s.Bucket, "public-read", "./frontend/dist/manifest.json", "themes/default/manifest.json", "application/json")
	if err != nil {
		return err
	}
	err = fileS3Copy(ses, s.Bucket, "public-read", "./frontend/dist/ui.css.map", "themes/default/ui.css.map", "application/json")
	if err != nil {
		return err
	}
	err = fileS3Copy(ses, s.Bucket, "public-read", "./frontend/dist/ui.js.map", "themes/default/ui.js.map", "application/json")
	if err != nil {
		return err
	}
	err = fileS3Copy(ses, s.Bucket, "public-read", "./frontend/dist/ui.css", "themes/default/ui.css", "text/css")
	if err != nil {
		return err
	}
	err = fileS3Copy(ses, s.Bucket, "public-read", "./frontend/dist/ui.js", "themes/default/ui.js", "application/javascript")
	if err != nil {
		return err
	}

	// Return no errors.
	return nil
}
