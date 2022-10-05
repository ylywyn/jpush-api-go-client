package jpushclient

import (
	"encoding/json"
	"errors"
)

const (
	IOS      = "ios"
	ANDROID  = "android"
	WINPHONE = "winphone"
	QUICKAPP = "quickapp"
)

type Platform struct {
	Os     interface{}
	osArry []string
}

func (this *Platform) All() *Platform {
	this.Os = "all"
	return this
}

func (this *Platform) add(os string) error {
	if this.Os == nil {
		this.osArry = make([]string, 0, 4)
	} else {
		switch this.Os.(type) {
		case string:
			return errors.New("platform is all")
		default:
		}
	}

	//判断是否重复
	for _, value := range this.osArry {
		if os == value {
			return nil
		}
	}

	switch os {
	case IOS:
		fallthrough
	case ANDROID:
		fallthrough
	case QUICKAPP:
		fallthrough
	case WINPHONE:
		this.osArry = append(this.osArry, os)
		this.Os = this.osArry
	default:
		return errors.New("unknow platform")
	}

	return nil
}

func (this *Platform) AddIOS() *Platform {
	this.add(IOS)
	return this
}

func (this *Platform) AddAndroid() *Platform {
	this.add(ANDROID)
	return this
}

func (this *Platform) AddWinPhone() *Platform {
	this.add(WINPHONE)
	return this
}

func (this *Platform) AddQuickApp() *Platform {
	this.add(QUICKAPP)
	return this
}

func (this *Platform) ToJson() (string, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (this *Platform) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return content, nil
}
