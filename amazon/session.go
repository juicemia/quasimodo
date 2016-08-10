package amazon

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3 is our augmented s3 service
type S3 struct {
	Bucket string
	Region string
	s3.S3
}

// NewS3Session creates a new session for quasimodo. It checks
// if the bucket is configured as a web site as well, since it's
// kinda pointless to do anything with a bucket that isn't.
func NewS3Session(region string, bucket string) (*S3, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "quasimodo",
		Config: aws.Config{
			Region: aws.String(region),
		},
	})
	if err != nil {
		return nil, err
	}

	svc := &S3{bucket, region, *s3.New(sess)}

	wwwConfigRequest := &s3.GetBucketWebsiteInput{
		Bucket: aws.String(svc.Bucket),
	}

	_, err = svc.GetBucketWebsite(wwwConfigRequest)
	if err != nil {
		return nil, err
	}

	return svc, nil
}
