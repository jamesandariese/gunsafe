package deliver

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/luksen/maildir"
)

var MessageDownloadError = errors.New("Message Download Error")

var apikey = flag.String("apikey", "", "Mailgun API key")
var maildirIn = flag.String("maildir", "mail", "Maildir pattern to deliver into; %s is username")

func Deliver(url string) error {
	if !flag.Parsed() {
		flag.Parse()
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	request.SetBasicAuth("api", *apikey)

	request.Header.Add("Accept", "message/rfc2822")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 || resp.StatusCode < 200 {
		return MessageDownloadError
	}

	decoded := map[string]interface{}{}

	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&decoded)
	fmt.Fprintf(os.Stderr, "%#v", decoded)

	md := maildir.Dir(*maildirIn)
	md.Create()
	if err := md.Clean(); err != nil {
		return err
	}
	delivery, err := md.NewDelivery()
	if err != nil {
		return err
	}
	defer delivery.Close()
	fmt.Fprintf(os.Stderr, "Delivering %d bytes", len(decoded["body-mime"].(string)))
	body := strings.Replace(decoded["body-mime"].(string), "\r\n", "\n", -1)
	if _, err := fmt.Fprintf(delivery, "%s", body); err != nil {
		delivery.Abort()
		return err
	}
	return nil
}
