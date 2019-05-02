package cmd

import "context"

type BSObjects struct {
	// TODO: If this service need to store states, this is the struct to store internal states.

}

func NewBlockStorageLayer() BlockStorageLayer {
	bs := &BSObjects{}
	return bs
}

// Shutdown - should be called when process shuts down.
func (fs *BSObjects) Shutdown(ctx context.Context) error {
	// TOOD: This is not bound yet.
	// Cleanup
	return nil
}

func (fs *BSObjects) DescribeVolumes(ctx context.Context) ([]Volume, error) {
	// TODO: Remove below constant return
	// TODO: Query underlay storage layer, get the list of volumes and return it.

	volumes := []Volume{
		{
			VolumeId:         "asdf",
			Size:             100,
			AvailabilityZone: "us-east-1",
			Status:           "in-use",
			CreateTime:       "2018-04-12T23:17:43.439Z",
			AttachmentSet: VolumeAttachmentSet{
				Volumes: []VolumeAttachment{
					{
						VolumeId:            "vol-0cc48e9546852d931",
						InstanceId:          "i-0df88542e60b86536",
						Device:              "/dev/xvdg",
						Status:              "attached",
						AttachTime:          "2019-03-07T21:49:43.000Z",
						DeleteOnTermination: "false",
					},
				},
			},
		},
	}
	return volumes, nil
}

func (fs *BSObjects) CreateVolume(ctx context.Context, size int64, availabilityZone string, snapshotId string, iops int64, kmsKeyId string, volumeType string) (Volume, error) {
	// TODO: Remove below constant return

	volume := Volume{
		VolumeId:         "asdf",
		Size:             100,
		AvailabilityZone: "us-east-1",
		Status:           "in-use",
		CreateTime:       "2018-04-12T23:17:43.439Z",
		AttachmentSet: VolumeAttachmentSet{
			Volumes: []VolumeAttachment{
				{
					VolumeId:            "vol-0cc48e9546852d931",
					InstanceId:          "i-0df88542e60b86536",
					Device:              "/dev/xvdg",
					Status:              "attached",
					AttachTime:          "2019-03-07T21:49:43.000Z",
					DeleteOnTermination: "false",
				},
			},
		},
	}
	return volume, nil
}

func (fs *BSObjects) DeleteVolume(ctx context.Context, volumeId string) error {
	// TODO: Remove below constant return
	return nil
}

func (fs *BSObjects) AttachVolume(ctx context.Context, volumeId string, instanceId string, device string) error {
	// TODO: Remove below constant return
	return nil
}
