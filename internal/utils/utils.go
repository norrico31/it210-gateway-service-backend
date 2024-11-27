package utils

import (
	"fmt"
	"strings"

	"github.com/norrico31/it210-gateway-service-backend/config"
)

func HandlePathV1(p string) string {
	// Check if the path already starts with /api/v1 and avoid duplicating it
	if strings.HasPrefix(p, "/api/v1") {
		return p
	}

	// Otherwise, add the /api/v1 prefix
	path := fmt.Sprintf(`/api/%s/%s`, config.Envs.AppVersion, p)
	return path
}
