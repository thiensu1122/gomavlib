//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package ualberta

import (
	"errors"
)

//
// Actions for reading/writing parameters between persistent and volatile storage when using MAV_CMD_PREFLIGHT_STORAGE.
// (Commonly parameters are loaded from persistent storage (flash/EEPROM) into volatile storage (RAM) on startup and written back when they are changed.)
type PREFLIGHT_STORAGE_PARAMETER_ACTION uint32

const (
	// Read all parameters from persistent storage. Replaces values in volatile storage.
	PARAM_READ_PERSISTENT PREFLIGHT_STORAGE_PARAMETER_ACTION = 0
	// Write all parameter values to persistent storage (flash/EEPROM)
	PARAM_WRITE_PERSISTENT PREFLIGHT_STORAGE_PARAMETER_ACTION = 1
	// Reset all user configurable parameters to their default value (including airframe selection, sensor calibration data, safety settings, and so on). Does not reset values that contain operation counters and vehicle computed statistics.
	PARAM_RESET_CONFIG_DEFAULT PREFLIGHT_STORAGE_PARAMETER_ACTION = 2
	// Reset only sensor calibration parameters to factory defaults (or firmware default if not available)
	PARAM_RESET_SENSOR_DEFAULT PREFLIGHT_STORAGE_PARAMETER_ACTION = 3
	// Reset all parameters, including operation counters, to default values
	PARAM_RESET_ALL_DEFAULT PREFLIGHT_STORAGE_PARAMETER_ACTION = 4
)

var labels_PREFLIGHT_STORAGE_PARAMETER_ACTION = map[PREFLIGHT_STORAGE_PARAMETER_ACTION]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e PREFLIGHT_STORAGE_PARAMETER_ACTION) MarshalText() ([]byte, error) {
	if l, ok := labels_PREFLIGHT_STORAGE_PARAMETER_ACTION[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_PREFLIGHT_STORAGE_PARAMETER_ACTION = map[string]PREFLIGHT_STORAGE_PARAMETER_ACTION{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *PREFLIGHT_STORAGE_PARAMETER_ACTION) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_PREFLIGHT_STORAGE_PARAMETER_ACTION[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e PREFLIGHT_STORAGE_PARAMETER_ACTION) String() string {
	if l, ok := labels_PREFLIGHT_STORAGE_PARAMETER_ACTION[e]; ok {
		return l
	}
	return "invalid value"
}
