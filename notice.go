package jpushclient

import "encoding/json"

type Notice struct {
	Alert    string          `json:"alert"`              // 通知的内容在各个平台上，都可能只有这一个最基本的属性 "alert"。
	Android  *AndroidNotice  `json:"android,omitempty"`  // Android通知
	IOS      *IOSNotice      `json:"ios,omitempty"`      // iOS通知
	VOIP     *VOIPNotice     `json:"voip,omitempty"`     //iOS VOIP功能。
	QuickApp *QuickAppNotice `json:"quickapp,omitempty"` //快应用通知
}

func (n2 *Notice) SetAlert(alert string) *Notice {
	n2.Alert = alert
	return n2
}

func (n2 *Notice) SetAndroidNotice(n *AndroidNotice) *Notice {
	n2.Android = n
	return n2
}

func (n2 *Notice) SetIOSNotice(n *IOSNotice) *Notice {
	n2.IOS = n
	return n2
}

func (n2 *Notice) SetVOIPNotice(n *VOIPNotice) *Notice {
	n2.VOIP = n
	return n2
}

func (n2 *Notice) SetQuickAppNotice(n *QuickAppNotice) *Notice {
	n2.QuickApp = n
	return n2
}

func (n2 *Notice) ToJson() (string, error) {
	content, err := json.Marshal(n2)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (n2 *Notice) ToBytes() ([]byte, error) {
	content, err := json.Marshal(n2)
	if err != nil {
		return nil, err
	}
	return content, nil
}

type AndroidNotice struct {
	Alert             string         `json:"alert"`                        // 通知内容
	Title             string         `json:"title,omitempty"`              // 通知标题
	BuilderId         int            `json:"builder_id,omitempty"`         // 通知栏样式 ID
	ChannelId         int            `json:"channel_id,omitempty"`         // android通知channel_id
	Category          string         `json:"category,omitempty"`           // 通知栏条目过滤或排序
	Priority          int            `json:"priority,omitempty"`           // 通知栏展示优先级, 默认为 0，范围为 -2～2。
	Style             int            `json:"style,omitempty"`              // 通知栏样式类型, 默认为 0，还有 1，2，3 可选
	AlertType         int            `json:"alert_type,omitempty"`         // 通知提醒方式, 可选范围为 -1～7
	BigText           string         `json:"big_text,omitempty"`           // 大文本通知栏样式, 当 style = 1 时可用，内容会被通知栏以大文本的形式展示出来
	Inbox             map[string]any `json:"inbox,omitempty"`              // 文本条目通知栏样式, 当 style = 2 时可用， json 的每个 key 对应的 value 会被当作文本条目逐条展示
	BigPicPath        string         `json:"big_pic_path,omitempty"`       // 大图片通知栏样式, 当 style = 3 时可用，可以是网络图片 url，或本地图片的 path
	Extras            map[string]any `json:"extras,omitempty"`             // 扩展字段
	LargeIcon         string         `json:"large_icon,omitempty"`         // 通知栏大图标
	SmallIconUri      string         `json:"small_icon_uri,omitempty"`     // 通知栏小图标
	Intent            map[string]any `json:"intent,omitempty"`             // 指定跳转页面
	UriActivity       string         `json:"uri_activity,omitempty"`       // 指定跳转页面, 该字段用于指定开发者想要打开的 activity，值为 activity 节点的 “android:name”属性值; 适配华为、小米、vivo厂商通道跳转；
	UriAction         string         `json:"uri_action,omitempty"`         // 指定跳转页面, 该字段用于指定开发者想要打开的 activity，值为 "activity"-"intent-filter"-"action" 节点的 "android:name" 属性值; 适配 oppo、fcm跳转；
	BadgeAddNum       int            `json:"badge_add_num,omitempty"`      // 角标数字，取值范围1-99
	BadgeSetNum       int            `json:"badge_set_num,omitempty"`      // 设置角标数字累加值，在原角标的基础上进行累加
	BadgeClass        string         `json:"badge_class,omitempty"`        // 桌面图标对应的应用入口Activity类， 配合badge_add_num使用，二者需要共存，缺少其一不可；
	Sound             string         `json:"sound,omitempty"`              // 填写Android工程中/res/raw/路径下铃声文件名称，无需文件名后缀
	ShowBeginTime     string         `json:"show_begin_time,omitempty"`    // 定时展示开始时间（yyyy-MM-dd HH:mm:ss）
	ShowEndTime       string         `json:"show_end_time,omitempty"`      // 定时展示结束时间（yyyy-MM-dd HH:mm:ss）
	DisplayForeground string         `json:"display_foreground,omitempty"` // APP在前台，通知是否展示, 值为 "1" 时，APP 在前台会弹出通知栏消息；值为 "0" 时，APP 在前台不会弹出通知栏消息。
}

func (a *AndroidNotice) SetAlert(alert string) *AndroidNotice {
	a.Alert = alert
	return a
}

func (a *AndroidNotice) SetTitle(title string) *AndroidNotice {
	a.Title = title
	return a
}

func (a *AndroidNotice) SetBuilderId(b int) *AndroidNotice {
	a.BuilderId = b
	return a
}

func (a *AndroidNotice) SetChannelId(b int) *AndroidNotice {
	a.ChannelId = b
	return a
}

func (a *AndroidNotice) SetCategory(category string) *AndroidNotice {
	a.Category = category
	return a
}

func (a *AndroidNotice) SetPriority(b int) *AndroidNotice {
	a.Priority = b
	return a
}

func (a *AndroidNotice) SetStyle(b int) *AndroidNotice {
	a.Style = b
	return a
}

func (a *AndroidNotice) SetAlertType(b int) *AndroidNotice {
	a.AlertType = b
	return a
}

func (a *AndroidNotice) SetBigText(bigText string) *AndroidNotice {
	a.BigText = bigText
	return a
}

func (a *AndroidNotice) AddInbox(key string, value any) *AndroidNotice {
	if a.Inbox == nil {
		a.Inbox = make(map[string]any)
	}
	a.Inbox[key] = value
	return a
}

func (a *AndroidNotice) SetBigPicPath(bigPicPath string) *AndroidNotice {
	a.BigPicPath = bigPicPath
	return a
}

func (a *AndroidNotice) AddExtras(key string, value any) *AndroidNotice {
	if a.Extras == nil {
		a.Extras = make(map[string]any)
	}
	a.Extras[key] = value
	return a
}

func (a *AndroidNotice) SetLargeIcon(largeIcon string) *AndroidNotice {
	a.LargeIcon = largeIcon
	return a
}

func (a *AndroidNotice) SetSmallIconUri(smallIconUri string) *AndroidNotice {
	a.SmallIconUri = smallIconUri
	return a
}

func (a *AndroidNotice) AddIntent(key string, value any) *AndroidNotice {
	if a.Intent == nil {
		a.Intent = make(map[string]any)
	}
	a.Intent[key] = value
	return a
}

func (a *AndroidNotice) SetUriActivity(uriActivity string) *AndroidNotice {
	a.UriActivity = uriActivity
	return a
}

func (a *AndroidNotice) SetUriAction(uriAction string) *AndroidNotice {
	a.UriAction = uriAction
	return a
}

func (a *AndroidNotice) SetBadgeAddNum(b int) *AndroidNotice {
	a.BadgeAddNum = b
	return a
}

func (a *AndroidNotice) SetBadgeSetNum(b int) *AndroidNotice {
	a.BadgeSetNum = b
	return a
}

func (a *AndroidNotice) SetBadgeClass(badgeClass string) *AndroidNotice {
	a.BadgeClass = badgeClass
	return a
}

func (a *AndroidNotice) SeSound(sound string) *AndroidNotice {
	a.Sound = sound
	return a
}

func (a *AndroidNotice) SetShowBeginTime(showBeginTime string) *AndroidNotice {
	a.ShowBeginTime = showBeginTime
	return a
}

func (a *AndroidNotice) SetShowEndTime(showEndTime string) *AndroidNotice {
	a.ShowEndTime = showEndTime
	return a
}

func (a *AndroidNotice) SetDisplayForeground(displayForeground string) *AndroidNotice {
	a.DisplayForeground = displayForeground
	return a
}

func (a *AndroidNotice) ToJson() (string, error) {
	content, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (a *AndroidNotice) ToBytes() ([]byte, error) {
	content, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return content, nil
}

type IOSNotice struct {
	Alert             any            `json:"alert"`                        // 通知内容
	Sound             any            `json:"sound,omitempty"`              // 通知提示声音或警告通知
	Badge             string         `json:"badge,omitempty"`              // 应用角标, 如果不填，表示不改变角标数字，否则把角标数字改为指定的数字；为 0 表示清除。
	ContentAvailable  bool           `json:"content-available,omitempty"`  // 推送唤醒
	MutableContent    bool           `json:"mutable-content,omitempty"`    // 通知扩展
	Category          string         `json:"category,omitempty"`           // 通知类别, IOS 8 才支持。设置 APNs payload 中的 "category" 字段值
	Extras            map[string]any `json:"extras,omitempty"`             // 扩展字段
	ThreadId          string         `json:"thread-id,omitempty"`          // 通知分组, ios 的远程通知通过该属性来对通知进行分组，同一个 thread-id 的通知归为一组。
	InterruptionLevel string         `json:"interruption-level,omitempty"` // 通知优先级和交付时间的中断级别, ios15 的通知级别，取值只能是active,critical,passive,timeSensitive中的一个。
}

func (ios *IOSNotice) SetAlert(alert any) *IOSNotice {
	ios.Alert = alert
	return ios
}

func (ios *IOSNotice) SetSound(n any) *IOSNotice {
	ios.Sound = n
	return ios
}

func (ios *IOSNotice) SetBadge(n string) *IOSNotice {
	ios.Badge = n
	return ios
}

func (ios *IOSNotice) SetContentAvailable(n bool) *IOSNotice {
	ios.ContentAvailable = n
	return ios
}

func (ios *IOSNotice) SetMutableContent(n bool) *IOSNotice {
	ios.MutableContent = n
	return ios
}

func (ios *IOSNotice) SetCategory(n string) *IOSNotice {
	ios.Category = n
	return ios
}

func (ios *IOSNotice) SetThreadId(n string) *IOSNotice {
	ios.ThreadId = n
	return ios
}

func (ios *IOSNotice) SetInterruptionLevel(n string) *IOSNotice {
	ios.InterruptionLevel = n
	return ios
}

func (ios *IOSNotice) AddExtras(key string, value any) *IOSNotice {
	if ios.Extras == nil {
		ios.Extras = make(map[string]any)
	}
	ios.Extras[key] = value
	return ios
}

func (ios *IOSNotice) ToJson() (string, error) {
	content, err := json.Marshal(ios)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (ios *IOSNotice) ToBytes() ([]byte, error) {
	content, err := json.Marshal(ios)
	if err != nil {
		return nil, err
	}
	return content, nil
}

type QuickAppNotice struct {
	Title  string         `json:"title"`            // 通知标题, 必填字段，快应用推送通知的标题
	Alert  string         `json:"alert"`            // 通知内容, 这里指定了，则会覆盖上级统一指定的 alert 信息。
	Page   string         `json:"page"`             // 通知跳转页面, 必填字段，快应用通知跳转地址。
	Extras map[string]any `json:"extras,omitempty"` // 扩展字段, 这里自定义 Key / value 信息，以供业务使用。
}

func (quick *QuickAppNotice) SetAlert(alert string) *QuickAppNotice {
	quick.Alert = alert
	return quick
}

func (quick *QuickAppNotice) SetTitle(title string) *QuickAppNotice {
	quick.Title = title
	return quick
}

func (quick *QuickAppNotice) SetPage(page string) *QuickAppNotice {
	quick.Page = page
	return quick
}

func (quick *QuickAppNotice) AddExtras(key string, value any) *QuickAppNotice {
	if quick.Extras == nil {
		quick.Extras = make(map[string]any)
	}
	quick.Extras[key] = value
	return quick
}

func (quick *QuickAppNotice) ToJson() (string, error) {
	content, err := json.Marshal(quick)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (quick *QuickAppNotice) ToBytes() ([]byte, error) {
	content, err := json.Marshal(quick)
	if err != nil {
		return nil, err
	}
	return content, nil
}
