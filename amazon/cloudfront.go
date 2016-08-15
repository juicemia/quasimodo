package amazon

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/pborman/uuid"
)

// BustCache triggers a CloudFront cache invalidation for
// the passed in distribution id. This is useful for sites
// who use CloudFront to provide HTTPS.
func (s *Service) BustCache(cfid string) error {
	svc := cloudfront.New(s.session)

	bustRequest := &cloudfront.CreateInvalidationInput{
		DistributionId: aws.String(cfid),
		InvalidationBatch: &cloudfront.InvalidationBatch{
			CallerReference: aws.String(uuid.New()),
			Paths: &cloudfront.Paths{
				Quantity: aws.Int64(1),
				Items: []*string{
					aws.String("/*"),
				},
			},
		},
	}

	_, err := svc.CreateInvalidation(bustRequest)
	if err != nil {
		return err
	}

	return nil
}
