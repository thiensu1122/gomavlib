//autogenerated:yes
//nolint:revive,misspell,govet,lll
package danny
// Modify the filter of what CAN messages to forward over the mavlink. This can be used to make CAN forwarding work well on low bandwidth links. The filtering is applied on bits 8 to 24 of the CAN id (2nd and 3rd bytes) which corresponds to the DroneCAN message ID for DroneCAN. Filters with more than 16 IDs can be constructed by sending multiple CAN_FILTER_MODIFY messages.
type MessageCanFilterModify struct {
    // System ID.
    TargetSystem uint8
    // Component ID.
    TargetComponent uint8
    // bus number
    Bus uint8
    // what operation to perform on the filter list. See CAN_FILTER_OP enum.
    Operation CAN_FILTER_OP `mavenum:"uint8"`
    // number of IDs in filter list
    NumIds uint8
    // filter IDs, length num_ids
    Ids [16]uint16
}

// GetID implements the message.Message interface.
func (*MessageCanFilterModify) GetID() uint32 {
    return 388
}
