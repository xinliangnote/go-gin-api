package uuid

import (
	"github.com/google/uuid"
)

func GenUUID() string {
	u, _ := uuid.NewRandom()
	return u.String()
}
