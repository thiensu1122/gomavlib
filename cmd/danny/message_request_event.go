//autogenerated:yes
//nolint:revive,misspell,govet,lll
package danny
// Request one or more events to be (re-)sent. If first_sequence==last_sequence, only a single event is requested. Note that first_sequence can be larger than last_sequence (because the sequence number can wrap). Each sequence will trigger an EVENT or EVENT_ERROR response.
type MessageRequestEvent struct {
    // System ID
    TargetSystem uint8
    // Component ID
    TargetComponent uint8
    // First sequence number of the requested event.
    FirstSequence uint16
    // Last sequence number of the requested event.
    LastSequence uint16
}

// GetID implements the message.Message interface.
func (*MessageRequestEvent) GetID() uint32 {
    return 412
}
