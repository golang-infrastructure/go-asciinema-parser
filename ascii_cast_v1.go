package asciinema_parser

import (
	"fmt"
)

// ------------------------------------------------- --------------------------------------------------------------------

// AsciiCastV1 v1格式的规范文档： https://github.com/asciinema/asciinema/blob/develop/doc/asciicast-v1.md
type AsciiCastV1 struct {
	Version Version `json:"version"`

	Width        uint              `json:"width" yaml:"width" mapstructure:"width"`
	Height       uint              `json:"height" yaml:"height" mapstructure:"height"`
	Duration     float64           `json:"duration" yaml:"duration" mapstructure:"duration"`
	Command      string            `json:"command" yaml:"command" mapstructure:"command"`
	Title        string            `json:"title" yaml:"title" mapstructure:"title"`
	Env          map[string]string `json:"env" yaml:"env" mapstructure:"env"`
	StdoutFrames []Frame           `json:"stdout" yaml:"stdout" mapstructure:"stdout"`
}

func (x *AsciiCastV1) Check() error {

	if x.Version != Version1 {
		return fmt.Errorf("AsciiCast version %v is not equals 1", x.Version)
	}

	return nil
}

// ------------------------------------------------- --------------------------------------------------------------------

type Frame []any

func (x Frame) GetDelayE() (float64, error) {
	if len(x) > 0 {
		v, ok := x[0].(float64)
		if !ok {
			return 0, fmt.Errorf("cast %s to float failed", x[0])
		}
		return v, nil
	} else {
		return 0, nil
	}
}

func (x Frame) GetDelay() float64 {
	v, _ := x.GetDelayE()
	return v
}

func (x Frame) GetDataE() (string, error) {
	if len(x) > 1 {
		v, ok := x[1].(string)
		if !ok {
			return "", fmt.Errorf("cast %s to string failed", x[0])
		}
		return v, nil
	} else {
		return "", nil
	}
}

func (x Frame) GetData() string {
	v, _ := x.GetDataE()
	return v
}

// ------------------------------------------------- --------------------------------------------------------------------
