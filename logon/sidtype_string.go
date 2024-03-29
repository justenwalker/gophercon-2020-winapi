// Code generated by "stringer -type=SidType"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SidTypeUser-1]
	_ = x[SidTypeGroup-2]
	_ = x[SidTypeDomain-3]
	_ = x[SidTypeAlias-4]
	_ = x[SidTypeWellKnownGroup-5]
	_ = x[SidTypeDeletedAccount-6]
	_ = x[SidTypeInvalid-7]
	_ = x[SidTypeUnknown-8]
	_ = x[SidTypeComputer-9]
	_ = x[SidTypeLabel-10]
	_ = x[SidTypeLogonSession-11]
}

const _SidType_name = "SidTypeUserSidTypeGroupSidTypeDomainSidTypeAliasSidTypeWellKnownGroupSidTypeDeletedAccountSidTypeInvalidSidTypeUnknownSidTypeComputerSidTypeLabelSidTypeLogonSession"

var _SidType_index = [...]uint8{0, 11, 23, 36, 48, 69, 90, 104, 118, 133, 145, 164}

func (i SidType) String() string {
	i -= 1
	if i >= SidType(len(_SidType_index)-1) {
		return "SidType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _SidType_name[_SidType_index[i]:_SidType_index[i+1]]
}
