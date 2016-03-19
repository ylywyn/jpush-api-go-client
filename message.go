package jpushclient

type Message struct {
	Content     string                 `json:"msg_content"`
	Title       string                 `json:"title,omitempty"`
	ContentType string                 `json:"content_type,omitempty"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
}

func (this *Message) SetContent(c string) {
	this.Content = c

}

func (this *Message) SetTitle(title string) {
	this.Title = title
}

func (this *Message) SetContentType(t string) {
	this.ContentType = t
}

func (this *Message) AddExtras(key string, value interface{}) {
	if this.Extras == nil {
		this.Extras = make(map[string]interface{})
	}
	this.Extras[key] = value
}
