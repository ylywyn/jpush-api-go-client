package jpushclient

import "encoding/json"

type Notice struct {
	Alert    string          `json:"alert,omitempty"`
	Android  *AndroidNotice  `json:"android,omitempty"`
	IOS      *IOSNotice      `json:"ios,omitempty"`
	VIOS     *VIOSNotice     `json:"vios,omitempty"`
	WINPhone *WinPhoneNotice `json:"winphone,omitempty"`
	QuickApp *QuickAppNotice `json:"quickapp,omitempty"`
}

type AndroidNotice struct {
	Alert             string                 `json:"alert"`
	Title             string                 `json:"title,omitempty"`
	BuilderId         int                    `json:"builder_id,int,omitempty"`
	ChannelId         int                    `json:"channel_id,int,omitempty"`
	Priority          int                    `json:"priority,omitempty"`
	Category          string                 `json:"category,omitempty"`
	Style             int                    `json:"style,int,omitempty"`
	AlertType         int                    `json:"alert_type,int,omitempty"`
	BigText           string                 `json:"big_text,omitempty"`
	Inbox             map[string]interface{} `json:"inbox,omitempty"`
	BigPicPath        string                 `json:"big_pic_path,omitempty"`
	Extras            map[string]interface{} `json:"extras,omitempty"`
	LargeIcon         string                 `json:"large_icon,omitempty"`
	SmallIconUri      string                 `json:"small_icon_uri,omitempty"`
	Intent            map[string]interface{} `json:"intent,omitempty"`
	UriActivity       string                 `json:"uri_activity,omitempty"`
	UriAction         string                 `json:"uri_action,omitempty"`
	BadgeAddNum       int                    `json:"badge_add_num,int,omitempty"`
	BadgeClass        string                 `json:"badge_class,omitempty"`
	Sound             string                 `json:"sound,omitempty"`
	ShowBeginTime     string                 `json:"show_begin_time,omitempty"`
	ShowEndTime       string                 `json:"show_end_time,omitempty"`
	DisplayForeground string                 `json:"display_foreground,omitempty"`
}

type IOSNotice struct {
	Alert            interface{}            `json:"alert"`
	Sound            string                 `json:"sound,omitempty"`
	Badge            string                 `json:"badge,omitempty"`
	ContentAvailable bool                   `json:"content-available,omitempty"`
	MutableContent   bool                   `json:"mutable-content,omitempty"`
	Category         string                 `json:"category,omitempty"`
	Extras           map[string]interface{} `json:"extras,omitempty"`
	ThreadId         string                 `json:"thread-id,omitempty"`
}

type VIOSNotice struct {
	Key string `json:"key,omitempty"`
}

type WinPhoneNotice struct {
	Alert    string                 `json:"alert"`
	Title    string                 `json:"title,omitempty"`
	OpenPage string                 `json:"_open_page,omitempty"`
	Extras   map[string]interface{} `json:"extras,omitempty"`
}

type QuickAppNotice struct {
	Title  string                 `json:"title"`
	Alert  string                 `json:"alert"`
	Page   string                 `json:"page"`
	Extras map[string]interface{} `json:"extras,omitempty"`
}

func (this *Notice) SetAlert(alert string) *Notice {
	this.Alert = alert
	return this
}

func (this *Notice) SetAndroidNotice(n *AndroidNotice) *Notice {
	this.Android = n
	return this
}

func (this *Notice) SetIOSNotice(n *IOSNotice) *Notice {
	this.IOS = n
	return this
}

func (this *Notice) SetVIOSNotice(n *VIOSNotice) *Notice {
	this.VIOS = n
	return this
}

func (this *Notice) SetWinPhoneNotice(n *WinPhoneNotice) *Notice {
	this.WINPhone = n
	return this
}

func (this *Notice) SetQuickAppNotice(n *QuickAppNotice) *Notice {
	this.QuickApp = n
	return this
}

func (this *Notice) ToJson() (string, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (this *Notice) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return content, nil
}
