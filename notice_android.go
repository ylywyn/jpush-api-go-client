package jpushclient

import "encoding/json"

type AndroidNotice struct {
	Alert             string                 `json:"alert"`
	Title             string                 `json:"title,omitempty"`
	BuilderId         int                    `json:"builder_id,int,omitempty"`
	ChannelId         int                    `json:"channel_id,int,omitempty"`
	Category          string                 `json:"category,omitempty"`
	Priority          int                    `json:"priority,omitempty"`
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
	BadgeSetNum       int                    `json:"badge_set_num,int,omitempty"`
	BadgeClass        string                 `json:"badge_class,omitempty"`
	Sound             string                 `json:"sound,omitempty"`
	ShowBeginTime     string                 `json:"show_begin_time,omitempty"`
	ShowEndTime       string                 `json:"show_end_time,omitempty"`
	DisplayForeground string                 `json:"display_foreground,omitempty"`
}

func (a *AndroidNotice) SetAlert(alert string) *AndroidNotice {
	a.Alert = alert
	return a
}

func (a *AndroidNotice) SetTitle(title string) *AndroidNotice {
	a.Title = title
	return a
}

func (a *AndroidNotice) SetBuilderId(b int) *AndroidNotice {
	a.BuilderId = b
	return a
}

func (a *AndroidNotice) SetChannelId(b int) *AndroidNotice {
	a.ChannelId = b
	return a
}

func (a *AndroidNotice) SetCategory(category string) *AndroidNotice {
	a.Category = category
	return a
}

func (a *AndroidNotice) SetPriority(b int) *AndroidNotice {
	a.Priority = b
	return a
}

func (a *AndroidNotice) SetStyle(b int) *AndroidNotice {
	a.Style = b
	return a
}

func (a *AndroidNotice) SetAlertType(b int) *AndroidNotice {
	a.AlertType = b
	return a
}

func (a *AndroidNotice) SetBigText(bigText string) *AndroidNotice {
	a.BigText = bigText
	return a
}

func (a *AndroidNotice) AddInbox(key string, value interface{}) *AndroidNotice {
	if a.Inbox == nil {
		a.Inbox = make(map[string]interface{})
	}
	a.Inbox[key] = value
	return a
}

func (a *AndroidNotice) SetBigPicPath(bigPicPath string) *AndroidNotice {
	a.BigPicPath = bigPicPath
	return a
}

func (a *AndroidNotice) AddExtras(key string, value interface{}) *AndroidNotice {
	if a.Extras == nil {
		a.Extras = make(map[string]interface{})
	}
	a.Extras[key] = value
	return a
}

func (a *AndroidNotice) SetLargeIcon(largeIcon string) *AndroidNotice {
	a.LargeIcon = largeIcon
	return a
}

func (a *AndroidNotice) SetSmallIconUri(smallIconUri string) *AndroidNotice {
	a.SmallIconUri = smallIconUri
	return a
}

func (a *AndroidNotice) AddIntent(key string, value interface{}) *AndroidNotice {
	if a.Intent == nil {
		a.Intent = make(map[string]interface{})
	}
	a.Intent[key] = value
	return a
}

func (a *AndroidNotice) SetUriActivity(uriActivity string) *AndroidNotice {
	a.UriActivity = uriActivity
	return a
}

func (a *AndroidNotice) SetUriAction(uriAction string) *AndroidNotice {
	a.UriAction = uriAction
	return a
}

func (a *AndroidNotice) SetBadgeAddNum(b int) *AndroidNotice {
	a.BadgeAddNum = b
	return a
}

func (a *AndroidNotice) SetBadgeSetNum(b int) *AndroidNotice {
	a.BadgeSetNum = b
	return a
}

func (a *AndroidNotice) SetBadgeClass(badgeClass string) *AndroidNotice {
	a.BadgeClass = badgeClass
	return a
}

func (a *AndroidNotice) SeSound(sound string) *AndroidNotice {
	a.Sound = sound
	return a
}

func (a *AndroidNotice) SetShowBeginTime(showBeginTime string) *AndroidNotice {
	a.ShowBeginTime = showBeginTime
	return a
}

func (a *AndroidNotice) SetShowEndTime(showEndTime string) *AndroidNotice {
	a.ShowEndTime = showEndTime
	return a
}

func (a *AndroidNotice) SetDisplayForeground(displayForeground string) *AndroidNotice {
	a.DisplayForeground = displayForeground
	return a
}

func (a *AndroidNotice) ToJson() (string, error) {
	content, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (a *AndroidNotice) ToBytes() ([]byte, error) {
	content, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return content, nil
}
