//autogenerated:yes
//nolint:revive,misspell,govet,lll
package danny
// The IMU readings in SI units in NED body frame
type MessageHilSensor struct {
    // Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
    TimeUsec uint64
    // X acceleration
    Xacc float32
    // Y acceleration
    Yacc float32
    // Z acceleration
    Zacc float32
    // Angular speed around X axis in body frame
    Xgyro float32
    // Angular speed around Y axis in body frame
    Ygyro float32
    // Angular speed around Z axis in body frame
    Zgyro float32
    // X Magnetic field
    Xmag float32
    // Y Magnetic field
    Ymag float32
    // Z Magnetic field
    Zmag float32
    // Absolute pressure
    AbsPressure float32
    // Differential pressure (airspeed)
    DiffPressure float32
    // Altitude calculated from pressure
    PressureAlt float32
    // Temperature
    Temperature float32
    // Bitmap for fields that have updated since last message
    FieldsUpdated HIL_SENSOR_UPDATED_FLAGS `mavenum:"uint32"`
    // Sensor ID (zero indexed). Used for multiple sensor inputs
    Id uint8 `mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageHilSensor) GetID() uint32 {
    return 107
}
