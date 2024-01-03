package jpushclient

import "encoding/json"

const (
	TAG            = "tag"
	TAGAND         = "tag_and"
	TAGNOT         = "tag_not"
	ALIAS          = "alias"
	ID             = "registration_id"
	SEGMENT        = "segment"
	ABTEST         = "abtest"
	LiveActivityId = "live_activity_id"
)

type Audience struct {
	Object   any
	audience map[string][]string
}

func (a *Audience) All() *Audience {
	a.Object = "all"
	return a
}

func (a *Audience) SetID(ids []string) *Audience {
	a.set(ID, ids)
	return a
}

func (a *Audience) SetTag(tags []string) *Audience {
	a.set(TAG, tags)
	return a
}

func (a *Audience) SetTagAnd(tags []string) *Audience {
	a.set(TAGAND, tags)
	return a
}

func (a *Audience) SetTagNot(tags []string) *Audience {
	a.set(TAGNOT, tags)
	return a
}

func (a *Audience) SetAlias(alias []string) *Audience {
	a.set(ALIAS, alias)
	return a
}

func (a *Audience) SetSegment(segment []string) *Audience {
	a.set(SEGMENT, segment)
	return a
}

func (a *Audience) SetABTest(abtest []string) *Audience {
	a.set(ABTEST, abtest)
	return a
}

func (a *Audience) SetLiveActivityId(liveActivityId []string) *Audience {
	a.set(LiveActivityId, liveActivityId)
	return a
}

func (a *Audience) set(key string, v []string) {
	if a.audience == nil {
		a.audience = make(map[string][]string)
		a.Object = a.audience
	}
	a.audience[key] = v
}

func (a *Audience) ToJson() (string, error) {
	content, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (a *Audience) ToBytes() ([]byte, error) {
	content, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return content, nil
}
