package jpushclient

import "encoding/json"

type LiveActivity struct {
	Event        string         `json:"event"`
	ContentState map[string]any `json:"content-state"`
	Alert        Alert          `json:"alert,omitempty"`
	BadgeAddNum  int            `json:"dismissal-date,omitempty"`
}
type Alert struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	Sound string `json:"sound,omitempty"`
}

func (a *LiveActivity) SetEvent(c string) *LiveActivity {
	a.Event = c
	return a
}

func (a *LiveActivity) SetBadgeAddNum(b int) *LiveActivity {
	a.BadgeAddNum = b
	return a
}

func (a *LiveActivity) AddContentState(key string, value any) *LiveActivity {
	if a.ContentState == nil {
		a.ContentState = make(map[string]any)
	}
	a.ContentState[key] = value
	return a
}

func (a *LiveActivity) SetAlert(alert Alert) *LiveActivity {
	a.Alert = alert
	return a
}

func (a *LiveActivity) ToJson() (string, error) {
	content, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (a *LiveActivity) ToBytes() ([]byte, error) {
	content, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return content, nil
}
