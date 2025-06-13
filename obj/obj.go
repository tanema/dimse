package obj

import (
	"github.com/tanema/dimse/obj/serviceobjectpair"
	"github.com/tanema/dimse/obj/transfersyntax"
)

var (
	VerificationClasses = serviceobjectpair.VerificationClasses
	StorageClasses      = serviceobjectpair.StorageClasses
	QRFindClasses       = []string{
		serviceobjectpair.PatientRootQueryRetrieveInformationModelFind,
		serviceobjectpair.StudyRootQueryRetrieveInformationModelFind,
		serviceobjectpair.PatientStudyOnlyQueryRetrieveInformationModelFind,
		serviceobjectpair.ModalityWorklistInformationFind,
	}
	QRMoveClasses = []string{
		serviceobjectpair.PatientRootQueryRetrieveInformationModelMove,
		serviceobjectpair.StudyRootQueryRetrieveInformationModelMove,
		serviceobjectpair.PatientStudyOnlyQueryRetrieveInformationModelMove,
	}
	QRGetClasses = append([]string{
		serviceobjectpair.PatientRootQueryRetrieveInformationModelGet,
		serviceobjectpair.StudyRootQueryRetrieveInformationModelGet,
		serviceobjectpair.PatientStudyOnlyQueryRetrieveInformationModelGet,
	}, StorageClasses...)
	StandardTransferSyntaxes = []string{
		transfersyntax.TSImplicitVRLittleEndian,
		transfersyntax.TSExplicitVRLittleEndian,
		transfersyntax.TSExplicitVRBigEndian,
		transfersyntax.TSDeflatedExplicitVRLittleEndian,
	}
	AllClasses = append([]string{serviceobjectpair.Verification}, serviceobjectpair.StorageClasses...)
)
