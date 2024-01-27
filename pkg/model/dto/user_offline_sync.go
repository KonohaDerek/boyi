package dto

import (
	"fmt"
)

const (
	UserOfflineSyncKey = "user:%d:device:%s:offline_sync"
)

// 產生key 值
func GenerateOfflineSyncKey(userID uint64, deviceUID string) string {
	return fmt.Sprintf(UserOfflineSyncKey, userID, deviceUID)
}
