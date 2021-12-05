package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_powerConsumption(t *testing.T) {
	lines := []string{
		"00100", //1
		"11110", //2
		"10110", //3
		"10111", //4
		"10101", //5
		"01111", //6
		"00111", //7
		"11100", //8
		"10000", //9
		"11001", //10
		"00010", //11
		"01010", //12
	}

	power := powerConsumption(lines)
	assert.Equal(t, int64(198), power)
}
