//autogenerated:yes
//nolint:revive,misspell,govet,lll
package danny
// Information about the status of a capture. Can be requested with a MAV_CMD_REQUEST_MESSAGE command.
type MessageCameraCaptureStatus struct {
    // Timestamp (time since system boot).
    TimeBootMs uint32
    // Current status of image capturing (0: idle, 1: capture in progress, 2: interval set but idle, 3: interval set and capture in progress)
    ImageStatus uint8
    // Current status of video capturing (0: idle, 1: capture in progress)
    VideoStatus uint8
    // Image capture interval
    ImageInterval float32
    // Elapsed time since recording started (0: Not supported/available). A GCS should compute recording time and use non-zero values of this field to correct any discrepancy.
    RecordingTimeMs uint32
    // Available storage capacity.
    AvailableCapacity float32
    // Total number of images captured ('forever', or until reset using MAV_CMD_STORAGE_FORMAT).
    ImageCount int32 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageCameraCaptureStatus) GetID() uint32 {
    return 262
}
