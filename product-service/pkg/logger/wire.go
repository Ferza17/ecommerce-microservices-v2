//go:build wireinject
// +build wireinject

package logger

import (
	"github.com/google/wire"
)

func ProvideLogger() IZapLogger {
	wire.Build(
		Set,
	)
	return nil
}
