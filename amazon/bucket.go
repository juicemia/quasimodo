package amazon

import (
	"github.com/HugoSTorres/quasimodo/fs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Clean checks the service's bucket for any content
// and clears it out. In AWS there's no way to just clear a
// bucket. We have to list out all the objects in the bucket
// and send delete requests for each one.
func (s *Service) Clean() error {
	svc := *s3.New(s.session)

	lsRequest := &s3.ListObjectsInput{
		Bucket: aws.String(s.Bucket),
	}

	lsResponse, err := svc.ListObjects(lsRequest)
	if err != nil {
		return err
	}

	var deleteObjects []*s3.ObjectIdentifier
	for _, obj := range lsResponse.Contents {
		deleteObjects = append(deleteObjects, &s3.ObjectIdentifier{
			Key: obj.Key,
		})
	}

	if len(deleteObjects) > 0 {
		deleteRequest := &s3.DeleteObjectsInput{
			Bucket: aws.String(s.Bucket),
			Delete: &s3.Delete{
				Objects: deleteObjects,
			},
		}

		_, err = svc.DeleteObjects(deleteRequest)
		if err != nil {
			return err
		}
	}

	return nil
}

// Publish uploads the hugo site's public folder to the S3 bucket.
func (s *Service) Publish() error {
	svc := *s3.New(s.session)

	f, err := fs.GetSite()
	if err != nil {
		return err
	}

	uploadRequests, err := getS3Inputs(s.Bucket, f)
	if err != nil {
		return err
	}

	for _, req := range uploadRequests {
		_, err := svc.PutObject(req)
		if err != nil {
			return err
		}
	}

	return nil
}

// CheckWWW checks the Bucket's website configuration. If the bucket
// isn't configured as a server, it will return an error.
func (s *Service) CheckWWW() error {
	svc := *s3.New(s.session)

	wwwConfigRequest := &s3.GetBucketWebsiteInput{
		Bucket: aws.String(s.Bucket),
	}

	_, err := svc.GetBucketWebsite(wwwConfigRequest)
	if err != nil {
		return err
	}

	return nil
}

func getS3Inputs(bucket string, files []fs.File) (acc []*s3.PutObjectInput, err error) {
	for _, f := range files {
		acc = append(acc, &s3.PutObjectInput{
			Bucket:      aws.String(bucket),
			Key:         aws.String(f.Path),
			Body:        f.Data,
			ContentType: aws.String(f.Mime),
		})
	}

	return
}
