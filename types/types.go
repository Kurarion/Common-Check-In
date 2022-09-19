package types

//人物定义
type CheckInData struct {
	Disable bool          `json:"disable"`
	Specs   []string      `json:"specs"`
	Url     string        `json:"url"`
	Method  string        `json:"method"`
	Headers []HeadersData `json:"headers"`
	Payload string        `json:"payload"`
}

type CheckInDatas []CheckInData

//Header type
type HeadersData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
