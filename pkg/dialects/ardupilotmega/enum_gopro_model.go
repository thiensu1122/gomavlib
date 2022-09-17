//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"errors"
)

type GOPRO_MODEL uint32

const (
	// Unknown gopro model.
	GOPRO_MODEL_UNKNOWN GOPRO_MODEL = 0
	// Hero 3+ Silver (HeroBus not supported by GoPro).
	GOPRO_MODEL_HERO_3_PLUS_SILVER GOPRO_MODEL = 1
	// Hero 3+ Black.
	GOPRO_MODEL_HERO_3_PLUS_BLACK GOPRO_MODEL = 2
	// Hero 4 Silver.
	GOPRO_MODEL_HERO_4_SILVER GOPRO_MODEL = 3
	// Hero 4 Black.
	GOPRO_MODEL_HERO_4_BLACK GOPRO_MODEL = 4
)

var labels_GOPRO_MODEL = map[GOPRO_MODEL]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GOPRO_MODEL) MarshalText() ([]byte, error) {
	if l, ok := labels_GOPRO_MODEL[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_GOPRO_MODEL = map[string]GOPRO_MODEL{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GOPRO_MODEL) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_GOPRO_MODEL[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e GOPRO_MODEL) String() string {
	if l, ok := labels_GOPRO_MODEL[e]; ok {
		return l
	}
	return "invalid value"
}
