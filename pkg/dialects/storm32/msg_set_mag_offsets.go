//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Set the magnetometer offsets
type MessageSetMagOffsets struct {
	// System ID.
	TargetSystem uint8
	// Component ID.
	TargetComponent uint8
	// Magnetometer X offset.
	MagOfsX int16
	// Magnetometer Y offset.
	MagOfsY int16
	// Magnetometer Z offset.
	MagOfsZ int16
}

// GetID implements the msg.Message interface.
func (*MessageSetMagOffsets) GetID() uint32 {
	return 151
}
