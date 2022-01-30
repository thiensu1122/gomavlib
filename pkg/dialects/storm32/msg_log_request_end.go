//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// Stop log transfer and resume normal logging
type MessageLogRequestEnd struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
}

// GetID implements the msg.Message interface.
func (*MessageLogRequestEnd) GetID() uint32 {
	return 122
}
