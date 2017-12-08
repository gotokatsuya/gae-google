package random

import (
	"encoding/base64"
	"strings"

	"github.com/google/uuid"
)

var escaper = strings.NewReplacer("9", "99", "-", "90", "_", "91")

func UUID() string {
	uid := uuid.Must(uuid.NewUUID()).String()
	return escaper.Replace(base64.RawURLEncoding.EncodeToString([]byte(uid)))
}
