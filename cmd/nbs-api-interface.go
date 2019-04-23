package cmd

import "context"

type BlockStorageLayer interface {
    Shutdown(ctx context.Context) error

	// TODO: We may need to create DescribeVolumes type, not the final response struct
	DescribeVolumes(ctx context.Context) (volumes DescribeVolumesResponse, err error)
}
