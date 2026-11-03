package reconciler

import "testing"

func TestIsVersionNewer(t *testing.T) {
	tests := []struct {
		newVer     string
		currentVer string
		expected   bool
	}{
		{"1.0.1", "1.0.0", true},
		{"1.0.0", "1.0.1", false},
		{"1.0.0", "1.0.0", false},
		{"1.0.0", "", true},
		{"", "1.0.0", false},
		{"", "", false},
		{"1.10.0", "1.9.0", true},
		{"2.0.0", "1.99.99", true},
		{"1.2.3", "1.2.3.4", false},
		{"1.2.3.4", "1.2.3", true},
		{"1.0.a", "1.0.0", true},
		{"1.0.0", "1.0.a", false},
		{"1.a.0", "1.b.0", false},
		{"1.b.0", "1.a.0", true},
	}

	for _, tt := range tests {
		result := IsVersionNewer(tt.newVer, tt.currentVer)
		if result != tt.expected {
			t.Errorf("IsVersionNewer(%q, %q) = %v; want %v", tt.newVer, tt.currentVer, result, tt.expected)
		}
	}
}

