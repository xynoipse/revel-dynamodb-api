package database

import (
	"fmt"
	"time"
)

// AppendPrefix appends prefix to string
func AppendPrefix(prefix string, s string) string {
	return fmt.Sprintf("%s#%s", prefix, s)
}

// GetCurrentTimestamp return epoch format timestamp
func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}
