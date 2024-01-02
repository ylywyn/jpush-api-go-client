package jpushclient

import "encoding/json"

type IOSNotice struct {
	Alert             interface{}            `json:"alert"`
	Sound             string                 `json:"sound,omitempty"`
	Badge             string                 `json:"badge,omitempty"`
	ContentAvailable  bool                   `json:"content-available,omitempty"`
	MutableContent    bool                   `json:"mutable-content,omitempty"`
	Category          string                 `json:"category,omitempty"`
	Extras            map[string]interface{} `json:"extras,omitempty"`
	ThreadId          string                 `json:"thread-id,omitempty"`
	InterruptionLevel string                 `json:"interruption-level,omitempty"`
}

func (ios *IOSNotice) SetAlert(alert interface{}) *IOSNotice {
	ios.Alert = alert
	return ios
}

func (ios *IOSNotice) SetSound(n string) *IOSNotice {
	ios.Sound = n
	return ios
}

func (ios *IOSNotice) SetBadge(n string) *IOSNotice {
	ios.Badge = n
	return ios
}

func (ios *IOSNotice) SetContentAvailable(n bool) *IOSNotice {
	ios.ContentAvailable = n
	return ios
}

func (ios *IOSNotice) SetMutableContent(n bool) *IOSNotice {
	ios.MutableContent = n
	return ios
}

func (ios *IOSNotice) SetCategory(n string) *IOSNotice {
	ios.Category = n
	return ios
}

func (ios *IOSNotice) SetThreadId(n string) *IOSNotice {
	ios.ThreadId = n
	return ios
}

func (ios *IOSNotice) SetInterruptionLevel(n string) *IOSNotice {
	ios.InterruptionLevel = n
	return ios
}

func (ios *IOSNotice) AddExtras(key string, value interface{}) *IOSNotice {
	if ios.Extras == nil {
		ios.Extras = make(map[string]interface{})
	}
	ios.Extras[key] = value
	return ios
}

func (ios *IOSNotice) ToJson() (string, error) {
	content, err := json.Marshal(ios)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (ios *IOSNotice) ToBytes() ([]byte, error) {
	content, err := json.Marshal(ios)
	if err != nil {
		return nil, err
	}
	return content, nil
}
