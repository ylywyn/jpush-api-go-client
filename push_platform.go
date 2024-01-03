package jpushclient

import (
	"encoding/json"
	"errors"
)

const (
	IOS      = "ios"
	ANDROID  = "android"
	QUICKAPP = "quickapp"
)

type Platform struct {
	Os     any
	osArry []string
}

func (plat *Platform) All() *Platform {
	plat.Os = "all"
	return plat
}

func (plat *Platform) add(os string) error {
	if plat.Os == nil {
		plat.osArry = make([]string, 0, 4)
	} else {
		switch plat.Os.(type) {
		case string:
			return errors.New("platform is all")
		default:
		}
	}

	//判断是否重复
	for _, value := range plat.osArry {
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

	default:
		return errors.New("unknow platform")
	}
}

func (plat *Platform) AddIOS() *Platform {
	err := plat.add(IOS)
	if err != nil {
		return &Platform{}
	}
	return plat
}

func (plat *Platform) AddAndroid() *Platform {
	err := plat.add(ANDROID)
	if err != nil {
		return &Platform{}
	}
	return plat
}

func (plat *Platform) AddQuickApp() *Platform {
	err := plat.add(QUICKAPP)
	if err != nil {
		return &Platform{}
	}
	return plat
}

func (plat *Platform) ToJson() (string, error) {
	content, err := json.Marshal(plat)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (plat *Platform) ToBytes() ([]byte, error) {
	content, err := json.Marshal(plat)
	if err != nil {
		return nil, err
	}
	return content, nil
}
