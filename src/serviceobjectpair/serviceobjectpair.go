package serviceobjectpair

type ServiceObjectPair string

const (
	ProceduralEventLogging                                             = "1.2.840.10008.1.40"
	SubstanceAdministrationLogging                                     = "1.2.840.10008.1.42"
	ModalityWorklistInformationFind                                    = "1.2.840.10008.5.1.4.31"
	ColorPaletteInformationModelFind                                   = "1.2.840.10008.5.1.4.39.2"
	ColorPaletteInformationModelMove                                   = "1.2.840.10008.5.1.4.39.3"
	ColorPaletteInformationModelGet                                    = "1.2.840.10008.5.1.4.39.4"
	DefinedProcedureProtocolInformationModelFind                       = "1.2.840.10008.5.1.4.20.1"
	DefinedProcedureProtocolInformationModelMove                       = "1.2.840.10008.5.1.4.20.2"
	DefinedProcedureProtocolInformationModelGet                        = "1.2.840.10008.5.1.4.20.3"
	DisplaySystem                                                      = "1.2.840.10008.5.1.1.40"
	HangingProtocolInformationModelFind                                = "1.2.840.10008.5.1.4.38.2"
	HangingProtocolInformationModelMove                                = "1.2.840.10008.5.1.4.38.3"
	HangingProtocolInformationModelGet                                 = "1.2.840.10008.5.1.4.38.4"
	GenericImplantTemplateInformationModelFind                         = "1.2.840.10008.5.1.4.43.2"
	GenericImplantTemplateInformationModelMove                         = "1.2.840.10008.5.1.4.43.3"
	GenericImplantTemplateInformationModelGet                          = "1.2.840.10008.5.1.4.43.4"
	ImplantAssemblyTemplateInformationModelFind                        = "1.2.840.10008.5.1.4.44.2"
	ImplantAssemblyTemplateInformationModelMove                        = "1.2.840.10008.5.1.4.44.3"
	ImplantAssemblyTemplateInformationModelGet                         = "1.2.840.10008.5.1.4.44.4"
	ImplantTemplateGroupInformationModelFind                           = "1.2.840.10008.5.1.4.45.2"
	ImplantTemplateGroupInformationModelMove                           = "1.2.840.10008.5.1.4.45.3"
	ImplantTemplateGroupInformationModelGet                            = "1.2.840.10008.5.1.4.45.4"
	InstanceAvailabilityNotification                                   = "1.2.840.10008.5.1.4.33"
	InventoryFind                                                      = "1.2.840.10008.5.1.4.1.1.201.2"
	InventoryMove                                                      = "1.2.840.10008.5.1.4.1.1.201.3"
	InventoryGet                                                       = "1.2.840.10008.5.1.4.1.1.201.4"
	MediaCreationManagement                                            = "1.2.840.10008.5.1.1.33"
	MediaStorageDirectoryStorage                                       = "1.2.840.10008.1.3.10"
	HangingProtocolStorage                                             = "1.2.840.10008.5.1.4.38.1"
	ColorPaletteStorage                                                = "1.2.840.10008.5.1.4.39.1"
	GenericImplantTemplateStorage                                      = "1.2.840.10008.5.1.4.43.1"
	ImplantAssemblyTemplateStorage                                     = "1.2.840.10008.5.1.4.44.1"
	ImplantTemplateGroupStorage                                        = "1.2.840.10008.5.1.4.45.1"
	CTDefinedProcedureProtocolStorage                                  = "1.2.840.10008.5.1.4.1.1.200.1"
	ProtocolApprovalStorage                                            = "1.2.840.10008.5.1.4.1.1.200.3"
	XADefinedProcedureProtocolStorage                                  = "1.2.840.10008.5.1.4.1.1.200.7"
	InventoryStorage                                                   = "1.2.840.10008.5.1.4.1.1.201.1"
	BasicFilmSession                                                   = "1.2.840.10008.5.1.1.1"
	BasicFilmBox                                                       = "1.2.840.10008.5.1.1.2"
	BasicGrayscaleImageBox                                             = "1.2.840.10008.5.1.1.4"
	BasicColorImageBox                                                 = "1.2.840.10008.5.1.1.4.1"
	PrintJob                                                           = "1.2.840.10008.5.1.1.14"
	BasicAnnotationBox                                                 = "1.2.840.10008.5.1.1.15"
	Printer                                                            = "1.2.840.10008.5.1.1.16"
	PrinterConfigurationRetrieval                                      = "1.2.840.10008.5.1.1.16.376"
	PresentationLUT                                                    = "1.2.840.10008.5.1.1.23"
	BasicGrayscalePrintManagementMeta                                  = "1.2.840.10008.5.1.1.9"
	BasicColorPrintManagementMeta                                      = "1.2.840.10008.5.1.1.18"
	ModalityPerformedProcedureStep                                     = "1.2.840.10008.3.1.2.3.3"
	ModalityPerformedProcedureStepRetrieve                             = "1.2.840.10008.3.1.2.3.4"
	ModalityPerformedProcedureStepNotification                         = "1.2.840.10008.3.1.2.3.5"
	ProtocolApprovalInformationModelFind                               = "1.2.840.10008.5.1.4.1.1.200.4"
	ProtocolApprovalInformationModelMove                               = "1.2.840.10008.5.1.4.1.1.200.5"
	ProtocolApprovalInformationModelGet                                = "1.2.840.10008.5.1.4.1.1.200.6"
	PatientRootQueryRetrieveInformationModelFind                       = "1.2.840.10008.5.1.4.1.2.1.1"
	PatientRootQueryRetrieveInformationModelMove                       = "1.2.840.10008.5.1.4.1.2.1.2"
	PatientRootQueryRetrieveInformationModelGet                        = "1.2.840.10008.5.1.4.1.2.1.3"
	StudyRootQueryRetrieveInformationModelFind                         = "1.2.840.10008.5.1.4.1.2.2.1"
	StudyRootQueryRetrieveInformationModelMove                         = "1.2.840.10008.5.1.4.1.2.2.2"
	StudyRootQueryRetrieveInformationModelGet                          = "1.2.840.10008.5.1.4.1.2.2.3"
	PatientStudyOnlyQueryRetrieveInformationModelFind                  = "1.2.840.10008.5.1.4.1.2.3.1"
	PatientStudyOnlyQueryRetrieveInformationModelMove                  = "1.2.840.10008.5.1.4.1.2.3.2"
	PatientStudyOnlyQueryRetrieveInformationModelGet                   = "1.2.840.10008.5.1.4.1.2.3.3"
	CompositeInstanceRootRetrieveMove                                  = "1.2.840.10008.5.1.4.1.2.4.2"
	CompositeInstanceRootRetrieveGet                                   = "1.2.840.10008.5.1.4.1.2.4.3"
	CompositeInstanceRetrieveWithoutBulkDataGet                        = "1.2.840.10008.5.1.4.1.2.5.3"
	RepositoryQuery                                                    = "1.2.840.10008.5.1.4.1.1.201.6"
	GeneralRelevantPatientInformationQuery                             = "1.2.840.10008.5.1.4.37.1"
	BreastImagingRelevantPatientInformationQuery                       = "1.2.840.10008.5.1.4.37.2"
	CardiacRelevantPatientInformationQuery                             = "1.2.840.10008.5.1.4.37.3"
	RTConventionalMachineVerification                                  = "1.2.840.10008.5.1.4.34.8"
	RTIonMachineVerification                                           = "1.2.840.10008.5.1.4.34.9"
	ComputedRadiographyImageStorage                                    = "1.2.840.10008.5.1.4.1.1.1"     // A.2
	DigitalXRayImageStorageForPresentation                             = "1.2.840.10008.5.1.4.1.1.1.1"   // A.26
	DigitalXRayImageStorageForProcessing                               = "1.2.840.10008.5.1.4.1.1.1.1.1" // A.26
	DigitalMammographyXRayImageStorageForPresentation                  = "1.2.840.10008.5.1.4.1.1.1.2"   // A.27
	DigitalMammographyXRayImageStorageForProcessing                    = "1.2.840.10008.5.1.4.1.1.1.2.1" // A.27
	DigitalIntraOralXRayImageStorageForPresentation                    = "1.2.840.10008.5.1.4.1.1.1.3"   // A.28
	DigitalIntraOralXRayImageStorageForProcessing                      = "1.2.840.10008.5.1.4.1.1.1.3.1" // A.28
	CTImageStorage                                                     = "1.2.840.10008.5.1.4.1.1.2"     // A.3
	EnhancedCTImageStorage                                             = "1.2.840.10008.5.1.4.1.1.2.1"   // A.38
	LegacyConvertedEnhancedCTImageStorage                              = "1.2.840.10008.5.1.4.1.1.2.2"   // A.70
	UltrasoundMultiFrameImageStorage                                   = "1.2.840.10008.5.1.4.1.1.3.1"   // A.7
	MRImageStorage                                                     = "1.2.840.10008.5.1.4.1.1.4"     // A.4
	EnhancedMRImageStorage                                             = "1.2.840.10008.5.1.4.1.1.4.1"   // A.36.2
	MRSpectroscopyStorage                                              = "1.2.840.10008.5.1.4.1.1.4.2"   // A.36.3
	EnhancedMRColorImageStorage                                        = "1.2.840.10008.5.1.4.1.1.4.3"   // A.36.4
	LegacyConvertedEnhancedMRImageStorage                              = "1.2.840.10008.5.1.4.1.1.4.4"   // A.71
	UltrasoundImageStorage                                             = "1.2.840.10008.5.1.4.1.1.6.1"   // A.6
	EnhancedUSVolumeStorage                                            = "1.2.840.10008.5.1.4.1.1.6.2"   // A.59
	PhotoacousticImageStorage                                          = "1.2.840.10008.5.1.4.1.1.6.3"
	SecondaryCaptureImageStorage                                       = "1.2.840.10008.5.1.4.1.1.7"     // A.8.1
	MultiFrameSingleBitSecondaryCaptureImageStorage                    = "1.2.840.10008.5.1.4.1.1.7.1"   // A.8.2
	MultiFrameGrayscaleByteSecondaryCaptureImageStorage                = "1.2.840.10008.5.1.4.1.1.7.2"   // A.8.3
	MultiFrameGrayscaleWordSecondaryCaptureImageStorage                = "1.2.840.10008.5.1.4.1.1.7.3"   // A.8.4
	MultiFrameTrueColorSecondaryCaptureImageStorage                    = "1.2.840.10008.5.1.4.1.1.7.4"   // A.8.5
	TwelveLeadECGWaveformStorage                                       = "1.2.840.10008.5.1.4.1.1.9.1.1" // A.34.3
	GeneralECGWaveformStorage                                          = "1.2.840.10008.5.1.4.1.1.9.1.2" // A.34.4
	AmbulatoryECGWaveformStorage                                       = "1.2.840.10008.5.1.4.1.1.9.1.3" // A.34.5
	General32bitECGWaveformStorage                                     = "1.2.840.10008.5.1.4.1.1.9.1.4"
	HemodynamicWaveformStorage                                         = "1.2.840.10008.5.1.4.1.1.9.2.1" // A.34.6
	CardiacElectrophysiologyWaveformStorage                            = "1.2.840.10008.5.1.4.1.1.9.3.1" // A.34.7
	BasicVoiceAudioWaveformStorage                                     = "1.2.840.10008.5.1.4.1.1.9.4.1" // A.34.2
	GeneralAudioWaveformStorage                                        = "1.2.840.10008.5.1.4.1.1.9.4.2" // A.34.10
	ArterialPulseWaveformStorage                                       = "1.2.840.10008.5.1.4.1.1.9.5.1" // A.34.8
	RespiratoryWaveformStorage                                         = "1.2.840.10008.5.1.4.1.1.9.6.1" // A.34.9
	MultichannelRespiratoryWaveformStorage                             = "1.2.840.10008.5.1.4.1.1.9.6.2" // A.34.16
	RoutineScalpElectroencephalogramWaveformStorage                    = "1.2.840.10008.5.1.4.1.1.9.7.1" // A.34.12
	ElectromyogramWaveformStorage                                      = "1.2.840.10008.5.1.4.1.1.9.7.2" // A.34.13
	ElectrooculogramWaveformStorage                                    = "1.2.840.10008.5.1.4.1.1.9.7.3" // A.34.14
	SleepElectroencephalogramWaveformStorage                           = "1.2.840.10008.5.1.4.1.1.9.7.4" // A.34.15
	BodyPositionWaveformStorage                                        = "1.2.840.10008.5.1.4.1.1.9.8.1" // A.34.17
	WaveformPresentationStateStorage                                   = "1.2.840.10008.5.1.4.1.1.9.100.1"
	WaveformAcquisitionPresentationStateStorage                        = "1.2.840.10008.5.1.4.1.1.9.100.2"
	GrayscaleSoftcopyPresentationStateStorage                          = "1.2.840.10008.5.1.4.1.1.11.1"  // A.33.1
	ColorSoftcopyPresentationStateStorage                              = "1.2.840.10008.5.1.4.1.1.11.2"  // A.33.2
	PseudoColorSoftcopyPresentationStageStorage                        = "1.2.840.10008.5.1.4.1.1.11.3"  // A.33.3
	BlendingSoftcopyPresentationStateStorage                           = "1.2.840.10008.5.1.4.1.1.11.4"  // A.33.4
	XAXRFGrayscaleSoftcopyPresentationStateStorage                     = "1.2.840.10008.5.1.4.1.1.11.5"  // A.33.6
	GrayscalePlanarMPRVolumetricPresentationStateStorage               = "1.2.840.10008.5.1.4.1.1.11.6"  // A.80.1
	CompositingPlanarMPRVolumetricPresentationStateStorage             = "1.2.840.10008.5.1.4.1.1.11.7"  // A.80.1
	AdvancedBlendingPresentationStateStorage                           = "1.2.840.10008.5.1.4.1.1.11.8"  // A.33.7
	VolumeRenderingVolumetricPresentationStateStorage                  = "1.2.840.10008.5.1.4.1.1.11.9"  // A.80.2
	SegmentedVolumeRenderingVolumetricPresentationStateStorage         = "1.2.840.10008.5.1.4.1.1.11.10" // A.80.2
	MultipleVolumeRenderingVolumetricPresentationStateStorage          = "1.2.840.10008.5.1.4.1.1.11.11" // A.80.2
	VariableModalityLUTSoftcopyPresentationStageStorage                = "1.2.840.10008.5.1.4.1.1.11.12"
	XRayAngiographicImageStorage                                       = "1.2.840.10008.5.1.4.1.1.12.1"   // A.14
	EnhancedXAImageStorage                                             = "1.2.840.10008.5.1.4.1.1.12.1.1" // A.47
	XRayRadiofluoroscopicImageStorage                                  = "1.2.840.10008.5.1.4.1.1.12.2"   // A.16
	EnhancedXRFImageStorage                                            = "1.2.840.10008.5.1.4.1.1.12.2.1" // A.48
	XRay3DAngiographicImageStorage                                     = "1.2.840.10008.5.1.4.1.1.13.1.1" // A.53
	XRay3DCraniofacialImageStorage                                     = "1.2.840.10008.5.1.4.1.1.13.1.2" // A.54
	BreastTomosynthesisImageStorage                                    = "1.2.840.10008.5.1.4.1.1.13.1.3" // A.55
	BreastProjectionXRayImageStorageForPresentation                    = "1.2.840.10008.5.1.4.1.1.13.1.4" // A.74
	BreastProjectionXRayImageStorageForProcessing                      = "1.2.840.10008.5.1.4.1.1.13.1.5" // A.74
	IntravascularOpticalCoherenceTomographyImageStorageForPresentation = "1.2.840.10008.5.1.4.1.1.14.1"   // A.66
	IntravascularOpticalCoherenceTomographyImageStorageForProcessing   = "1.2.840.10008.5.1.4.1.1.14.2"   // A.66
	NuclearMedicineImageStorage                                        = "1.2.840.10008.5.1.4.1.1.20"     // A.5
	ParametricMapStorage                                               = "1.2.840.10008.5.1.4.1.1.30"     // A.75
	RawDataStorage                                                     = "1.2.840.10008.5.1.4.1.1.66"     // A.37
	SpatialRegistrationStorage                                         = "1.2.840.10008.5.1.4.1.1.66.1"   // A.39.1
	SpatialFiducialsStorage                                            = "1.2.840.10008.5.1.4.1.1.66.2"   // A.40
	DeformableSpatialRegistrationStorage                               = "1.2.840.10008.5.1.4.1.1.66.3"   // A.39.2
	SegmentationStorage                                                = "1.2.840.10008.5.1.4.1.1.66.4"   // A.51
	SurfaceSegmentationStorage                                         = "1.2.840.10008.5.1.4.1.1.66.5"   // A.57
	TractographyResultsStorage                                         = "1.2.840.10008.5.1.4.1.1.66.6"   // A.78
	LabelMapSegmentationStorage                                        = "1.2.840.10008.5.1.4.1.1.66.7"
	HeightMapSegmentationStorage                                       = "1.2.840.10008.5.1.4.1.1.66.8"
	RealWorldValueMappingStorage                                       = "1.2.840.10008.5.1.4.1.1.67"       // A.46
	SurfaceScanMeshStorage                                             = "1.2.840.10008.5.1.4.1.1.68.1"     // A.68
	SurfaceScanPointCloudStorage                                       = "1.2.840.10008.5.1.4.1.1.68.2"     // A.69
	VLEndoscopicImageStorage                                           = "1.2.840.10008.5.1.4.1.1.77.1.1"   // A.32.1
	VideoEndoscopicImageStorage                                        = "1.2.840.10008.5.1.4.1.1.77.1.1.1" // A.32.5
	VLMicroscopicImageStorage                                          = "1.2.840.10008.5.1.4.1.1.77.1.2"   // A.32.2
	VideoMicroscopicImageStorage                                       = "1.2.840.10008.5.1.4.1.1.77.1.2.1" // A.32.6
	VLSlideCoordinatesMicroscopicImageStorage                          = "1.2.840.10008.5.1.4.1.1.77.1.3"   // A.32.3
	VLPhotographicImageStorage                                         = "1.2.840.10008.5.1.4.1.1.77.1.4"   // A.32.4
	VideoPhotographicImageStorage                                      = "1.2.840.10008.5.1.4.1.1.77.1.4.1" // A.32.7
	OphthalmicPhotography8BitImageStorage                              = "1.2.840.10008.5.1.4.1.1.77.1.5.1" // A.41
	OphthalmicPhotography16BitImageStorage                             = "1.2.840.10008.5.1.4.1.1.77.1.5.2" // A.42
	StereometricRelationshipStorage                                    = "1.2.840.10008.5.1.4.1.1.77.1.5.3" // A.43
	OphthalmicTomographyImageStorage                                   = "1.2.840.10008.5.1.4.1.1.77.1.5.4" // A.52
	WideFieldOphthalmicPhotographyStereographicProjectionImageStorage  = "1.2.840.10008.5.1.4.1.1.77.1.5.5" // A.76
	WideFieldOphthalmicPhotography3DCoordinatesImageStorage            = "1.2.840.10008.5.1.4.1.1.77.1.5.6" // A.77
	OphthalmicOpticalCoherenceTomographyEnFaceImageStorage             = "1.2.840.10008.5.1.4.1.1.77.1.5.7" // A.83
	OphthlamicOpticalCoherenceTomographyBscanVolumeAnalysisStorage     = "1.2.840.10008.5.1.4.1.1.77.1.5.8" // A.84
	VLWholeSlideMicroscopyImageStorage                                 = "1.2.840.10008.5.1.4.1.1.77.1.6"   // A.32.8
	DermoscopicPhotographyImageStorage                                 = "1.2.840.10008.5.1.4.1.1.77.1.7"   // A.32.11
	ConfocalMicroscopyImageStorage                                     = "1.2.840.10008.5.1.4.1.1.77.1.8"
	ConfocalMicroscopyTiledPyramidalImageStorage                       = "1.2.840.10008.5.1.4.1.1.77.1.9"
	LensometryMeasurementsStorage                                      = "1.2.840.10008.5.1.4.1.1.78.1"  // A.60.1
	AutorefractionMeasurementsStorage                                  = "1.2.840.10008.5.1.4.1.1.78.2"  // A.60.2
	KeratometryMeasurementsStorage                                     = "1.2.840.10008.5.1.4.1.1.78.3"  // A.60.3
	SubjectiveRefractionMeasurementsStorage                            = "1.2.840.10008.5.1.4.1.1.78.4"  // A.60.4
	VisualAcuityMeasurementsStorage                                    = "1.2.840.10008.5.1.4.1.1.78.5"  // A.60.5
	SpectaclePrescriptionReportStorage                                 = "1.2.840.10008.5.1.4.1.1.78.6"  // A.35.9
	OphthalmicAxialMeasurementsStorage                                 = "1.2.840.10008.5.1.4.1.1.78.7"  // A.60.6
	IntraocularLensCalculationsStorage                                 = "1.2.840.10008.5.1.4.1.1.78.8"  // A.60.7
	MacularGridThicknessAndVolumeReportStorage                         = "1.2.840.10008.5.1.4.1.1.79.1"  // A.35.11
	OphthalmicVisualFieldStaticPerimetryMeasurementsStorage            = "1.2.840.10008.5.1.4.1.1.80.1"  // A.65
	OphthalmicThicknessMapStorage                                      = "1.2.840.10008.5.1.4.1.1.81.1"  // A.67
	CornealTopographyMapStorage                                        = "1.2.840.10008.5.1.4.1.1.82.1"  // A.73
	BasicTextSRStorage                                                 = "1.2.840.10008.5.1.4.1.1.88.11" // A.35.1
	EnhancedSRStorage                                                  = "1.2.840.10008.5.1.4.1.1.88.22" // A.35.2
	ComprehensiveSRStorage                                             = "1.2.840.10008.5.1.4.1.1.88.33" // A.35.3
	Comprehensive3DSRStorage                                           = "1.2.840.10008.5.1.4.1.1.88.34" // A.35.13
	ExtensibleSRStorage                                                = "1.2.840.10008.5.1.4.1.1.88.35" // A.35.15
	ProcedureLogStorage                                                = "1.2.840.10008.5.1.4.1.1.88.40" // A.35.7
	MammographyCADSRStorage                                            = "1.2.840.10008.5.1.4.1.1.88.50" // A.35.5
	KeyObjectSelectionDocumentStorage                                  = "1.2.840.10008.5.1.4.1.1.88.59" // A.35.4
	ChestCADSRStorage                                                  = "1.2.840.10008.5.1.4.1.1.88.65" // A.35.6
	XRayRadiationDoseSRStorage                                         = "1.2.840.10008.5.1.4.1.1.88.67" // A.35.8
	RadiopharmaceuticalRadiationDoseSRStorage                          = "1.2.840.10008.5.1.4.1.1.88.68" // A.35.14
	ColonCADSRStorage                                                  = "1.2.840.10008.5.1.4.1.1.88.69" // A.35.10
	ImplantationPlanSRStorage                                          = "1.2.840.10008.5.1.4.1.1.88.70" // A.35.12
	AcquisitionContextSRStorage                                        = "1.2.840.10008.5.1.4.1.1.88.71" // A.35.16
	SimplifiedAdultEchoSRStorage                                       = "1.2.840.10008.5.1.4.1.1.88.72" // A.35.17
	PatientRadiationDoseSRStorage                                      = "1.2.840.10008.5.1.4.1.1.88.73" // A.35.18
	PlannedImagingAgentAdministrationSRStorage                         = "1.2.840.10008.5.1.4.1.1.88.74" // A.35.19
	PerformedImagingAgentAdministrationSRStorage                       = "1.2.840.10008.5.1.4.1.1.88.75" // A.35.20
	EnhancedXRayRadiationDoseSRStorage                                 = "1.2.840.10008.5.1.4.1.1.88.76" // A.35.
	WaveformAnnotationSRStorage                                        = "1.2.840.10008.5.1.4.1.1.88.77"
	ContentAssessmentResultsStorage                                    = "1.2.840.10008.5.1.4.1.1.90.1" // A.81
	MicroscopyBulkSimpleAnnotationsStorage                             = "1.2.840.10008.5.1.4.1.1.91.1"
	EncapsulatedPDFStorage                                             = "1.2.840.10008.5.1.4.1.1.104.1" // A.45.1
	EncapsulatedCDAStorage                                             = "1.2.840.10008.5.1.4.1.1.104.2" // A.45.2
	EncapsulatedSTLStorage                                             = "1.2.840.10008.5.1.4.1.1.104.3" // A.85.1
	EncapsulatedOBJStorage                                             = "1.2.840.10008.5.1.4.1.1.104.4" // A.85.2
	EncapsulatedMTLStorage                                             = "1.2.840.10008.5.1.4.1.1.104.5" // A.85.3
	PositronEmissionTomographyImageStorage                             = "1.2.840.10008.5.1.4.1.1.128"   // A.21
	LegacyConvertedEnhancedPETImageStorage                             = "1.2.840.10008.5.1.4.1.1.128.1" // A.72
	EnhancedPETImageStorage                                            = "1.2.840.10008.5.1.4.1.1.130"   // A.56
	BasicStructuredDisplayStorage                                      = "1.2.840.10008.5.1.4.1.1.131"   // A.33.5
	CTPerformedProcedureProtocolStorage                                = "1.2.840.10008.5.1.4.1.1.200.2" // A.82.1
	XAPerformedProcedureProtocolStorage                                = "1.2.840.10008.5.1.4.1.1.200.8"
	RTImageStorage                                                     = "1.2.840.10008.5.1.4.1.1.481.1"  // A.17
	RTDoseStorage                                                      = "1.2.840.10008.5.1.4.1.1.481.2"  // A.18
	RTStructureSetStorage                                              = "1.2.840.10008.5.1.4.1.1.481.3"  // A.19
	RTBeamsTreatmentRecordStorage                                      = "1.2.840.10008.5.1.4.1.1.481.4"  // A.29
	RTPlanStorage                                                      = "1.2.840.10008.5.1.4.1.1.481.5"  // A.20
	RTBrachyTreatmentRecordStorage                                     = "1.2.840.10008.5.1.4.1.1.481.6"  // A.20
	RTTreatmentSummaryRecordStorage                                    = "1.2.840.10008.5.1.4.1.1.481.7"  // A.31
	RTIonPlanStorage                                                   = "1.2.840.10008.5.1.4.1.1.481.8"  // A.49
	RTIonBeamsTreatmentRecordStorage                                   = "1.2.840.10008.5.1.4.1.1.481.9"  // A.50
	RTPhysicianIntentStorage                                           = "1.2.840.10008.5.1.4.1.1.481.10" // A.86.1.2
	RTSegmentAnnotationStorage                                         = "1.2.840.10008.5.1.4.1.1.481.11" // A.86.1.3
	RTRadiationSetStorage                                              = "1.2.840.10008.5.1.4.1.1.481.12" // A.86.1.4
	CArmPhotonElectronRadiationStorage                                 = "1.2.840.10008.5.1.4.1.1.481.13" // A.86.1.5
	TomotherapeuticRadiationStorage                                    = "1.2.840.10008.5.1.4.1.1.481.14" // A.86.1.6
	RoboticArmRadiationStorage                                         = "1.2.840.10008.5.1.4.1.1.481.15" // A.86.1.7
	RTRadiationRecordSetStorage                                        = "1.2.840.10008.5.1.4.1.1.481.16" // A.86.1.8
	RTRadiationSalvageRecordStorage                                    = "1.2.840.10008.5.1.4.1.1.481.17" // A.86.1.9
	TomotherapeuticRadiationRecordStorage                              = "1.2.840.10008.5.1.4.1.1.481.18" // A.86.1.10
	CArmPhotonElectronRadiationRecordStorage                           = "1.2.840.10008.5.1.4.1.1.481.19" // A.86.1.11
	RoboticArmRadiationRecordStorage                                   = "1.2.840.10008.5.1.4.1.1.481.20" // A.86.1.12
	RTRadiationSetDeliveryInstructionStorage                           = "1.2.840.10008.5.1.4.1.1.481.21"
	RTTreatmentPreparationStorage                                      = "1.2.840.10008.5.1.4.1.1.481.22"
	EnhancedRTImageStorage                                             = "1.2.840.10008.5.1.4.1.1.481.23"
	EnhancedContinuousRTImageStorage                                   = "1.2.840.10008.5.1.4.1.1.481.24"
	RTPatientPositionAcquisitionInstructionStorage                     = "1.2.840.10008.5.1.4.1.1.481.25"
	RTBeamsDeliveryInstructionStorage                                  = "1.2.840.10008.5.1.4.34.7"  // A.64
	RTBrachyApplicationSetupDeliveryInstructionsStorage                = "1.2.840.10008.5.1.4.34.10" // A.79
	StorageCommitmentPushModel                                         = "1.2.840.10008.1.20.1"
	InventoryCreation                                                  = "1.2.840.10008.5.1.4.1.1.201.5"
	ProductCharacteristicsQuery                                        = "1.2.840.10008.5.1.4.41"
	SubstanceApprovalQuery                                             = "1.2.840.10008.5.1.4.42"
	UnifiedProcedureStepPush                                           = "1.2.840.10008.5.1.4.34.6.1"
	UnifiedProcedureStepWatch                                          = "1.2.840.10008.5.1.4.34.6.2"
	UnifiedProcedureStepPull                                           = "1.2.840.10008.5.1.4.34.6.3"
	UnifiedProcedureStepEvent                                          = "1.2.840.10008.5.1.4.34.6.4"
	UnifiedProcedureStepQuery                                          = "1.2.840.10008.5.1.4.34.6.5"
	Verification                                                       = "1.2.840.10008.1.1"
)

var (
	AppEventClasses             = []string{ProceduralEventLogging, SubstanceAdministrationLogging}
	BasicWorklistClasses        = []string{ModalityWorklistInformationFind}
	ColorPaletteClasses         = []string{ColorPaletteInformationModelFind, ColorPaletteInformationModelMove, ColorPaletteInformationModelGet}
	DefinedProcedureClasses     = []string{DefinedProcedureProtocolInformationModelFind, DefinedProcedureProtocolInformationModelMove, DefinedProcedureProtocolInformationModelGet}
	DisplaySystemClasses        = []string{DisplaySystem}
	InstanceAvailabilityClasses = []string{InstanceAvailabilityNotification}
	HangingProtocolClasses      = []string{HangingProtocolInformationModelFind, HangingProtocolInformationModelMove, HangingProtocolInformationModelGet}
	ImplantTemplateClasses      = []string{GenericImplantTemplateInformationModelFind, GenericImplantTemplateInformationModelMove,
		GenericImplantTemplateInformationModelGet, ImplantAssemblyTemplateInformationModelFind, ImplantAssemblyTemplateInformationModelMove,
		ImplantAssemblyTemplateInformationModelGet, ImplantTemplateGroupInformationModelFind, ImplantTemplateGroupInformationModelMove,
		ImplantTemplateGroupInformationModelGet}
	InventoryClasses        = []string{InventoryFind, InventoryMove, InventoryGet}
	MediaCreationClasses    = []string{MediaCreationManagement}
	MediaStorageClasses     = []string{MediaStorageDirectoryStorage}
	NonPatientObjectClasses = []string{HangingProtocolStorage, ColorPaletteStorage, GenericImplantTemplateStorage,
		ImplantAssemblyTemplateStorage, ImplantTemplateGroupStorage, CTDefinedProcedureProtocolStorage, ProtocolApprovalStorage,
		XADefinedProcedureProtocolStorage, InventoryStorage}
	PrintManagementClasses = []string{BasicFilmSession, BasicFilmBox, BasicGrayscaleImageBox, BasicColorImageBox,
		PrintJob, BasicAnnotationBox, Printer, PrinterConfigurationRetrieval, PresentationLUT, BasicGrayscalePrintManagementMeta,
		BasicColorPrintManagementMeta}
	ProcedureStepClasses    = []string{ModalityPerformedProcedureStep, ModalityPerformedProcedureStepRetrieve, ModalityPerformedProcedureStepNotification}
	ProtocolApprovalClasses = []string{ProtocolApprovalInformationModelFind, ProtocolApprovalInformationModelMove, ProtocolApprovalInformationModelGet}
	QRClasses               = []string{
		PatientRootQueryRetrieveInformationModelFind,
		PatientRootQueryRetrieveInformationModelMove,
		PatientRootQueryRetrieveInformationModelGet,
		StudyRootQueryRetrieveInformationModelFind,
		StudyRootQueryRetrieveInformationModelMove,
		StudyRootQueryRetrieveInformationModelGet,
		PatientStudyOnlyQueryRetrieveInformationModelFind,
		PatientStudyOnlyQueryRetrieveInformationModelMove,
		PatientStudyOnlyQueryRetrieveInformationModelGet,
		CompositeInstanceRootRetrieveMove,
		CompositeInstanceRootRetrieveGet,
		CompositeInstanceRetrieveWithoutBulkDataGet,
		RepositoryQuery,
	}
	RelevantPatientQueryClasses = []string{
		GeneralRelevantPatientInformationQuery,
		BreastImagingRelevantPatientInformationQuery,
		CardiacRelevantPatientInformationQuery,
	}
	RTMachineVerificationClasses = []string{
		RTConventionalMachineVerification,
		RTIonMachineVerification,
	}
	StorageClasses = []string{
		ComputedRadiographyImageStorage,
		DigitalXRayImageStorageForPresentation,
		DigitalXRayImageStorageForProcessing,
		DigitalMammographyXRayImageStorageForPresentation,
		DigitalMammographyXRayImageStorageForProcessing,
		DigitalIntraOralXRayImageStorageForPresentation,
		DigitalIntraOralXRayImageStorageForProcessing,
		CTImageStorage,
		EnhancedCTImageStorage,
		LegacyConvertedEnhancedCTImageStorage,
		UltrasoundMultiFrameImageStorage,
		MRImageStorage,
		EnhancedMRImageStorage,
		MRSpectroscopyStorage,
		EnhancedMRColorImageStorage,
		LegacyConvertedEnhancedMRImageStorage,
		UltrasoundImageStorage,
		EnhancedUSVolumeStorage,
		PhotoacousticImageStorage,
		SecondaryCaptureImageStorage,
		MultiFrameSingleBitSecondaryCaptureImageStorage,
		MultiFrameGrayscaleByteSecondaryCaptureImageStorage,
		MultiFrameGrayscaleWordSecondaryCaptureImageStorage,
		MultiFrameTrueColorSecondaryCaptureImageStorage,
		TwelveLeadECGWaveformStorage,
		GeneralECGWaveformStorage,
		AmbulatoryECGWaveformStorage,
		General32bitECGWaveformStorage,
		HemodynamicWaveformStorage,
		CardiacElectrophysiologyWaveformStorage,
		BasicVoiceAudioWaveformStorage,
		GeneralAudioWaveformStorage,
		ArterialPulseWaveformStorage,
		RespiratoryWaveformStorage,
		MultichannelRespiratoryWaveformStorage,
		RoutineScalpElectroencephalogramWaveformStorage,
		ElectromyogramWaveformStorage,
		ElectrooculogramWaveformStorage,
		SleepElectroencephalogramWaveformStorage,
		BodyPositionWaveformStorage,
		WaveformPresentationStateStorage,
		WaveformAcquisitionPresentationStateStorage,
		GrayscaleSoftcopyPresentationStateStorage,
		ColorSoftcopyPresentationStateStorage,
		PseudoColorSoftcopyPresentationStageStorage,
		BlendingSoftcopyPresentationStateStorage,
		XAXRFGrayscaleSoftcopyPresentationStateStorage,
		GrayscalePlanarMPRVolumetricPresentationStateStorage,
		CompositingPlanarMPRVolumetricPresentationStateStorage,
		AdvancedBlendingPresentationStateStorage,
		VolumeRenderingVolumetricPresentationStateStorage,
		SegmentedVolumeRenderingVolumetricPresentationStateStorage,
		MultipleVolumeRenderingVolumetricPresentationStateStorage,
		VariableModalityLUTSoftcopyPresentationStageStorage,
		XRayAngiographicImageStorage,
		EnhancedXAImageStorage,
		XRayRadiofluoroscopicImageStorage,
		EnhancedXRFImageStorage,
		XRay3DAngiographicImageStorage,
		XRay3DCraniofacialImageStorage,
		BreastTomosynthesisImageStorage,
		BreastProjectionXRayImageStorageForPresentation,
		BreastProjectionXRayImageStorageForProcessing,
		IntravascularOpticalCoherenceTomographyImageStorageForPresentation,
		IntravascularOpticalCoherenceTomographyImageStorageForProcessing,
		NuclearMedicineImageStorage,
		ParametricMapStorage,
		RawDataStorage,
		SpatialRegistrationStorage,
		SpatialFiducialsStorage,
		DeformableSpatialRegistrationStorage,
		SegmentationStorage,
		SurfaceSegmentationStorage,
		TractographyResultsStorage,
		LabelMapSegmentationStorage,
		HeightMapSegmentationStorage,
		RealWorldValueMappingStorage,
		SurfaceScanMeshStorage,
		SurfaceScanPointCloudStorage,
		VLEndoscopicImageStorage,
		VideoEndoscopicImageStorage,
		VLMicroscopicImageStorage,
		VideoMicroscopicImageStorage,
		VLSlideCoordinatesMicroscopicImageStorage,
		VLPhotographicImageStorage,
		VideoPhotographicImageStorage,
		OphthalmicPhotography8BitImageStorage,
		OphthalmicPhotography16BitImageStorage,
		StereometricRelationshipStorage,
		OphthalmicTomographyImageStorage,
		WideFieldOphthalmicPhotographyStereographicProjectionImageStorage,
		WideFieldOphthalmicPhotography3DCoordinatesImageStorage,
		OphthalmicOpticalCoherenceTomographyEnFaceImageStorage,
		OphthlamicOpticalCoherenceTomographyBscanVolumeAnalysisStorage,
		VLWholeSlideMicroscopyImageStorage,
		DermoscopicPhotographyImageStorage,
		ConfocalMicroscopyImageStorage,
		ConfocalMicroscopyTiledPyramidalImageStorage,
		LensometryMeasurementsStorage,
		AutorefractionMeasurementsStorage,
		KeratometryMeasurementsStorage,
		SubjectiveRefractionMeasurementsStorage,
		VisualAcuityMeasurementsStorage,
		SpectaclePrescriptionReportStorage,
		OphthalmicAxialMeasurementsStorage,
		IntraocularLensCalculationsStorage,
		MacularGridThicknessAndVolumeReportStorage,
		OphthalmicVisualFieldStaticPerimetryMeasurementsStorage,
		OphthalmicThicknessMapStorage,
		CornealTopographyMapStorage,
		BasicTextSRStorage,
		EnhancedSRStorage,
		ComprehensiveSRStorage,
		Comprehensive3DSRStorage,
		ExtensibleSRStorage,
		ProcedureLogStorage,
		MammographyCADSRStorage,
		KeyObjectSelectionDocumentStorage,
		ChestCADSRStorage,
		XRayRadiationDoseSRStorage,
		RadiopharmaceuticalRadiationDoseSRStorage,
		ColonCADSRStorage,
		ImplantationPlanSRStorage,
		AcquisitionContextSRStorage,
		SimplifiedAdultEchoSRStorage,
		PatientRadiationDoseSRStorage,
		PlannedImagingAgentAdministrationSRStorage,
		PerformedImagingAgentAdministrationSRStorage,
		EnhancedXRayRadiationDoseSRStorage,
		WaveformAnnotationSRStorage,
		ContentAssessmentResultsStorage,
		MicroscopyBulkSimpleAnnotationsStorage,
		EncapsulatedPDFStorage,
		EncapsulatedCDAStorage,
		EncapsulatedSTLStorage,
		EncapsulatedOBJStorage,
		EncapsulatedMTLStorage,
		PositronEmissionTomographyImageStorage,
		LegacyConvertedEnhancedPETImageStorage,
		EnhancedPETImageStorage,
		BasicStructuredDisplayStorage,
		CTPerformedProcedureProtocolStorage,
		XAPerformedProcedureProtocolStorage,
		RTImageStorage,
		RTDoseStorage,
		RTStructureSetStorage,
		RTBeamsTreatmentRecordStorage,
		RTPlanStorage,
		RTBrachyTreatmentRecordStorage,
		RTTreatmentSummaryRecordStorage,
		RTIonPlanStorage,
		RTIonBeamsTreatmentRecordStorage,
		RTPhysicianIntentStorage,
		RTSegmentAnnotationStorage,
		RTRadiationSetStorage,
		CArmPhotonElectronRadiationStorage,
		TomotherapeuticRadiationStorage,
		RoboticArmRadiationStorage,
		RTRadiationRecordSetStorage,
		RTRadiationSalvageRecordStorage,
		TomotherapeuticRadiationRecordStorage,
		CArmPhotonElectronRadiationRecordStorage,
		RoboticArmRadiationRecordStorage,
		RTRadiationSetDeliveryInstructionStorage,
		RTTreatmentPreparationStorage,
		EnhancedRTImageStorage,
		EnhancedContinuousRTImageStorage,
		RTPatientPositionAcquisitionInstructionStorage,
		RTBeamsDeliveryInstructionStorage,
		RTBrachyApplicationSetupDeliveryInstructionsStorage,
	}
	StorageCommitmentClasses       = []string{StorageCommitmentPushModel}
	StorageManagementClasses       = []string{InventoryCreation}
	SubstanceAdministrationClasses = []string{ProductCharacteristicsQuery, SubstanceApprovalQuery}
	UnifiedProcedureStepClasses    = []string{UnifiedProcedureStepPush, UnifiedProcedureStepWatch, UnifiedProcedureStepPull, UnifiedProcedureStepEvent, UnifiedProcedureStepQuery}
	VerificationClasses            = []string{Verification}
	QRFindClasses                  = []string{
		PatientRootQueryRetrieveInformationModelFind,
		StudyRootQueryRetrieveInformationModelFind,
		PatientStudyOnlyQueryRetrieveInformationModelFind,
		ModalityWorklistInformationFind,
	}
	QRMoveClasses = []string{
		PatientRootQueryRetrieveInformationModelMove,
		StudyRootQueryRetrieveInformationModelMove,
		PatientStudyOnlyQueryRetrieveInformationModelMove,
	}
	QRGetClasses = append([]string{
		PatientRootQueryRetrieveInformationModelGet,
		StudyRootQueryRetrieveInformationModelGet,
		PatientStudyOnlyQueryRetrieveInformationModelGet,
	}, StorageClasses...)
	AllClasses = append([]string{Verification}, StorageClasses...)
)
