package test

import (
	"os"
	"testing"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/requests"
)

func TestGetSingleAccount(t *testing.T) {
	token := os.Getenv("CANVAS_API_TOKEN")
	canvasURL := "atomicjolt.instructure.com"
	testAccountID := "578"

	getSingleAccount := requests.GetSingleAccount{}
	getSingleAccount.Path.ID = testAccountID

	type fields struct {
		AccessToken string
		CanvasURL   string
		UserAgent   string
	}
	type args struct {
		canvasRequest canvasapi.CanvasRequest
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantSuccess bool
	}{
		{
			name: "Get Single Account",
			fields: fields{
				AccessToken: token,
				CanvasURL:   canvasURL,
			},
			args: args{
				canvasRequest: &getSingleAccount,
			},
			wantSuccess: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			canvas := canvasapi.New(tt.fields.AccessToken, tt.fields.CanvasURL)
			account, err := getSingleAccount.Do(&canvas)

			if err != nil && tt.wantSuccess {
				t.Errorf("getSingleAccount failed %v", err)
			}

			t.Logf("getSingleAccount returned %v", account)
		})
	}
}
