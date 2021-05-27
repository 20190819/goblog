package types

import (
	"strconv"

	"github.com/yangliang4488/goblog/pkg/logger"
)

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		logger.LogError(err)
	}
	return i
}
