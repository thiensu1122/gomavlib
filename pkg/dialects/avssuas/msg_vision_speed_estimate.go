//autogenerated:yes
//nolint:revive,misspell,govet,lll
package avssuas

// Speed estimate from a vision source.
type MessageVisionSpeedEstimate struct {
	// Timestamp (UNIX time or time since system boot)
	Usec uint64
	// Global X speed
	X float32
	// Global Y speed
	Y float32
	// Global Z speed
	Z float32
	// Row-major representation of 3x3 linear velocity covariance matrix (states: vx, vy, vz; 1st three entries - 1st row, etc.). If unknown, assign NaN value to first element in the array.
	Covariance [9]float32 `mavext:"true"`
	// Estimate reset counter. This should be incremented when the estimate resets in any of the dimensions (position, velocity, attitude, angular speed). This is designed to be used when e.g an external SLAM system detects a loop-closure and the estimate jumps.
	ResetCounter uint8 `mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageVisionSpeedEstimate) GetID() uint32 {
	return 103
}
