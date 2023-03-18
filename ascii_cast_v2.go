package asciinema_parser

import (
	"fmt"
	"strings"
	"time"
)

// 格式规范文档： https://github.com/asciinema/asciinema/blob/develop/doc/asciicast-v2.md
type AsciiCastV2 struct {

	// 表示文件的格式版本
	Version Version `json:"version" yaml:"version" mapstructure:"version"`

	// 屏幕的宽度，但是是列
	Width uint `json:"width" yaml:"width" mapstructure:"width"`

	// 屏幕的高度，但是是行
	Height uint `json:"height" yaml:"height" mapstructure:"height"`

	// 录制开始时间
	Timestamp uint64 `json:"timestamp" yaml:"timestamp" mapstructure:"timestamp"`
	time      time.Time

	// 录制持续时间
	Duration float64 `json:"duration" yaml:"duration" mapstructure:"duration"`

	Command string `json:"command" yaml:"command" mapstructure:"command"`

	Title string `json:"title" yaml:"title" mapstructure:"title"`

	Theme *Theme `json:"theme" yaml:"theme" mapstructure:"theme"`

	//
	IdleTimeLimit float64 `json:"idle_time_limit" yaml:"idle_time_limit" mapstructure:"idle_time_limit"`

	// 相关环境变量
	EnvMap map[string]string `json:"env" yaml:"env" mapstructure:"env"`

	EventStream []*Event `json:"event_stream" yaml:"event_stream" mapstructure:"event_stream"`
}

func (x *AsciiCastV2) GetTime() time.Time {
	if x.time.IsZero() {
		x.time = time.Unix(int64(x.Timestamp), 0)
	}
	return x.time
}

// ------------------------------------------------- --------------------------------------------------------------------

// Theme 可以设置一个初始化的主题
type Theme struct {

	// 主题的前景色
	FrontColor string `json:"fg" yaml:"fg" mapstructure:"fg"`

	// 主题的背景色
	BackgroundColor string `json:"bg" yaml:"bg" mapstructure:"bg"`

	// 8或者16个颜色的调色板
	Palette string `json:"palette" yaml:"palette" mapstructure:"palette"`
}

// ------------------------------------------------- --------------------------------------------------------------------

type EventType string

const (
	EventTypeUnknown = ""
	EventTypeOutput  = "o"
	EventTypeInput   = "i"
)

func ParseEventType(eventType string) (EventType, error) {
	switch strings.ToLower(eventType) {
	case EventTypeOutput:
		return EventTypeOutput, nil
	case EventTypeInput:
		return EventTypeInput, nil
	default:
		return EventTypeUnknown, fmt.Errorf("value %s is not a valid event type", eventType)
	}
}

// ------------------------------------------------- --------------------------------------------------------------------

// Event 事件流中的某一个事件
type Event struct {
	cast *AsciiCastV2

	// 时间发生的时间距开始时间的偏移
	Delay float64 `json:"delay" yaml:"delay" mapstructure:"delay"`

	// 事件的类型，v2的话是有output和input
	EventType EventType `json:"event_type" yaml:"event_type" mapstructure:"event_type"`

	// 事件的数据，通常是显示的内容
	EventData string `json:"event_data" yaml:"event_data" mapstructure:"event_data"`
}

func (x *Event) GetEventTime() time.Time {
	return x.cast.GetTime().Add(time.Millisecond * time.Duration(x.Delay*1000))
}

func (x *Event) GetAsciiCast() *AsciiCastV2 {
	return x.cast
}

// ------------------------------------------------- --------------------------------------------------------------------
