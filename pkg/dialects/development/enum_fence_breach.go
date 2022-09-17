//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package development

import (
	"errors"
)

type FENCE_BREACH uint32

const (
	// No last fence breach
	FENCE_BREACH_NONE FENCE_BREACH = 0
	// Breached minimum altitude
	FENCE_BREACH_MINALT FENCE_BREACH = 1
	// Breached maximum altitude
	FENCE_BREACH_MAXALT FENCE_BREACH = 2
	// Breached fence boundary
	FENCE_BREACH_BOUNDARY FENCE_BREACH = 3
)

var labels_FENCE_BREACH = map[FENCE_BREACH]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e FENCE_BREACH) MarshalText() ([]byte, error) {
	if l, ok := labels_FENCE_BREACH[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_FENCE_BREACH = map[string]FENCE_BREACH{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *FENCE_BREACH) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_FENCE_BREACH[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e FENCE_BREACH) String() string {
	if l, ok := labels_FENCE_BREACH[e]; ok {
		return l
	}
	return "invalid value"
}
