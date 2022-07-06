package gct

import (
	"testing"
)

func TestIsBlank(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{" \t\n\r\v\f\x00　", true},
		{"0", false},
		{"hello", false},
	}
	for _, test := range tests {
		actual := StringUtils.IsBlank(test.param)
		if actual != test.expected {
			t.Errorf("%s must be %t", test.param, test.expected)
		}
	}
}

func TestIsNotBlank(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{" \t\n\r\v\f\x00　", false},
		{"0", true},
		{"hello", true},
	}
	for _, test := range tests {
		actual := StringUtils.IsNotBlank(test.param)
		if actual != test.expected {
			t.Errorf("%s must be %t", test.param, test.expected)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
	}
	for _, test := range tests {
		actual := StringUtils.IsEmpty(test.param)
		if actual != test.expected {
			t.Errorf("%s must be %t", test.param, test.expected)
		}
	}
}

func TestIsNotEmpty(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"0", true},
	}
	for _, test := range tests {
		actual := StringUtils.IsNotEmpty(test.param)
		if actual != test.expected {
			t.Errorf("%s must be %t", test.param, test.expected)
		}
	}
}

func TestIsBlankOrEmpty(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{" \t\n\r\v\f\x00　", true},
		{"0", false},
		{"hello", false},
	}
	for _, test := range tests {
		actual := StringUtils.IsBlankOrEmpty(test.param)
		if actual != test.expected {
			t.Errorf("%s must be %t", test.param, test.expected)
		}
	}
}

func TestIsNotBlankOrEmpty(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{" \t\n\r\v\f\x00　", false},
		{"0", true},
		{"hello", true},
	}
	for _, test := range tests {
		actual := StringUtils.IsNotBlankOrEmpty(test.param)
		if actual != test.expected {
			t.Errorf("%s must be %t", test.param, test.expected)
		}
	}
}
