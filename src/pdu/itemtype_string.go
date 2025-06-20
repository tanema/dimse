// Code generated by "stringer -type ItemType -trimprefix ItemType"; DO NOT EDIT.

package pdu

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ItemTypeApplicationContext-16]
	_ = x[ItemTypePresentationContextRequest-32]
	_ = x[ItemTypePresentationContextResponse-33]
	_ = x[ItemTypeAbstractSyntax-48]
	_ = x[ItemTypeTransferSyntax-64]
	_ = x[ItemTypeUserInformation-80]
	_ = x[ItemTypeUserInformationMaximumLength-81]
	_ = x[ItemTypeImplementationClassUID-82]
	_ = x[ItemTypeAsynchronousOperationsWindow-83]
	_ = x[ItemTypeRoleSelection-84]
	_ = x[ItemTypeImplementationVersionName-85]
}

const (
	_ItemType_name_0 = "ApplicationContext"
	_ItemType_name_1 = "PresentationContextRequestPresentationContextResponse"
	_ItemType_name_2 = "AbstractSyntax"
	_ItemType_name_3 = "TransferSyntax"
	_ItemType_name_4 = "UserInformationUserInformationMaximumLengthImplementationClassUIDAsynchronousOperationsWindowRoleSelectionImplementationVersionName"
)

var (
	_ItemType_index_1 = [...]uint8{0, 26, 53}
	_ItemType_index_4 = [...]uint8{0, 15, 43, 65, 93, 106, 131}
)

func (i ItemType) String() string {
	switch {
	case i == 16:
		return _ItemType_name_0
	case 32 <= i && i <= 33:
		i -= 32
		return _ItemType_name_1[_ItemType_index_1[i]:_ItemType_index_1[i+1]]
	case i == 48:
		return _ItemType_name_2
	case i == 64:
		return _ItemType_name_3
	case 80 <= i && i <= 85:
		i -= 80
		return _ItemType_name_4[_ItemType_index_4[i]:_ItemType_index_4[i+1]]
	default:
		return "ItemType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
