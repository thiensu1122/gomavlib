//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Read configured OSD parameter reply.
type MessageOsdParamShowConfigReply struct {
	// Request ID - copied from request.
	RequestId uint32
	// Config error type.
	Result OSD_PARAM_CONFIG_ERROR `mavenum:"uint8"`
	// Onboard parameter id, terminated by NULL if the length is less than 16 human-readable chars and WITHOUT null termination (NULL) byte if the length is exactly 16 chars - applications have to provide 16+1 bytes storage if the ID is stored as string
	ParamId string `mavlen:"16"`
	// Config type.
	ConfigType OSD_PARAM_CONFIG_TYPE `mavenum:"uint8"`
	// OSD parameter minimum value.
	MinValue float32
	// OSD parameter maximum value.
	MaxValue float32
	// OSD parameter increment.
	Increment float32
}

// GetID implements the msg.Message interface.
func (*MessageOsdParamShowConfigReply) GetID() uint32 {
	return 11036
}
