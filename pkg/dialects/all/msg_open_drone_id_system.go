//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

// Data for filling the OpenDroneID System message. The System Message contains general system information including the operator location and possible aircraft group information.
type MessageOpenDroneIdSystem struct {
	// System ID (0 for broadcast).
	TargetSystem uint8
	// Component ID (0 for broadcast).
	TargetComponent uint8
	// Only used for drone ID data received from other UAs. See detailed description at https://mavlink.io/en/services/opendroneid.html.
	IdOrMac [20]uint8
	// Specifies the operator location type.
	OperatorLocationType MAV_ODID_OPERATOR_LOCATION_TYPE `mavenum:"uint8"`
	// Specifies the classification type of the UA.
	ClassificationType MAV_ODID_CLASSIFICATION_TYPE `mavenum:"uint8"`
	// Latitude of the operator. If unknown: 0 (both Lat/Lon).
	OperatorLatitude int32
	// Longitude of the operator. If unknown: 0 (both Lat/Lon).
	OperatorLongitude int32
	// Number of aircraft in the area, group or formation (default 1).
	AreaCount uint16
	// Radius of the cylindrical area of the group or formation (default 0).
	AreaRadius uint16
	// Area Operations Ceiling relative to WGS84. If unknown: -1000 m.
	AreaCeiling float32
	// Area Operations Floor relative to WGS84. If unknown: -1000 m.
	AreaFloor float32
	// When classification_type is MAV_ODID_CLASSIFICATION_TYPE_EU, specifies the category of the UA.
	CategoryEu MAV_ODID_CATEGORY_EU `mavenum:"uint8"`
	// When classification_type is MAV_ODID_CLASSIFICATION_TYPE_EU, specifies the class of the UA.
	ClassEu MAV_ODID_CLASS_EU `mavenum:"uint8"`
	// Geodetic altitude of the operator relative to WGS84. If unknown: -1000 m.
	OperatorAltitudeGeo float32
}

// GetID implements the msg.Message interface.
func (*MessageOpenDroneIdSystem) GetID() uint32 {
	return 12904
}
