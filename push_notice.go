package jpushclient

import "encoding/json"

type Notice struct {
	Alert    string          `json:"alert"`
	Android  *AndroidNotice  `json:"android,omitempty"`
	IOS      *IOSNotice      `json:"ios,omitempty"`
	VOIP     *VOIPNotice     `json:"voip,omitempty"`
	QuickApp *QuickAppNotice `json:"quickapp,omitempty"`
}

func (n2 *Notice) SetAlert(alert string) *Notice {
	n2.Alert = alert
	return n2
}

func (n2 *Notice) SetAndroidNotice(n *AndroidNotice) *Notice {
	n2.Android = n
	return n2
}

func (n2 *Notice) SetIOSNotice(n *IOSNotice) *Notice {
	n2.IOS = n
	return n2
}

func (n2 *Notice) SetVOIPNotice(n *VOIPNotice) *Notice {
	n2.VOIP = n
	return n2
}

func (n2 *Notice) SetQuickAppNotice(n *QuickAppNotice) *Notice {
	n2.QuickApp = n
	return n2
}

func (n2 *Notice) ToJson() (string, error) {
	content, err := json.Marshal(n2)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (n2 *Notice) ToBytes() ([]byte, error) {
	content, err := json.Marshal(n2)
	if err != nil {
		return nil, err
	}
	return content, nil
}
