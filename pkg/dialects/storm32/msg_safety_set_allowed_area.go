//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Set a safety zone (volume), which is defined by two corners of a cube. This message can be used to tell the MAV which setpoints/waypoints to accept and which to reject. Safety areas are often enforced by national or competition regulations.
type MessageSafetySetAllowedArea struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Coordinate frame. Can be either global, GPS, right-handed with Z axis up or local, right handed, Z axis down.
	Frame MAV_FRAME `mavenum:"uint8"`
	// x position 1 / Latitude 1
	P1x float32
	// y position 1 / Longitude 1
	P1y float32
	// z position 1 / Altitude 1
	P1z float32
	// x position 2 / Latitude 2
	P2x float32
	// y position 2 / Longitude 2
	P2y float32
	// z position 2 / Altitude 2
	P2z float32
}

// GetID implements the msg.Message interface.
func (*MessageSafetySetAllowedArea) GetID() uint32 {
	return 54
}
