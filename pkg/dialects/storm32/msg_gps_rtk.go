//autogenerated:yes
//nolint:revive,misspell,govet,lll
package storm32

// RTK GPS data. Gives information on the relative baseline calculation the GPS is reporting
type MessageGpsRtk struct {
	// Time since boot of last baseline message received.
	TimeLastBaselineMs uint32
	// Identification of connected RTK receiver.
	RtkReceiverId uint8
	// GPS Week Number of last baseline
	Wn uint16
	// GPS Time of Week of last baseline
	Tow uint32
	// GPS-specific health report for RTK data.
	RtkHealth uint8
	// Rate of baseline messages being received by GPS
	RtkRate uint8
	// Current number of sats used for RTK calculation.
	Nsats uint8
	// Coordinate system of baseline
	BaselineCoordsType RTK_BASELINE_COORDINATE_SYSTEM `mavenum:"uint8"`
	// Current baseline in ECEF x or NED north component.
	BaselineAMm int32
	// Current baseline in ECEF y or NED east component.
	BaselineBMm int32
	// Current baseline in ECEF z or NED down component.
	BaselineCMm int32
	// Current estimate of baseline accuracy.
	Accuracy uint32
	// Current number of integer ambiguity hypotheses.
	IarNumHypotheses int32
}

// GetID implements the msg.Message interface.
func (*MessageGpsRtk) GetID() uint32 {
	return 127
}
