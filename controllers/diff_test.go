package controllers

import (
	"testing"
)

func boolPtr(b bool) *bool {
	return &b
}

func boolPtrPtr(b bool) **bool {
	p := &b
	return &p
}

func strPtr(s string) *string {
	return &s
}

type dummyModel struct {
	Name    string
	Enabled *bool
	Double  **bool
	Rate    *int64
}

func TestNeedsUpdate(t *testing.T) {
	bTrue1 := true
	bTrue2 := true
	bFalse := false

	tests := []struct {
		name     string
		broker   interface{}
		state    interface{}
		expected bool
	}{
		{
			name: "Identical pointers with same value should not need update",
			broker: dummyModel{
				Name:    "test",
				Enabled: &bTrue1,
			},
			state: dummyModel{
				Name:    "test",
				Enabled: &bTrue2,
			},
			expected: false,
		},
		{
			name: "Different pointer values should need update",
			broker: dummyModel{
				Name:    "test",
				Enabled: &bTrue1,
			},
			state: dummyModel{
				Name:    "test",
				Enabled: &bFalse,
			},
			expected: true,
		},
		{
			name: "Nested double pointer identical value",
			broker: dummyModel{
				Name:   "test",
				Double: boolPtrPtr(true),
			},
			state: dummyModel{
				Name:   "test",
				Double: boolPtrPtr(true),
			},
			expected: false,
		},
		{
			name: "Nested double pointer different value",
			broker: dummyModel{
				Name:   "test",
				Double: boolPtrPtr(true),
			},
			state: dummyModel{
				Name:   "test",
				Double: boolPtrPtr(false),
			},
			expected: true,
		},
		{
			name: "State omits field (zero value / nil) - should not trigger update",
			broker: dummyModel{
				Name:    "test",
				Enabled: &bTrue1,
				Rate:    nil,
			},
			state: dummyModel{
				Name: "test", // Enabled is nil in state
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NeedsUpdate(tt.broker, tt.state)
			if got != tt.expected {
				t.Errorf("NeedsUpdate() = %v, want %v", got, tt.expected)
			}
		})
	}
}

