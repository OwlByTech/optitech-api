package service

import (
	"bytes"
	"io"
	cnf "optitech/internal/config"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func createSessionS3() (*session.Session, error) {
	s3Config := cnf.GetS3Config()
	sess, err := session.NewSession(s3Config)
	if err != nil {
		return nil, err
	}

	return sess, nil
}

func DownloadDocumentWithFilename(path string, filename string) (*string, error) {
	sess, err := createSessionS3()
	if err != nil {
		return nil, err
	}

	c := s3.New(sess)
	rcd := "attachment; filename =\"" + filename + "\""
	req, _ := c.GetObjectRequest(&s3.GetObjectInput{
		Bucket:                     aws.String(cnf.Env.DigitalOceanBucket),
		Key:                        aws.String(path),
		ResponseContentDisposition: &rcd,
	})

	signUrl, err := req.Presign(15 * time.Minute)
	if err != nil {
		return nil, err
	}

	return &signUrl, nil
}

func DownloadDocument(path string) (*string, error) {
	sess, err := createSessionS3()
	if err != nil {
		return nil, err
	}

	c := s3.New(sess)
	req, _ := c.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(cnf.Env.DigitalOceanBucket),
		Key:    aws.String(path),
	})

	urlStr, err := req.Presign(15 * time.Minute)
	if err != nil {
		return nil, err
	}
	return &urlStr, nil
}

func DownloadDocumentByte(path string) ([]byte, error) {
	sess, err := createSessionS3()

	if err != nil {
		return nil, err
	}

	c := s3.New(sess)
	result, err := c.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(cnf.Env.DigitalOceanBucket),
		Key:    aws.String(path),
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

func UploadDocument(fileBytes []byte, path string) error {
	sess, err := createSessionS3()
	if err != nil {
		return nil
	}

	uploader := s3manager.NewUploader(sess)
	fileReader := bytes.NewReader(fileBytes)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(cnf.Env.DigitalOceanBucket),
		Key:    aws.String(path),
		Body:   fileReader,
	})

	if err != nil {
		return err
	}

	return nil
}
