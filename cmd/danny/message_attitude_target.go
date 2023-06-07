//autogenerated:yes
//nolint:revive,misspell,govet,lll
package danny
// Reports the current commanded attitude of the vehicle as specified by the autopilot. This should match the commands sent in a SET_ATTITUDE_TARGET message if the vehicle is being controlled this way.
type MessageAttitudeTarget struct {
    // Timestamp (time since system boot).
    TimeBootMs uint32
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
}

// GetID implements the message.Message interface.
func (*MessageAttitudeTarget) GetID() uint32 {
    return 83
}
