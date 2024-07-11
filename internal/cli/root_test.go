package cli

import "testing"

func TestRun(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		hasError bool
	}{
		{"No sub-command specified", []string{}, true},
		{"Unknown sub-command", []string{"unknown"}, true},
		{"Valid sub-command", []string{"model"}, false},
		{"Valid sub-command with arguments", []string{"service", "-name", "Test"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Run(tt.args)
			if (err != nil) != tt.hasError {
				t.Errorf("expected error %v, got error %v", tt.hasError, err)
			}
		})
	}
}
