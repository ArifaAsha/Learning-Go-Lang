package main

import (
	"fmt"

	fb "github.com/huandu/facebook/v2"
)

func main() {

	res, _ := fb.Get("/1247558609120429", fb.Params{ //1st parameter = Facebook Developer App Id
		"fields":       "name, id",
		"access_token": "EAARupdZB7rK0BANgH04UrBHcC4RD3QgtMjVqHL7G0vIXx1hl95zqo15nNZC6hNh9O0ZASDxbMm7gvyYh3XVjCqRxNfvhm2Ni6JEacIaZBme5WxPYbrBo072SM9iNMnkM3T4Fq7lBrvZBZBAHThWlEmZCpwZC2G6n8wfoUD1ZAJ9QOJwMA3F9nOGrGJB4WmgYsAm7bLzqDstFNVYDspyb3HwV5YJavuc1lhiYikb44ZCS3ZAjAOtO4EQ9qzifFZAnbQwLHQUZD", //Generated Access Token
	})

	fmt.Println("Here is my Facebook first name:", res["name"])
	fmt.Println("Here is my Facebook first name:", res["id"])
}
