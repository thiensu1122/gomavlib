//autogenerated:yes
//nolint:revive,misspell,govet,lll
package danny
// Status generated by radio and injected into MAVLink stream.
type MessageRadioStatus struct {
    // Local (message sender) received signal strength indication in device-dependent units/scale. Values: [0-254], UINT8_MAX: invalid/unknown.
    Rssi uint8
    // Remote (message receiver) signal strength indication in device-dependent units/scale. Values: [0-254], UINT8_MAX: invalid/unknown.
    Remrssi uint8
    // Remaining free transmitter buffer space.
    Txbuf uint8
    // Local background noise level. These are device dependent RSSI values (scale as approx 2x dB on SiK radios). Values: [0-254], UINT8_MAX: invalid/unknown.
    Noise uint8
    // Remote background noise level. These are device dependent RSSI values (scale as approx 2x dB on SiK radios). Values: [0-254], UINT8_MAX: invalid/unknown.
    Remnoise uint8
    // Count of radio packet receive errors (since boot).
    Rxerrors uint16
    // Count of error corrected radio packets (since boot).
    Fixed uint16
}

// GetID implements the message.Message interface.
func (*MessageRadioStatus) GetID() uint32 {
    return 109
}
