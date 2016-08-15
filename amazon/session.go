package amazon

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Service is our augmented s3 service
type Service struct {
	Bucket string
	Region string

	session *session.Session
}

// NewService creates a new session for quasimodo. It checks
// if the bucket is configured as a web site as well, since it's
// kinda pointless to do anything with a bucket that isn't.
func NewService(region string, bucket string) (*Service, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "quasimodo",
		Config: aws.Config{
			Region: aws.String(region),
		},
	})
	if err != nil {
		return nil, err
	}

	return &Service{
		Bucket:  bucket,
		Region:  region,
		session: sess,
	}, nil
}
