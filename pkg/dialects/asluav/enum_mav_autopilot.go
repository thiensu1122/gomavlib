//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package asluav

import (
	"errors"
)

// Micro air vehicle / autopilot classes. This identifies the individual model.
type MAV_AUTOPILOT uint32

const (
	// Generic autopilot, full support for everything
	MAV_AUTOPILOT_GENERIC MAV_AUTOPILOT = 0
	// Reserved for future use.
	MAV_AUTOPILOT_RESERVED MAV_AUTOPILOT = 1
	// SLUGS autopilot, http://slugsuav.soe.ucsc.edu
	MAV_AUTOPILOT_SLUGS MAV_AUTOPILOT = 2
	// ArduPilot - Plane/Copter/Rover/Sub/Tracker, https://ardupilot.org
	MAV_AUTOPILOT_ARDUPILOTMEGA MAV_AUTOPILOT = 3
	// OpenPilot, http://openpilot.org
	MAV_AUTOPILOT_OPENPILOT MAV_AUTOPILOT = 4
	// Generic autopilot only supporting simple waypoints
	MAV_AUTOPILOT_GENERIC_WAYPOINTS_ONLY MAV_AUTOPILOT = 5
	// Generic autopilot supporting waypoints and other simple navigation commands
	MAV_AUTOPILOT_GENERIC_WAYPOINTS_AND_SIMPLE_NAVIGATION_ONLY MAV_AUTOPILOT = 6
	// Generic autopilot supporting the full mission command set
	MAV_AUTOPILOT_GENERIC_MISSION_FULL MAV_AUTOPILOT = 7
	// No valid autopilot, e.g. a GCS or other MAVLink component
	MAV_AUTOPILOT_INVALID MAV_AUTOPILOT = 8
	// PPZ UAV - http://nongnu.org/paparazzi
	MAV_AUTOPILOT_PPZ MAV_AUTOPILOT = 9
	// UAV Dev Board
	MAV_AUTOPILOT_UDB MAV_AUTOPILOT = 10
	// FlexiPilot
	MAV_AUTOPILOT_FP MAV_AUTOPILOT = 11
	// PX4 Autopilot - http://px4.io/
	MAV_AUTOPILOT_PX4 MAV_AUTOPILOT = 12
	// SMACCMPilot - http://smaccmpilot.org
	MAV_AUTOPILOT_SMACCMPILOT MAV_AUTOPILOT = 13
	// AutoQuad -- http://autoquad.org
	MAV_AUTOPILOT_AUTOQUAD MAV_AUTOPILOT = 14
	// Armazila -- http://armazila.com
	MAV_AUTOPILOT_ARMAZILA MAV_AUTOPILOT = 15
	// Aerob -- http://aerob.ru
	MAV_AUTOPILOT_AEROB MAV_AUTOPILOT = 16
	// ASLUAV autopilot -- http://www.asl.ethz.ch
	MAV_AUTOPILOT_ASLUAV MAV_AUTOPILOT = 17
	// SmartAP Autopilot - http://sky-drones.com
	MAV_AUTOPILOT_SMARTAP MAV_AUTOPILOT = 18
	// AirRails - http://uaventure.com
	MAV_AUTOPILOT_AIRRAILS MAV_AUTOPILOT = 19
	// Fusion Reflex - https://fusion.engineering
	MAV_AUTOPILOT_REFLEX MAV_AUTOPILOT = 20
)

var labels_MAV_AUTOPILOT = map[MAV_AUTOPILOT]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_AUTOPILOT) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_AUTOPILOT[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_AUTOPILOT = map[string]MAV_AUTOPILOT{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_AUTOPILOT) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_AUTOPILOT[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_AUTOPILOT) String() string {
	if l, ok := labels_MAV_AUTOPILOT[e]; ok {
		return l
	}
	return "invalid value"
}
