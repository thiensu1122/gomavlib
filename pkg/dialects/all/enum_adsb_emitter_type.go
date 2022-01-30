//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package all

import (
	"errors"
)

// ADSB classification for the type of vehicle emitting the transponder signal
type ADSB_EMITTER_TYPE int

const (
	ADSB_EMITTER_TYPE_NO_INFO           ADSB_EMITTER_TYPE = 0
	ADSB_EMITTER_TYPE_LIGHT             ADSB_EMITTER_TYPE = 1
	ADSB_EMITTER_TYPE_SMALL             ADSB_EMITTER_TYPE = 2
	ADSB_EMITTER_TYPE_LARGE             ADSB_EMITTER_TYPE = 3
	ADSB_EMITTER_TYPE_HIGH_VORTEX_LARGE ADSB_EMITTER_TYPE = 4
	ADSB_EMITTER_TYPE_HEAVY             ADSB_EMITTER_TYPE = 5
	ADSB_EMITTER_TYPE_HIGHLY_MANUV      ADSB_EMITTER_TYPE = 6
	ADSB_EMITTER_TYPE_ROTOCRAFT         ADSB_EMITTER_TYPE = 7
	ADSB_EMITTER_TYPE_UNASSIGNED        ADSB_EMITTER_TYPE = 8
	ADSB_EMITTER_TYPE_GLIDER            ADSB_EMITTER_TYPE = 9
	ADSB_EMITTER_TYPE_LIGHTER_AIR       ADSB_EMITTER_TYPE = 10
	ADSB_EMITTER_TYPE_PARACHUTE         ADSB_EMITTER_TYPE = 11
	ADSB_EMITTER_TYPE_ULTRA_LIGHT       ADSB_EMITTER_TYPE = 12
	ADSB_EMITTER_TYPE_UNASSIGNED2       ADSB_EMITTER_TYPE = 13
	ADSB_EMITTER_TYPE_UAV               ADSB_EMITTER_TYPE = 14
	ADSB_EMITTER_TYPE_SPACE             ADSB_EMITTER_TYPE = 15
	ADSB_EMITTER_TYPE_UNASSGINED3       ADSB_EMITTER_TYPE = 16
	ADSB_EMITTER_TYPE_EMERGENCY_SURFACE ADSB_EMITTER_TYPE = 17
	ADSB_EMITTER_TYPE_SERVICE_SURFACE   ADSB_EMITTER_TYPE = 18
	ADSB_EMITTER_TYPE_POINT_OBSTACLE    ADSB_EMITTER_TYPE = 19
)

var labels_ADSB_EMITTER_TYPE = map[ADSB_EMITTER_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e ADSB_EMITTER_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_ADSB_EMITTER_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_ADSB_EMITTER_TYPE = map[string]ADSB_EMITTER_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *ADSB_EMITTER_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_ADSB_EMITTER_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e ADSB_EMITTER_TYPE) String() string {
	if l, ok := labels_ADSB_EMITTER_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
