package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_string_returns_uint_value(t *testing.T) {
	assert.Equal(t, uint64(30), ParseInt("30", 10))
}

func Test_empty_string_must_return_default_value(t *testing.T) {
	assert.Equal(t, uint64(10), ParseInt("", 10))
}

func Test_json_string_to_json_map(t *testing.T) {
	// GIVEN
	jsonString := `{"test": "yes", "hello": "no", "duck": "bird"}`

	// WHEN
	obj, _ := ConvertJsonToMap(jsonString)

	// THEN
	assert.Equal(t, "yes", obj["test"])
	assert.Equal(t, "no", obj["hello"])
	assert.NotEqual(t, "donald", obj["duck"])
}
