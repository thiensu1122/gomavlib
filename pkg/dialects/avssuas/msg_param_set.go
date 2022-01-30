//autogenerated:yes
//nolint:revive,misspell,govet,lll
package avssuas

// Set a parameter value (write new value to permanent storage).
// The receiving component should acknowledge the new parameter value by broadcasting a PARAM_VALUE message (broadcasting ensures that multiple GCS all have an up-to-date list of all parameters). If the sending GCS did not receive a PARAM_VALUE within its timeout time, it should re-send the PARAM_SET message. The parameter microservice is documented at https://mavlink.io/en/services/parameter.html.
// PARAM_SET may also be called within the context of a transaction (started with MAV_CMD_PARAM_TRANSACTION). Within a transaction the receiving component should respond with PARAM_ACK_TRANSACTION to the setter component (instead of broadcasting PARAM_VALUE), and PARAM_SET should be re-sent if this is ACK not received.
type MessageParamSet struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Onboard parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamId string `mavlen:"16"`
	// Onboard parameter value
	ParamValue float32
	// Onboard parameter type.
	ParamType MAV_PARAM_TYPE `mavenum:"uint8"`
}

// GetID implements the msg.Message interface.
func (*MessageParamSet) GetID() uint32 {
	return 23
}
