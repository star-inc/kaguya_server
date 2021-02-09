/*
   Copyright 2020 Star Inc.(https://starinc.xyz)
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
	"github.com/gin-gonic/gin"
	Kernel "github.com/star-inc/kaguya_kernel"
	TalkService "github.com/star-inc/kaguya_kernel/service/talk"
	Rethink "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func main() {
	router := gin.Default()

	dbConfig := Kernel.RethinkConfig{
		ConnectConfig: Rethink.ConnectOpts{Address: "localhost"},
		DatabaseName:  "Kaguya",
	}

	router.GET("/talk", func(c *gin.Context) {
		service := TalkService.NewServiceInterface(dbConfig, "talk")
		cookie, err := c.Request.Cookie("kaguya_token")
		if err != nil {
			panic(err)
		}
		handler := Kernel.Run(service, newAuth(cookie.Value))
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
