//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"errors"
)

// Navigational status of AIS vessel, enum duplicated from AIS standard, https://gpsd.gitlab.io/gpsd/AIVDM.html
type AIS_NAV_STATUS uint32

const (
	// Under way using engine.
	UNDER_WAY                           AIS_NAV_STATUS = 0
	AIS_NAV_ANCHORED                    AIS_NAV_STATUS = 1
	AIS_NAV_UN_COMMANDED                AIS_NAV_STATUS = 2
	AIS_NAV_RESTRICTED_MANOEUVERABILITY AIS_NAV_STATUS = 3
	AIS_NAV_DRAUGHT_CONSTRAINED         AIS_NAV_STATUS = 4
	AIS_NAV_MOORED                      AIS_NAV_STATUS = 5
	AIS_NAV_AGROUND                     AIS_NAV_STATUS = 6
	AIS_NAV_FISHING                     AIS_NAV_STATUS = 7
	AIS_NAV_SAILING                     AIS_NAV_STATUS = 8
	AIS_NAV_RESERVED_HSC                AIS_NAV_STATUS = 9
	AIS_NAV_RESERVED_WIG                AIS_NAV_STATUS = 10
	AIS_NAV_RESERVED_1                  AIS_NAV_STATUS = 11
	AIS_NAV_RESERVED_2                  AIS_NAV_STATUS = 12
	AIS_NAV_RESERVED_3                  AIS_NAV_STATUS = 13
	// Search And Rescue Transponder.
	AIS_NAV_AIS_SART AIS_NAV_STATUS = 14
	// Not available (default).
	AIS_NAV_UNKNOWN AIS_NAV_STATUS = 15
)

var labels_AIS_NAV_STATUS = map[AIS_NAV_STATUS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e AIS_NAV_STATUS) MarshalText() ([]byte, error) {
	if l, ok := labels_AIS_NAV_STATUS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_AIS_NAV_STATUS = map[string]AIS_NAV_STATUS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *AIS_NAV_STATUS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_AIS_NAV_STATUS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e AIS_NAV_STATUS) String() string {
	if l, ok := labels_AIS_NAV_STATUS[e]; ok {
		return l
	}
	return "invalid value"
}
