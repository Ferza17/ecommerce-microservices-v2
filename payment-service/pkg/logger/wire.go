//go:build wireinject
// +build wireinject

package logger

import (
	"github.com/google/wire"
)

// ProvideZapLogger wires dependencies for IZapLogger
func ProvideZapLogger() IZapLogger {
	wire.Build(Set)
	return nil
}
