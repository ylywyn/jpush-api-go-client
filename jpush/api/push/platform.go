package push

const (
	IOS      = "ios"
	ANDROID  = "android"
	WINPHONE = "winphone"
)

type Platform struct {
	Object interface{}
}

func (this *Platform) All() {
	this.Object = "all"
}

func (this *Platform) Add(os string) {

	if this.Object == nil {
		s := []string{os}
		this.Object = s
	} else {
		switch this.Object.(type) {
		case []string:
			this.Object = append(this.Object.([]string), os)
		default:
		}

	}
}
