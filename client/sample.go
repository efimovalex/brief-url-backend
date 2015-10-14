package client

type SomeResponse struct {
	SomeCode  string  `json:"code"`
	SomeType  string  `json:"type"`
	SomeValue float32 `json:"value"`
}
