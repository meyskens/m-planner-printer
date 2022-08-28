package api

//go:generate easyjson -all $GOFILE

import (
	"fmt"
	"io/ioutil"

	"github.com/mailru/easyjson"
	"tinygo.org/x/drivers/net/http"
)

var buf [40960]byte

type API struct {
	key      string
	endpoint string
}

func NewApi(endpoint, key string) *API {
	return &API{
		key:      key,
		endpoint: endpoint,
	}
}

type response struct {
	Data []printJob `json:"data"`
}

type printJob struct {
	User       string `json:"user"`
	EscposData []byte `json:"escposData"`
}

func (a *API) GetPrintJobs() ([][]byte, error) {
	http.SetBuf(buf[:])

	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/tiny/printjobs", a.endpoint), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", a.key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	data := response{}
	err = easyjson.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	out := make([][]byte, 0)
	for _, d := range data.Data {
		out = append(out, []byte(d.EscposData))
	}

	return out, nil
}
