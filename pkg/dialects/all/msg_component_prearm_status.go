//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Message reporting the status of the prearm checks. The flags are component specific.
type MessageComponentPrearmStatus struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Currently enabled prearm checks. 0 means no checks are being performed, UINT32_MAX means not known.
	EnabledFlags uint32
	// Currently not passed prearm checks. 0 means all checks have been passed.
	FailFlags uint32
}

// GetID implements the msg.Message interface.
func (*MessageComponentPrearmStatus) GetID() uint32 {
	return 60025
}
