//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package development

import (
	"errors"
)

// Indicates the severity level, generally used for status messages to indicate their relative urgency. Based on RFC-5424 using expanded definitions at: http://www.kiwisyslog.com/kb/info:-syslog-message-levels/.
type MAV_SEVERITY int

const (
	// System is unusable. This is a "panic" condition.
	MAV_SEVERITY_EMERGENCY MAV_SEVERITY = 0
	// Action should be taken immediately. Indicates error in non-critical systems.
	MAV_SEVERITY_ALERT MAV_SEVERITY = 1
	// Action must be taken immediately. Indicates failure in a primary system.
	MAV_SEVERITY_CRITICAL MAV_SEVERITY = 2
	// Indicates an error in secondary/redundant systems.
	MAV_SEVERITY_ERROR MAV_SEVERITY = 3
	// Indicates about a possible future error if this is not resolved within a given timeframe. Example would be a low battery warning.
	MAV_SEVERITY_WARNING MAV_SEVERITY = 4
	// An unusual event has occurred, though not an error condition. This should be investigated for the root cause.
	MAV_SEVERITY_NOTICE MAV_SEVERITY = 5
	// Normal operational messages. Useful for logging. No action is required for these messages.
	MAV_SEVERITY_INFO MAV_SEVERITY = 6
	// Useful non-operational messages that can assist in debugging. These should not occur during normal operation.
	MAV_SEVERITY_DEBUG MAV_SEVERITY = 7
)

var labels_MAV_SEVERITY = map[MAV_SEVERITY]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_SEVERITY) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_SEVERITY[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_SEVERITY = map[string]MAV_SEVERITY{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_SEVERITY) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_SEVERITY[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_SEVERITY) String() string {
	if l, ok := labels_MAV_SEVERITY[e]; ok {
		return l
	}
	return "invalid value"
}
