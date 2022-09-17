//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"errors"
)

type MAV_ODID_CATEGORY_EU uint32

const (
	// The category for the UA, according to the EU specification, is undeclared.
	MAV_ODID_CATEGORY_EU_UNDECLARED MAV_ODID_CATEGORY_EU = 0
	// The category for the UA, according to the EU specification, is the Open category.
	MAV_ODID_CATEGORY_EU_OPEN MAV_ODID_CATEGORY_EU = 1
	// The category for the UA, according to the EU specification, is the Specific category.
	MAV_ODID_CATEGORY_EU_SPECIFIC MAV_ODID_CATEGORY_EU = 2
	// The category for the UA, according to the EU specification, is the Certified category.
	MAV_ODID_CATEGORY_EU_CERTIFIED MAV_ODID_CATEGORY_EU = 3
)

var labels_MAV_ODID_CATEGORY_EU = map[MAV_ODID_CATEGORY_EU]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ODID_CATEGORY_EU) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_ODID_CATEGORY_EU[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_ODID_CATEGORY_EU = map[string]MAV_ODID_CATEGORY_EU{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ODID_CATEGORY_EU) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_ODID_CATEGORY_EU[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_ODID_CATEGORY_EU) String() string {
	if l, ok := labels_MAV_ODID_CATEGORY_EU[e]; ok {
		return l
	}
	return "invalid value"
}
