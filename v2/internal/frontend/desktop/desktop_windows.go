//go:build windows
// +build windows

package desktop

import (
	"context"
	"github.com/wailsapp/zappie/v2/internal/binding"
	"github.com/wailsapp/zappie/v2/internal/frontend"
	"github.com/wailsapp/zappie/v2/internal/frontend/desktop/windows"
	"github.com/wailsapp/zappie/v2/internal/logger"
	"github.com/wailsapp/zappie/v2/pkg/options"
)

func NewFrontend(ctx context.Context, appoptions *options.App, logger *logger.Logger, appBindings *binding.Bindings, dispatcher frontend.Dispatcher) frontend.Frontend {
	return windows.NewFrontend(ctx, appoptions, logger, appBindings, dispatcher)
}
