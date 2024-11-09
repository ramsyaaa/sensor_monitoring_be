package models

func (Sensor) TableName() string {
	return "sensors"
}

type Sensor struct {
	AccessToken   string `json:"accessToken"`
	AppKey        string `json:"appkey"`
	Bias          string `json:"bias"`
	DecimalPlaces string `json:"decimalPlacse"`
	DeviceId      int    `json:"deviceId"`
	DeviceName    string `json:"deviceName"`
	FailTime      string `json:"fialtime"`
	Flag          string `json:"flag"`
	HeartbeatDate string `json:"heartbeatDate"`
	Hls           string `json:"hls"`
	Id            int64  `json:"id"`
	IocUrl        string `json:"iocUrl"`
	IsAlarms      int    `json:"isAlarms"`
	IsDelete      int    `json:"isDelete"`
	IsLine        int    `json:"isLine"`
	IsMapping     int    `json:"isMapping"`
	Lat           string `json:"lat"`
	Live          string `json:"live"`
	Lng           string `json:"lng"`
	OpenysId      int    `json:"openysId"`
	OrderNum      int    `json:"ordernum"`
	Replay        string `json:"replay"`
	Rtmp          string `json:"rtmp"`
	Secret        string `json:"secret"`
	SendDataType  string `json:"send_data_type"`
	SendValue     string `json:"send_value"`
	SensorName    string `json:"sensorName"`
	SensorTypeId  int    `json:"sensorTypeId"`
	SwitchOff     string `json:"switch_off"`
	SwitchOn      string `json:"switch_on"`
	Switcher      string `json:"switcher"`
	TpFlag        string `json:"tp_flag"`
	Unit          string `json:"unit"`
	UpdateDate    string `json:"updateDate"`
	UserId        int    `json:"userId"`
	UserName      string `json:"userName"`
	Value         string `json:"value"`
}
