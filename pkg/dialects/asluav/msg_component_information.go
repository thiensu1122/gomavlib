//autogenerated:yes
//nolint:revive,misspell,govet,lll
package asluav

// Information about a component. For camera components instead use CAMERA_INFORMATION, and for autopilots additionally use AUTOPILOT_VERSION. Components including GCSes should consider supporting requests of this message via MAV_CMD_REQUEST_MESSAGE.
type MessageComponentInformation struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// CRC32 of the TYPE_GENERAL file (can be used by a GCS for file caching). The general metadata file contains URLs to other files of different type according to COMP_METADATA_TYPE.
	GeneralMetadataFileCrc uint32
	// Component definition URI for TYPE_GENERAL. This must be a MAVLink FTP URI and the file might be compressed with xz.
	GeneralMetadataUri string `mavlen:"100"`
	// CRC32 of the TYPE_PERIPHERALS file (can be used by a GCS for file caching).
	PeripheralsMetadataFileCrc uint32
	// (Optional) Component definition URI for TYPE_PERIPHERALS. This must be a MAVLink FTP URI and the file might be compressed with xz. Peripherals are listed as a separate URL and not included in the general metadata file because it will likely be generated at runtime while the general (and other referenced files) might be generated at compile time.
	PeripheralsMetadataUri string `mavlen:"100"`
}

// GetID implements the msg.Message interface.
func (*MessageComponentInformation) GetID() uint32 {
	return 395
}
