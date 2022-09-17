//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"errors"
)

type GOPRO_PROTUNE_EXPOSURE uint32

const (
	// -5.0 EV (Hero 3+ Only).
	GOPRO_PROTUNE_EXPOSURE_NEG_5_0 GOPRO_PROTUNE_EXPOSURE = 0
	// -4.5 EV (Hero 3+ Only).
	GOPRO_PROTUNE_EXPOSURE_NEG_4_5 GOPRO_PROTUNE_EXPOSURE = 1
	// -4.0 EV (Hero 3+ Only).
	GOPRO_PROTUNE_EXPOSURE_NEG_4_0 GOPRO_PROTUNE_EXPOSURE = 2
	// -3.5 EV (Hero 3+ Only).
	GOPRO_PROTUNE_EXPOSURE_NEG_3_5 GOPRO_PROTUNE_EXPOSURE = 3
	// -3.0 EV (Hero 3+ Only).
	GOPRO_PROTUNE_EXPOSURE_NEG_3_0 GOPRO_PROTUNE_EXPOSURE = 4
	// -2.5 EV (Hero 3+ Only).
	GOPRO_PROTUNE_EXPOSURE_NEG_2_5 GOPRO_PROTUNE_EXPOSURE = 5
	// -2.0 EV.
	GOPRO_PROTUNE_EXPOSURE_NEG_2_0 GOPRO_PROTUNE_EXPOSURE = 6
	// -1.5 EV.
	GOPRO_PROTUNE_EXPOSURE_NEG_1_5 GOPRO_PROTUNE_EXPOSURE = 7
	// -1.0 EV.
	GOPRO_PROTUNE_EXPOSURE_NEG_1_0 GOPRO_PROTUNE_EXPOSURE = 8
	// -0.5 EV.
	GOPRO_PROTUNE_EXPOSURE_NEG_0_5 GOPRO_PROTUNE_EXPOSURE = 9
	// 0.0 EV.
	GOPRO_PROTUNE_EXPOSURE_ZERO GOPRO_PROTUNE_EXPOSURE = 10
	// +0.5 EV.
	GOPRO_PROTUNE_EXPOSURE_POS_0_5 GOPRO_PROTUNE_EXPOSURE = 11
	// +1.0 EV.
	GOPRO_PROTUNE_EXPOSURE_POS_1_0 GOPRO_PROTUNE_EXPOSURE = 12
	// +1.5 EV.
	GOPRO_PROTUNE_EXPOSURE_POS_1_5 GOPRO_PROTUNE_EXPOSURE = 13
	// +2.0 EV.
	GOPRO_PROTUNE_EXPOSURE_POS_2_0 GOPRO_PROTUNE_EXPOSURE = 14
	// +2.5 EV (Hero 3+ Only).
	GOPRO_PROTUNE_EXPOSURE_POS_2_5 GOPRO_PROTUNE_EXPOSURE = 15
	// +3.0 EV (Hero 3+ Only).
	GOPRO_PROTUNE_EXPOSURE_POS_3_0 GOPRO_PROTUNE_EXPOSURE = 16
	// +3.5 EV (Hero 3+ Only).
	GOPRO_PROTUNE_EXPOSURE_POS_3_5 GOPRO_PROTUNE_EXPOSURE = 17
	// +4.0 EV (Hero 3+ Only).
	GOPRO_PROTUNE_EXPOSURE_POS_4_0 GOPRO_PROTUNE_EXPOSURE = 18
	// +4.5 EV (Hero 3+ Only).
	GOPRO_PROTUNE_EXPOSURE_POS_4_5 GOPRO_PROTUNE_EXPOSURE = 19
	// +5.0 EV (Hero 3+ Only).
	GOPRO_PROTUNE_EXPOSURE_POS_5_0 GOPRO_PROTUNE_EXPOSURE = 20
)

var labels_GOPRO_PROTUNE_EXPOSURE = map[GOPRO_PROTUNE_EXPOSURE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GOPRO_PROTUNE_EXPOSURE) MarshalText() ([]byte, error) {
	if l, ok := labels_GOPRO_PROTUNE_EXPOSURE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_GOPRO_PROTUNE_EXPOSURE = map[string]GOPRO_PROTUNE_EXPOSURE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GOPRO_PROTUNE_EXPOSURE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_GOPRO_PROTUNE_EXPOSURE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e GOPRO_PROTUNE_EXPOSURE) String() string {
	if l, ok := labels_GOPRO_PROTUNE_EXPOSURE[e]; ok {
		return l
	}
	return "invalid value"
}
