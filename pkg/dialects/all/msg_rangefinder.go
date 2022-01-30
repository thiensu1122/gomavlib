//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Rangefinder reporting.
type MessageRangefinder struct {
	// Distance.
	Distance float32
	// Raw voltage if available, zero otherwise.
	Voltage float32
}

// GetID implements the msg.Message interface.
func (*MessageRangefinder) GetID() uint32 {
	return 173
}
