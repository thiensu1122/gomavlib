//autogenerated:yes
//nolint:revive,misspell,govet,lll
package danny
// Vibration levels and accelerometer clipping
type MessageVibration struct {
    // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
    TimeUsec uint64
    // Vibration levels on X-axis
    VibrationX float32
    // Vibration levels on Y-axis
    VibrationY float32
    // Vibration levels on Z-axis
    VibrationZ float32
    // first accelerometer clipping count
    Clipping_0 uint32
    // second accelerometer clipping count
    Clipping_1 uint32
    // third accelerometer clipping count
    Clipping_2 uint32
}

// GetID implements the message.Message interface.
func (*MessageVibration) GetID() uint32 {
    return 241
}
