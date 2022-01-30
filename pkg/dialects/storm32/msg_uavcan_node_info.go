//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// General information describing a particular UAVCAN node. Please refer to the definition of the UAVCAN service "uavcan.protocol.GetNodeInfo" for the background information. This message should be emitted by the system whenever a new node appears online, or an existing node reboots. Additionally, it can be emitted upon request from the other end of the MAVLink channel (see MAV_CMD_UAVCAN_GET_NODE_INFO). It is also not prohibited to emit this message unconditionally at a low frequency. The UAVCAN specification is available at http://uavcan.org.
type MessageUavcanNodeInfo struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Time since the start-up of the node.
	UptimeSec uint32
	// Node name string. For example, "sapog.px4.io".
	Name string `mavlen:"80"`
	// Hardware major version number.
	HwVersionMajor uint8
	// Hardware minor version number.
	HwVersionMinor uint8
	// Hardware unique 128-bit ID.
	HwUniqueId [16]uint8
	// Software major version number.
	SwVersionMajor uint8
	// Software minor version number.
	SwVersionMinor uint8
	// Version control system (VCS) revision identifier (e.g. git short commit hash). 0 if unknown.
	SwVcsCommit uint32
}

// GetID implements the msg.Message interface.
func (*MessageUavcanNodeInfo) GetID() uint32 {
	return 311
}
