package jpushclient

const (
	TAG     = "tag"
	TAGAND  = "tag_and"
	TAGNOT  = "tag_not"
	ALIAS   = "alias"
	ID      = "registration_id"
	SEGMENT = "segment"
	ABTEST  = "abtest"
)

type Audience struct {
	Object   interface{}
	audience map[string][]string
}

func (this *Audience) All() {
	this.Object = "all"
}

func (this *Audience) SetID(ids []string) {
	this.set(ID, ids)
}

func (this *Audience) SetTag(tags []string) {
	this.set(TAG, tags)
}

func (this *Audience) SetTagAnd(tags []string) {
	this.set(TAGAND, tags)
}

func (this *Audience) SetTagNot(tags []string) {
	this.set(TAGNOT, tags)
}

func (this *Audience) SetAlias(alias []string) {
	this.set(ALIAS, alias)
}

func (this *Audience) SetSegment(segment []string) {
	this.set(SEGMENT, segment)
}

func (this *Audience) SetABTest(abtest []string) {
	this.set(ABTEST, abtest)
}

func (this *Audience) set(key string, v []string) {
	if this.audience == nil {
		this.audience = make(map[string][]string)
		this.Object = this.audience
	}

	//v, ok = this.audience[key]
	//if ok {
	//	return
	//}
	this.audience[key] = v
}
