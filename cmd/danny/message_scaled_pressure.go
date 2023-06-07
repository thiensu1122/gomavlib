//autogenerated:yes
//nolint:revive,misspell,govet,lll
package danny
// The pressure readings for the typical setup of one absolute and differential pressure sensor. The units are as specified in each field.
type MessageScaledPressure struct {
    // Timestamp (time since system boot).
    TimeBootMs uint32
    // Absolute pressure
    PressAbs float32
    // Differential pressure 1
    PressDiff float32
    // Absolute pressure temperature
    Temperature int16
    // Differential pressure temperature (0, if not available). Report values of 0 (or 1) as 1 cdegC.
    TemperaturePressDiff int16 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageScaledPressure) GetID() uint32 {
    return 29
}
