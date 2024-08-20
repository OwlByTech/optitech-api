package service

import (
	"bytes"
	"fmt"
	"io"
	cnf "optitech/internal/config"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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

func DownloadDocumentByte(route string, directory string) ([]byte, error) {
	s3Config := cnf.GetS3Config()

	sess, err := session.NewSession(s3Config)
	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	result, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(cnf.Env.DigitalOceanBucket),
		Key:    aws.String(fmt.Sprintf("%s/%s", directory, route)),
	})
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, result.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func UploadDocument(fileBytes []byte, name string, institutionName string) (*string, error) {

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(cnf.Env.DigitalOceanKey, cnf.Env.DigitalOceanSecret, ""),
		Endpoint:         aws.String(cnf.Env.DigitalOceanEndpoint),
		S3ForcePathStyle: aws.Bool(false),
		Region:           aws.String(cnf.Env.DigitalOceanRegion),
	}

	sess, err := session.NewSession(s3Config)
	if err != nil {
		return nil, err
	}

	uploader := s3manager.NewUploader(sess)
	fileReader := bytes.NewReader(fileBytes)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(cnf.Env.DigitalOceanBucket),
		Key:    aws.String(fmt.Sprintf("%s/%s", institutionName, name)),
		Body:   fileReader,
	})
	if err != nil {
		return nil, err
	}

	return &name, nil
}
