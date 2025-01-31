/*
 * Todolist RESTful API
 *
 * OpenAPI for Todolist RESTful API
 *
 * API version: 1
 * Contact: ariefkardityahrmwn@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CreateOrUpdateTodolist struct {
	Name string `json:"name,omitempty"`
	Priority int32 `json:"priority,omitempty"`
	Tags []string `json:"tags,omitempty"`
}
