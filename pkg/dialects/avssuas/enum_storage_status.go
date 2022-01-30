//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package avssuas

import (
	"errors"
)

// Flags to indicate the status of camera storage.
type STORAGE_STATUS int

const (
	// Storage is missing (no microSD card loaded for example.)
	STORAGE_STATUS_EMPTY STORAGE_STATUS = 0
	// Storage present but unformatted.
	STORAGE_STATUS_UNFORMATTED STORAGE_STATUS = 1
	// Storage present and ready.
	STORAGE_STATUS_READY STORAGE_STATUS = 2
	// Camera does not supply storage status information. Capacity information in STORAGE_INFORMATION fields will be ignored.
	STORAGE_STATUS_NOT_SUPPORTED STORAGE_STATUS = 3
)

var labels_STORAGE_STATUS = map[STORAGE_STATUS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e STORAGE_STATUS) MarshalText() ([]byte, error) {
	if l, ok := labels_STORAGE_STATUS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_STORAGE_STATUS = map[string]STORAGE_STATUS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *STORAGE_STATUS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_STORAGE_STATUS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e STORAGE_STATUS) String() string {
	if l, ok := labels_STORAGE_STATUS[e]; ok {
		return l
	}
	return "invalid value"
}
