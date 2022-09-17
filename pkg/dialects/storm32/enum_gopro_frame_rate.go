//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package storm32

import (
	"errors"
)

type GOPRO_FRAME_RATE uint32

const (
	// 12 FPS.
	GOPRO_FRAME_RATE_12 GOPRO_FRAME_RATE = 0
	// 15 FPS.
	GOPRO_FRAME_RATE_15 GOPRO_FRAME_RATE = 1
	// 24 FPS.
	GOPRO_FRAME_RATE_24 GOPRO_FRAME_RATE = 2
	// 25 FPS.
	GOPRO_FRAME_RATE_25 GOPRO_FRAME_RATE = 3
	// 30 FPS.
	GOPRO_FRAME_RATE_30 GOPRO_FRAME_RATE = 4
	// 48 FPS.
	GOPRO_FRAME_RATE_48 GOPRO_FRAME_RATE = 5
	// 50 FPS.
	GOPRO_FRAME_RATE_50 GOPRO_FRAME_RATE = 6
	// 60 FPS.
	GOPRO_FRAME_RATE_60 GOPRO_FRAME_RATE = 7
	// 80 FPS.
	GOPRO_FRAME_RATE_80 GOPRO_FRAME_RATE = 8
	// 90 FPS.
	GOPRO_FRAME_RATE_90 GOPRO_FRAME_RATE = 9
	// 100 FPS.
	GOPRO_FRAME_RATE_100 GOPRO_FRAME_RATE = 10
	// 120 FPS.
	GOPRO_FRAME_RATE_120 GOPRO_FRAME_RATE = 11
	// 240 FPS.
	GOPRO_FRAME_RATE_240 GOPRO_FRAME_RATE = 12
	// 12.5 FPS.
	GOPRO_FRAME_RATE_12_5 GOPRO_FRAME_RATE = 13
)

var labels_GOPRO_FRAME_RATE = map[GOPRO_FRAME_RATE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GOPRO_FRAME_RATE) MarshalText() ([]byte, error) {
	if l, ok := labels_GOPRO_FRAME_RATE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_GOPRO_FRAME_RATE = map[string]GOPRO_FRAME_RATE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GOPRO_FRAME_RATE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_GOPRO_FRAME_RATE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e GOPRO_FRAME_RATE) String() string {
	if l, ok := labels_GOPRO_FRAME_RATE[e]; ok {
		return l
	}
	return "invalid value"
}
