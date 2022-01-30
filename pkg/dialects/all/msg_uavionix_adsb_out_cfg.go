//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Static data to configure the ADS-B transponder (send within 10 sec of a POR and every 10 sec thereafter)
type MessageUavionixAdsbOutCfg struct {
	// Vehicle address (24 bit)
	Icao uint32 `mavname:"ICAO"`
	// Vehicle identifier (8 characters, null terminated, valid characters are A-Z, 0-9, " " only)
	Callsign string `mavlen:"9"`
	// Transmitting vehicle type. See ADSB_EMITTER_TYPE enum
	Emittertype ADSB_EMITTER_TYPE `mavenum:"uint8" mavname:"emitterType"`
	// Aircraft length and width encoding (table 2-35 of DO-282B)
	Aircraftsize UAVIONIX_ADSB_OUT_CFG_AIRCRAFT_SIZE `mavenum:"uint8" mavname:"aircraftSize"`
	// GPS antenna lateral offset (table 2-36 of DO-282B)
	Gpsoffsetlat UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT `mavenum:"uint8" mavname:"gpsOffsetLat"`
	// GPS antenna longitudinal offset from nose [if non-zero, take position (in meters) divide by 2 and add one] (table 2-37 DO-282B)
	Gpsoffsetlon UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LON `mavenum:"uint8" mavname:"gpsOffsetLon"`
	// Aircraft stall speed in cm/s
	Stallspeed uint16 `mavname:"stallSpeed"`
	// ADS-B transponder reciever and transmit enable flags
	Rfselect UAVIONIX_ADSB_OUT_RF_SELECT `mavenum:"uint8" mavname:"rfSelect"`
}

// GetID implements the msg.Message interface.
func (*MessageUavionixAdsbOutCfg) GetID() uint32 {
	return 10001
}
