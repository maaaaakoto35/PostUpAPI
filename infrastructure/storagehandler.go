package infrastructure

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

// GetFederationToken this func is return token.
func GetFederationToken(userID string) (result *sts.GetFederationTokenOutput, err error) {
	svc := sts.New(session.New())
	policy := aws.String(getPolicy(userID))

	input := &sts.GetFederationTokenInput{
		DurationSeconds: aws.Int64(3600),
		Name:            aws.String(userID),
		Policy:          policy,
		Tags: []*sts.Tag{
			{
				Key:   aws.String("Project"),
				Value: aws.String("Pegausus"),
			},
			{
				Key:   aws.String("Cost-Center"),
				Value: aws.String("98765"),
			},
		},
	}

	result, err = svc.GetFederationToken(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case sts.ErrCodeMalformedPolicyDocumentException:
				fmt.Println(sts.ErrCodeMalformedPolicyDocumentException, aerr.Error())
			case sts.ErrCodePackedPolicyTooLargeException:
				fmt.Println(sts.ErrCodePackedPolicyTooLargeException, aerr.Error())
			case sts.ErrCodeRegionDisabledException:
				fmt.Println(sts.ErrCodeRegionDisabledException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}
	return
}

func getPolicy(userID string) string {
	return fmt.Sprintf(`{
		\"Version\": \"2012-10-17\",
		\"Statement\": [
			{
			\"Effect\": \"Allow\",
			\"Action\": [\"s3:ListBucket\"],
			\"Resource\": [\"arn:aws:s3:::post\"]
			},
			{
			\"Effect\": \"Allow\",
			\"Action\": [
				\"s3:PutObject\",
			],
			\"Resource\": [\"arn:aws:s3:::post/%s/*\"]
			}
		]
	}`, userID)
}
