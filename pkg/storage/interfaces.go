package storage

import (
	"context"

	metav1 "ooneko.github.com/vehicle-insight/pkg/apis/meta/v1"
)

type Interface interface {
	Create(ctx context.Context, key string, obj metav1.Object, ttl uint64) error

	Delete(ctx context.Context, key string, out interface{}) error

	Get(ctx context.Context, key string, resourceVersion string, out metav1.Object, ignoreNotFound bool)

	List(ctx context.Context, key string, resourceVersion string, listObj interface{}) error

	Update(ctx context.Context, key string, obj, out interface{}) error
}
