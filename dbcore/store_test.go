package dbcore

import "testing"

func TestCanInitialize(t *testing.T) {
	if _, err := NewStore(); err != nil {
		t.Errorf("Error in : %v", err)
	}
}
