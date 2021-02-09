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
	Kernel "github.com/star-inc/kaguya_kernel"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Authorize struct {
	me *Kernel.User
}

func newAuth(cookie string) Kernel.AuthorizeInterface {
	authorize := new(Authorize)
	me := new(Kernel.User)
	response, err := http.PostForm(
		"http://dev.localhost:5000/api/verify",
		url.Values{"authToken": {cookie}},
	)
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
	me.Identity = result["reason"].(string)
	authorize.me = me
	return authorize
}

func (authorize Authorize) Me() *Kernel.User {
	return authorize.me
}

func (authorize Authorize) Permission(tableName string) bool {
	if tableName != "" {
		return true
	}
	return false
}
