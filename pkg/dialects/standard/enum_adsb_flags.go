//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package standard

import (
	"errors"
)

// These flags indicate status such as data validity of each data source. Set = data valid
type ADSB_FLAGS uint32

const (
	ADSB_FLAGS_VALID_COORDS            ADSB_FLAGS = 1
	ADSB_FLAGS_VALID_ALTITUDE          ADSB_FLAGS = 2
	ADSB_FLAGS_VALID_HEADING           ADSB_FLAGS = 4
	ADSB_FLAGS_VALID_VELOCITY          ADSB_FLAGS = 8
	ADSB_FLAGS_VALID_CALLSIGN          ADSB_FLAGS = 16
	ADSB_FLAGS_VALID_SQUAWK            ADSB_FLAGS = 32
	ADSB_FLAGS_SIMULATED               ADSB_FLAGS = 64
	ADSB_FLAGS_VERTICAL_VELOCITY_VALID ADSB_FLAGS = 128
	ADSB_FLAGS_BARO_VALID              ADSB_FLAGS = 256
	ADSB_FLAGS_SOURCE_UAT              ADSB_FLAGS = 32768
)

var labels_ADSB_FLAGS = map[ADSB_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e ADSB_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_ADSB_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_ADSB_FLAGS = map[string]ADSB_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *ADSB_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_ADSB_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e ADSB_FLAGS) String() string {
	if l, ok := labels_ADSB_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}
