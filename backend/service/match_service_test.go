package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchService_isAvailable(t *testing.T) {
	// GIVEN
	quota := 32

	// WHEN
	result := Match.isAvailable(quota)

	// THEN
	assert.NoError(t, result)
}

func TestMatchService_isAvailable_invalid(t *testing.T) {
	// GIVEN
	quota := 48

	// WHEN
	result := Match.isAvailable(quota)

	// THEN
	assert.Error(t, result)
}

func TestMatchService_isSuitablePayload(t *testing.T) {
	// GIVEN
	sizeOfCompetitors := 64
	quota := 64

	// WHEN
	result := Match.isSuitablePayload(sizeOfCompetitors, quota)

	// THEN
	assert.True(t, result)
}

func TestMatchService_isSuitablePayload_invalid(t *testing.T) {
	// GIVEN
	sizeOfCompetitors := 62
	quota := 64

	// WHEN
	result := Match.isSuitablePayload(sizeOfCompetitors, quota)

	// THEN
	assert.False(t, result)
}
