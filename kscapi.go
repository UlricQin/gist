package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/ulricqin/awsauth"
)

var (
	service = "kec"
	region  = "cn-beijing-6"
	ak      = "***"
	sk      = "***"
)

func main() {
	endpoint := fmt.Sprintf("http://%s.%s.api.ksyun.com/", service, region)

	query := url.Values{
		"Action":       []string{"MonitorInstances"},
		"Version":      []string{"2016-03-04"},
		"InstanceId.1": []string{"472ddf6c-6ab5-4f0e-ae94-35ad7723317c"},
	}

	url := endpoint + "?" + query.Encode()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("new request fail", err)
		return
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	awsauth.Sign4(req, awsauth.Credentials{AccessKeyID: ak, SecretAccessKey: sk}, region, service)

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Println("http get fail", err)
		return
	}

	defer res.Body.Close()

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("read response fail", err)
		return
	}

	fmt.Println(string(bs))
}
