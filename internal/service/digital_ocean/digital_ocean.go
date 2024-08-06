package service

import (
	"fmt"
	"mime/multipart"
	cnf "optitech/internal/config"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gofiber/fiber/v2/log"
)

func DownloadDocument(route string, directory string) (string, error) {

	s3Config := cnf.GetS3Config()

	sess, err := session.NewSession(s3Config)
	if err != nil {
		return "", err
	}

	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(cnf.Env.DigitalOceanBucket),
		Key:    aws.String(fmt.Sprintf("%s/%s", directory, route)),
	})
	urlStr, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}

	return urlStr, nil
}

func UploadDocument(fileHeader *multipart.FileHeader, name string, institutionName string) (string, error) {

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(cnf.Env.DigitalOceanKey, cnf.Env.DigitalOceanSecret, ""),
		Endpoint:         aws.String(cnf.Env.DigitalOceanEndpoint),
		S3ForcePathStyle: aws.Bool(false),
		Region:           aws.String(cnf.Env.DigitalOceanRegion),
	}

	sess, err := session.NewSession(s3Config)

	if err != nil {
		return "", err
	}
	uploader := s3manager.NewUploader(sess)

	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}

	log.Info(name, institutionName)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(cnf.Env.DigitalOceanBucket),
		Key:    aws.String(fmt.Sprintf("%s/%s", institutionName, name)),
		Body:   file,
	})
	if err != nil {
		return "", err
	}
	defer file.Close()
	return name, nil
}
