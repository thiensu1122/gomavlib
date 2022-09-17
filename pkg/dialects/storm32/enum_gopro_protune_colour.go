//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package storm32

import (
	"errors"
)

type GOPRO_PROTUNE_COLOUR uint32

const (
	// Auto.
	GOPRO_PROTUNE_COLOUR_STANDARD GOPRO_PROTUNE_COLOUR = 0
	// Neutral.
	GOPRO_PROTUNE_COLOUR_NEUTRAL GOPRO_PROTUNE_COLOUR = 1
)

var labels_GOPRO_PROTUNE_COLOUR = map[GOPRO_PROTUNE_COLOUR]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GOPRO_PROTUNE_COLOUR) MarshalText() ([]byte, error) {
	if l, ok := labels_GOPRO_PROTUNE_COLOUR[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_GOPRO_PROTUNE_COLOUR = map[string]GOPRO_PROTUNE_COLOUR{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GOPRO_PROTUNE_COLOUR) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_GOPRO_PROTUNE_COLOUR[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e GOPRO_PROTUNE_COLOUR) String() string {
	if l, ok := labels_GOPRO_PROTUNE_COLOUR[e]; ok {
		return l
	}
	return "invalid value"
}
