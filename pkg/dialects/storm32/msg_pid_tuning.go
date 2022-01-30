//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// PID tuning information.
type MessagePidTuning struct {
	// Axis.
	Axis PID_TUNING_AXIS `mavenum:"uint8"`
	// Desired rate.
	Desired float32
	// Achieved rate.
	Achieved float32
	// FF component.
	Ff float32 `mavname:"FF"`
	// P component.
	P float32 `mavname:"P"`
	// I component.
	I float32 `mavname:"I"`
	// D component.
	D float32 `mavname:"D"`
	// Slew rate.
	Srate float32 `mavext:"true" mavname:"SRate"`
	// P/D oscillation modifier.
	Pdmod float32 `mavext:"true" mavname:"PDmod"`
}

// GetID implements the msg.Message interface.
func (*MessagePidTuning) GetID() uint32 {
	return 194
}
