//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// A rally point. Used to set a point when from GCS -> MAV. Also used to return a point from MAV -> GCS.
type MessageRallyPoint struct {
	// System ID.
	TargetSystem uint8
	// Component ID.
	TargetComponent uint8
	// Point index (first point is 0).
	Idx uint8
	// Total number of points (for sanity checking).
	Count uint8
	// Latitude of point.
	Lat int32
	// Longitude of point.
	Lng int32
	// Transit / loiter altitude relative to home.
	Alt int16
	// Break altitude relative to home.
	BreakAlt int16
	// Heading to aim for when landing.
	LandDir uint16
	// Configuration flags.
	Flags RALLY_FLAGS `mavenum:"uint8"`
}

// GetID implements the msg.Message interface.
func (*MessageRallyPoint) GetID() uint32 {
	return 175
}
