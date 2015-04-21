package push

type Notice struct {
	Alert string `json:"alert"`
}

type AndroidNotice struct {
	Object NoticeAndroid `json:"android"`
}

type NoticeAndroid struct {
	Alert     string                 `json:"alert"`
	Title     string                 `json:"title"`
	BuilderId int                    `json:"builder_id"`
	Extras    map[string]interface{} `json:"extras"`
}

func (this *AndroidNotice) SetAlert(alert string) {
	this.Object.Alert = alert

}

func (this *AndroidNotice) SetTitle(title string) {
	this.Object.Title = title
}

func (this *AndroidNotice) SetBuilderId(id int) {
	this.Object.BuilderId = id
}

func (this *AndroidNotice) AddExtras(key string, value interface{}) {
	if this.Object.Extras == nil {
		this.Object.Extras = make(map[string]interface{})
	}
	this.Object.Extras[key] = value
}
