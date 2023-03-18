package asciinema_parser

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"strconv"
)

// ------------------------------------------------- --------------------------------------------------------------------

type Version uint

const (
	VersionUnknown Version = 0
	Version1       Version = 1
	Version2       Version = 2
)

// ------------------------------------------------- --------------------------------------------------------------------

func DetectVersion(ctx context.Context, asciiCastString string) (Version, error) {
	versionResult := gjson.Get(asciiCastString, "version")
	atoi, err := strconv.Atoi(versionResult.String())
	if err != nil {
		return VersionUnknown, err
	}
	switch Version(atoi) {
	case Version1:
		return Version1, nil
	case Version2:
		return Version2, nil
	default:
		return VersionUnknown, fmt.Errorf("unknown version: %s", versionResult.String())
	}
}

// ------------------------------------------------- --------------------------------------------------------------------

// ParseV1 解析V1版本的Ascii Cast 文件
// https://github.com/asciinema/asciinema/blob/develop/doc/asciicast-v1.md
func ParseV1(ctx context.Context, asciiCastV1Bytes []byte) (*AsciiCastV1, error) {
	cast := &AsciiCastV1{}
	err := json.Unmarshal(asciiCastV1Bytes, cast)
	if err != nil {
		return nil, err
	}
	if err := cast.Check(); err != nil {
		return nil, err
	}
	return cast, nil
}

// ------------------------------------------------- --------------------------------------------------------------------

// ParseV2 解析V2版本的Ascii Cast文件
// v2格式规范的文档： https://github.com/asciinema/asciinema/blob/develop/doc/asciicast-v2.md
func ParseV2(ctx context.Context, asciiCastV2Bytes []byte) (*AsciiCastV2, error) {
	scanner := bufio.NewScanner(bytes.NewBuffer(asciiCastV2Bytes))

	// 读取第一行
	if !scanner.Scan() {
		return nil, fmt.Errorf("asciiCastV2Bytes %s is invalid", string(asciiCastV2Bytes))
	}
	firstLineBytes := scanner.Bytes()
	cast := &AsciiCastV2{}
	err := json.Unmarshal(firstLineBytes, cast)
	if err != nil {
		return nil, err
	}

	// 读取后续的行，追加到数组中，这里暂时不考虑放不下的情况
	lineNum := 2
	for scanner.Scan() {

		lineBytes := scanner.Bytes()
		slice := make([]any, 0)
		err := json.Unmarshal(lineBytes, &slice)
		if err != nil {
			return nil, err
		}

		if len(slice) != 3 {
			return nil, fmt.Errorf("line %d is invalid event: %s", lineNum, string(lineBytes))
		}

		delay, ok := slice[0].(float64)
		if !ok {
			return nil, fmt.Errorf("line %d is invalid event: %s", lineNum, string(lineBytes))
		}

		eventTypeString, ok := slice[1].(string)
		if !ok {
			return nil, fmt.Errorf("line %d is invalid event: %s", lineNum, string(lineBytes))
		}
		eventType, err := ParseEventType(eventTypeString)
		if err != nil {
			return nil, err
		}

		data, ok := slice[2].(string)
		if !ok {
			return nil, fmt.Errorf("line %d is invalid event: %s", lineNum, string(lineBytes))
		}
		cast.EventStream = append(cast.EventStream, &Event{
			cast:      cast,
			Delay:     delay,
			EventType: eventType,
			EventData: data,
		})
	}

	return cast, nil
}

// ------------------------------------------------- --------------------------------------------------------------------
