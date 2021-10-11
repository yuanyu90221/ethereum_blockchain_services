package tests

import (
	"reflect"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/yuanyu90221/ethereum_blockchain_services/configs"
)

func TestGetEnvConfig(t *testing.T) {
	tests := []struct {
		name string
		want configs.Config
	}{
		{
			name: "default",
			want: configs.Config{PORT: 5566},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := configs.GetEnvConfig(); !reflect.DeepEqual(got.PORT, tt.want.PORT) {
				t.Errorf("GetEnvConfig().PORT = %v, want %v", got, tt.want.PORT)
			}
		})
	}
}
