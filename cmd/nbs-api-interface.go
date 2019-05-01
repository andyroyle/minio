package cmd

import "context"

type BlockStorageLayer interface {
	Shutdown(ctx context.Context) error

	DescribeVolumes(ctx context.Context) (volumes []Volume, err error)
	CreateVolume(ctx context.Context, size int64, availabilityZone string, snapshotId string, iops int64, kmsKeyId string, volumeType string) (Volume, error)
	DeleteVolume(ctx context.Context, volumeId string) error
	AttachVolume(ctx context.Context, volumeId string, instanceId string, device string) error
}
