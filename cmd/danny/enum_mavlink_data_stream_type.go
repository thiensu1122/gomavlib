//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package danny

import (
    "strings"
    "fmt"
)
type MAVLINK_DATA_STREAM_TYPE uint32

const (
    MAVLINK_DATA_STREAM_IMG_JPEG MAVLINK_DATA_STREAM_TYPE = 0
    MAVLINK_DATA_STREAM_IMG_BMP MAVLINK_DATA_STREAM_TYPE = 1
    MAVLINK_DATA_STREAM_IMG_RAW8U MAVLINK_DATA_STREAM_TYPE = 2
    MAVLINK_DATA_STREAM_IMG_RAW32U MAVLINK_DATA_STREAM_TYPE = 3
    MAVLINK_DATA_STREAM_IMG_PGM MAVLINK_DATA_STREAM_TYPE = 4
    MAVLINK_DATA_STREAM_IMG_PNG MAVLINK_DATA_STREAM_TYPE = 5
)

var labels_MAVLINK_DATA_STREAM_TYPE = map[MAVLINK_DATA_STREAM_TYPE]string{
    MAVLINK_DATA_STREAM_IMG_JPEG: "MAVLINK_DATA_STREAM_IMG_JPEG",
    MAVLINK_DATA_STREAM_IMG_BMP: "MAVLINK_DATA_STREAM_IMG_BMP",
    MAVLINK_DATA_STREAM_IMG_RAW8U: "MAVLINK_DATA_STREAM_IMG_RAW8U",
    MAVLINK_DATA_STREAM_IMG_RAW32U: "MAVLINK_DATA_STREAM_IMG_RAW32U",
    MAVLINK_DATA_STREAM_IMG_PGM: "MAVLINK_DATA_STREAM_IMG_PGM",
    MAVLINK_DATA_STREAM_IMG_PNG: "MAVLINK_DATA_STREAM_IMG_PNG",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAVLINK_DATA_STREAM_TYPE) MarshalText() ([]byte, error) {
    var names []string
    for mask, label := range labels_MAVLINK_DATA_STREAM_TYPE {
        if e&mask == mask {
            names = append(names, label)
        }
    }
    return []byte(strings.Join(names, " | ")), nil
}


// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAVLINK_DATA_STREAM_TYPE) UnmarshalText(text []byte) error {
    labels := strings.Split(string(text), " | ")
    var mask MAVLINK_DATA_STREAM_TYPE
    for _, label := range labels {
        found := false
        for value, l := range labels_MAVLINK_DATA_STREAM_TYPE {
            if l == label {
                mask |= value
                found = true
                break
            }
        }
        if !found {
            return fmt.Errorf("invalid label '%s'", label)
        }
    }
    *e = mask
    return nil
}

// String implements the fmt.Stringer interface.
func (e MAVLINK_DATA_STREAM_TYPE) String() string {
    val, _ := e.MarshalText()
    return string(val)
}
