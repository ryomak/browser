package transmitter

import (
	"net/http"
	"io/ioutil"
	"strings"
	"github.com/ryomak/browser/src/datatype"
)

func Transmit(url string) (*datatype.Response, error) {
	if strings.Index("http",url) != -1 {
		return GetLocal(url)
	}
	return Fetch(url)
}

func GetLocal(path string)(*datatype.Response,error){
	body ,err := ioutil.ReadFile(path)
	if err != nil{
		return &datatype.Response{},err
	}
	res := &datatype.Response{
		Body:string(body),
	}
	return res,nil
}

func Fetch(url string)(*datatype.Response, error){
	resp,err := http.Get(url)
	if err != nil {
		return &datatype.Response{},err
	}
	defer resp.Body.Close()
	body ,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &datatype.Response{},err
	}
	res := &datatype.Response{
		Header:resp.Header,
		Body:string(body),
	}
	return res ,nil
}
