package exchange

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitstamp_SetPairs(t *testing.T) {
	bitstamp := Bitstamp{}
	bitstamp.SetPairs()
	pairs := bitstamp.GetConfig().Pairs

	assert.Contains(t, pairs, &Pair{"BTC", "USD"})
	assert.Contains(t, pairs, &Pair{"ETH", "EUR"})
}

/* Disabled due to rate limit
func TestBitstamp_GetResponse(t *testing.T) {
	bitstamp := Bitstamp{}
	price, err := bitstamp.GetResponse("BTC", "USD")
	if err != nil {
		log.Fatal(err)
	}
	assert.True(t, price.Price > 0, "price from Bitstamp isn't greater than 0")
	assert.True(t, price.Volume > 0, "volume from Bitstamp isn't greater than 0")
}
*/