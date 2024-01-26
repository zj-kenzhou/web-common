package response

type DataResult[V any] struct {
	Total int64 `json:"total"`
	Data  V     `json:"data"`
}

type OptionResult[V any, L any] struct {
	Value V `json:"value"`
	Label L `json:"label"`
}
