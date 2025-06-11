package dimse

import (
	"github.com/suyashkumar/dicom/pkg/tag"
)

var (
	CommandGroupLength                   = tag.Tag{Group: 0x0000, Element: 0x0000}
	CommandField                         = tag.Tag{Group: 0x0000, Element: 0x0100}
	MessageID                            = tag.Tag{Group: 0x0000, Element: 0x0110}
	CommandDataSetType                   = tag.Tag{Group: 0x0000, Element: 0x0800}
	MessageIDBeingRespondedTo            = tag.Tag{Group: 0x0000, Element: 0x0120}
	AffectedSOPClassUID                  = tag.Tag{Group: 0x0000, Element: 0x0002}
	Priority                             = tag.Tag{Group: 0x0000, Element: 0x0700}
	MoveDestination                      = tag.Tag{Group: 0x0000, Element: 0x0600}
	NumberOfRemainingSuboperations       = tag.Tag{Group: 0x0000, Element: 0x1020}
	NumberOfCompletedSuboperations       = tag.Tag{Group: 0x0000, Element: 0x1021}
	NumberOfFailedSuboperations          = tag.Tag{Group: 0x0000, Element: 0x1022}
	NumberOfWarningSuboperations         = tag.Tag{Group: 0x0000, Element: 0x1023}
	AffectedSOPInstanceUID               = tag.Tag{Group: 0x0000, Element: 0x1000}
	MoveOriginatorApplicationEntityTitle = tag.Tag{Group: 0x0000, Element: 0x1030}
	MoveOriginatorMessageID              = tag.Tag{Group: 0x0000, Element: 0x1031}
	StatusTag                            = tag.Tag{Group: 0x0000, Element: 0x0900}
	ErrorComment                         = tag.Tag{Group: 0x0000, Element: 0x0902}

	CommandGroupLengthInfo                   = tag.Info{Tag: CommandGroupLength, VRs: []string{"UL"}, VM: "1", Name: "Command Group Length", Keyword: "CommandGroupLength"}
	CommandFieldInfo                         = tag.Info{Tag: CommandField, VRs: []string{"US"}, VM: "1", Name: "Command Field", Keyword: "CommandField"}
	MessageIDInfo                            = tag.Info{Tag: MessageID, VRs: []string{"US"}, VM: "1", Name: "Message ID", Keyword: "MessageID"}
	CommandDataSetTypeInfo                   = tag.Info{Tag: CommandDataSetType, VRs: []string{"US"}, VM: "1", Name: "CommandDataSetType", Keyword: "CommandDataSetType"}
	MessageIDBeingRespondedToInfo            = tag.Info{Tag: MessageIDBeingRespondedTo, VRs: []string{"US"}, VM: "1", Name: "Message ID Being Responded To", Keyword: "MessageIDBeingRespondedTo"}
	AffectedSOPClassUIDInfo                  = tag.Info{Tag: AffectedSOPClassUID, VRs: []string{"UI"}, VM: "1", Name: "AffectedSOPClassUID", Keyword: "AffectedSOPClassUID"}
	PriorityInfo                             = tag.Info{Tag: Priority, VRs: []string{"US"}, VM: "1", Name: "Priority", Keyword: "Priority"}
	MoveDestinationInfo                      = tag.Info{Tag: MoveDestination, VRs: []string{"AE"}, VM: "1", Name: "MoveDestination", Keyword: "MoveDestination"}
	NumberOfRemainingSuboperationsInfo       = tag.Info{Tag: NumberOfRemainingSuboperations, VRs: []string{"US"}, VM: "1", Name: "NumberOfRemainingSuboperations", Keyword: "NumberOfRemainingSuboperations"}
	NumberOfCompletedSuboperationsInfo       = tag.Info{Tag: NumberOfCompletedSuboperations, VRs: []string{"US"}, VM: "1", Name: "NumberOfCompletedSuboperations", Keyword: "NumberOfCompletedSuboperations"}
	NumberOfFailedSuboperationsInfo          = tag.Info{Tag: NumberOfFailedSuboperations, VRs: []string{"US"}, VM: "1", Name: "NumberOfFailedSuboperations", Keyword: "NumberOfFailedSuboperations"}
	NumberOfWarningSuboperationsInfo         = tag.Info{Tag: NumberOfWarningSuboperations, VRs: []string{"US"}, VM: "1", Name: "NumberOfWarningSuboperations", Keyword: "NumberOfWarningSuboperations"}
	AffectedSOPInstanceUIDInfo               = tag.Info{Tag: AffectedSOPInstanceUID, VRs: []string{"UI"}, VM: "1", Name: "AffectedSOPInstanceUID", Keyword: "AffectedSOPInstanceUID"}
	MoveOriginatorApplicationEntityTitleInfo = tag.Info{Tag: MoveOriginatorApplicationEntityTitle, VRs: []string{"AE"}, VM: "1", Name: "MoveOriginatorApplicationEntityTitle", Keyword: "MoveOriginatorApplicationEntityTitle"}
	MoveOriginatorMessageIDInfo              = tag.Info{Tag: MoveOriginatorMessageID, VRs: []string{"US"}, VM: "1", Name: "MoveOriginatorMessageID", Keyword: "MoveOriginatorMessageID"}
	StatusTagInfo                            = tag.Info{Tag: StatusTag, VRs: []string{"US"}, VM: "1", Name: "StatusTag", Keyword: "StatusTag"}
	ErrorCommentInfo                         = tag.Info{Tag: ErrorComment, VRs: []string{"LO"}, VM: "1", Name: "ErrorComment", Keyword: "ErrorComment"}
)

func init() {
	tag.Add(CommandGroupLengthInfo, true)
	tag.Add(CommandFieldInfo, true)
	tag.Add(MessageIDInfo, true)
	tag.Add(CommandDataSetTypeInfo, true)
	tag.Add(MessageIDBeingRespondedToInfo, true)
	tag.Add(AffectedSOPClassUIDInfo, true)
	tag.Add(PriorityInfo, true)
	tag.Add(MoveDestinationInfo, true)
	tag.Add(NumberOfRemainingSuboperationsInfo, true)
	tag.Add(NumberOfCompletedSuboperationsInfo, true)
	tag.Add(NumberOfFailedSuboperationsInfo, true)
	tag.Add(NumberOfWarningSuboperationsInfo, true)
	tag.Add(AffectedSOPInstanceUIDInfo, true)
	tag.Add(MoveOriginatorApplicationEntityTitleInfo, true)
	tag.Add(MoveOriginatorMessageIDInfo, true)
	tag.Add(StatusTagInfo, true)
	tag.Add(ErrorCommentInfo, true)
}
