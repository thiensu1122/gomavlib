//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package uavionix

import (
	"errors"
)

// RC type
type RC_TYPE uint32

const (
	// Spektrum DSM2
	RC_TYPE_SPEKTRUM_DSM2 RC_TYPE = 0
	// Spektrum DSMX
	RC_TYPE_SPEKTRUM_DSMX RC_TYPE = 1
)

var labels_RC_TYPE = map[RC_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e RC_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_RC_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_RC_TYPE = map[string]RC_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *RC_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_RC_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e RC_TYPE) String() string {
	if l, ok := labels_RC_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
