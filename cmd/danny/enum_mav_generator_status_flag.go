//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package danny

import (
    "strings"
    "fmt"
)
// Flags to report status/failure cases for a power generator (used in GENERATOR_STATUS). Note that FAULTS are conditions that cause the generator to fail. Warnings are conditions that require attention before the next use (they indicate the system is not operating properly).
type MAV_GENERATOR_STATUS_FLAG uint32

const (
    // Generator is off.
    MAV_GENERATOR_STATUS_FLAG_OFF MAV_GENERATOR_STATUS_FLAG = 1
    // Generator is ready to start generating power.
    MAV_GENERATOR_STATUS_FLAG_READY MAV_GENERATOR_STATUS_FLAG = 2
    // Generator is generating power.
    MAV_GENERATOR_STATUS_FLAG_GENERATING MAV_GENERATOR_STATUS_FLAG = 4
    // Generator is charging the batteries (generating enough power to charge and provide the load).
    MAV_GENERATOR_STATUS_FLAG_CHARGING MAV_GENERATOR_STATUS_FLAG = 8
    // Generator is operating at a reduced maximum power.
    MAV_GENERATOR_STATUS_FLAG_REDUCED_POWER MAV_GENERATOR_STATUS_FLAG = 16
    // Generator is providing the maximum output.
    MAV_GENERATOR_STATUS_FLAG_MAXPOWER MAV_GENERATOR_STATUS_FLAG = 32
    // Generator is near the maximum operating temperature, cooling is insufficient.
    MAV_GENERATOR_STATUS_FLAG_OVERTEMP_WARNING MAV_GENERATOR_STATUS_FLAG = 64
    // Generator hit the maximum operating temperature and shutdown.
    MAV_GENERATOR_STATUS_FLAG_OVERTEMP_FAULT MAV_GENERATOR_STATUS_FLAG = 128
    // Power electronics are near the maximum operating temperature, cooling is insufficient.
    MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_OVERTEMP_WARNING MAV_GENERATOR_STATUS_FLAG = 256
    // Power electronics hit the maximum operating temperature and shutdown.
    MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_OVERTEMP_FAULT MAV_GENERATOR_STATUS_FLAG = 512
    // Power electronics experienced a fault and shutdown.
    MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_FAULT MAV_GENERATOR_STATUS_FLAG = 1024
    // The power source supplying the generator failed e.g. mechanical generator stopped, tether is no longer providing power, solar cell is in shade, hydrogen reaction no longer happening.
    MAV_GENERATOR_STATUS_FLAG_POWERSOURCE_FAULT MAV_GENERATOR_STATUS_FLAG = 2048
    // Generator controller having communication problems.
    MAV_GENERATOR_STATUS_FLAG_COMMUNICATION_WARNING MAV_GENERATOR_STATUS_FLAG = 4096
    // Power electronic or generator cooling system error.
    MAV_GENERATOR_STATUS_FLAG_COOLING_WARNING MAV_GENERATOR_STATUS_FLAG = 8192
    // Generator controller power rail experienced a fault.
    MAV_GENERATOR_STATUS_FLAG_POWER_RAIL_FAULT MAV_GENERATOR_STATUS_FLAG = 16384
    // Generator controller exceeded the overcurrent threshold and shutdown to prevent damage.
    MAV_GENERATOR_STATUS_FLAG_OVERCURRENT_FAULT MAV_GENERATOR_STATUS_FLAG = 32768
    // Generator controller detected a high current going into the batteries and shutdown to prevent battery damage.
    MAV_GENERATOR_STATUS_FLAG_BATTERY_OVERCHARGE_CURRENT_FAULT MAV_GENERATOR_STATUS_FLAG = 65536
    // Generator controller exceeded it's overvoltage threshold and shutdown to prevent it exceeding the voltage rating.
    MAV_GENERATOR_STATUS_FLAG_OVERVOLTAGE_FAULT MAV_GENERATOR_STATUS_FLAG = 131072
    // Batteries are under voltage (generator will not start).
    MAV_GENERATOR_STATUS_FLAG_BATTERY_UNDERVOLT_FAULT MAV_GENERATOR_STATUS_FLAG = 262144
    // Generator start is inhibited by e.g. a safety switch.
    MAV_GENERATOR_STATUS_FLAG_START_INHIBITED MAV_GENERATOR_STATUS_FLAG = 524288
    // Generator requires maintenance.
    MAV_GENERATOR_STATUS_FLAG_MAINTENANCE_REQUIRED MAV_GENERATOR_STATUS_FLAG = 1048576
    // Generator is not ready to generate yet.
    MAV_GENERATOR_STATUS_FLAG_WARMING_UP MAV_GENERATOR_STATUS_FLAG = 2097152
    // Generator is idle.
    MAV_GENERATOR_STATUS_FLAG_IDLE MAV_GENERATOR_STATUS_FLAG = 4194304
)

var labels_MAV_GENERATOR_STATUS_FLAG = map[MAV_GENERATOR_STATUS_FLAG]string{
    MAV_GENERATOR_STATUS_FLAG_OFF: "MAV_GENERATOR_STATUS_FLAG_OFF",
    MAV_GENERATOR_STATUS_FLAG_READY: "MAV_GENERATOR_STATUS_FLAG_READY",
    MAV_GENERATOR_STATUS_FLAG_GENERATING: "MAV_GENERATOR_STATUS_FLAG_GENERATING",
    MAV_GENERATOR_STATUS_FLAG_CHARGING: "MAV_GENERATOR_STATUS_FLAG_CHARGING",
    MAV_GENERATOR_STATUS_FLAG_REDUCED_POWER: "MAV_GENERATOR_STATUS_FLAG_REDUCED_POWER",
    MAV_GENERATOR_STATUS_FLAG_MAXPOWER: "MAV_GENERATOR_STATUS_FLAG_MAXPOWER",
    MAV_GENERATOR_STATUS_FLAG_OVERTEMP_WARNING: "MAV_GENERATOR_STATUS_FLAG_OVERTEMP_WARNING",
    MAV_GENERATOR_STATUS_FLAG_OVERTEMP_FAULT: "MAV_GENERATOR_STATUS_FLAG_OVERTEMP_FAULT",
    MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_OVERTEMP_WARNING: "MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_OVERTEMP_WARNING",
    MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_OVERTEMP_FAULT: "MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_OVERTEMP_FAULT",
    MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_FAULT: "MAV_GENERATOR_STATUS_FLAG_ELECTRONICS_FAULT",
    MAV_GENERATOR_STATUS_FLAG_POWERSOURCE_FAULT: "MAV_GENERATOR_STATUS_FLAG_POWERSOURCE_FAULT",
    MAV_GENERATOR_STATUS_FLAG_COMMUNICATION_WARNING: "MAV_GENERATOR_STATUS_FLAG_COMMUNICATION_WARNING",
    MAV_GENERATOR_STATUS_FLAG_COOLING_WARNING: "MAV_GENERATOR_STATUS_FLAG_COOLING_WARNING",
    MAV_GENERATOR_STATUS_FLAG_POWER_RAIL_FAULT: "MAV_GENERATOR_STATUS_FLAG_POWER_RAIL_FAULT",
    MAV_GENERATOR_STATUS_FLAG_OVERCURRENT_FAULT: "MAV_GENERATOR_STATUS_FLAG_OVERCURRENT_FAULT",
    MAV_GENERATOR_STATUS_FLAG_BATTERY_OVERCHARGE_CURRENT_FAULT: "MAV_GENERATOR_STATUS_FLAG_BATTERY_OVERCHARGE_CURRENT_FAULT",
    MAV_GENERATOR_STATUS_FLAG_OVERVOLTAGE_FAULT: "MAV_GENERATOR_STATUS_FLAG_OVERVOLTAGE_FAULT",
    MAV_GENERATOR_STATUS_FLAG_BATTERY_UNDERVOLT_FAULT: "MAV_GENERATOR_STATUS_FLAG_BATTERY_UNDERVOLT_FAULT",
    MAV_GENERATOR_STATUS_FLAG_START_INHIBITED: "MAV_GENERATOR_STATUS_FLAG_START_INHIBITED",
    MAV_GENERATOR_STATUS_FLAG_MAINTENANCE_REQUIRED: "MAV_GENERATOR_STATUS_FLAG_MAINTENANCE_REQUIRED",
    MAV_GENERATOR_STATUS_FLAG_WARMING_UP: "MAV_GENERATOR_STATUS_FLAG_WARMING_UP",
    MAV_GENERATOR_STATUS_FLAG_IDLE: "MAV_GENERATOR_STATUS_FLAG_IDLE",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_GENERATOR_STATUS_FLAG) MarshalText() ([]byte, error) {
    var names []string
    for mask, label := range labels_MAV_GENERATOR_STATUS_FLAG {
        if e&mask == mask {
            names = append(names, label)
        }
    }
    return []byte(strings.Join(names, " | ")), nil
}


// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_GENERATOR_STATUS_FLAG) UnmarshalText(text []byte) error {
    labels := strings.Split(string(text), " | ")
    var mask MAV_GENERATOR_STATUS_FLAG
    for _, label := range labels {
        found := false
        for value, l := range labels_MAV_GENERATOR_STATUS_FLAG {
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
func (e MAV_GENERATOR_STATUS_FLAG) String() string {
    val, _ := e.MarshalText()
    return string(val)
}
