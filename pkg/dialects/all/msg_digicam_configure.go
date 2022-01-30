//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Configure on-board Camera Control System.
type MessageDigicamConfigure struct {
	// System ID.
	TargetSystem uint8
	// Component ID.
	TargetComponent uint8
	// Mode enumeration from 1 to N //P, TV, AV, M, etc. (0 means ignore).
	Mode uint8
	// Divisor number //e.g. 1000 means 1/1000 (0 means ignore).
	ShutterSpeed uint16
	// F stop number x 10 //e.g. 28 means 2.8 (0 means ignore).
	Aperture uint8
	// ISO enumeration from 1 to N //e.g. 80, 100, 200, Etc (0 means ignore).
	Iso uint8
	// Exposure type enumeration from 1 to N (0 means ignore).
	ExposureType uint8
	// Command Identity (incremental loop: 0 to 255). //A command sent multiple times will be executed or pooled just once.
	CommandId uint8
	// Main engine cut-off time before camera trigger (0 means no cut-off).
	EngineCutOff uint8
	// Extra parameters enumeration (0 means ignore).
	ExtraParam uint8
	// Correspondent value to given extra_param.
	ExtraValue float32
}

// GetID implements the msg.Message interface.
func (*MessageDigicamConfigure) GetID() uint32 {
	return 154
}
