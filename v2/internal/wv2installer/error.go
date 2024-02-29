//go:build windows && wv2runtime.error
// +build windows,wv2runtime.error

package wv2installer

import (
	"fmt"
	"github.com/wailsapp/zappie/v2/internal/webview2runtime"
	"github.com/wailsapp/zappie/v2/pkg/options/windows"
)

func doInstallationStrategy(installStatus installationStatus, messages *windows.Messages) error {
	_ = webview2runtime.Error(messages.ContactAdmin, messages.Error)
	return fmt.Errorf(messages.Webview2NotInstalled)
}
