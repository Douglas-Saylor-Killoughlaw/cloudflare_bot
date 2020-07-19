package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    client := &http.Client {
    }

    req, _ := http.NewRequest("GET", "https://api.cloudflare.com/client/v4/user/tokens/verify", nil)
    req.Header.Add("Authorization", "Bearer xp-trskKUw9zSx1cw3SFoyc2p0Bq-4sB4N5mD5lK")
    req.Header.Add("Content-Type", "application/json")
    resp, _ := client.Do(req)
    data, _ := ioutil.ReadAll(resp.Body)
    fmt.Print(string(data)+"\n")
}
