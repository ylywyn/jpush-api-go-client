package jpushclient

import (
	"encoding/json"
	"time"
)

type ScheduleRequest struct {
	Cid     string           `json:"cid"`
	Name    string           `json:"name"`
	Enabled bool             `json:"enabled"`
	Push    *PushRequest     `json:"push"`
	Trigger *ScheduleTrigger `json:"trigger"`
}
type ScheduleTrigger struct {
	Single     *ScheduleTriggerSingle     `json:"single,omitempty"`
	Periodical *ScheduleTriggerPeriodical `json:"periodical,omitempty"`
}

type ScheduleTriggerSingle struct {
	Timer string `json:"time,omitempty"`
}

type ScheduleTriggerPeriodical struct {
	Start     string `json:"start,omitempty"`
	End       string `json:"end,omitempty"`
	Time      string `json:"time,omitempty"`
	TimeUnit  string `json:"time_unit,omitempty"`
	Frequency int    `json:"frequency,omitempty"`
	Point     any    `json:"point,omitempty"`
}

func NewSchedule(name, cid string, enabled bool, push *PushRequest) *ScheduleRequest {
	return &ScheduleRequest{
		Cid:     cid,
		Name:    name,
		Enabled: enabled,
		Push:    push,
	}
}
func (s *ScheduleRequest) SingleTrigger(t time.Time) {
	s.Trigger = &ScheduleTrigger{
		Single: &ScheduleTriggerSingle{
			Timer: t.Format("2006-01-02 15:04:05"),
		},
	}
}

func (s *ScheduleRequest) PeriodicalTrigger(start, end time.Time, timeUnit string, frequency int, point []string) {
	s.Trigger = &ScheduleTrigger{
		Periodical: &ScheduleTriggerPeriodical{
			Start:     start.Format("2006-01-02 15:04:05"),
			End:       end.Format("2006-01-02 15:04:05"),
			Time:      start.Format("15:04:05"),
			TimeUnit:  timeUnit,
			Frequency: frequency,
			Point:     point,
		},
	}
}

func (s *ScheduleRequest) ToJson() (string, error) {
	content, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (s *ScheduleRequest) ToBytes() ([]byte, error) {
	content, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return content, nil
}
