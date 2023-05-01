package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"insomnia/internal/api"
	"insomnia/internal/config"
	"insomnia/internal/platform/etcd"
	"insomnia/pkg/errors"
	"insomnia/pkg/rule_engine"
	"time"
)

type RequestForPostByProjectID struct {
	Rule string `json:"rule"`
}

type RequestForValidate struct {
	Spec string `json:"spec"`
}

// PostByProjectID godoc
// @Summary      Update linting rules by project id
// @Description  If there is admin grant, allow creating or updating linting rules, otherwise not allowed
// @Tags         LintingRule
// @Accept       json
// @Produce      json
// @Param        project_id path string true "unique ID for a project"
// @Param        token query string true "session token"
// @Param        {object} body RequestForPostByProjectID true "custom linting rule in yaml"
// @Router       /api/linting_rule/{project_id} [post]
// @success      200 {object} handler.ResponseProtocol{}
func PostByProjectID(c *gin.Context) {
	projectID := c.GetString(api.XProjectIDKey)
	var request RequestForPostByProjectID
	if err := c.ShouldBindJSON(&request); err != nil {
		Response(c, errors.BadRequest, fmt.Sprintf("read body error: %s", err.Error()), nil)
	}

	if _, err := rule_engine.GenerateSchemaFromYaml([]byte(request.Rule)); err != nil {
		Response(c, errors.BadRequest, fmt.Sprintf("rule illega: %s", err.Error()), nil)
		return
	}
	ctx, cancel := context.WithTimeout(c, time.Second)
	defer cancel()

	//todo version check
	if _, err := etcd.FakePut(ctx, fmt.Sprintf("%s/%s", config.Config.EtcdLintingRulePrefix, projectID), request.Rule); err != nil {
		//if _, err := etcd.Client.Put(ctx, fmt.Sprintf("%s/%s", config.Config.EtcdLintingRulePrefix, projectID), request.Rule); err != nil {
		Response(c, errors.InternalServerError, fmt.Sprintf("rule save error: %s", err.Error()), nil)
		return
	}

	ResponseOK(c, nil)
}

// Validate godoc
// @Summary      Validate spec
// @Description  Validate spec with linting rule for the project
// @Tags         LintingRule
// @Accept       json
// @Produce      json
// @Param        project_id path string true "unique ID for a project"
// @Param        token query string true "session token"
// @Param         {object} body RequestForValidate true "Specification to be validated "
// @Router       /api/linting_rule/{project_id}/validate [post]
// @success      200 {object} handler.ResponseProtocol{data=[]rule_engine.ValidationResult} "validate result"
func Validate(c *gin.Context) {
	projectID := c.GetString(api.XProjectIDKey)
	var request RequestForValidate
	if err := c.ShouldBindJSON(&request); err != nil {
		Response(c, errors.BadRequest, fmt.Sprintf("read body error: %s", err.Error()), nil)
	}

	result, err := rule_engine.Render(projectID, []byte(request.Spec))
	if err != nil {
		Response(c, errors.InternalServerError, fmt.Sprintf("read body error: %s", err.Error()), nil)
		return
	}

	ResponseOK(c, result)
}
