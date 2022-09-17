//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package all

import (
	"errors"
)

type GOPRO_REQUEST_STATUS uint32

const (
	// The write message with ID indicated succeeded.
	GOPRO_REQUEST_SUCCESS GOPRO_REQUEST_STATUS = 0
	// The write message with ID indicated failed.
	GOPRO_REQUEST_FAILED GOPRO_REQUEST_STATUS = 1
)

var labels_GOPRO_REQUEST_STATUS = map[GOPRO_REQUEST_STATUS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GOPRO_REQUEST_STATUS) MarshalText() ([]byte, error) {
	if l, ok := labels_GOPRO_REQUEST_STATUS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_GOPRO_REQUEST_STATUS = map[string]GOPRO_REQUEST_STATUS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GOPRO_REQUEST_STATUS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_GOPRO_REQUEST_STATUS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e GOPRO_REQUEST_STATUS) String() string {
	if l, ok := labels_GOPRO_REQUEST_STATUS[e]; ok {
		return l
	}
	return "invalid value"
}
