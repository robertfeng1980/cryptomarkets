package bitstamp

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func (e *Bitstamp) getAuthValues() url.Values {
	nonce := strconv.FormatInt(time.Now().Unix(), 10)

	message := nonce + e.CustomerID + e.ApiKey

	mac := hmac.New(sha256.New, []byte(e.Secret))
	mac.Write([]byte(message))
	macSum := mac.Sum(nil)
	sig := strings.ToUpper(hex.EncodeToString(macSum))

	return url.Values{
		"key":       {e.ApiKey},
		"signature": {sig},
		"nonce":     {nonce},
	}
}
