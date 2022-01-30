//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package all

import (
	"errors"
)

type HEADING_TYPE int

const (
	HEADING_TYPE_COURSE_OVER_GROUND HEADING_TYPE = 0
	HEADING_TYPE_HEADING            HEADING_TYPE = 1
)

var labels_HEADING_TYPE = map[HEADING_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e HEADING_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_HEADING_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_HEADING_TYPE = map[string]HEADING_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *HEADING_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_HEADING_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e HEADING_TYPE) String() string {
	if l, ok := labels_HEADING_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
