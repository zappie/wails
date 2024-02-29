package runtime

import (
	"context"

	"github.com/wailsapp/zappie/v2/internal/frontend"
)

type Screen = frontend.Screen

// ScreenGetAllScreens returns all screens
func ScreenGetAll(ctx context.Context) ([]Screen, error) {
	appFrontend := getFrontend(ctx)
	return appFrontend.ScreenGetAll()
}
