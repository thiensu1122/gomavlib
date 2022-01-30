//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Message reporting the current status of a gimbal device. This message should be broadcasted by a gimbal device component at a low regular rate (e.g. 4 Hz). For higher rates it should be emitted with a target.
type MessageStorm32GimbalDeviceStatus struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Gimbal device flags currently applied.
	Flags MAV_STORM32_GIMBAL_DEVICE_FLAGS `mavenum:"uint16"`
	// Quaternion components, w, x, y, z (1 0 0 0 is the null-rotation). The frame depends on the STORM32_GIMBAL_DEVICE_FLAGS_YAW_ABSOLUTE flag.
	Q [4]float32
	// X component of angular velocity (NaN if unknown).
	AngularVelocityX float32
	// Y component of angular velocity (NaN if unknown).
	AngularVelocityY float32
	// Z component of angular velocity (the frame depends on the STORM32_GIMBAL_DEVICE_FLAGS_YAW_ABSOLUTE flag, NaN if unknown).
	AngularVelocityZ float32
	// Yaw in absolute frame relative to Earth's North, north is 0 (NaN if unknown).
	YawAbsolute float32
	// Failure flags (0 for no failure).
	FailureFlags GIMBAL_DEVICE_ERROR_FLAGS `mavenum:"uint16"`
}

// GetID implements the msg.Message interface.
func (*MessageStorm32GimbalDeviceStatus) GetID() uint32 {
	return 60001
}
