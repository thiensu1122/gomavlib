//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package development

import (
	"errors"
)

type MAV_ODID_AUTH_TYPE uint32

const (
	// No authentication type is specified.
	MAV_ODID_AUTH_TYPE_NONE MAV_ODID_AUTH_TYPE = 0
	// Signature for the UAS (Unmanned Aircraft System) ID.
	MAV_ODID_AUTH_TYPE_UAS_ID_SIGNATURE MAV_ODID_AUTH_TYPE = 1
	// Signature for the Operator ID.
	MAV_ODID_AUTH_TYPE_OPERATOR_ID_SIGNATURE MAV_ODID_AUTH_TYPE = 2
	// Signature for the entire message set.
	MAV_ODID_AUTH_TYPE_MESSAGE_SET_SIGNATURE MAV_ODID_AUTH_TYPE = 3
	// Authentication is provided by Network Remote ID.
	MAV_ODID_AUTH_TYPE_NETWORK_REMOTE_ID MAV_ODID_AUTH_TYPE = 4
	// The exact authentication type is indicated by the first byte of authentication_data and these type values are managed by ICAO.
	MAV_ODID_AUTH_TYPE_SPECIFIC_AUTHENTICATION MAV_ODID_AUTH_TYPE = 5
)

var labels_MAV_ODID_AUTH_TYPE = map[MAV_ODID_AUTH_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ODID_AUTH_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_ODID_AUTH_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_ODID_AUTH_TYPE = map[string]MAV_ODID_AUTH_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ODID_AUTH_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_ODID_AUTH_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_ODID_AUTH_TYPE) String() string {
	if l, ok := labels_MAV_ODID_AUTH_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
