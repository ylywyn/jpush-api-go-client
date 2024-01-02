package jpushclient

import "encoding/json"

type QuickAppNotice struct {
	Title  string                 `json:"title"`
	Alert  string                 `json:"alert"`
	Page   string                 `json:"page"`
	Extras map[string]interface{} `json:"extras,omitempty"`
}

func (quick *QuickAppNotice) SetAlert(alert string) *QuickAppNotice {
	quick.Alert = alert
	return quick
}

func (quick *QuickAppNotice) SetTitle(title string) *QuickAppNotice {
	quick.Title = title
	return quick
}

func (quick *QuickAppNotice) SetPage(page string) *QuickAppNotice {
	quick.Page = page
	return quick
}

func (quick *QuickAppNotice) AddExtras(key string, value interface{}) *QuickAppNotice {
	if quick.Extras == nil {
		quick.Extras = make(map[string]interface{})
	}
	quick.Extras[key] = value
	return quick
}

func (quick *QuickAppNotice) ToJson() (string, error) {
	content, err := json.Marshal(quick)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (quick *QuickAppNotice) ToBytes() ([]byte, error) {
	content, err := json.Marshal(quick)
	if err != nil {
		return nil, err
	}
	return content, nil
}
