package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
	"strings"
)

func Connect() {

}

func Post(url string, content []byte) ([]byte, error) {
	res, err := http.Post(url, "application/x-www-form-urlencoded", bytes.NewBuffer(content))
	if err != nil {
		// handle error
		beego.Debug("erro : ", err)

		return nil, err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		beego.Debug("erro : ", err)

		return nil, err
	}

	return resBody, err
}

func Get(url string) ([]byte, error) {
	// send parameter to env service
	res, err := http.Get(url)
	if err != nil {
		// handle error
		beego.Debug("erro : ", err)
		return nil, err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		beego.Debug("erro : ", err)
		return nil, err
	}

	return resBody, err
}

func Put(url string, content []byte) ([]byte, error) {
	client := http.Client{}
	req, _ := http.NewRequest("PUT", url, strings.NewReader(string(content)))

	res, err := client.Do(req)

	if err != nil {
		// handle error
		beego.Debug("erro : ", err)
		return nil, err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		beego.Debug("erro : ", err)
		return nil, err
	}

	return resBody, err

}

func Delete(url string, content []byte) ([]byte, error) {
	client := http.Client{}
	req, _ := http.NewRequest("DELETE", url, strings.NewReader(string(content)))

	res, err := client.Do(req)

	if err != nil {
		// handle error
		beego.Debug("erro : ", err)
		return nil, err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		beego.Debug("erro : ", err)
		return nil, err
	}

	return resBody, err

}

// func Post(ip string, port string, content []byte) ([]byte, error) {
// 	// send parameter to env service
// 	res, err := http.Post(ip+":"+port, "application/x-www-form-urlencoded", bytes.NewBuffer(content))
// 	//res, err := http.Post(ip+":"+port, "application/json", bytes.NewBuffer(content))

// 	//body := bytes.NewBuffer(content)
// 	//res, err := http.Post(ip+":"+port, "application/json;charset=utf-8", body)
// 	if err != nil {
// 		// handle error
// 		beego.Debug("erro : ", err)

// 		return nil, err
// 	}
// 	defer res.Body.Close()
// 	resBody, err := ioutil.ReadAll(res.Body)

// 	if err != nil {
// 		beego.Debug("erro : ", err)

// 		return nil, err
// 	}

// 	return resBody, err
// }

// func Get(ip string, port string, content []byte) ([]byte, error) {
// 	// send parameter to env service
// 	res, err := http.Get(ip + ":" + port + string(content))
// 	if err != nil {
// 		// handle error
// 		beego.Debug("erro : ", err)
// 		return nil, err
// 	}
// 	defer res.Body.Close()
// 	resBody, err := ioutil.ReadAll(res.Body)

// 	if err != nil {
// 		beego.Debug("erro : ", err)
// 		return nil, err
// 	}

// 	return resBody, err
// }

func Close() {

}
