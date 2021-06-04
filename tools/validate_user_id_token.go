package tools

import (
	auth "bitbucket.org/edgelabsolutions/insights-core/auth/jwt"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (t *Tools) ValidateDataIDToken(c *gin.Context, name string, value string, rd auth.AuthInterface, tk auth.TokenInterface) bool {
	// CHECK DATA FROM JWT
	metadata, err := tk.ExtractTokenMetadata(c.Request)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("something occur when extractTokenMeta auth :%s", err)
		return false
	}
	tokenUserID, err := rd.FetchAuth(metadata.TokenUuid)
	if err != nil {
		log.WithFields(log.Fields{
			"redis-dsn": os.Getenv("REDIS_DSN"),
			"uuid":      metadata.TokenUuid,
			"userid":    metadata.UserId,
			"error":     err,
		}).Errorf("something occur when Fetch auth :%s", err)
		return false
	}

	log.WithFields(log.Fields{
		"user_type": tokenUserID.UserType,
		"userId": tokenUserID.UserID,
	}).Infof("user verification")

	switch name {
	case "user_id":
		if tokenUserID.UserID == value {
			return true
		}
	case "user_type":
		if tokenUserID.UserType == value {
			return true
		}
	}

	return false
}
