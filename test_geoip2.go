package main

import "github.com/savaki/geoip2"
import "encoding/json"
import "os"

func main() {
	gl := geoip2.New("116626", "sjFkCDGvY5im")
	resp, _ := gl.Insights(nil, "47.186.133.28")
    json.NewEncoder(os.Stdout).Encode(resp)
}
