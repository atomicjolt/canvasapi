package canvasapi

type CanvasRequest interface {
	GetMethod() string
	GetURLPath() string
	GetQuery() (string, error)
	GetBody() (string, error)
	HasErrors() error
}

type CanvasModel interface {
	HasError() error
}
