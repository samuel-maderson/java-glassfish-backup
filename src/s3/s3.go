package s3

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Upload(cfg aws.Config, bucketName string, filename string, filecontent string) {

	file, err := os.Open(filecontent)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	client := s3.NewFromConfig(cfg)

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		fmt.Println("failed to upload object", err)
		return
	}

}
