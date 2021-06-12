package canvasapi

import (
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		accessToken string
		canvasURL   string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "create Canvas",
			args: args{
				accessToken: "test",
				canvasURL:   "https://canvas.example.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			canvas := New(tt.args.accessToken, tt.args.canvasURL)
			if canvas.AccessToken != tt.args.accessToken || canvas.CanvasURL != tt.args.canvasURL {
				t.Errorf("expected New to create a canvas instance with accessToken %v and canvasURL %v", tt.args.accessToken, tt.args.canvasURL)
			}
		})
	}
}
