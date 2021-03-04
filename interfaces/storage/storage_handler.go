package storage

import "github.com/aws/aws-sdk-go/service/sts"

// StorageHandler this interface is connecting infrastructure/storage.
type StorageHandler interface {
	GetFederationToken(string) (*sts.GetFederationTokenOutput, error)
}
