//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package avssuas

import (
	"errors"
)

//
// Actions for reading and writing plan information (mission, rally points, geofence) between persistent and volatile storage when using MAV_CMD_PREFLIGHT_STORAGE.
// (Commonly missions are loaded from persistent storage (flash/EEPROM) into volatile storage (RAM) on startup and written back when they are changed.)
type PREFLIGHT_STORAGE_MISSION_ACTION uint32

const (
	// Read current mission data from persistent storage
	MISSION_READ_PERSISTENT PREFLIGHT_STORAGE_MISSION_ACTION = 0
	// Write current mission data to persistent storage
	MISSION_WRITE_PERSISTENT PREFLIGHT_STORAGE_MISSION_ACTION = 1
	// Erase all mission data stored on the vehicle (both persistent and volatile storage)
	MISSION_RESET_DEFAULT PREFLIGHT_STORAGE_MISSION_ACTION = 2
)

var labels_PREFLIGHT_STORAGE_MISSION_ACTION = map[PREFLIGHT_STORAGE_MISSION_ACTION]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e PREFLIGHT_STORAGE_MISSION_ACTION) MarshalText() ([]byte, error) {
	if l, ok := labels_PREFLIGHT_STORAGE_MISSION_ACTION[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_PREFLIGHT_STORAGE_MISSION_ACTION = map[string]PREFLIGHT_STORAGE_MISSION_ACTION{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *PREFLIGHT_STORAGE_MISSION_ACTION) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_PREFLIGHT_STORAGE_MISSION_ACTION[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e PREFLIGHT_STORAGE_MISSION_ACTION) String() string {
	if l, ok := labels_PREFLIGHT_STORAGE_MISSION_ACTION[e]; ok {
		return l
	}
	return "invalid value"
}
