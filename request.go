package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tgBotIP/types"
)

func request() (types.Site, error) {
	{
		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://ifconfig.co", nil)
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Accept", "application/json")

		res, _ := client.Do(req)
		defer func(Body io.ReadCloser) {
			var err = Body.Close()
			if err != nil {
			}
		}(res.Body)

		body, err := io.ReadAll(res.Body)
		var info types.Site
		err = json.Unmarshal(body, &info)
		if err != nil {
			return types.Site{}, err
		}
		return info, nil
	}
}
