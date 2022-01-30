//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// The global position, as returned by the Global Positioning System (GPS). This is
// NOT the global position estimate of the sytem, but rather a RAW sensor value. See message GLOBAL_POSITION_INT for the global position estimate.
type MessageHilGps struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// 0-1: no fix, 2: 2D fix, 3: 3D fix. Some applications will not use the value of this field unless it is at least two, so always correctly fill in the fix.
	FixType uint8
	// Latitude (WGS84)
	Lat int32
	// Longitude (WGS84)
	Lon int32
	// Altitude (MSL). Positive for up.
	Alt int32
	// GPS HDOP horizontal dilution of position (unitless * 100). If unknown, set to: UINT16_MAX
	Eph uint16
	// GPS VDOP vertical dilution of position (unitless * 100). If unknown, set to: UINT16_MAX
	Epv uint16
	// GPS ground speed. If unknown, set to: UINT16_MAX
	Vel uint16
	// GPS velocity in north direction in earth-fixed NED frame
	Vn int16
	// GPS velocity in east direction in earth-fixed NED frame
	Ve int16
	// GPS velocity in down direction in earth-fixed NED frame
	Vd int16
	// Course over ground (NOT heading, but direction of movement), 0.0..359.99 degrees. If unknown, set to: UINT16_MAX
	Cog uint16
	// Number of satellites visible. If unknown, set to UINT8_MAX
	SatellitesVisible uint8
	// GPS ID (zero indexed). Used for multiple GPS inputs
	Id uint8 `mavext:"true"`
	// Yaw of vehicle relative to Earth's North, zero means not available, use 36000 for north
	Yaw uint16 `mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageHilGps) GetID() uint32 {
	return 113
}
