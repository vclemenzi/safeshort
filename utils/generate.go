package utils

import (
	"crypto/md5"
	"encoding/base64"
	"strconv"
	"strings"
	"time"
)

func GenerateID() string {
	milliseconds := time.Now().UnixNano() / int64(time.Millisecond)

	hasher := md5.New()
	hasher.Write([]byte(strconv.FormatInt(milliseconds, 10)))
	md5Hash := hasher.Sum(nil)

	base64Str := base64.RawURLEncoding.EncodeToString(md5Hash)

	base64Str = strings.ReplaceAll(base64Str, "+", "-")
	base64Str = strings.ReplaceAll(base64Str, "/", "_")

	shortID := base64Str[:9]

	return shortID
}
