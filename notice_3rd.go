package jpushclient

import "encoding/json"

type Notice3rd struct {
	Title       string         `json:"title,omitempty"`         // 补发通知标题，如果为空则默认为应用名称
	Content     string         `json:"content"`                 // 补发通知的内容，如果存在 notification_3rd 这个key，content 字段不能为空，且值不能为空字符串。
	Intent      map[string]any `json:"intent,omitempty"`        // 使用 intent 里的 url 指定点击通知栏后跳转的目标页面
	UriActivity string         `json:"uri_activity,omitempty"`  // 该字段用于指定开发者想要打开的 activity，值为 activity 节点的 “android:name”属性值;适配华为、小米、vivo厂商通道跳转；针对 VIP 厂商通道用户使用生效。
	UriAction   string         `json:"uri_action,omitempty"`    // 指定跳转页面；该字段用于指定开发者想要打开的 activity，值为 "activity"-"intent-filter"-"action" 节点的 "android:name" 属性值;适配 oppo、fcm跳转；针对 VIP 厂商通道用户使用生效。
	BadgeAddNum int            `json:"badge_add_num,omitempty"` // 角标数字，取值范围1-99
	BadgeSetNum int            `json:"badge_set_num,omitempty"` // 设置角标数字累加值，在原角标的基础上进行累加
	BadgeClass  string         `json:"badge_class,omitempty"`   // 桌面图标对应的应用入口Activity类， 比如“com.test.badge.MainActivity；
	Sound       string         `json:"sound,omitempty"`         // 填写Android工程中/res/raw/路径下铃声文件名称，无需文件名后缀；注意：针对Android 8.0以上，当传递了channel_id 时，此属性不生效。
	ChannelId   string         `json:"channel_id,omitempty"`    // 根据 channel ID 来指定通知栏展示效果，不超过 1000 字节
	Extras      map[string]any `json:"extras,omitempty"`        // 扩展字段；这里自定义 JSON 格式的 Key / Value 信息，以供业务使用。
}

func (notice *Notice3rd) SetContent(c string) *Notice3rd {
	notice.Content = c
	return notice
}

func (notice *Notice3rd) SetTitle(title string) *Notice3rd {
	notice.Title = title
	return notice
}

func (notice *Notice3rd) SetChannelId(id string) *Notice3rd {
	notice.ChannelId = id
	return notice
}

func (notice *Notice3rd) SetUriActivity(uri string) *Notice3rd {
	notice.UriActivity = uri
	return notice
}

func (notice *Notice3rd) SetUriAction(uri string) *Notice3rd {
	notice.UriAction = uri
	return notice
}

func (notice *Notice3rd) SetBadgeAddNum(b int) *Notice3rd {
	notice.BadgeAddNum = b
	return notice
}

func (notice *Notice3rd) SetBadgeSetNum(b int) *Notice3rd {
	notice.BadgeSetNum = b
	return notice
}

func (notice *Notice3rd) SetBadgeClass(b string) *Notice3rd {
	notice.BadgeClass = b
	return notice
}

func (notice *Notice3rd) SetSound(s string) *Notice3rd {
	notice.Sound = s
	return notice
}

func (notice *Notice3rd) AddIntent(key string, value any) *Notice3rd {
	if notice.Intent == nil {
		notice.Intent = make(map[string]any)
	}
	notice.Intent[key] = value
	return notice
}

func (notice *Notice3rd) AddExtras(key string, value any) *Notice3rd {
	if notice.Extras == nil {
		notice.Extras = make(map[string]any)
	}
	notice.Extras[key] = value
	return notice
}

func (notice *Notice3rd) ToJson() (string, error) {
	content, err := json.Marshal(notice)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (notice *Notice3rd) ToBytes() ([]byte, error) {
	content, err := json.Marshal(notice)
	if err != nil {
		return nil, err
	}
	return content, nil
}
