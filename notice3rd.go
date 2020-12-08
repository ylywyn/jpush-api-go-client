package jpushclient

import "encoding/json"

type Notice3rd struct {
	Title       string                 `json:"title,omitempty"`
	Content     string                 `json:"content"`
	ChannelId   string                 `json:"channel_id,omitempty"`
	UriActivity string                 `json:"uri_activity,omitempty"`
	UriAction   string                 `json:"uri_action,omitempty"`
	BadgeAddNum string                 `json:"badge_add_num,omitempty"`
	BadgeClass  string                 `json:"badge_class,omitempty"`
	Sound       string                 `json:"sound,omitempty"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
}

func (this *Notice3rd) SetContent(c string) *Notice3rd {
	this.Content = c
	return this
}

func (this *Notice3rd) SetTitle(title string) *Notice3rd {
	this.Title = title
	return this
}

func (this *Notice3rd) SetChannelId(id string) *Notice3rd {
	this.ChannelId = id
	return this
}

func (this *Notice3rd) SetUriActivity(uri string) *Notice3rd {
	this.UriActivity = uri
	return this
}

func (this *Notice3rd) SetUriAction(uri string) *Notice3rd {
	this.UriAction = uri
	return this
}

func (this *Notice3rd) SetBadgeAddNum(b string) *Notice3rd {
	this.BadgeAddNum = b
	return this
}

func (this *Notice3rd) SetBadgeClass(b string) *Notice3rd {
	this.BadgeClass = b
	return this
}

func (this *Notice3rd) SetSound(s string) *Notice3rd {
	this.Sound = s
	return this
}

func (this *Notice3rd) AddExtras(key string, value interface{}) *Notice3rd {
	if this.Extras == nil {
		this.Extras = make(map[string]interface{})
	}
	this.Extras[key] = value
	return this
}

func (this *Notice3rd) ToJson() (string, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (this *Notice3rd) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return content, nil
}
