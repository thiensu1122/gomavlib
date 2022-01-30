//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// The location and information of an AIS vessel
type MessageAisVessel struct {
	// Mobile Marine Service Identifier, 9 decimal digits
	Mmsi uint32 `mavname:"MMSI"`
	// Latitude
	Lat int32
	// Longitude
	Lon int32
	// Course over ground
	Cog uint16 `mavname:"COG"`
	// True heading
	Heading uint16
	// Speed over ground
	Velocity uint16
	// Turn rate
	TurnRate int8
	// Navigational status
	NavigationalStatus AIS_NAV_STATUS `mavenum:"uint8"`
	// Type of vessels
	Type AIS_TYPE `mavenum:"uint8"`
	// Distance from lat/lon location to bow
	DimensionBow uint16
	// Distance from lat/lon location to stern
	DimensionStern uint16
	// Distance from lat/lon location to port side
	DimensionPort uint8
	// Distance from lat/lon location to starboard side
	DimensionStarboard uint8
	// The vessel callsign
	Callsign string `mavlen:"7"`
	// The vessel name
	Name string `mavlen:"20"`
	// Time since last communication in seconds
	Tslc uint16
	// Bitmask to indicate various statuses including valid data fields
	Flags AIS_FLAGS `mavenum:"uint16"`
}

// GetID implements the msg.Message interface.
func (*MessageAisVessel) GetID() uint32 {
	return 301
}
