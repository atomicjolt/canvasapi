package canvasapi

import "net/url"

type CanvasRequest interface {
	GetMethod() string
	GetURLPath() string
	GetQuery() (string, error)
	GetBody() (url.Values, error)
	GetJSON() ([]byte, error)
	HasErrors() error
}

type CanvasModel interface {
	HasError() error
}
