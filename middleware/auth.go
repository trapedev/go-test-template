package middleware

import "net/http"

type AuthGroup string

const (
	JUKSLAdmin   AuthGroup = `juksl-admin`
	CompanyAdmin AuthGroup = `company-admin`
	SchoolAdmin  AuthGroup = `school-admin`
	Staff        AuthGroup = `staff`
	Teacher      AuthGroup = `teacher`
	Student      AuthGroup = `student`
	None         AuthGroup = `none`
)

var MessageList map[AuthGroup]string = map[AuthGroup]string{
	JUKSLAdmin:   `サービス管理者`,
	CompanyAdmin: `会社管理者`,
	SchoolAdmin:  `校舎管理者`,
	Staff:        `スタッフ`,
	Teacher:      `講師`,
	Student:      `生徒`,
	None:         ``,
}

var AuthWeightList map[AuthGroup]uint = map[AuthGroup]uint{
	JUKSLAdmin:   32,
	CompanyAdmin: 16,
	SchoolAdmin:  8,
	Staff:        4,
	Teacher:      2,
	Student:      1,
	None:         0,
}

type AuthGroupConfig struct {
	ReadAuth  []http.Handler
	WriteAuth []http.Handler
}

// ユーザーのもつ権限が，必要最低限の権限を満たしているか判定
func JudgeUserAuth(minimunAllowAuth AuthGroup, userAuth AuthGroup) (bool, string) {
	if AuthWeightList[minimunAllowAuth] > AuthWeightList[userAuth] {
		return false, MessageList[minimunAllowAuth]
	}
	return true, "OK"
}
