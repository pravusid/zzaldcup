package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchService_isAvailable(t *testing.T) {
	// GIVEN
	quota := 32

	// WHEN
	result := MatchService.isAvailable(quota)

	// THEN
	assert.True(t, result)
}

func TestMatchService_isAvailable_invalid(t *testing.T) {
	// GIVEN
	quota := 48

	// WHEN
	result := MatchService.isAvailable(quota)

	// THEN
	assert.False(t, result)
}

func TestMatchService_isSuitablePayload(t *testing.T) {
	// GIVEN
	sizeOfCompetitors := 64
	quota := 64

	// WHEN
	result := MatchService.isSuitablePayload(sizeOfCompetitors, quota)

	// THEN
	assert.True(t, result)
}

func TestMatchService_isSuitablePayload_valid_when_zero_competitor(t *testing.T) {
	// GIVEN
	sizeOfCompetitors := 0
	quota := 64

	// WHEN
	result := MatchService.isSuitablePayload(sizeOfCompetitors, quota)

	// THEN
	assert.True(t, result)
}

func TestMatchService_isSuitablePayload_invalid(t *testing.T) {
	// GIVEN
	sizeOfCompetitors := 62
	quota := 64

	// WHEN
	result := MatchService.isSuitablePayload(sizeOfCompetitors, quota)

	// THEN
	assert.False(t, result)
}