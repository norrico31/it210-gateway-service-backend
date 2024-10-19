package utils

import (
	"fmt"

	"github.com/norrico31/it210-gateway-service-backend/config"
)

func HandlePathV1(p string) string {
	path := fmt.Sprintf(`/api/%s/%s`, config.Envs.AppVersion1, p)
	return path
}
