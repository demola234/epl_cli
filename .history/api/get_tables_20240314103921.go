package api 

import(
	"http/net"
	"encoding/json"

)

func getData() {
	
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://example.com/tables", nil
}