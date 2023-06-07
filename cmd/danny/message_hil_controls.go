//autogenerated:yes
//nolint:revive,misspell,govet,lll
package danny
// Sent from autopilot to simulation. Hardware in the loop control outputs
type MessageHilControls struct {
    // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
    TimeUsec uint64
    // Control output -1 .. 1
    RollAilerons float32
    // Control output -1 .. 1
    PitchElevator float32
    // Control output -1 .. 1
    YawRudder float32
    // Throttle 0 .. 1
    Throttle float32
    // Aux 1, -1 .. 1
    Aux1 float32
    // Aux 2, -1 .. 1
    Aux2 float32
    // Aux 3, -1 .. 1
    Aux3 float32
    // Aux 4, -1 .. 1
    Aux4 float32
    // System mode.
    Mode MAV_MODE `mavenum:"uint8"`
    // Navigation mode (MAV_NAV_MODE)
    NavMode uint8
}

// GetID implements the message.Message interface.
func (*MessageHilControls) GetID() uint32 {
    return 91
}
