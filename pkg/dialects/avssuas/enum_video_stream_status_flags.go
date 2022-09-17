//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package avssuas

import (
	"errors"
)

// Stream status flags (Bitmap)
type VIDEO_STREAM_STATUS_FLAGS uint32

const (
	// Stream is active (running)
	VIDEO_STREAM_STATUS_FLAGS_RUNNING VIDEO_STREAM_STATUS_FLAGS = 1
	// Stream is thermal imaging
	VIDEO_STREAM_STATUS_FLAGS_THERMAL VIDEO_STREAM_STATUS_FLAGS = 2
)

var labels_VIDEO_STREAM_STATUS_FLAGS = map[VIDEO_STREAM_STATUS_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e VIDEO_STREAM_STATUS_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_VIDEO_STREAM_STATUS_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_VIDEO_STREAM_STATUS_FLAGS = map[string]VIDEO_STREAM_STATUS_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *VIDEO_STREAM_STATUS_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_VIDEO_STREAM_STATUS_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e VIDEO_STREAM_STATUS_FLAGS) String() string {
	if l, ok := labels_VIDEO_STREAM_STATUS_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}
