package jpushclient

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

func (this *Notice3rd) SetContent(c string) {
	this.Content = c
}

func (this *Notice3rd) SetTitle(title string) {
	this.Title = title
}

func (this *Notice3rd) SetChannelId(id string) {
	this.ChannelId = id
}

func (this *Notice3rd) SetUriActivity(uri string) {
	this.UriActivity = uri
}

func (this *Notice3rd) SetUriAction(uri string) {
	this.UriAction = uri
}

func (this *Notice3rd) SetBadgeAddNum(b string) {
	this.BadgeAddNum = b
}

func (this *Notice3rd) SetBadgeClass(b string) {
	this.BadgeClass = b
}

func (this *Notice3rd) SetSound(s string) {
	this.Sound = s
}
func (this *Notice3rd) AddExtras(key string, value interface{}) {
	if this.Extras == nil {
		this.Extras = make(map[string]interface{})
	}
	this.Extras[key] = value
}
