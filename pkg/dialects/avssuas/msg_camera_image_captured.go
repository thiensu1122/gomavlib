//autogenerated:yes
//nolint:revive,misspell,govet,lll
package avssuas

// Information about a captured image. This is emitted every time a message is captured.
// MAV_CMD_REQUEST_MESSAGE can be used to (re)request this message for a specific sequence number or range of sequence numbers:
// MAV_CMD_REQUEST_MESSAGE.param2 indicates the sequence number the first image to send, or set to -1 to send the message for all sequence numbers.
// MAV_CMD_REQUEST_MESSAGE.param3 is used to specify a range of messages to send:
// set to 0 (default) to send just the the message for the sequence number in param 2,
// set to -1 to send the message for the sequence number in param 2 and all the following sequence numbers,
// set to the sequence number of the final message in the range.
type MessageCameraImageCaptured struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Timestamp (time since UNIX epoch) in UTC. 0 for unknown.
	TimeUtc uint64
	// Deprecated/unused. Component IDs are used to differentiate multiple cameras.
	CameraId uint8
	// Latitude where image was taken
	Lat int32
	// Longitude where capture was taken
	Lon int32
	// Altitude (MSL) where image was taken
	Alt int32
	// Altitude above ground
	RelativeAlt int32
	// Quaternion of camera orientation (w, x, y, z order, zero-rotation is 1, 0, 0, 0)
	Q [4]float32
	// Zero based index of this image (i.e. a new image will have index CAMERA_CAPTURE_STATUS.image count -1)
	ImageIndex int32
	// Boolean indicating success (1) or failure (0) while capturing this image.
	CaptureResult int8
	// URL of image taken. Either local storage or http://foo.jpg if camera provides an HTTP interface.
	FileUrl string `mavlen:"205"`
}

// GetID implements the msg.Message interface.
func (*MessageCameraImageCaptured) GetID() uint32 {
	return 263
}
