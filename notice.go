package jpushclient

type Notice struct {
	Alert    string          `json:"alert,omitempty"`
	Android  *AndroidNotice  `json:"android,omitempty"`
	IOS      *IOSNotice      `json:"ios,omitempty"`
	WINPhone *WinPhoneNotice `json:"winphone,omitempty"`
}

type AndroidNotice struct {
	Alert     string                 `json:"alert"`
	Title     string                 `json:"title,omitempty"`
	BuilderId int                    `json:"builder_id,omitempty"`
	Extras    map[string]interface{} `json:"extras,omitempty"`
	Intent    map[string]interface{} `json:"intent,omitempty"`
}

type IOSNotice struct {
	Alert            interface{}            `json:"alert"`
	Sound            string                 `json:"sound,omitempty"`
	Badge            string                 `json:"badge,omitempty"`
	ContentAvailable bool                   `json:"content-available,omitempty"`
	MutableContent   bool                   `json:"mutable-content,omitempty"`
	Category         string                 `json:"category,omitempty"`
	Extras           map[string]interface{} `json:"extras,omitempty"`
}

type WinPhoneNotice struct {
	Alert    string                 `json:"alert"`
	Title    string                 `json:"title,omitempty"`
	OpenPage string                 `json:"_open_page,omitempty"`
	Extras   map[string]interface{} `json:"extras,omitempty"`
}

func (this *Notice) SetAlert(alert string) {
	this.Alert = alert
}

func (this *Notice) SetAndroidNotice(n *AndroidNotice) {
	this.Android = n
}

func (this *Notice) SetIOSNotice(n *IOSNotice) {
	this.IOS = n
}

func (this *Notice) SetWinPhoneNotice(n *WinPhoneNotice) {
	this.WINPhone = n
}
