//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Emitted during mission execution when control reaches MAV_CMD_GROUP_END.
type MessageGroupEnd struct {
	// Mission-unique group id (from MAV_CMD_GROUP_END).
	GroupId uint32
	// CRC32 checksum of current plan for MAV_MISSION_TYPE_ALL. As defined in MISSION_CHECKSUM message.
	MissionChecksum uint32
	// Timestamp (UNIX Epoch time or time since system boot).
	// The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
}

// GetID implements the msg.Message interface.
func (*MessageGroupEnd) GetID() uint32 {
	return 415
}
