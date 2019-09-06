package jpushclient

import (
	"errors"
)

const (
	IOS      = "ios"
	ANDROID  = "android"
	WINPHONE = "winphone"
)

type Platform struct {
	Os     interface{}
	osArry []string
}

func (this *Platform) All() {
	this.Os = "all"
}

func (this *Platform) Add(os string) error {
	if this.Os == nil {
		this.osArry = make([]string, 0, 3)
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
	case WINPHONE:
		this.osArry = append(this.osArry, os)
		this.Os = this.osArry
	default:
		return errors.New("unknow platform")
	}

	return nil
}

func (this *Platform) AddIOS() {
	this.Add(IOS)
}

func (this *Platform) AddAndrid() {
	this.Add(ANDROID)
}

func (this *Platform) AddWinphone() {
	this.Add(WINPHONE)
}
