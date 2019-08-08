package jpushclient

const (
	TAG     = "tag"
	TAG_AND = "tag_and"
	TAG_NOT = "tag_not"
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
	this.set(TAG_AND, tags)
}

func (this *Audience) SetTagNot(tags []string) {
	this.set(TAG_NOT, tags)
}

func (this *Audience) SetAlias(alias []string) {
	this.set(ALIAS, alias)
}

func (this *Audience) SetSegment(segIDs []string) {
	this.set(SEGMENT, segIDs)
}

func (this *Audience) SetABTest(testIDs []string) {
	this.set(ABTEST, testIDs)
}

func (this *Audience) set(key string, v []string) {
	if this.Object == nil {
		this.audience = map[string][]string{key: v}
		this.Object = this.audience
	}

	//v, ok = this.audience[key]
	//if ok {
	//	return
	//}
	this.audience[key] = v
}
