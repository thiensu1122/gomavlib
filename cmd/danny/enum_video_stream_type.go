//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package danny

import (
    "strings"
    "fmt"
)
// Video stream types
type VIDEO_STREAM_TYPE uint32

const (
    // Stream is RTSP
    VIDEO_STREAM_TYPE_RTSP VIDEO_STREAM_TYPE = 0
    // Stream is RTP UDP (URI gives the port number)
    VIDEO_STREAM_TYPE_RTPUDP VIDEO_STREAM_TYPE = 1
    // Stream is MPEG on TCP
    VIDEO_STREAM_TYPE_TCP_MPEG VIDEO_STREAM_TYPE = 2
    // Stream is h.264 on MPEG TS (URI gives the port number)
    VIDEO_STREAM_TYPE_MPEG_TS_H264 VIDEO_STREAM_TYPE = 3
)

var labels_VIDEO_STREAM_TYPE = map[VIDEO_STREAM_TYPE]string{
    VIDEO_STREAM_TYPE_RTSP: "VIDEO_STREAM_TYPE_RTSP",
    VIDEO_STREAM_TYPE_RTPUDP: "VIDEO_STREAM_TYPE_RTPUDP",
    VIDEO_STREAM_TYPE_TCP_MPEG: "VIDEO_STREAM_TYPE_TCP_MPEG",
    VIDEO_STREAM_TYPE_MPEG_TS_H264: "VIDEO_STREAM_TYPE_MPEG_TS_H264",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e VIDEO_STREAM_TYPE) MarshalText() ([]byte, error) {
    var names []string
    for mask, label := range labels_VIDEO_STREAM_TYPE {
        if e&mask == mask {
            names = append(names, label)
        }
    }
    return []byte(strings.Join(names, " | ")), nil
}


// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *VIDEO_STREAM_TYPE) UnmarshalText(text []byte) error {
    labels := strings.Split(string(text), " | ")
    var mask VIDEO_STREAM_TYPE
    for _, label := range labels {
        found := false
        for value, l := range labels_VIDEO_STREAM_TYPE {
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
func (e VIDEO_STREAM_TYPE) String() string {
    val, _ := e.MarshalText()
    return string(val)
}
