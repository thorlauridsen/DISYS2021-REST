/*
 * ITU REST API
 *
 * This is documentation for our REST api for mandatory exercise 1 in the distributed systems course
 *
 * API version: 1.0.1½
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type Teacher struct {
	Id int64 `json:"id"`

	Name string `json:"name"`

	Score int64 `json:"score,omitempty"`
}

func parseTeacher(body io.ReadCloser) Teacher {
	resp, _ := ioutil.ReadAll(body)

	var teacher Teacher
	json.Unmarshal(resp, &teacher)
	return teacher
}
