//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Sets a desired vehicle attitude. Used by an external controller to command the vehicle (manual controller or other system).
type MessageSetAttitudeTarget struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Bitmap to indicate which dimensions should be ignored by the vehicle.
	TypeMask ATTITUDE_TARGET_TYPEMASK `mavenum:"uint8"`
	// Attitude quaternion (w, x, y, z order, zero-rotation is 1, 0, 0, 0)
	Q [4]float32
	// Body roll rate
	BodyRollRate float32
	// Body pitch rate
	BodyPitchRate float32
	// Body yaw rate
	BodyYawRate float32
	// Collective thrust, normalized to 0 .. 1 (-1 .. 1 for vehicles capable of reverse trust)
	Thrust float32
	// 3D thrust setpoint in the body NED frame, normalized to -1 .. 1
	ThrustBody [3]float32 `mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageSetAttitudeTarget) GetID() uint32 {
	return 82
}
