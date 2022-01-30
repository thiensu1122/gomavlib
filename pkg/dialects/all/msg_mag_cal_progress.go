//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Reports progress of compass calibration.
type MessageMagCalProgress struct {
	// Compass being calibrated.
	CompassId uint8
	// Bitmask of compasses being calibrated.
	CalMask uint8
	// Calibration Status.
	CalStatus MAG_CAL_STATUS `mavenum:"uint8"`
	// Attempt number.
	Attempt uint8
	// Completion percentage.
	CompletionPct uint8
	// Bitmask of sphere sections (see http://en.wikipedia.org/wiki/Geodesic_grid).
	CompletionMask [10]uint8
	// Body frame direction vector for display.
	DirectionX float32
	// Body frame direction vector for display.
	DirectionY float32
	// Body frame direction vector for display.
	DirectionZ float32
}

// GetID implements the msg.Message interface.
func (*MessageMagCalProgress) GetID() uint32 {
	return 191
}
