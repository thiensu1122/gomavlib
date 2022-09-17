//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package storm32

import (
	"errors"
)

// Possible remote log data block statuses.
type MAV_REMOTE_LOG_DATA_BLOCK_STATUSES uint32

const (
	// This block has NOT been received.
	MAV_REMOTE_LOG_DATA_BLOCK_NACK MAV_REMOTE_LOG_DATA_BLOCK_STATUSES = 0
	// This block has been received.
	MAV_REMOTE_LOG_DATA_BLOCK_ACK MAV_REMOTE_LOG_DATA_BLOCK_STATUSES = 1
)

var labels_MAV_REMOTE_LOG_DATA_BLOCK_STATUSES = map[MAV_REMOTE_LOG_DATA_BLOCK_STATUSES]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_REMOTE_LOG_DATA_BLOCK_STATUSES) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_REMOTE_LOG_DATA_BLOCK_STATUSES[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_REMOTE_LOG_DATA_BLOCK_STATUSES = map[string]MAV_REMOTE_LOG_DATA_BLOCK_STATUSES{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_REMOTE_LOG_DATA_BLOCK_STATUSES) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_REMOTE_LOG_DATA_BLOCK_STATUSES[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_REMOTE_LOG_DATA_BLOCK_STATUSES) String() string {
	if l, ok := labels_MAV_REMOTE_LOG_DATA_BLOCK_STATUSES[e]; ok {
		return l
	}
	return "invalid value"
}
