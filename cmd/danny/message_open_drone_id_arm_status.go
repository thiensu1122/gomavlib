//autogenerated:yes
//nolint:revive,misspell,govet,lll
package danny
// Transmitter (remote ID system) is enabled and ready to start sending location and other required information. This is streamed by transmitter. A flight controller uses it as a condition to arm.
type MessageOpenDroneIdArmStatus struct {
    // Status level indicating if arming is allowed.
    Status MAV_ODID_ARM_STATUS `mavenum:"uint8"`
    // Text error message, should be empty if status is good to arm. Fill with nulls in unused portion.
    Error string `mavlen:"50"`
}

// GetID implements the message.Message interface.
func (*MessageOpenDroneIdArmStatus) GetID() uint32 {
    return 12918
}
