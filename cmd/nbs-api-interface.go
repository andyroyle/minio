package cmd

import "context"

type BlockStorageLayer interface {
    Shutdown(ctx context.Context) error

	DescribeVolumes(ctx context.Context) (volumes []Volume, err error)
}
