package db

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"gophkeeper/internal/logger"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type s3Storage struct {
	client *s3.Client
}

/*
 	1. install minIO. brew install minio/stable/minio
	2. set data directory. mkdir -p ~/minio-data
	3. start minio. minio server ~/minio-data

	/Users/dglazkov/gophkeeper/text.txt - to save text in bucket
*/

func NewS3() (*s3Storage, error) {
	endpoint := "http://127.0.0.1:9000"
	accessKey := "minioadmin"
	secretKey := "minioadmin"

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:           endpoint,
				SigningRegion: "us-east-1",
			}, nil
		})),
	)
	if err != nil {
		logger.Log.Errorf("failed to load configuration: %v", err)
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	return &s3Storage{client: client}, nil
}

func (s *s3Storage) SaveBinaryData(ctx context.Context, userId, key string, data []byte, metadata string) error {
	bucketName := fmt.Sprintf("userbucket-%s", userId)

	_, err := s.client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})

	if err != nil {
		var noSuchBucket *types.NoSuchBucket
		if errors.As(err, &noSuchBucket) {
			_, err = s.client.CreateBucket(ctx, &s3.CreateBucketInput{
				Bucket: aws.String(bucketName),
			})

			if err != nil {
				logger.Log.Errorf("failed to create bucket: %v", err)
				return err
			}
		} else {
			logger.Log.Errorf("failed to check bucket: %v", err)
		}
	}

	_, err = s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
		Metadata: map[string]string{
			"default": metadata,
		},
	})

	return err
}

func (s *s3Storage) GetBinaryData(ctx context.Context, userId, key string) ([]byte, string, error) {
	bucketName := fmt.Sprintf("userbucket-%s", userId)

	res, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, "", err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)

	if err != nil {
		logger.Log.Errorf("failed to read object: %v", err)
		return nil, "", err
	}

	return buf.Bytes(), res.Metadata["default"], nil
}

func (s *s3Storage) DeleteBinaryData(ctx context.Context, userId, key string) error {
	bucketName := fmt.Sprintf("userbucket-%s", userId)

	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})

	return err
}
