package token

import (
	//"github.com/dgrijalva/jwt-go"
)
type ROLE string
const (
	READER ROLE = "READER"
	OPERATOR ROLE = "OPERATOR"
)

type Privilege struct {
	/* 角色 */
	role ROLE

	/* 拥有 role 的服务 */
	service string

	/* 环境 */
	env string
}

type User struct {
	id string
	Type string
	privileges []*Privilege
}

/* token 用于 api 调用认证 */
type JwtToken struct {
	userInfo User
	issueAt int64
	issueBy string
}

/* 签发token */
func (t *JwtToken) IssueToken() JwtToken {

	return JwtToken{}
}

/* 解析 token */
func (t *JwtToken) Decode(token string) JwtToken {
	return JwtToken{}
}

/* 编码token */
func (t *JwtToken) Encode(user JwtToken) string {
	return ""
}

/* 校验 token */
func (t *JwtToken) Valid() bool {
	return true
}
