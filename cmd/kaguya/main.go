// Copyright 2021 Star Inc.(https://starinc.xyz)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/gin-gonic/gin"
	Rethink "gopkg.in/rethinkdb/rethinkdb-go.v6"
	Kernel "gopkg.in/star-inc/kaguyakernel.v1"
	TalkService "gopkg.in/star-inc/kaguyakernel.v1/service/talk"
)

func main() {
	router := gin.Default()
	const authCookie = "kaguya_token"

	dbConfig := Kernel.RethinkConfig{
		DatabaseName:  "Kaguya",
		ConnectConfig: Rethink.ConnectOpts{Address: "localhost"},
	}

	router.GET("/talk/:target", func(c *gin.Context) {
		manager := TalkService.NewManager(dbConfig, c.Param("target"))
		if !manager.Check() {
			manager.Create()
		}
		service := TalkService.NewServiceInterface(
			dbConfig,
			c.Param("target"),
			func(contentType int, content string) bool {
				return true
			},
		)
		cookie, err := c.Request.Cookie(authCookie)
		if err != nil {
			panic(err)
		}
		handler := Kernel.Run(service, newAuth(cookie.Value), "kaguya_test")
		err = handler.HandleRequest(c.Writer, c.Request)
		if err != nil {
			panic(err)
		}
	})

	err := router.Run(":10101")
	if err != nil {
		panic(err)
	}
}
