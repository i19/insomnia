package middleware

import (
	"github.com/gin-gonic/gin"
	"insomnia/internal/api"

	"insomnia/internal/api/handler"
	"insomnia/pkg/errors"
	"insomnia/pkg/grant"
)

const (
	sessionKeyForQuery = "session"
	projectIDForParam  = "project_id"
)

func GetGrantMiddleware(ctx *gin.Context) {
	sessionString := ctx.Query(sessionKeyForQuery)
	projectID := ctx.Param(projectIDForParam)

	// todo check ...
	//if len(tokenString) != 32 && len(projectID) != 32 {
	//	handler.Response(
	//		ctx, errors.BadRequest, "illegal token or projectID", nil,
	//	)
	//	return
	//}

	g, err := grant.GetBySession(sessionString)
	if err != nil {
		handler.Response(
			ctx, errors.Unauthorized, err.Error(), nil,
		)
		return
	}

	ctx.Set(api.XSessionKey, sessionString)
	ctx.Set(api.XProjectIDKey, projectID)
	ctx.Set(api.XGrantKey, g)
	ctx.Next()
}

func IsGrantedForThisProject(ctx *gin.Context) {
	tg, _ := ctx.Get(api.XGrantKey)
	projectID, _ := ctx.Get(api.XProjectIDKey)

	g := tg.(*grant.Grant)
	if !g.HasProject(projectID.(string)) {
		handler.Response(
			ctx,
			errors.Forbidden,
			"this session is not granted to this project",
			nil,
		)
		ctx.Abort()
		return
	}
	ctx.Next()
}
func IsAdminForThisProject(ctx *gin.Context) {
	tg, _ := ctx.Get(api.XGrantKey)
	projectID, _ := ctx.Get(api.XProjectIDKey)

	g := tg.(*grant.Grant)

	if !g.HasProject(projectID.(string)) || !g.IsAdmin() {
		handler.Response(
			ctx,
			errors.Forbidden,
			"this session is not admin for this project",
			nil,
		)
		ctx.Abort()
		return
	}
	ctx.Next()
}
