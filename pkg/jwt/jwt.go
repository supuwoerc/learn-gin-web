package jwt

import (
	"gin-web/models"
	"gin-web/pkg/constant"
	"gin-web/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

type TokenClaimsBasic struct {
	UID      uint
	Email    string
	NickName string
	Gender   models.UserGender
	About    string
	Birthday string
}

type TokenClaims struct {
	jwt.RegisteredClaims
	User *TokenClaimsBasic
}

type JwtBuilder struct {
	ctx *gin.Context
}

var jwtBuilder *JwtBuilder

func NewJwtBuilder(ctx *gin.Context) *JwtBuilder {
	if jwtBuilder == nil {
		jwtBuilder = &JwtBuilder{
			ctx: ctx,
		}
	}
	return jwtBuilder
}

// 生成token
func (j *JwtBuilder) generateToken(user *TokenClaimsBasic, createAt time.Time, duration time.Duration) (string, error) {
	claims := TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    viper.GetString("jwt.issuer"),
			IssuedAt:  jwt.NewNumericDate(createAt),
			ExpiresAt: jwt.NewNumericDate(createAt.Add(duration)),
		},
		User: user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(viper.GetString("jwt.secret")))
}

// 生成短token
func (j *JwtBuilder) generateAccessToken(user *TokenClaimsBasic, createAt time.Time) (string, error) {
	return j.generateToken(user, createAt, viper.GetDuration("jwt.expires")*time.Minute)
}

// 生成长token
func (j *JwtBuilder) generateRefreshToken(createAt time.Time) (string, error) {
	return j.generateToken(&TokenClaimsBasic{}, createAt, viper.GetDuration("jwt.refreshTokenExpires")*time.Minute)
}

type RefreshTokenCallback func(claims *TokenClaims, accessToken, refreshToken string) error

// 校验并生成长短token
func (j *JwtBuilder) ReGenerateAccessAndRefreshToken(accessToken, refreshToken string, callback RefreshTokenCallback) (string, string, error) {
	if _, err := j.ParseToken(refreshToken); err != nil {
		return "", "", constant.GetError(j.ctx, response.InvalidRefreshToken)
	}
	claims, err := j.ParseToken(accessToken)
	if err == nil {
		return "", "", constant.GetError(j.ctx, response.UnnecessaryRefreshToken)
	}
	if err != constant.GetError(j.ctx, response.InvalidToken) {
		return "", "", err
	}
	createAt := time.Now()
	newAccessToken, err := j.generateAccessToken(&TokenClaimsBasic{
		UID:      claims.User.UID,
		Email:    claims.User.Email,
		NickName: claims.User.NickName,
		Gender:   claims.User.Gender,
		About:    claims.User.About,
		Birthday: claims.User.Birthday,
	}, createAt)
	if err != nil {
		return "", "", err
	}
	newRefreshToken, err := j.generateRefreshToken(createAt)
	if err != nil {
		return "", "", err
	}
	if callback != nil {
		callbackErr := callback(claims, newAccessToken, newRefreshToken)
		if callbackErr != nil {
			return "", "", callbackErr
		}
	}
	return newAccessToken, newRefreshToken, nil
}

// 生成长短token
func (j *JwtBuilder) GenerateAccessAndRefreshToken(user *TokenClaimsBasic) (string, string, error) {
	createAt := time.Now()
	newAccessToken, err := j.generateAccessToken(user, createAt)
	if err != nil {
		return "", "", err
	}
	newRefreshToken, err := j.generateRefreshToken(createAt)
	if err != nil {
		return "", "", err
	}
	return newAccessToken, newRefreshToken, nil
}

// 解析token
func (j *JwtBuilder) ParseToken(tokenString string) (*TokenClaims, error) {
	claims := TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("jwt.secret")), nil
	})
	if err != nil || !token.Valid {
		return &claims, constant.GetError(j.ctx, response.InvalidToken)
	}
	return &claims, nil
}

// 获取上下文的claims
func GetUserClaims(ctx *gin.Context) *TokenClaims {
	value, exists := ctx.Get(constant.ClaimKeyContext)
	if exists {
		return value.(*TokenClaims)
	}
	return nil
}
