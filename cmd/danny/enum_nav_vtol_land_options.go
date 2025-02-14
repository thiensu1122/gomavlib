//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package danny

import (
    "strings"
    "fmt"
)
type NAV_VTOL_LAND_OPTIONS uint32

const (
    // Default autopilot landing behaviour.
    NAV_VTOL_LAND_OPTIONS_DEFAULT NAV_VTOL_LAND_OPTIONS = 0
    // Descend in fixed wing mode, transitioning to multicopter mode for vertical landing when close to the ground.
    // The fixed wing descent pattern is at the discretion of the vehicle (e.g. transition altitude, loiter direction, radius, and speed, etc.).
    NAV_VTOL_LAND_OPTIONS_FW_DESCENT NAV_VTOL_LAND_OPTIONS = 1
    // Land in multicopter mode on reaching the landing coordinates (the whole landing is by "hover descent").
    NAV_VTOL_LAND_OPTIONS_HOVER_DESCENT NAV_VTOL_LAND_OPTIONS = 2
)

var labels_NAV_VTOL_LAND_OPTIONS = map[NAV_VTOL_LAND_OPTIONS]string{
    NAV_VTOL_LAND_OPTIONS_DEFAULT: "NAV_VTOL_LAND_OPTIONS_DEFAULT",
    NAV_VTOL_LAND_OPTIONS_FW_DESCENT: "NAV_VTOL_LAND_OPTIONS_FW_DESCENT",
    NAV_VTOL_LAND_OPTIONS_HOVER_DESCENT: "NAV_VTOL_LAND_OPTIONS_HOVER_DESCENT",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e NAV_VTOL_LAND_OPTIONS) MarshalText() ([]byte, error) {
    var names []string
    for mask, label := range labels_NAV_VTOL_LAND_OPTIONS {
        if e&mask == mask {
            names = append(names, label)
        }
    }
    return []byte(strings.Join(names, " | ")), nil
}


// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *NAV_VTOL_LAND_OPTIONS) UnmarshalText(text []byte) error {
    labels := strings.Split(string(text), " | ")
    var mask NAV_VTOL_LAND_OPTIONS
    for _, label := range labels {
        found := false
        for value, l := range labels_NAV_VTOL_LAND_OPTIONS {
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
func (e NAV_VTOL_LAND_OPTIONS) String() string {
    val, _ := e.MarshalText()
    return string(val)
}
