//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Message appropriate for high latency connections like Iridium
type MessageHighLatency struct {
	// Bitmap of enabled system modes.
	BaseMode MAV_MODE_FLAG `mavenum:"uint8"`
	// A bitfield for use for autopilot-specific flags.
	CustomMode uint32
	// The landed state. Is set to MAV_LANDED_STATE_UNDEFINED if landed state is unknown.
	LandedState MAV_LANDED_STATE `mavenum:"uint8"`
	// roll
	Roll int16
	// pitch
	Pitch int16
	// heading
	Heading uint16
	// throttle (percentage)
	Throttle int8
	// heading setpoint
	HeadingSp int16
	// Latitude
	Latitude int32
	// Longitude
	Longitude int32
	// Altitude above mean sea level
	AltitudeAmsl int16
	// Altitude setpoint relative to the home position
	AltitudeSp int16
	// airspeed
	Airspeed uint8
	// airspeed setpoint
	AirspeedSp uint8
	// groundspeed
	Groundspeed uint8
	// climb rate
	ClimbRate int8
	// Number of satellites visible. If unknown, set to UINT8_MAX
	GpsNsat uint8
	// GPS Fix type.
	GpsFixType GPS_FIX_TYPE `mavenum:"uint8"`
	// Remaining battery (percentage)
	BatteryRemaining uint8
	// Autopilot temperature (degrees C)
	Temperature int8
	// Air temperature (degrees C) from airspeed sensor
	TemperatureAir int8
	// failsafe (each bit represents a failsafe where 0=ok, 1=failsafe active (bit0:RC, bit1:batt, bit2:GPS, bit3:GCS, bit4:fence)
	Failsafe uint8
	// current waypoint number
	WpNum uint8
	// distance to target
	WpDistance uint16
}

// GetID implements the msg.Message interface.
func (*MessageHighLatency) GetID() uint32 {
	return 234
}
