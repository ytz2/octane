package handler

import "go.uber.org/fx"

// Module ...
var Module = fx.Provide(
	NewLotto,
)
