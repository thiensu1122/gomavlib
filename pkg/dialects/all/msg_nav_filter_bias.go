//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Accelerometer and Gyro biases from the navigation filter
type MessageNavFilterBias struct {
	// Timestamp (microseconds)
	Usec uint64
	// b_f[0]
	Accel_0 float32
	// b_f[1]
	Accel_1 float32
	// b_f[2]
	Accel_2 float32
	// b_f[0]
	Gyro_0 float32
	// b_f[1]
	Gyro_1 float32
	// b_f[2]
	Gyro_2 float32
}

// GetID implements the msg.Message interface.
func (*MessageNavFilterBias) GetID() uint32 {
	return 220
}
