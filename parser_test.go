package asciinema_parser

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseV1(t *testing.T) {
	asciiCastV1Bytes, err := os.ReadFile("./test_data/ascii_cast_v1.cast")
	assert.Nil(t, err)
	v1, err := ParseV1(context.Background(), asciiCastV1Bytes)
	assert.Nil(t, err)
	assert.NotNil(t, v1)

	for _, frame := range v1.StdoutFrames {

		data, err := frame.GetDataE()
		assert.Nil(t, err)
		assert.NotEmpty(t, data)

		_, err = frame.GetDelayE()
		assert.Nil(t, err)
	}
}

func TestParseV2(t *testing.T) {
	asciiCastV2Bytes, err := os.ReadFile("./test_data/ascii_cast_v2.cast")
	assert.Nil(t, err)
	v2, err := ParseV2(context.Background(), asciiCastV2Bytes)
	assert.Nil(t, err)
	assert.NotNil(t, v2)
}

func TestDetectVersion(t *testing.T) {

	asciiCastV1Bytes, err := os.ReadFile("./test_data/ascii_cast_v1.cast")
	assert.Nil(t, err)
	version, err := DetectVersion(context.Background(), string(asciiCastV1Bytes))
	assert.Nil(t, err)
	assert.Equal(t, Version1, version)

	asciiCastV2Bytes, err := os.ReadFile("./test_data/ascii_cast_v2.cast")
	assert.Nil(t, err)
	version, err = DetectVersion(context.Background(), string(asciiCastV2Bytes))
	assert.Nil(t, err)
	assert.Equal(t, Version2, version)

}
