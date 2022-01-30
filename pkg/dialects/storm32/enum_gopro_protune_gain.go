//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package storm32

import (
	"errors"
)

type GOPRO_PROTUNE_GAIN int

const (
	// ISO 400.
	GOPRO_PROTUNE_GAIN_400 GOPRO_PROTUNE_GAIN = 0
	// ISO 800 (Only Hero 4).
	GOPRO_PROTUNE_GAIN_800 GOPRO_PROTUNE_GAIN = 1
	// ISO 1600.
	GOPRO_PROTUNE_GAIN_1600 GOPRO_PROTUNE_GAIN = 2
	// ISO 3200 (Only Hero 4).
	GOPRO_PROTUNE_GAIN_3200 GOPRO_PROTUNE_GAIN = 3
	// ISO 6400.
	GOPRO_PROTUNE_GAIN_6400 GOPRO_PROTUNE_GAIN = 4
)

var labels_GOPRO_PROTUNE_GAIN = map[GOPRO_PROTUNE_GAIN]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GOPRO_PROTUNE_GAIN) MarshalText() ([]byte, error) {
	if l, ok := labels_GOPRO_PROTUNE_GAIN[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_GOPRO_PROTUNE_GAIN = map[string]GOPRO_PROTUNE_GAIN{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GOPRO_PROTUNE_GAIN) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_GOPRO_PROTUNE_GAIN[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e GOPRO_PROTUNE_GAIN) String() string {
	if l, ok := labels_GOPRO_PROTUNE_GAIN[e]; ok {
		return l
	}
	return "invalid value"
}
