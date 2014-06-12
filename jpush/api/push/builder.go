package push

//Base Builder
type Builder struct {
	Platform interface{} `json:"platform"`
	Audience interface{} `json:"audience"`
	Options  *Option     `json:"options"`
}

//MessageBuilder
type MessageBuilder struct {
	Builder
	Message interface{} `json:"message"`
}

//NoticeBuilder
type NoticeBuilder struct {
	Builder
	Notification interface{} `json:"notification"`
}

//MessageAndNotice
type MessageAndNoticeBuilder struct {
	Builder
	Notification interface{} `json:"notification"`
	Message      interface{} `json:"message"`
}

//---------------------MessageBuilder --------------------
func NewMessageBuilder() *MessageBuilder {
	mb := &MessageBuilder{}
	o := &Option{1, 86400, true}
	mb.Options = o
	return mb
}

func (this *MessageBuilder) SetMessage(m *Message) {
	this.Message = m
}

func (this *MessageBuilder) SetPlatForm(pf *PlatForm) {
	this.Platform = pf.Object
}

func (this *MessageBuilder) SetAudience(ad *Audience) {
	this.Audience = ad.Object
}

func (this *MessageBuilder) SetOptions(o *Option) {
	this.Options = o
}

//---------------------NoticeBuilder----------------------
func NewNoticeBuilder() *NoticeBuilder {
	nb := &NoticeBuilder{}
	o := &Option{1, 86400, true}
	nb.Options = o
	return nb
}

func (this *NoticeBuilder) SetSimpleNotice(notice string) {
	s := &Notice{notice}
	this.Notification = s
}

func (this *NoticeBuilder) SetAndroidNotice(notice *AndroidNotice) {
	this.Notification = notice
}

func (this *NoticeBuilder) SetPlatForm(pf *PlatForm) {
	this.Platform = pf.Object
}

func (this *NoticeBuilder) SetAudience(ad *Audience) {
	this.Audience = ad.Object
}

func (this *NoticeBuilder) SetOptions(o *Option) {
	this.Options = o
}

//------------------MessageAndNoticeBuilder------------------
func NewMessageAndNoticeBuilder() *MessageAndNoticeBuilder {
	mnb := &MessageAndNoticeBuilder{}
	o := &Option{1, 86400, true}
	mnb.Options = o
	return mnb
}

func (this *MessageAndNoticeBuilder) SetPlatForm(pf *PlatForm) {
	this.Platform = pf.Object
}

func (this *MessageAndNoticeBuilder) SetAudience(ad *Audience) {
	this.Audience = ad.Object
}

func (this *MessageAndNoticeBuilder) SetOptions(o *Option) {
	this.Options = o
}

func (this *MessageAndNoticeBuilder) SetMessage(m *Message) {
	this.Message = m
}

func (this *MessageAndNoticeBuilder) SetSimpleNotice(notice string) {
	s := &Notice{notice}
	this.Notification = s
}

func (this *MessageAndNoticeBuilder) SetAndroidNotice(notice *AndroidNotice) {
	this.Notification = notice
}
