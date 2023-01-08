package util

import (
	"go.uber.org/zap"
)

var Logger, _ = zap.NewProduction()
