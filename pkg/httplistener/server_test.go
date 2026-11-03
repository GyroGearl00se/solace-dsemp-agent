package httplistener

import "testing"

func TestConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{
			name:    "disabled listener does not require credentials",
			config:  Config{},
			wantErr: false,
		},
		{
			name: "enabled listener requires both credentials",
			config: Config{
				Enabled: true,
			},
			wantErr: true,
		},
		{
			name: "enabled listener rejects whitespace credentials",
			config: Config{
				Enabled:  true,
				Username: "   ",
				Password: "secret",
			},
			wantErr: true,
		},
		{
			name: "enabled listener accepts credentials",
			config: Config{
				Enabled:  true,
				Username: "agent",
				Password: "secret",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Fatalf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
