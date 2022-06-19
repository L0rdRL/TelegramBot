package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host string 
	basePath string
	client http.Client
}

const (
	getUpdatesMethod="getUpdates"
	SendMassageMethod="sendMessage"
)

func New(host string, token string) Client{
	return Client{
		host: host,
		basePath: newBasePath(token),
		client: http.Clent{}
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) Updates(offset int, limit int)  ([]Update, error){
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	// integer to ASCII
	q.Add("limit", strconv.Itoa(limit))

	// do request
	data, err:=c.doRequset(getUpdatesMethod, q)
	if err != nil {
		return nil, err
	}

	var res UpdatesRespons

	if err:=json.Unmarshal(data,&res); err != nil{ 
		return nil, err 
	}

	return res.Result, nil
}

func (c *Client) SendMassage(chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_,err:=c.doRequset(SendMassageMethod, q)
	if err != nil {
		return e.Wrap("can't send message", err)
	}

	return nil

}


func (c *Client) doRequset(method string, query url.Values) (data []byte, err error) {
	defer func() {err = e.WrapIfErr("can't do request", err)} ()

	u:=url.URL{
		Scheme: "https",
		Host: c.host
		Path: path.Join(c.basePath, method),
	}

	req,err :=	http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	resp,err:=c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {_=resp.Body.Close()}()

	body,err:=io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

