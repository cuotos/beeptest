package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWPMCalc(t *testing.T) {
	tcs := []struct {
		inputWpm         int
		expectedDuration time.Duration
	}{
		{30, time.Millisecond * 40},
		{1, time.Millisecond * 1200},
	}

	for _, tc := range tcs {
		actual := wpm(tc.inputWpm)
		assert.Equal(t, tc.expectedDuration, actual)
	}
}
