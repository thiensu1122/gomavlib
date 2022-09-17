//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"errors"
)

// Flags to indicate the type of storage.
type STORAGE_TYPE uint32

const (
	// Storage type is not known.
	STORAGE_TYPE_UNKNOWN STORAGE_TYPE = 0
	// Storage type is USB device.
	STORAGE_TYPE_USB_STICK STORAGE_TYPE = 1
	// Storage type is SD card.
	STORAGE_TYPE_SD STORAGE_TYPE = 2
	// Storage type is microSD card.
	STORAGE_TYPE_MICROSD STORAGE_TYPE = 3
	// Storage type is CFast.
	STORAGE_TYPE_CF STORAGE_TYPE = 4
	// Storage type is CFexpress.
	STORAGE_TYPE_CFE STORAGE_TYPE = 5
	// Storage type is XQD.
	STORAGE_TYPE_XQD STORAGE_TYPE = 6
	// Storage type is HD mass storage type.
	STORAGE_TYPE_HD STORAGE_TYPE = 7
	// Storage type is other, not listed type.
	STORAGE_TYPE_OTHER STORAGE_TYPE = 254
)

var labels_STORAGE_TYPE = map[STORAGE_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e STORAGE_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_STORAGE_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_STORAGE_TYPE = map[string]STORAGE_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *STORAGE_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_STORAGE_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e STORAGE_TYPE) String() string {
	if l, ok := labels_STORAGE_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
