// File: wire.go

//go:build wireinject
// +build wireinject

package logger

import (
	"github.com/google/wire"
)

// InitializeLogger provides an instance of IZapLogger using Google Wire.
func ProvideLogger() IZapLogger {
	wire.Build(NewZapLogger)
	return nil // Wire will replace this during code generation.
}
