// Code generated by sdkgen. DO NOT EDIT.

package redis

import (
	"context"

	"google.golang.org/grpc"
)

// Redis provides access to "redis" component of Yandex.Cloud
type Redis struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewRedis creates instance of Redis
func NewRedis(g func(ctx context.Context) (*grpc.ClientConn, error)) *Redis {
	return &Redis{g}
}

// Cluster gets ClusterService client
func (r *Redis) Cluster() *ClusterServiceClient {
	return &ClusterServiceClient{getConn: r.getConn}
}

// ResourcePreset gets ResourcePresetService client
func (r *Redis) ResourcePreset() *ResourcePresetServiceClient {
	return &ResourcePresetServiceClient{getConn: r.getConn}
}

// Backup gets BackupService client
func (r *Redis) Backup() *BackupServiceClient {
	return &BackupServiceClient{getConn: r.getConn}
}

// User gets UserService client
func (r *Redis) User() *UserServiceClient {
	return &UserServiceClient{getConn: r.getConn}
}
