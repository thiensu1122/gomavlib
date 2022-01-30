//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Extended EKF state estimates for ASLUAVs
type MessageEkfExt struct {
	// Time since system start
	Timestamp uint64
	// Magnitude of wind velocity (in lateral inertial plane)
	Windspeed float32 `mavname:"Windspeed"`
	// Wind heading angle from North
	Winddir float32 `mavname:"WindDir"`
	// Z (Down) component of inertial wind velocity
	Windz float32 `mavname:"WindZ"`
	// Magnitude of air velocity
	Airspeed float32 `mavname:"Airspeed"`
	// Sideslip angle
	Beta float32
	// Angle of attack
	Alpha float32
}

// GetID implements the msg.Message interface.
func (*MessageEkfExt) GetID() uint32 {
	return 8007
}
