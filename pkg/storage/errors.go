package storage

import "fmt"

const (
	ErrCodeKeyNotFound int = iota + 1
	ErrCodeKeyExists
)

var errorCodeToMessage = map[int]string{
	ErrCodeKeyNotFound: "key not found",
	ErrCodeKeyExists:   "key exists",
}

func NewKeyNotFoundError(key string, rv uint64) *StorageError {
	return &StorageError{
		Code:            ErrCodeKeyNotFound,
		Key:             key,
		ResourceVersion: rv,
	}
}

func NewKeyExistError(key string, rv uint64) *StorageError {
	return &StorageError{
		Code:            ErrCodeKeyExists,
		Key:             key,
		ResourceVersion: rv,
	}
}

type StorageError struct {
	Code            int
	Key             string
	ResourceVersion uint64
	AdditionErrMsg  string
}

func (e *StorageError) Error() string {
	return fmt.Sprintf("StorageError: %s, Code: %d, Key: %s, ResourceVersion: %d, AdditionalErrorMsg: %s",
		errorCodeToMessage[e.Code], e.Code, e.Key, e.ResourceVersion, e.AdditionErrMsg)
}
