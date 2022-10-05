package jpushclient

import "encoding/json"

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

func (this *Audience) All() *Audience {
	this.Object = "all"
	return this
}

func (this *Audience) SetID(ids []string) *Audience {
	this.set(ID, ids)
	return this
}

func (this *Audience) SetTag(tags []string) *Audience {
	this.set(TAG, tags)
	return this
}

func (this *Audience) SetTagAnd(tags []string) *Audience {
	this.set(TAGAND, tags)
	return this
}

func (this *Audience) SetTagNot(tags []string) *Audience {
	this.set(TAGNOT, tags)
	return this
}

func (this *Audience) SetAlias(alias []string) *Audience {
	this.set(ALIAS, alias)
	return this
}

func (this *Audience) SetSegment(segment []string) *Audience {
	this.set(SEGMENT, segment)
	return this
}

func (this *Audience) SetABTest(abtest []string) *Audience {
	this.set(ABTEST, abtest)
	return this
}

func (this *Audience) set(key string, v []string) {
	if this.audience == nil {
		this.audience = make(map[string][]string)
		this.Object = this.audience
	}
	this.audience[key] = v
}

func (this *Audience) ToJson() (string, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (this *Audience) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return content, nil
}
