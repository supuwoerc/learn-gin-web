// Code generated by "stringer -type=StatusCode -linecomment -output code_string.go"; DO NOT EDIT.

package response

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Ok-10000]
	_ = x[Error-10001]
	_ = x[InvalidParams-10002]
	_ = x[InvalidToken-10003]
	_ = x[CancelRequest-10004]
	_ = x[RecoveryError-10005]
	_ = x[InvalidRefreshToken-10006]
	_ = x[UnnecessaryRefreshToken-10007]
	_ = x[AuthErr-10008]
	_ = x[NoAuthority-10009]
	_ = x[TimeoutErr-10010]
	_ = x[Busy-10011]
	_ = x[UserCreateDuplicateEmail-20000]
	_ = x[UserLoginEmailNotFound-20001]
	_ = x[UserLoginFail-20002]
	_ = x[UserLoginTokenPairCacheErr-20003]
	_ = x[PasswordValidErr-20004]
	_ = x[UserNotExist-20005]
	_ = x[CaptchaVerifyFail-30000]
	_ = x[RoleCreateDuplicateName-40000]
	_ = x[NoValidRoles-40001]
	_ = x[PermissionCreateDuplicate-50000]
	_ = x[PermissionNotExist-50001]
	_ = x[InvalidAttachmentLength-50000]
}

const (
	_StatusCode_name_0 = "okerrorinvalidParamsinvalidTokencancelRequestrecoveryErrorinvalidRefreshTokenunnecessaryRefreshTokenauthErrnoAuthoritytimeoutErrbusy"
	_StatusCode_name_1 = "userCreateDuplicateEmailuserLoginEmailNotFounduserLoginFailuserLoginTokenPairCacheErrpasswordValidErruserNotExist"
	_StatusCode_name_2 = "captchaVerifyFail"
	_StatusCode_name_3 = "roleCreateDuplicateNamenoValidRoles"
	_StatusCode_name_4 = "permissionCreateDuplicatepermissionNotExist"
)

var (
	_StatusCode_index_0 = [...]uint8{0, 2, 7, 20, 32, 45, 58, 77, 100, 107, 118, 128, 132}
	_StatusCode_index_1 = [...]uint8{0, 24, 46, 59, 85, 101, 113}
	_StatusCode_index_3 = [...]uint8{0, 23, 35}
	_StatusCode_index_4 = [...]uint8{0, 25, 43}
)

func (i StatusCode) String() string {
	switch {
	case 10000 <= i && i <= 10011:
		i -= 10000
		return _StatusCode_name_0[_StatusCode_index_0[i]:_StatusCode_index_0[i+1]]
	case 20000 <= i && i <= 20005:
		i -= 20000
		return _StatusCode_name_1[_StatusCode_index_1[i]:_StatusCode_index_1[i+1]]
	case i == 30000:
		return _StatusCode_name_2
	case 40000 <= i && i <= 40001:
		i -= 40000
		return _StatusCode_name_3[_StatusCode_index_3[i]:_StatusCode_index_3[i+1]]
	case 50000 <= i && i <= 50001:
		i -= 50000
		return _StatusCode_name_4[_StatusCode_index_4[i]:_StatusCode_index_4[i+1]]
	default:
		return "StatusCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
