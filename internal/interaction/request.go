package interaction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tgBotIP/internal/env"
	"tgBotIP/internal/types"
)

func Request() (types.Site, error) {
	{
		client := &http.Client{}
		parseEnv := env.ParseEnv("IP_REQUEST_ADDRESS")
		req, err := http.NewRequest("GET", parseEnv, nil)
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
