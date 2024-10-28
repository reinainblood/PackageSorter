package sorter

import (
	"testing"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		name          string
		width, height float64
		length, mass  float64
		expected      string
	}{
		{"Standard package", 10, 10, 10, 10, "STANDARD"},
		{"Heavy package", 10, 10, 10, 25, "SPECIAL"},
		{"Bulky package", 100, 200, 100, 10, "SPECIAL"},
		{"Rejected package", 100, 100, 200, 30, "REJECTED"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Sort(tc.width, tc.height, tc.length, tc.mass)
			if result != tc.expected {
				t.Errorf("Sort(%v, %v, %v, %v) = %v; want %v",
					tc.width, tc.height, tc.length, tc.mass,
					result, tc.expected)
			}
		})
	}
}

func TestIsHeavy(t *testing.T) {
	testCases := []struct {
		name     string
		dims     Dimensions
		expected bool
	}{
		{"Not heavy", Dimensions{10, 10, 10, 10}, false},
		{"Heavy", Dimensions{10, 10, 10, 25}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsHeavy(tc.dims)
			if result != tc.expected {
				t.Errorf("IsHeavy(%v) = %v; want %v",
					tc.dims, result, tc.expected)
			}
		})
	}
}

func TestIsBulky(t *testing.T) {
	testCases := []struct {
		name     string
		dims     Dimensions
		expected bool
	}{
		{"Not bulky", Dimensions{10, 10, 10, 10}, false},
		{"Bulky volume", Dimensions{100, 100, 1000, 100}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsBulky(tc.dims)
			if result != tc.expected {
				t.Errorf("IsBulky(%v) = %v; want %v",
					tc.dims, result, tc.expected)
			}
		})
	}
}
