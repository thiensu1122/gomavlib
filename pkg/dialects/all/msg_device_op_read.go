//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Read registers for a device.
type MessageDeviceOpRead struct {
	// System ID.
	TargetSystem uint8
	// Component ID.
	TargetComponent uint8
	// Request ID - copied to reply.
	RequestId uint32
	// The bus type.
	Bustype DEVICE_OP_BUSTYPE `mavenum:"uint8"`
	// Bus number.
	Bus uint8
	// Bus address.
	Address uint8
	// Name of device on bus (for SPI).
	Busname string `mavlen:"40"`
	// First register to read.
	Regstart uint8
	// Count of registers to read.
	Count uint8
	// Bank number.
	Bank uint8 `mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageDeviceOpRead) GetID() uint32 {
	return 11000
}
