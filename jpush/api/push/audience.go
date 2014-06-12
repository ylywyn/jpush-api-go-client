package push

const (
	TAG     = "tag"
	TAG_AND = "tag_and"
	ALIAS   = "alias"
	ID      = "registration_id"
)

type Audience struct {
	Object interface{}
}

func (this *Audience) All() {
	this.Object = "all"
}

func (this *Audience) SetID(Object []string) {
	this.set(ID, Object)
}

func (this *Audience) SetTag(Object []string) {
	this.set(TAG, Object)
}

func (this *Audience) SetTagAnd(Object []string) {
	this.set(TAG_AND, Object)
}

func (this *Audience) SetAlias(Object []string) {
	this.set(ALIAS, Object)
}

func (this *Audience) set(key string, Object []string) {
	if this.Object == nil {
		s := map[string][]string{key: Object}
		this.Object = s
	} else {
		switch this.Object.(type) {
		case map[string][]string:
			this.Object.(map[string][]string)[key] = Object
		default:
		}
	}
}
