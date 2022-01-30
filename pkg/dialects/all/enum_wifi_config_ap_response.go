//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package all

import (
	"errors"
)

// Possible responses from a WIFI_CONFIG_AP message.
type WIFI_CONFIG_AP_RESPONSE int

const (
	// Undefined response. Likely an indicative of a system that doesn't support this request.
	WIFI_CONFIG_AP_RESPONSE_UNDEFINED WIFI_CONFIG_AP_RESPONSE = 0
	// Changes accepted.
	WIFI_CONFIG_AP_RESPONSE_ACCEPTED WIFI_CONFIG_AP_RESPONSE = 1
	// Changes rejected.
	WIFI_CONFIG_AP_RESPONSE_REJECTED WIFI_CONFIG_AP_RESPONSE = 2
	// Invalid Mode.
	WIFI_CONFIG_AP_RESPONSE_MODE_ERROR WIFI_CONFIG_AP_RESPONSE = 3
	// Invalid SSID.
	WIFI_CONFIG_AP_RESPONSE_SSID_ERROR WIFI_CONFIG_AP_RESPONSE = 4
	// Invalid Password.
	WIFI_CONFIG_AP_RESPONSE_PASSWORD_ERROR WIFI_CONFIG_AP_RESPONSE = 5
)

var labels_WIFI_CONFIG_AP_RESPONSE = map[WIFI_CONFIG_AP_RESPONSE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e WIFI_CONFIG_AP_RESPONSE) MarshalText() ([]byte, error) {
	if l, ok := labels_WIFI_CONFIG_AP_RESPONSE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_WIFI_CONFIG_AP_RESPONSE = map[string]WIFI_CONFIG_AP_RESPONSE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *WIFI_CONFIG_AP_RESPONSE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_WIFI_CONFIG_AP_RESPONSE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e WIFI_CONFIG_AP_RESPONSE) String() string {
	if l, ok := labels_WIFI_CONFIG_AP_RESPONSE[e]; ok {
		return l
	}
	return "invalid value"
}
