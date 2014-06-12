package push

type Message struct {
	Content     string `json:"msg_content"`
	Title       string `json:"title"`
	ContentType string `json:"content_type"`
	//Extras      map[string]string `json:"extras"`
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

//func (this *Message) AddExtras(key, value string) {
//	if this.Extras == nil {
//		this.Extras = make(map[string]string)
//	}
//	this.Extras[key] = value
//}
