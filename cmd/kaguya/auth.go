/*
   Copyright 2021 Star Inc.(https://starinc.xyz)
   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at
       http://www.apache.org/licenses/LICENSE-2.0
   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package main

import (
	"encoding/json"
	Kernel "gopkg.in/star-inc/kaguyakernel.v1"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Authorize struct {
	cookie   string
	identity string
}

func newAuth(cookie string) Kernel.AuthorizeInterface {
	authorize := new(Authorize)
	authorize.cookie = cookie
	result := postRequest(
		"http://localhost:5000/api/verify",
		url.Values{"authToken": {cookie}},
	)
	authorize.identity = result["reason"].(string)
	return authorize
}

func postRequest(url string, formData url.Values) map[string]interface{} {
	response, err := http.PostForm(url, formData)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalln(err)
	}
	return result
}

func (authorize *Authorize) Me() string {
	return authorize.identity
}

func (authorize *Authorize) Permission(tableName string) bool {
	if tableName != "" {
		return true
	}
	return false
}
