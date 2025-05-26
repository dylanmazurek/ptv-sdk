package ptv

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"net/url"
	"strings"
)

type Credentials struct {
	Key    string
	UserID string
}

func (c *Credentials) AuthRequestUrl(reqUrl url.URL) (url.URL, error) {
	authedUrl := reqUrl

	q := authedUrl.Query()
	q.Set("devid", c.UserID)

	stringToSign := authedUrl.Path
	paramsForSigningEncoded := q.Encode()
	if paramsForSigningEncoded != "" {
		stringToSign += "?" + paramsForSigningEncoded
	}

	hash := hmac.New(sha1.New, []byte(c.Key))
	_, err := hash.Write([]byte(stringToSign))
	if err != nil {
		return url.URL{}, err
	}
	hashBytes := hash.Sum(nil)
	hashString := strings.ToUpper(hex.EncodeToString(hashBytes))

	q.Set("signature", hashString)

	authedUrl.RawQuery = q.Encode()
	authedUrl.ForceQuery = true

	return authedUrl, nil
}
