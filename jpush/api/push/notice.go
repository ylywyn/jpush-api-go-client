package push

type Notice struct {
	Alert string `json:"alert"`
}

type AndroidNotice struct {
	Object NoticeAndroid `json:"android"`
}

type NoticeAndroid struct {
	Alert     string `json:"alert"`
	Title     string `json:"title"`
	BuilderId int    `json:"builder_id"`
	//Extras    map[string]string `json:"extras"`
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

//func (this *AndroidNotice) AddExtras(key, value string) {
//	if this.notice.Extras == nil {
//		this.notice.Extras = make(map[string]string)
//	}
//	this.notice.Extras[key] = value
//}
