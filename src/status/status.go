package status

//go:generate stringer -type Status
//go:generate stringer -type Level
type (
	Status int
	Level  int
)

const (
	Successful                          Status = 0x0000
	WarnReqOptionalAttrNotSupported     Status = 0x0001
	Cancel                              Status = 0xFE00
	Pending                             Status = 0xFF00
	Continue                            Status = 0xFF01
	FailNoSuchAttributeValue            Status = 0x0105
	FailInvalidAttributeValue           Status = 0x0106
	WarnAttributeListError              Status = 0x0107
	FailProcessingError                 Status = 0x0110
	FailDuplicateSOP                    Status = 0x0111
	FailNoSuchSOPInstance               Status = 0x0112
	FailNoSuchEventType                 Status = 0x0113
	FailNoSuchArgument                  Status = 0x0114
	FailInvalidArgumentValue            Status = 0x0115
	WarnAttributeValueOutOfRage         Status = 0x0116
	FailInvalidObjectInstance           Status = 0x0117
	FailNoSuchSOPClass                  Status = 0x0118
	FailClassInstanceConflict           Status = 0x0119
	FailMissingAttribute                Status = 0x0120
	FailMissingAttributeValue           Status = 0x0121
	FailSOPClassNotSupported            Status = 0x0122
	FailNoSuchAction                    Status = 0x0123
	FailNotAuthorized                   Status = 0x0124
	FailDuplicateInvocation             Status = 0x0210
	FailUnrecognizedOperation           Status = 0x0211
	FailMistypedArgument                Status = 0x0212
	FailResourceLimitation              Status = 0x0212
	FailMediaCreateAlreadRecv           Status = 0xA510
	FailOutOfResources                  Status = 0xA700
	FailOutOfResourcesMatches           Status = 0xA701
	FailOutOfResourcesSubOps            Status = 0xA702
	FailInvalidPriorRecordKey           Status = 0xA710
	FailMoveDestUnknown                 Status = 0xA801
	FailIdentifierDoesNotMatchSOPClass  Status = 0xA900
	FailNoFramesReqWereFound            Status = 0xAA00
	FailUnableToCreateNewObj            Status = 0xAA01
	FailUnableToExtractFrames           Status = 0xAA02
	FailTimeReqMismatchWithSOP          Status = 0xAA03
	FailInvalidRequest                  Status = 0xAA04
	WarnSubOpFailure                    Status = 0xB000
	WarnResponseLimitReached            Status = 0xB001
	WarnElementDiscarded                Status = 0xB006
	WarnDSDoesNotMatchSOPClass          Status = 0xB007
	WarnAttrListErr                     Status = 0xB010
	WarnSpecifiedFrameDoesntMatch       Status = 0xB101
	WarnEvtLoggedunderDiffUID           Status = 0xB102
	WarnIDsInconsistent                 Status = 0xB104
	WarnUPSCreatedWithModifications     Status = 0xB300
	WarnDeletionLockNotGranted          Status = 0xB301
	WarnUPSAlreadyCancelled             Status = 0xB304
	WarnCoercedInvalidValues            Status = 0xB305
	WarnUPSAlreadyComplete              Status = 0xB306
	WarnMemoryAllocNotSupported         Status = 0xB600
	WarnFilmSessPrintNotSupported       Status = 0xB601
	WarnFilmSessSOPInstNoImage          Status = 0xB602
	WarnFilmBoxSOPInstNoImage           Status = 0xB603
	WarnImgLargerThanBox                Status = 0xB604
	WarnReqDensityOutOfRange            Status = 0xB605
	WarnImgLargerThanBoxCropped         Status = 0xB609
	WarnImgCmdPrintLargerThanBox        Status = 0xB60A
	FailUnableToProcess                 Status = 0xC000
	FailMoreThanOneMatch                Status = 0xC100
	FailProcLoggingNotAvail             Status = 0xC101
	FailEventInfoDoesntMatchTmpl        Status = 0xC102
	FailCannotMatchEventToStudy         Status = 0xC103
	FailPatientCannotBeIdentified       Status = 0xC110
	FailUpdateOfMedAdminFailed          Status = 0xC111
	FailIDsInconsistentStudy            Status = 0xC104
	FailOpNotAuthedToAddEntry           Status = 0xC10E
	FailNoObjectInstance                Status = 0xC112
	FailUnableToSupportReq              Status = 0xC200
	FailMediaCreateReqAlreadyComplete   Status = 0xC201
	FailMediaCreateReqAlreadyInProgress Status = 0xC202
	FailCancellationDeniedForNoReason   Status = 0xC202
	FailREferencedFracGroupNotExist     Status = 0xC221
	FailNoBeamsExist                    Status = 0xC222
	FailSCUAlreadVerifying              Status = 0xC223
	FailRefBeamNumberNotFound           Status = 0xC224
	FailRefDeviceNotSupported           Status = 0xC225
	FailRefDeviceNotFound               Status = 0xC226
	FailNoSuchObjectInstance            Status = 0xC227
	FailUPSMayNoLongerBeUpdated         Status = 0xC300
	FailTransUIDNotProvided             Status = 0xC301
	FailUPSAlreadyInProgress            Status = 0xC302
	FailUPSMayOnlyBeScheduled           Status = 0xC303
	FailUPSNotMetFinalStateReq          Status = 0xC304
	FailSOPInstUIDNotExist              Status = 0xC307
	FailAETitleIsUnknownToSCP           Status = 0xC308
	FailUPSStateNotScheduled            Status = 0xC309
	FailUPSNotYetInProgress             Status = 0xC310
	FailUPSAlreadyComplete2             Status = 0xC311
	FailPerformerCannotBeContacted      Status = 0xC312
	FailPerformerChoosesNotToCancel     Status = 0xC313
	FailSpecifiedActionNotAppropriate   Status = 0xC314
	FailEventReportNotSupported         Status = 0xC315
	FailFilmSessNoFilmBox               Status = 0xC600
	FailPrintQueueFull                  Status = 0xC601
	FailPrintQueueFull2                 Status = 0xC602
	FailImageSizeLargerThanBox          Status = 0xC603
	FailOOMStoreImage                   Status = 0xC605
	FailCombinedImageTooLarge           Status = 0xC613
	FailExitingFilmBox                  Status = 0xC616

	Success Level = iota
	Warning
	Failure
)

// large sets that are not yet supported
// for _code in range(0xA700, 0xA7FF + 1): STORAGE_SERVICE_CLASS_STATUS[_code] = (STATUS_FAILURE, "Refused: Out of Resources")
// for _code in range(0xA900, 0xA9FF + 1): STORAGE_SERVICE_CLASS_STATUS[_code] = ( STATUS_FAILURE, "Data Set Does Not Match SOP Class",)
// for _code in range(0xC000, 0xCFFF + 1): QR_FIND_SERVICE_CLASS_STATUS[_code] = (STATUS_FAILURE, "Unable to Process")

func StatusLevel(code Status) Level {
	if (code >= 0xA000 && code < 0xB000) || (code >= 0xC000 && code < 0xD000) {
		return Failure
	} else if code == 0x0107 || code == 0x0116 || code == 0x0001 || (code >= 0xB000 && code < 0xC000) {
		return Warning
	}
	return Success
}
