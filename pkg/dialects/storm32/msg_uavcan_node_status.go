//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// General status information of an UAVCAN node. Please refer to the definition of the UAVCAN message "uavcan.protocol.NodeStatus" for the background information. The UAVCAN specification is available at http://uavcan.org.
type MessageUavcanNodeStatus struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Time since the start-up of the node.
	UptimeSec uint32
	// Generalized node health status.
	Health UAVCAN_NODE_HEALTH `mavenum:"uint8"`
	// Generalized operating mode.
	Mode UAVCAN_NODE_MODE `mavenum:"uint8"`
	// Not used currently.
	SubMode uint8
	// Vendor-specific status information.
	VendorSpecificStatusCode uint16
}

// GetID implements the msg.Message interface.
func (*MessageUavcanNodeStatus) GetID() uint32 {
	return 310
}
