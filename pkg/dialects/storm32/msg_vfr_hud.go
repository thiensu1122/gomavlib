//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Metrics typically displayed on a HUD for fixed wing aircraft.
type MessageVfrHud struct {
	// Vehicle speed in form appropriate for vehicle type. For standard aircraft this is typically calibrated airspeed (CAS) or indicated airspeed (IAS) - either of which can be used by a pilot to estimate stall speed.
	Airspeed float32
	// Current ground speed.
	Groundspeed float32
	// Current heading in compass units (0-360, 0=north).
	Heading int16
	// Current throttle setting (0 to 100).
	Throttle uint16
	// Current altitude (MSL).
	Alt float32
	// Current climb rate.
	Climb float32
}

// GetID implements the msg.Message interface.
func (*MessageVfrHud) GetID() uint32 {
	return 74
}
