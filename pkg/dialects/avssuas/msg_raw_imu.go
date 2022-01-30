//autogenerated:yes
//nolint:revive,misspell,govet,lll
package avssuas

// The RAW IMU readings for a 9DOF sensor, which is identified by the id (default IMU1). This message should always contain the true raw values without any scaling to allow data capture and system debugging.
type MessageRawImu struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// X acceleration (raw)
	Xacc int16
	// Y acceleration (raw)
	Yacc int16
	// Z acceleration (raw)
	Zacc int16
	// Angular speed around X axis (raw)
	Xgyro int16
	// Angular speed around Y axis (raw)
	Ygyro int16
	// Angular speed around Z axis (raw)
	Zgyro int16
	// X Magnetic field (raw)
	Xmag int16
	// Y Magnetic field (raw)
	Ymag int16
	// Z Magnetic field (raw)
	Zmag int16
	// Id. Ids are numbered from 0 and map to IMUs numbered from 1 (e.g. IMU1 will have a message with id=0)
	Id uint8 `mavext:"true"`
	// Temperature, 0: IMU does not provide temperature values. If the IMU is at 0C it must send 1 (0.01C).
	Temperature int16 `mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageRawImu) GetID() uint32 {
	return 27
}
