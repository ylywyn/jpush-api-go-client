package jpushclient

import (
	"io/ioutil"
	"net/http"
)

const (
	HOST_DEVICE = "https://device.jpush.cn"
)

type DeviceClient struct {
	MasterSecret string
	AppKey       string
	AuthCode     string
	BaseUrl      string
}

func NewDeviceClient(secret, appKey string) *DeviceClient {
	//base64
	auth := "Basic " + base64Coder.EncodeToString([]byte(appKey+":"+secret))
	devicer := &DeviceClient{secret, appKey, auth, HOST_DEVICE}
	return devicer
}

func (this *DeviceClient) DeleteAlias(alias string) (string, error) {
	path := "/v3/aliases/" + alias
	url := this.BaseUrl + path
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Add("Authorization", this.AuthCode)
	req.Header.Add("Accept", CONTENT_TYPE_JSON)
	resp, err := client.Do(req)

	if err != nil {
		if resp != nil {
			resp.Body.Close()
		}
		return "", err
	}
	if resp == nil {
		return "", nil
	}

	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(r), nil
}
