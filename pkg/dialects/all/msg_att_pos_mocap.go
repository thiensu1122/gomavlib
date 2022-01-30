//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Motion capture attitude and position
type MessageAttPosMocap struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Attitude quaternion (w, x, y, z order, zero-rotation is 1, 0, 0, 0)
	Q [4]float32
	// X position (NED)
	X float32
	// Y position (NED)
	Y float32
	// Z position (NED)
	Z float32
	// Row-major representation of a pose 6x6 cross-covariance matrix upper right triangle (states: x, y, z, roll, pitch, yaw; first six entries are the first ROW, next five entries are the second ROW, etc.). If unknown, assign NaN value to first element in the array.
	Covariance [21]float32 `mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageAttPosMocap) GetID() uint32 {
	return 138
}
