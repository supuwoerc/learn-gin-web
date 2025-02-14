// Code generated by "stringer -type=Prefix -linecomment -output prefix_string.go"; DO NOT EDIT.

package constant

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PermissionIdPrefix-1]
	_ = x[PermissionNamePrefix-2]
	_ = x[PermissionResourcePrefix-3]
	_ = x[RoleIdPrefix-4]
	_ = x[RoleNamePrefix-5]
	_ = x[SignUpEmailPrefix-6]
	_ = x[DepartmentIdPrefix-7]
	_ = x[DepartmentNamePrefix-8]
}

const _Prefix_name = "lock:permission:idlock:permission:namelock:permission:resourcelock:role:idlock:role:namelock:signup:emaillock:department:idlock:department:name"

var _Prefix_index = [...]uint8{0, 18, 38, 62, 74, 88, 105, 123, 143}

func (i Prefix) String() string {
	i -= 1
	if i < 0 || i >= Prefix(len(_Prefix_index)-1) {
		return "Prefix(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Prefix_name[_Prefix_index[i]:_Prefix_index[i+1]]
}
