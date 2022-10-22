package utils

import (
	"strconv"
	"strings"
	"time"
)

func ParseDuration(d string) (time.Duration, error) {
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	if strings.HasSuffix(d, "d") {
		day := strings.TrimSuffix(d, "d")
		d, _ := strconv.Atoi(day)
		dr = 3600 * 24 * time.Duration(d) * time.Second
		return dr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}
