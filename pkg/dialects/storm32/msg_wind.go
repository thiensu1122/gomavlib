//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Wind estimation.
type MessageWind struct {
	// Wind direction (that wind is coming from).
	Direction float32
	// Wind speed in ground plane.
	Speed float32
	// Vertical wind speed.
	SpeedZ float32
}

// GetID implements the msg.Message interface.
func (*MessageWind) GetID() uint32 {
	return 168
}
