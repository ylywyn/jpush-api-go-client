package jpushclient

import (
	"encoding/json"
)

type TagPayLoad struct {
	Tags   interface{} `json:"tags"`
	Alias  interface{} `json:"alias,omitempty"`
	Mobile interface{} `json:"mobile,omitempty"`
}

func NewPushTagPayLoad() *TagPayLoad {
	pl := &TagPayLoad{}
	return pl
}

func (this *TagPayLoad) SetTags(v map[string][]string) {
	this.Tags = v
}

func (this *TagPayLoad) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return content, nil
}
