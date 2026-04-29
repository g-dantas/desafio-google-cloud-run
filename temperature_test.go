package main

import (
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		in, want float64
	}{
		{0, 32},
		{100, 212},
		{28.5, 83.3},
		{-40, -40},
	}
	for _, tt := range tests {
		got := CelsiusToFahrenheit(tt.in)
		if !almostEqual(got, tt.want) {
			t.Errorf("CelsiusToFahrenheit(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	tests := []struct {
		in, want float64
	}{
		{0, 273},
		{28.5, 301.5},
		{-273, 0},
	}
	for _, tt := range tests {
		got := CelsiusToKelvin(tt.in)
		if !almostEqual(got, tt.want) {
			t.Errorf("CelsiusToKelvin(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
