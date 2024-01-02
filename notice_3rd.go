package jpushclient

import "encoding/json"

type Notice3rd struct {
	Title       string                 `json:"title,omitempty"`
	Content     string                 `json:"content"`
	Intent      map[string]interface{} `json:"intent,omitempty"`
	UriActivity string                 `json:"uri_activity,omitempty"`
	UriAction   string                 `json:"uri_action,omitempty"`
	BadgeAddNum int                    `json:"badge_add_num,omitempty"`
	BadgeSetNum int                    `json:"badge_set_num,omitempty"`
	BadgeClass  string                 `json:"badge_class,omitempty"`
	Sound       string                 `json:"sound,omitempty"`
	ChannelId   string                 `json:"channel_id,omitempty"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
}

func (notice *Notice3rd) SetContent(c string) *Notice3rd {
	notice.Content = c
	return notice
}

func (notice *Notice3rd) SetTitle(title string) *Notice3rd {
	notice.Title = title
	return notice
}

func (notice *Notice3rd) SetChannelId(id string) *Notice3rd {
	notice.ChannelId = id
	return notice
}

func (notice *Notice3rd) SetUriActivity(uri string) *Notice3rd {
	notice.UriActivity = uri
	return notice
}

func (notice *Notice3rd) SetUriAction(uri string) *Notice3rd {
	notice.UriAction = uri
	return notice
}

func (notice *Notice3rd) SetBadgeAddNum(b int) *Notice3rd {
	notice.BadgeAddNum = b
	return notice
}

func (notice *Notice3rd) SetBadgeSetNum(b int) *Notice3rd {
	notice.BadgeSetNum = b
	return notice
}

func (notice *Notice3rd) SetBadgeClass(b string) *Notice3rd {
	notice.BadgeClass = b
	return notice
}

func (notice *Notice3rd) SetSound(s string) *Notice3rd {
	notice.Sound = s
	return notice
}

func (notice *Notice3rd) AddIntent(key string, value interface{}) *Notice3rd {
	if notice.Intent == nil {
		notice.Intent = make(map[string]interface{})
	}
	notice.Intent[key] = value
	return notice
}

func (notice *Notice3rd) AddExtras(key string, value interface{}) *Notice3rd {
	if notice.Extras == nil {
		notice.Extras = make(map[string]interface{})
	}
	notice.Extras[key] = value
	return notice
}

func (notice *Notice3rd) ToJson() (string, error) {
	content, err := json.Marshal(notice)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (notice *Notice3rd) ToBytes() ([]byte, error) {
	content, err := json.Marshal(notice)
	if err != nil {
		return nil, err
	}
	return content, nil
}
