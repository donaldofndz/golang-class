package cryptocurrency

var (
	bitcoinPrice float32
)

func init() {
	bitcoinPrice = 38.02
}

// ConvertDollarsToBitcoin is a function that convert dolars to bitcoin
func ConvertDollarsToBitcoin(dollars float32) float32 {
	return bitcoinPrice * dollars
}
