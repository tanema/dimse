package tags

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

	commandGroupLengthInfo                   = tag.Info{Tag: CommandGroupLength, VRs: []string{"UL"}, VM: "1", Name: "Command Group Length", Keyword: "CommandGroupLength"}
	commandFieldInfo                         = tag.Info{Tag: CommandField, VRs: []string{"US"}, VM: "1", Name: "Command Field", Keyword: "CommandField"}
	messageIDInfo                            = tag.Info{Tag: MessageID, VRs: []string{"US"}, VM: "1", Name: "Message ID", Keyword: "MessageID"}
	commandDataSetTypeInfo                   = tag.Info{Tag: CommandDataSetType, VRs: []string{"US"}, VM: "1", Name: "CommandDataSetType", Keyword: "CommandDataSetType"}
	messageIDBeingRespondedToInfo            = tag.Info{Tag: MessageIDBeingRespondedTo, VRs: []string{"US"}, VM: "1", Name: "Message ID Being Responded To", Keyword: "MessageIDBeingRespondedTo"}
	affectedSOPClassUIDInfo                  = tag.Info{Tag: AffectedSOPClassUID, VRs: []string{"UI"}, VM: "1", Name: "AffectedSOPClassUID", Keyword: "AffectedSOPClassUID"}
	priorityInfo                             = tag.Info{Tag: Priority, VRs: []string{"US"}, VM: "1", Name: "Priority", Keyword: "Priority"}
	moveDestinationInfo                      = tag.Info{Tag: MoveDestination, VRs: []string{"AE"}, VM: "1", Name: "MoveDestination", Keyword: "MoveDestination"}
	numberOfRemainingSuboperationsInfo       = tag.Info{Tag: NumberOfRemainingSuboperations, VRs: []string{"US"}, VM: "1", Name: "NumberOfRemainingSuboperations", Keyword: "NumberOfRemainingSuboperations"}
	numberOfCompletedSuboperationsInfo       = tag.Info{Tag: NumberOfCompletedSuboperations, VRs: []string{"US"}, VM: "1", Name: "NumberOfCompletedSuboperations", Keyword: "NumberOfCompletedSuboperations"}
	numberOfFailedSuboperationsInfo          = tag.Info{Tag: NumberOfFailedSuboperations, VRs: []string{"US"}, VM: "1", Name: "NumberOfFailedSuboperations", Keyword: "NumberOfFailedSuboperations"}
	numberOfWarningSuboperationsInfo         = tag.Info{Tag: NumberOfWarningSuboperations, VRs: []string{"US"}, VM: "1", Name: "NumberOfWarningSuboperations", Keyword: "NumberOfWarningSuboperations"}
	affectedSOPInstanceUIDInfo               = tag.Info{Tag: AffectedSOPInstanceUID, VRs: []string{"UI"}, VM: "1", Name: "AffectedSOPInstanceUID", Keyword: "AffectedSOPInstanceUID"}
	moveOriginatorApplicationEntityTitleInfo = tag.Info{Tag: MoveOriginatorApplicationEntityTitle, VRs: []string{"AE"}, VM: "1", Name: "MoveOriginatorApplicationEntityTitle", Keyword: "MoveOriginatorApplicationEntityTitle"}
	moveOriginatorMessageIDInfo              = tag.Info{Tag: MoveOriginatorMessageID, VRs: []string{"US"}, VM: "1", Name: "MoveOriginatorMessageID", Keyword: "MoveOriginatorMessageID"}
	statusTagInfo                            = tag.Info{Tag: StatusTag, VRs: []string{"US"}, VM: "1", Name: "StatusTag", Keyword: "StatusTag"}
	errorCommentInfo                         = tag.Info{Tag: ErrorComment, VRs: []string{"LO"}, VM: "1", Name: "ErrorComment", Keyword: "ErrorComment"}
)

func init() {
	tag.Add(commandGroupLengthInfo, true)
	tag.Add(commandFieldInfo, true)
	tag.Add(messageIDInfo, true)
	tag.Add(commandDataSetTypeInfo, true)
	tag.Add(messageIDBeingRespondedToInfo, true)
	tag.Add(affectedSOPClassUIDInfo, true)
	tag.Add(priorityInfo, true)
	tag.Add(moveDestinationInfo, true)
	tag.Add(numberOfRemainingSuboperationsInfo, true)
	tag.Add(numberOfCompletedSuboperationsInfo, true)
	tag.Add(numberOfFailedSuboperationsInfo, true)
	tag.Add(numberOfWarningSuboperationsInfo, true)
	tag.Add(affectedSOPInstanceUIDInfo, true)
	tag.Add(moveOriginatorApplicationEntityTitleInfo, true)
	tag.Add(moveOriginatorMessageIDInfo, true)
	tag.Add(statusTagInfo, true)
	tag.Add(errorCommentInfo, true)
}
