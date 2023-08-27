package openfigi

type FilterJob struct {
	Query                   string     `json:"query"`
	Start                   string     `json:"start,omitempty"`
	ExchCode                string     `json:"exchCode,omitempty"`
	MicCode                 string     `json:"micCode,omitempty"`
	Currency                string     `json:"currency,omitempty"`
	MarketSecDes            string     `json:"marketSecDes,omitempty"`
	SecurityType            string     `json:"securityType,omitempty"`
	SecurityType2           string     `json:"securityType2,omitempty"`
	IncludeUnlistedEquities bool       `json:"includeUnlistedEquities,omitempty"`
	OptionType              string     `json:"optionType,omitempty"`
	Strike                  *[]float64 `json:"strike,omitempty"`
	ContractSize            *[]float64 `json:"contractSize,omitempty"`
	Coupon                  *[]float64 `json:"coupon,omitempty"`
	Expiration              *[]string  `json:"expiration,omitempty"`
	Maturity                *[]string  `json:"maturity,omitempty"`
	StateCode               string     `json:"stateCode,omitempty"`
}
