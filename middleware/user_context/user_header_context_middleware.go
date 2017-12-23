package user_middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gocms-io/gocms/context/consts"
	"github.com/gocms-io/gocms/domain/user/user_model"
)

// UserHeaderContext
func UserHeaderContext() gin.HandlerFunc {
	return userHeaderContext
}

func userHeaderContext(c *gin.Context) {

	// check if header exists
	userContextString := c.Request.Header.Get(consts.GOCMS_HEADER_USER_CONTEXT_KEY)
	if userContextString == "" {
		c.Next()
		return
	}

	// if it does add it to the context
	userContext := user_model.UnmarshalUserContextHeader([]byte(userContextString))
	c.Set(consts.USER_KEY_FOR_GIN_CONTEXT, *userContext)
	c.Next()
}

func GetUserFromContext(c *gin.Context) (*user_model.UserContextHeader, bool) {
	// get user from context
	if userContext, ok := c.Get(consts.USER_KEY_FOR_GIN_CONTEXT); ok {
		if userHeaderContext, ok := userContext.(user_model.UserContextHeader); ok {
			return &userHeaderContext, true
		}
	}
	return nil, false
}
