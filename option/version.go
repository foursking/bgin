package option

import (
	"github.com/foursking/bgin/types"
)

func WithVersionKey(key string) App {
	return func(app *types.App) {
		app.Version.Key = key
	}
}

func WithVersionLatest(latest string) App {
	return func(app *types.App) {
		app.Version.Latest = latest
	}
}
