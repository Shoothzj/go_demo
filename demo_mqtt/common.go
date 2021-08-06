package demo_mqtt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func clientId(deviceId string) string {
	now := time.Now()
	return fmt.Sprintf("%s_0_0_%s", deviceId, now.Format("2006010215"))
}

func eventTime() string {
	return time.Now().Format("20060102T150405Z")
}

func computeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
