package jpushclient

import (
	"encoding/json"
	"time"
)

type Schedule struct {
	Name    string                 `json:"name"`
	Enabled bool                   `json:"enabled"`
	Trigger map[string]interface{} `json:"trigger"`
	Push    *PayLoad               `json:"push"`
}

func NewSchedule(name string, enabled bool, push *PayLoad) *Schedule {
	return &Schedule{
		Name:    name,
		Enabled: enabled,
		Push:    push,
	}
}
func (s *Schedule) SingleTrigger(t time.Time) {
	s.Trigger = map[string]interface{}{
		"single": map[string]interface{}{
			"time": t.Format("2006-01-02 15:04:05"),
		},
	}

}

func (s *Schedule) PeriodicalTrigger(start time.Time, end time.Time, time time.Time, timeUnit string, frequency int, point []string) {
	s.Trigger = map[string]interface{}{
		"periodical": map[string]interface{}{
			"start":     start.Format("2006-01-02 15:04:05"),
			"end":       start.Format("2006-01-02 15:04:05"),
			"time":      start.Format("15:04:05"),
			"time_unit": timeUnit,
			"frequency": frequency,
			"point":     point,
		},
	}
}
func (this *Schedule) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return content, nil
}
