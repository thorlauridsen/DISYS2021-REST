/*
 * ITU REST API
 *
 * This is documentation for our REST api for mandatory exercise 1 in the distributed systems course
 *
 * API version: 1.0.1½
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Course struct {

	Id int64 `json:"id"`

	Name string `json:"name"`

	Ects float32 `json:"ects"`

	TeacherID int64 `json:"teacherID"`

	Rating int64 `json:"rating,omitempty"`

	Students []Student `json:"students,omitempty"`
}
