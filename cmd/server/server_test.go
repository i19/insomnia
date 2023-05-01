package server

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/parnurzeal/gorequest"

	"insomnia/internal/api/handler"
	"insomnia/internal/api/router"
	"insomnia/pkg/errors"
	"insomnia/pkg/rule_engine"
)

var ruleYaml = `
type: object
properties:
  okDescription:
    type: string
    format: description
  failDescription:
    type: string
    format: description
`

var spec = `{"okDescription": "description hello world","failDescription": "hello world"}`

const (
	adminUser   = "admin_session"
	grantUser   = "grant_session"
	guest       = "guest_session"
	testProject = "project_1"
)

func TestUpdateRuleByDifferentGrants(t *testing.T) {
	server := router.Get()
	request := gorequest.New()
	cases := map[string]errors.ErrorCode{
		adminUser: errors.OK,
		grantUser: errors.Forbidden,
		guest:     errors.Forbidden,
	}
	for user, want := range cases {
		req, err := request.Post(fmt.Sprintf("/api/linting_rule/%s", testProject)).
			Query(map[string]string{"session": user}).
			Send(&handler.RequestForPostByProjectID{Rule: ruleYaml}).
			MakeRequest()
		if err != nil {
			t.Fatal(err)
		}
		recorder := httptest.NewRecorder()
		server.ServeHTTP(recorder, req)
		if recorder.Code != 200 {
			t.Errorf("http code should be 200 all the time, user %s get %d", user, recorder.Code)
		}
		var resp handler.ResponseProtocol
		if err := json.Unmarshal(recorder.Body.Bytes(), &resp); err != nil {
			t.Errorf("http response protocol error")
		}
		if resp.Code != want {
			t.Errorf("user %s didn't get expeceted code %d , get %d", user, want, resp.Code)
		}
	}
}

func TestValidate(t *testing.T) {
	server := router.Get()
	request := gorequest.New()
	cases := map[string][]int{
		adminUser: {int(errors.OK), 1},
		grantUser: {int(errors.OK), 1},
		guest:     {int(errors.Forbidden), 0},
	}

	rule_engine.InitByRaw(testProject, []byte(ruleYaml))
	for user, want := range cases {
		req, err := request.Post(fmt.Sprintf("/api/linting_rule/%s/validate", testProject)).
			Query(map[string]string{"session": user}).
			Send(&handler.RequestForValidate{Spec: spec}).
			MakeRequest()
		if err != nil {
			t.Fatal(err)
		}
		recorder := httptest.NewRecorder()
		server.ServeHTTP(recorder, req)
		if recorder.Code != 200 {
			t.Errorf("http code should be 200 all the time, user %s get %d", user, recorder.Code)
			continue
		}
		var resp handler.ResponseProtocol
		if err := json.Unmarshal(recorder.Body.Bytes(), &resp); err != nil {
			t.Errorf("http response protocol error")
			continue
		}
		if int(resp.Code) != want[0] {
			t.Errorf("user %s get wrong response code  %s", user, recorder.Body.String())
			continue
		}

		var result rule_engine.ValidationResult
		resultInByte, _ := json.Marshal(resp.Data)
		if err := json.Unmarshal(resultInByte, &result); err != nil {
			t.Errorf("user %s get wrong response structure, should be rule_engine.ValidationResult get %s", user, string(resultInByte))
			continue
		}
		if resp.Code == errors.OK && len(result.Errors) != want[1] {
			t.Errorf("user %s get wrong result size, expect '(root).failDescription ... ' get %d |%s|", user, len(result.Errors), strings.Join(result.Errors, "\n"))
		}
	}
}
