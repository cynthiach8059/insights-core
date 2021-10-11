package tools

import (
	auth "github.com/cynthiach8059/insights-core/auth/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

//GetUserFromJWTToken gets userID from jwt token
func (t *Tools) GetUserFromJWTToken(c *gin.Context, rd auth.AuthInterface, tk auth.TokenInterface) (int, error) {
	// CHECK DATA FROM JWT
	metadata, err := tk.ExtractTokenMetadata(c.Request)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("something occur when extractTokenMeta auth :%s", err)
		return 0,fmt.Errorf("something occur when extractTokenMeta auth :%s", err)
	}
	tokenUserID, err := rd.FetchAuth(metadata.TokenUuid)
	if err != nil {
		log.WithFields(log.Fields{
			"redis-dsn": os.Getenv("REDIS_DSN"),
			"uuid":      metadata.TokenUuid,
			"userid":    metadata.UserId,
			"error":     err,
		}).Errorf("something occur when Fetch auth :%s", err)
		return 0,fmt.Errorf("something occur when Fetch auth :%s", err)
	}



	userID, err:= strconv.Atoi(tokenUserID.UserID)
	if err!=nil{
		log.WithFields(log.Fields{
			"redis-dsn": os.Getenv("REDIS_DSN"),
			"uuid":      metadata.TokenUuid,
			"userid":    metadata.UserId,
			"error":     err,
		}).Errorf("cannot parse userID from token :%s", err)
		return 0,fmt.Errorf("cannot parse userID from token :%s", err)
	}


	return userID, nil
}
