package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const prismDriveApiUrl = "https://app.prismdrive.com/api/v1"

// ------------------ name  type size  url   path
const fileEntryHead = "Name      \tType  \tSize                \tURL                     \tPath                    \n--------------------------------------------------------------\n"
const fmtFileEntry = `%10s\t%6s\t%20d\t%24s\t%24s\n`

type Client struct {
	baseUrl string
	headers http.Header
	c       *http.Client
}

func NewClient(token string) *Client {
	hc := &http.Client{}
	c := &Client{
		baseUrl: prismDriveApiUrl,
		headers: make(http.Header),
		c:       hc,
	}
	if token != "" {
		c.headers.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	return c
}

func Login(user, pass string) error {
	log.Printf("user: %s pass: %s\n", user, pass)
	b, err := json.Marshal(ReqLogin{Email: user, Password: pass, DeviceName: ""})
	if err != nil {
		return err
	}
	res, err := http.Post(prismDriveApiUrl+"/auth/login", "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	resb, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	var data ResLoginOK
	if err = json.Unmarshal(resb, &data); err != nil {
		return err
	}
	log.Printf("%v\n", data)
	fmt.Println(data.User.Token)
	return nil
}

func (c *Client) Ls() error {
	res, err := c.get("file-entries")
	if err != nil {
		return err
	}
	resb, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	var data FileEntries
	if err = json.Unmarshal(resb, &data); err != nil {
		return err
	}

	fmt.Printf(fileEntryHead)
	for _, fe := range data {
		// ------------------ name  type size  url   path
		fmt.Printf(fmtFileEntry, fe.Name, fe.Type, fe.Size, fe.Url, fe.Path)
	}
	return nil
}

func (c *Client) get(route string) (*http.Response, error) {
	return c.c.Get(c.baseUrl + "/" + route)
}
