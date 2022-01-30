//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package all

import (
	"errors"
)

type GOPRO_PHOTO_RESOLUTION int

const (
	// 5MP Medium.
	GOPRO_PHOTO_RESOLUTION_5MP_MEDIUM GOPRO_PHOTO_RESOLUTION = 0
	// 7MP Medium.
	GOPRO_PHOTO_RESOLUTION_7MP_MEDIUM GOPRO_PHOTO_RESOLUTION = 1
	// 7MP Wide.
	GOPRO_PHOTO_RESOLUTION_7MP_WIDE GOPRO_PHOTO_RESOLUTION = 2
	// 10MP Wide.
	GOPRO_PHOTO_RESOLUTION_10MP_WIDE GOPRO_PHOTO_RESOLUTION = 3
	// 12MP Wide.
	GOPRO_PHOTO_RESOLUTION_12MP_WIDE GOPRO_PHOTO_RESOLUTION = 4
)

var labels_GOPRO_PHOTO_RESOLUTION = map[GOPRO_PHOTO_RESOLUTION]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GOPRO_PHOTO_RESOLUTION) MarshalText() ([]byte, error) {
	if l, ok := labels_GOPRO_PHOTO_RESOLUTION[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_GOPRO_PHOTO_RESOLUTION = map[string]GOPRO_PHOTO_RESOLUTION{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GOPRO_PHOTO_RESOLUTION) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_GOPRO_PHOTO_RESOLUTION[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e GOPRO_PHOTO_RESOLUTION) String() string {
	if l, ok := labels_GOPRO_PHOTO_RESOLUTION[e]; ok {
		return l
	}
	return "invalid value"
}
