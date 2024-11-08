package ms_gateway

// IndicatorParams は指標のパラメータを表す構造体
type IndicatorParams struct {
	Type   string            `json:"type"`   // 指標の種類 (SMA, EMA, RSIなど)
	Params map[string]string `json:"params"` // 指標のパラメータ (期間など)
}
