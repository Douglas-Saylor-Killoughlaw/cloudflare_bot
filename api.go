package cloudflare_api
/*
import (
		"net/http"
		"log"
		"io/ioutil"
)

curl -X GET "https://api.cloudflare.com/client/v4/user/tokens/verify" \
     -H "Authorization: Bearer EYK6GGHLsHld02oGbyjRl9-piJ93WL4gjPwdijEL" \
	      -H "Content-Type:application/json"

*/
type user struct {
	api		string
	token	string
}


