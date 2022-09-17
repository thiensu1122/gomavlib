//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package development

import (
	"errors"
)

// Camera tracking target data (shows where tracked target is within image)
type CAMERA_TRACKING_TARGET_DATA uint32

const (
	// No target data
	CAMERA_TRACKING_TARGET_DATA_NONE CAMERA_TRACKING_TARGET_DATA = 0
	// Target data embedded in image data (proprietary)
	CAMERA_TRACKING_TARGET_DATA_EMBEDDED CAMERA_TRACKING_TARGET_DATA = 1
	// Target data rendered in image
	CAMERA_TRACKING_TARGET_DATA_RENDERED CAMERA_TRACKING_TARGET_DATA = 2
	// Target data within status message (Point or Rectangle)
	CAMERA_TRACKING_TARGET_DATA_IN_STATUS CAMERA_TRACKING_TARGET_DATA = 4
)

var labels_CAMERA_TRACKING_TARGET_DATA = map[CAMERA_TRACKING_TARGET_DATA]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e CAMERA_TRACKING_TARGET_DATA) MarshalText() ([]byte, error) {
	if l, ok := labels_CAMERA_TRACKING_TARGET_DATA[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_CAMERA_TRACKING_TARGET_DATA = map[string]CAMERA_TRACKING_TARGET_DATA{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *CAMERA_TRACKING_TARGET_DATA) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_CAMERA_TRACKING_TARGET_DATA[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e CAMERA_TRACKING_TARGET_DATA) String() string {
	if l, ok := labels_CAMERA_TRACKING_TARGET_DATA[e]; ok {
		return l
	}
	return "invalid value"
}
