package cmd

import (
	"github.com/markcheno/go-talib"
	"github.com/nntaoli-project/goex"
	"github.com/nntaoli-project/goex/okex"
	"github.com/nntaoli-project/goex_talib"
	"net/http"
	"testing"
)

var api = okex.NewOKEx(&goex.APIConfig{
	HttpClient: http.DefaultClient,
	Endpoint:   "https://www.okex.me",
})

func Test_talib(t *testing.T) {
	data, _ := api.GetKlineRecords(goex.BTC_USDT, goex.KLINE_PERIOD_1H, 300, 0)
	t.Log(goex_talib.Macd(data, 12, 26, 9, goex_talib.InClose))
	t.Log(goex_talib.Ma(data, 60, talib.EMA, goex_talib.InClose))
	t.Log(goex_talib.Atr(data, 20))
}
