//autogenerated:yes
//nolint:revive,misspell,govet,lll
package avssuas

// Bind a RC channel to a parameter. The parameter should change according to the RC channel value.
type MessageParamMapRc struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Onboard parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamId string `mavlen:"16"`
	// Parameter index. Send -1 to use the param ID field as identifier (else the param id will be ignored), send -2 to disable any existing map for this rc_channel_index.
	ParamIndex int16
	// Index of parameter RC channel. Not equal to the RC channel id. Typically corresponds to a potentiometer-knob on the RC.
	ParameterRcChannelIndex uint8
	// Initial parameter value
	ParamValue0 float32
	// Scale, maps the RC range [-1, 1] to a parameter value
	Scale float32
	// Minimum param value. The protocol does not define if this overwrites an onboard minimum value. (Depends on implementation)
	ParamValueMin float32
	// Maximum param value. The protocol does not define if this overwrites an onboard maximum value. (Depends on implementation)
	ParamValueMax float32
}

// GetID implements the msg.Message interface.
func (*MessageParamMapRc) GetID() uint32 {
	return 50
}
