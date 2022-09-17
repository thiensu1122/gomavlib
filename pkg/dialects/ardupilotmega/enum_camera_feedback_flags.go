//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"errors"
)

type CAMERA_FEEDBACK_FLAGS uint32

const (
	// Shooting photos, not video.
	CAMERA_FEEDBACK_PHOTO CAMERA_FEEDBACK_FLAGS = 0
	// Shooting video, not stills.
	CAMERA_FEEDBACK_VIDEO CAMERA_FEEDBACK_FLAGS = 1
	// Unable to achieve requested exposure (e.g. shutter speed too low).
	CAMERA_FEEDBACK_BADEXPOSURE CAMERA_FEEDBACK_FLAGS = 2
	// Closed loop feedback from camera, we know for sure it has successfully taken a picture.
	CAMERA_FEEDBACK_CLOSEDLOOP CAMERA_FEEDBACK_FLAGS = 3
	// Open loop camera, an image trigger has been requested but we can't know for sure it has successfully taken a picture.
	CAMERA_FEEDBACK_OPENLOOP CAMERA_FEEDBACK_FLAGS = 4
)

var labels_CAMERA_FEEDBACK_FLAGS = map[CAMERA_FEEDBACK_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e CAMERA_FEEDBACK_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_CAMERA_FEEDBACK_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_CAMERA_FEEDBACK_FLAGS = map[string]CAMERA_FEEDBACK_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *CAMERA_FEEDBACK_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_CAMERA_FEEDBACK_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e CAMERA_FEEDBACK_FLAGS) String() string {
	if l, ok := labels_CAMERA_FEEDBACK_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}
