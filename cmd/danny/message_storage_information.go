//autogenerated:yes
//nolint:revive,misspell,govet,lll
package danny
// Information about a storage medium. This message is sent in response to a request with MAV_CMD_REQUEST_MESSAGE and whenever the status of the storage changes (STORAGE_STATUS). Use MAV_CMD_REQUEST_MESSAGE.param2 to indicate the index/id of requested storage: 0 for all, 1 for first, 2 for second, etc.
type MessageStorageInformation struct {
    // Timestamp (time since system boot).
    TimeBootMs uint32
    // Storage ID (1 for first, 2 for second, etc.)
    StorageId uint8
    // Number of storage devices
    StorageCount uint8
    // Status of storage
    Status STORAGE_STATUS `mavenum:"uint8"`
    // Total capacity. If storage is not ready (STORAGE_STATUS_READY) value will be ignored.
    TotalCapacity float32
    // Used capacity. If storage is not ready (STORAGE_STATUS_READY) value will be ignored.
    UsedCapacity float32
    // Available storage capacity. If storage is not ready (STORAGE_STATUS_READY) value will be ignored.
    AvailableCapacity float32
    // Read speed.
    ReadSpeed float32
    // Write speed.
    WriteSpeed float32
    // Type of storage
    Type STORAGE_TYPE `mavenum:"uint8" mavext:"true"`
    // Textual storage name to be used in UI (microSD 1, Internal Memory, etc.) This is a NULL terminated string. If it is exactly 32 characters long, add a terminating NULL. If this string is empty, the generic type is shown to the user.
    Name string `mavext:"true" mavlen:"32"`
    // Flags indicating whether this instance is preferred storage for photos, videos, etc.
    // Note: Implementations should initially set the flags on the system-default storage id used for saving media (if possible/supported).
    // This setting can then be overridden using MAV_CMD_SET_STORAGE_USAGE.
    // If the media usage flags are not set, a GCS may assume storage ID 1 is the default storage for all media types.
    StorageUsage STORAGE_USAGE_FLAG `mavenum:"uint8" mavext:"true"`
}

// GetID implements the message.Message interface.
func (*MessageStorageInformation) GetID() uint32 {
    return 261
}
