package infrastructure

import (
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sts"
)

// StorageHandler struct
type StorageHandler struct{}

// GetFederationToken this func is return token.
func GetFederationToken(userID string) (result *sts.GetFederationTokenOutput, err error) {
	svc := sts.New(session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})))
	policy := aws.String(getPolicy(userID))

	input := &sts.GetFederationTokenInput{
		DurationSeconds: aws.Int64(3600),
		Name:            aws.String(userID),
		Policy:          policy,
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
			\"Resource\": [\"arn:aws:s3:::post/output/%s/*\"]
			}
		]
	}`, userID)
}

// GetPresignedURL this func is to return pre-signed URL.
func GetPresignedURL(userID string, num int) (url string, err error) {
	svc := s3.New(session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})))

	// request
	fileHash, err := makeRandomStr(10)
	key := fmt.Sprintf("post/output/%s/%d-%s", userID, num, fileHash)
	c, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String("post"),
		Key:    aws.String(key),
	})

	// url
	url, err = c.Presign(15 * time.Minute)
	if err != nil {
		fmt.Println("error presigning request", err)
		return
	}
	return
}

// makeRandomStr this func is making random string.
func makeRandomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}
