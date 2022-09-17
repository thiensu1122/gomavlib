//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"errors"
)

type GOPRO_VIDEO_SETTINGS_FLAGS uint32

const (
	// 0=NTSC, 1=PAL.
	GOPRO_VIDEO_SETTINGS_TV_MODE GOPRO_VIDEO_SETTINGS_FLAGS = 1
)

var labels_GOPRO_VIDEO_SETTINGS_FLAGS = map[GOPRO_VIDEO_SETTINGS_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GOPRO_VIDEO_SETTINGS_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_GOPRO_VIDEO_SETTINGS_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_GOPRO_VIDEO_SETTINGS_FLAGS = map[string]GOPRO_VIDEO_SETTINGS_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GOPRO_VIDEO_SETTINGS_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_GOPRO_VIDEO_SETTINGS_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e GOPRO_VIDEO_SETTINGS_FLAGS) String() string {
	if l, ok := labels_GOPRO_VIDEO_SETTINGS_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}
