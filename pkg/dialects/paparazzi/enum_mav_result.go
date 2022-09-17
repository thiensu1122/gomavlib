//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package paparazzi

import (
	"errors"
)

// Result from a MAVLink command (MAV_CMD)
type MAV_RESULT uint32

const (
	// Command is valid (is supported and has valid parameters), and was executed.
	MAV_RESULT_ACCEPTED MAV_RESULT = 0
	// Command is valid, but cannot be executed at this time. This is used to indicate a problem that should be fixed just by waiting (e.g. a state machine is busy, can't arm because have not got GPS lock, etc.). Retrying later should work.
	MAV_RESULT_TEMPORARILY_REJECTED MAV_RESULT = 1
	// Command is invalid (is supported but has invalid parameters). Retrying same command and parameters will not work.
	MAV_RESULT_DENIED MAV_RESULT = 2
	// Command is not supported (unknown).
	MAV_RESULT_UNSUPPORTED MAV_RESULT = 3
	// Command is valid, but execution has failed. This is used to indicate any non-temporary or unexpected problem, i.e. any problem that must be fixed before the command can succeed/be retried. For example, attempting to write a file when out of memory, attempting to arm when sensors are not calibrated, etc.
	MAV_RESULT_FAILED MAV_RESULT = 4
	// Command is valid and is being executed. This will be followed by further progress updates, i.e. the component may send further COMMAND_ACK messages with result MAV_RESULT_IN_PROGRESS (at a rate decided by the implementation), and must terminate by sending a COMMAND_ACK message with final result of the operation. The COMMAND_ACK.progress field can be used to indicate the progress of the operation.
	MAV_RESULT_IN_PROGRESS MAV_RESULT = 5
	// Command has been cancelled (as a result of receiving a COMMAND_CANCEL message).
	MAV_RESULT_CANCELLED MAV_RESULT = 6
)

var labels_MAV_RESULT = map[MAV_RESULT]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_RESULT) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_RESULT[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_RESULT = map[string]MAV_RESULT{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_RESULT) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_RESULT[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_RESULT) String() string {
	if l, ok := labels_MAV_RESULT[e]; ok {
		return l
	}
	return "invalid value"
}
