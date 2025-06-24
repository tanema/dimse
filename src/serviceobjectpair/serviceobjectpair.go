package serviceobjectpair

type UID string

const (
	ProceduralEventLogging                                             UID = "1.2.840.10008.1.40"
	SubstanceAdministrationLogging                                     UID = "1.2.840.10008.1.42"
	ModalityWorklistInformationFind                                    UID = "1.2.840.10008.5.1.4.31"
	ColorPaletteInformationModelFind                                   UID = "1.2.840.10008.5.1.4.39.2"
	ColorPaletteInformationModelMove                                   UID = "1.2.840.10008.5.1.4.39.3"
	ColorPaletteInformationModelGet                                    UID = "1.2.840.10008.5.1.4.39.4"
	DefinedProcedureProtocolInformationModelFind                       UID = "1.2.840.10008.5.1.4.20.1"
	DefinedProcedureProtocolInformationModelMove                       UID = "1.2.840.10008.5.1.4.20.2"
	DefinedProcedureProtocolInformationModelGet                        UID = "1.2.840.10008.5.1.4.20.3"
	DisplaySystem                                                      UID = "1.2.840.10008.5.1.1.40"
	HangingProtocolInformationModelFind                                UID = "1.2.840.10008.5.1.4.38.2"
	HangingProtocolInformationModelMove                                UID = "1.2.840.10008.5.1.4.38.3"
	HangingProtocolInformationModelGet                                 UID = "1.2.840.10008.5.1.4.38.4"
	GenericImplantTemplateInformationModelFind                         UID = "1.2.840.10008.5.1.4.43.2"
	GenericImplantTemplateInformationModelMove                         UID = "1.2.840.10008.5.1.4.43.3"
	GenericImplantTemplateInformationModelGet                          UID = "1.2.840.10008.5.1.4.43.4"
	ImplantAssemblyTemplateInformationModelFind                        UID = "1.2.840.10008.5.1.4.44.2"
	ImplantAssemblyTemplateInformationModelMove                        UID = "1.2.840.10008.5.1.4.44.3"
	ImplantAssemblyTemplateInformationModelGet                         UID = "1.2.840.10008.5.1.4.44.4"
	ImplantTemplateGroupInformationModelFind                           UID = "1.2.840.10008.5.1.4.45.2"
	ImplantTemplateGroupInformationModelMove                           UID = "1.2.840.10008.5.1.4.45.3"
	ImplantTemplateGroupInformationModelGet                            UID = "1.2.840.10008.5.1.4.45.4"
	InstanceAvailabilityNotification                                   UID = "1.2.840.10008.5.1.4.33"
	InventoryFind                                                      UID = "1.2.840.10008.5.1.4.1.1.201.2"
	InventoryMove                                                      UID = "1.2.840.10008.5.1.4.1.1.201.3"
	InventoryGet                                                       UID = "1.2.840.10008.5.1.4.1.1.201.4"
	MediaCreationManagement                                            UID = "1.2.840.10008.5.1.1.33"
	MediaStorageDirectoryStorage                                       UID = "1.2.840.10008.1.3.10"
	HangingProtocolStorage                                             UID = "1.2.840.10008.5.1.4.38.1"
	ColorPaletteStorage                                                UID = "1.2.840.10008.5.1.4.39.1"
	GenericImplantTemplateStorage                                      UID = "1.2.840.10008.5.1.4.43.1"
	ImplantAssemblyTemplateStorage                                     UID = "1.2.840.10008.5.1.4.44.1"
	ImplantTemplateGroupStorage                                        UID = "1.2.840.10008.5.1.4.45.1"
	CTDefinedProcedureProtocolStorage                                  UID = "1.2.840.10008.5.1.4.1.1.200.1"
	ProtocolApprovalStorage                                            UID = "1.2.840.10008.5.1.4.1.1.200.3"
	XADefinedProcedureProtocolStorage                                  UID = "1.2.840.10008.5.1.4.1.1.200.7"
	InventoryStorage                                                   UID = "1.2.840.10008.5.1.4.1.1.201.1"
	BasicFilmSession                                                   UID = "1.2.840.10008.5.1.1.1"
	BasicFilmBox                                                       UID = "1.2.840.10008.5.1.1.2"
	BasicGrayscaleImageBox                                             UID = "1.2.840.10008.5.1.1.4"
	BasicColorImageBox                                                 UID = "1.2.840.10008.5.1.1.4.1"
	PrintJob                                                           UID = "1.2.840.10008.5.1.1.14"
	BasicAnnotationBox                                                 UID = "1.2.840.10008.5.1.1.15"
	Printer                                                            UID = "1.2.840.10008.5.1.1.16"
	PrinterConfigurationRetrieval                                      UID = "1.2.840.10008.5.1.1.16.376"
	PresentationLUT                                                    UID = "1.2.840.10008.5.1.1.23"
	BasicGrayscalePrintManagementMeta                                  UID = "1.2.840.10008.5.1.1.9"
	BasicColorPrintManagementMeta                                      UID = "1.2.840.10008.5.1.1.18"
	ModalityPerformedProcedureStep                                     UID = "1.2.840.10008.3.1.2.3.3"
	ModalityPerformedProcedureStepRetrieve                             UID = "1.2.840.10008.3.1.2.3.4"
	ModalityPerformedProcedureStepNotification                         UID = "1.2.840.10008.3.1.2.3.5"
	ProtocolApprovalInformationModelFind                               UID = "1.2.840.10008.5.1.4.1.1.200.4"
	ProtocolApprovalInformationModelMove                               UID = "1.2.840.10008.5.1.4.1.1.200.5"
	ProtocolApprovalInformationModelGet                                UID = "1.2.840.10008.5.1.4.1.1.200.6"
	PatientRootQueryRetrieveInformationModelFind                       UID = "1.2.840.10008.5.1.4.1.2.1.1"
	PatientRootQueryRetrieveInformationModelMove                       UID = "1.2.840.10008.5.1.4.1.2.1.2"
	PatientRootQueryRetrieveInformationModelGet                        UID = "1.2.840.10008.5.1.4.1.2.1.3"
	StudyRootQueryRetrieveInformationModelFind                         UID = "1.2.840.10008.5.1.4.1.2.2.1"
	StudyRootQueryRetrieveInformationModelMove                         UID = "1.2.840.10008.5.1.4.1.2.2.2"
	StudyRootQueryRetrieveInformationModelGet                          UID = "1.2.840.10008.5.1.4.1.2.2.3"
	PatientStudyOnlyQueryRetrieveInformationModelFind                  UID = "1.2.840.10008.5.1.4.1.2.3.1"
	PatientStudyOnlyQueryRetrieveInformationModelMove                  UID = "1.2.840.10008.5.1.4.1.2.3.2"
	PatientStudyOnlyQueryRetrieveInformationModelGet                   UID = "1.2.840.10008.5.1.4.1.2.3.3"
	CompositeInstanceRootRetrieveMove                                  UID = "1.2.840.10008.5.1.4.1.2.4.2"
	CompositeInstanceRootRetrieveGet                                   UID = "1.2.840.10008.5.1.4.1.2.4.3"
	CompositeInstanceRetrieveWithoutBulkDataGet                        UID = "1.2.840.10008.5.1.4.1.2.5.3"
	RepositoryQuery                                                    UID = "1.2.840.10008.5.1.4.1.1.201.6"
	GeneralRelevantPatientInformationQuery                             UID = "1.2.840.10008.5.1.4.37.1"
	BreastImagingRelevantPatientInformationQuery                       UID = "1.2.840.10008.5.1.4.37.2"
	CardiacRelevantPatientInformationQuery                             UID = "1.2.840.10008.5.1.4.37.3"
	RTConventionalMachineVerification                                  UID = "1.2.840.10008.5.1.4.34.8"
	RTIonMachineVerification                                           UID = "1.2.840.10008.5.1.4.34.9"
	ComputedRadiographyImageStorage                                    UID = "1.2.840.10008.5.1.4.1.1.1"     // A.2
	DigitalXRayImageStorageForPresentation                             UID = "1.2.840.10008.5.1.4.1.1.1.1"   // A.26
	DigitalXRayImageStorageForProcessing                               UID = "1.2.840.10008.5.1.4.1.1.1.1.1" // A.26
	DigitalMammographyXRayImageStorageForPresentation                  UID = "1.2.840.10008.5.1.4.1.1.1.2"   // A.27
	DigitalMammographyXRayImageStorageForProcessing                    UID = "1.2.840.10008.5.1.4.1.1.1.2.1" // A.27
	DigitalIntraOralXRayImageStorageForPresentation                    UID = "1.2.840.10008.5.1.4.1.1.1.3"   // A.28
	DigitalIntraOralXRayImageStorageForProcessing                      UID = "1.2.840.10008.5.1.4.1.1.1.3.1" // A.28
	CTImageStorage                                                     UID = "1.2.840.10008.5.1.4.1.1.2"     // A.3
	EnhancedCTImageStorage                                             UID = "1.2.840.10008.5.1.4.1.1.2.1"   // A.38
	LegacyConvertedEnhancedCTImageStorage                              UID = "1.2.840.10008.5.1.4.1.1.2.2"   // A.70
	UltrasoundMultiFrameImageStorage                                   UID = "1.2.840.10008.5.1.4.1.1.3.1"   // A.7
	MRImageStorage                                                     UID = "1.2.840.10008.5.1.4.1.1.4"     // A.4
	EnhancedMRImageStorage                                             UID = "1.2.840.10008.5.1.4.1.1.4.1"   // A.36.2
	MRSpectroscopyStorage                                              UID = "1.2.840.10008.5.1.4.1.1.4.2"   // A.36.3
	EnhancedMRColorImageStorage                                        UID = "1.2.840.10008.5.1.4.1.1.4.3"   // A.36.4
	LegacyConvertedEnhancedMRImageStorage                              UID = "1.2.840.10008.5.1.4.1.1.4.4"   // A.71
	UltrasoundImageStorage                                             UID = "1.2.840.10008.5.1.4.1.1.6.1"   // A.6
	EnhancedUSVolumeStorage                                            UID = "1.2.840.10008.5.1.4.1.1.6.2"   // A.59
	PhotoacousticImageStorage                                          UID = "1.2.840.10008.5.1.4.1.1.6.3"
	SecondaryCaptureImageStorage                                       UID = "1.2.840.10008.5.1.4.1.1.7"     // A.8.1
	MultiFrameSingleBitSecondaryCaptureImageStorage                    UID = "1.2.840.10008.5.1.4.1.1.7.1"   // A.8.2
	MultiFrameGrayscaleByteSecondaryCaptureImageStorage                UID = "1.2.840.10008.5.1.4.1.1.7.2"   // A.8.3
	MultiFrameGrayscaleWordSecondaryCaptureImageStorage                UID = "1.2.840.10008.5.1.4.1.1.7.3"   // A.8.4
	MultiFrameTrueColorSecondaryCaptureImageStorage                    UID = "1.2.840.10008.5.1.4.1.1.7.4"   // A.8.5
	TwelveLeadECGWaveformStorage                                       UID = "1.2.840.10008.5.1.4.1.1.9.1.1" // A.34.3
	GeneralECGWaveformStorage                                          UID = "1.2.840.10008.5.1.4.1.1.9.1.2" // A.34.4
	AmbulatoryECGWaveformStorage                                       UID = "1.2.840.10008.5.1.4.1.1.9.1.3" // A.34.5
	General32bitECGWaveformStorage                                     UID = "1.2.840.10008.5.1.4.1.1.9.1.4"
	HemodynamicWaveformStorage                                         UID = "1.2.840.10008.5.1.4.1.1.9.2.1" // A.34.6
	CardiacElectrophysiologyWaveformStorage                            UID = "1.2.840.10008.5.1.4.1.1.9.3.1" // A.34.7
	BasicVoiceAudioWaveformStorage                                     UID = "1.2.840.10008.5.1.4.1.1.9.4.1" // A.34.2
	GeneralAudioWaveformStorage                                        UID = "1.2.840.10008.5.1.4.1.1.9.4.2" // A.34.10
	ArterialPulseWaveformStorage                                       UID = "1.2.840.10008.5.1.4.1.1.9.5.1" // A.34.8
	RespiratoryWaveformStorage                                         UID = "1.2.840.10008.5.1.4.1.1.9.6.1" // A.34.9
	MultichannelRespiratoryWaveformStorage                             UID = "1.2.840.10008.5.1.4.1.1.9.6.2" // A.34.16
	RoutineScalpElectroencephalogramWaveformStorage                    UID = "1.2.840.10008.5.1.4.1.1.9.7.1" // A.34.12
	ElectromyogramWaveformStorage                                      UID = "1.2.840.10008.5.1.4.1.1.9.7.2" // A.34.13
	ElectrooculogramWaveformStorage                                    UID = "1.2.840.10008.5.1.4.1.1.9.7.3" // A.34.14
	SleepElectroencephalogramWaveformStorage                           UID = "1.2.840.10008.5.1.4.1.1.9.7.4" // A.34.15
	BodyPositionWaveformStorage                                        UID = "1.2.840.10008.5.1.4.1.1.9.8.1" // A.34.17
	WaveformPresentationStateStorage                                   UID = "1.2.840.10008.5.1.4.1.1.9.100.1"
	WaveformAcquisitionPresentationStateStorage                        UID = "1.2.840.10008.5.1.4.1.1.9.100.2"
	GrayscaleSoftcopyPresentationStateStorage                          UID = "1.2.840.10008.5.1.4.1.1.11.1"  // A.33.1
	ColorSoftcopyPresentationStateStorage                              UID = "1.2.840.10008.5.1.4.1.1.11.2"  // A.33.2
	PseudoColorSoftcopyPresentationStageStorage                        UID = "1.2.840.10008.5.1.4.1.1.11.3"  // A.33.3
	BlendingSoftcopyPresentationStateStorage                           UID = "1.2.840.10008.5.1.4.1.1.11.4"  // A.33.4
	XAXRFGrayscaleSoftcopyPresentationStateStorage                     UID = "1.2.840.10008.5.1.4.1.1.11.5"  // A.33.6
	GrayscalePlanarMPRVolumetricPresentationStateStorage               UID = "1.2.840.10008.5.1.4.1.1.11.6"  // A.80.1
	CompositingPlanarMPRVolumetricPresentationStateStorage             UID = "1.2.840.10008.5.1.4.1.1.11.7"  // A.80.1
	AdvancedBlendingPresentationStateStorage                           UID = "1.2.840.10008.5.1.4.1.1.11.8"  // A.33.7
	VolumeRenderingVolumetricPresentationStateStorage                  UID = "1.2.840.10008.5.1.4.1.1.11.9"  // A.80.2
	SegmentedVolumeRenderingVolumetricPresentationStateStorage         UID = "1.2.840.10008.5.1.4.1.1.11.10" // A.80.2
	MultipleVolumeRenderingVolumetricPresentationStateStorage          UID = "1.2.840.10008.5.1.4.1.1.11.11" // A.80.2
	VariableModalityLUTSoftcopyPresentationStageStorage                UID = "1.2.840.10008.5.1.4.1.1.11.12"
	XRayAngiographicImageStorage                                       UID = "1.2.840.10008.5.1.4.1.1.12.1"   // A.14
	EnhancedXAImageStorage                                             UID = "1.2.840.10008.5.1.4.1.1.12.1.1" // A.47
	XRayRadiofluoroscopicImageStorage                                  UID = "1.2.840.10008.5.1.4.1.1.12.2"   // A.16
	EnhancedXRFImageStorage                                            UID = "1.2.840.10008.5.1.4.1.1.12.2.1" // A.48
	XRay3DAngiographicImageStorage                                     UID = "1.2.840.10008.5.1.4.1.1.13.1.1" // A.53
	XRay3DCraniofacialImageStorage                                     UID = "1.2.840.10008.5.1.4.1.1.13.1.2" // A.54
	BreastTomosynthesisImageStorage                                    UID = "1.2.840.10008.5.1.4.1.1.13.1.3" // A.55
	BreastProjectionXRayImageStorageForPresentation                    UID = "1.2.840.10008.5.1.4.1.1.13.1.4" // A.74
	BreastProjectionXRayImageStorageForProcessing                      UID = "1.2.840.10008.5.1.4.1.1.13.1.5" // A.74
	IntravascularOpticalCoherenceTomographyImageStorageForPresentation UID = "1.2.840.10008.5.1.4.1.1.14.1"   // A.66
	IntravascularOpticalCoherenceTomographyImageStorageForProcessing   UID = "1.2.840.10008.5.1.4.1.1.14.2"   // A.66
	NuclearMedicineImageStorage                                        UID = "1.2.840.10008.5.1.4.1.1.20"     // A.5
	ParametricMapStorage                                               UID = "1.2.840.10008.5.1.4.1.1.30"     // A.75
	RawDataStorage                                                     UID = "1.2.840.10008.5.1.4.1.1.66"     // A.37
	SpatialRegistrationStorage                                         UID = "1.2.840.10008.5.1.4.1.1.66.1"   // A.39.1
	SpatialFiducialsStorage                                            UID = "1.2.840.10008.5.1.4.1.1.66.2"   // A.40
	DeformableSpatialRegistrationStorage                               UID = "1.2.840.10008.5.1.4.1.1.66.3"   // A.39.2
	SegmentationStorage                                                UID = "1.2.840.10008.5.1.4.1.1.66.4"   // A.51
	SurfaceSegmentationStorage                                         UID = "1.2.840.10008.5.1.4.1.1.66.5"   // A.57
	TractographyResultsStorage                                         UID = "1.2.840.10008.5.1.4.1.1.66.6"   // A.78
	LabelMapSegmentationStorage                                        UID = "1.2.840.10008.5.1.4.1.1.66.7"
	HeightMapSegmentationStorage                                       UID = "1.2.840.10008.5.1.4.1.1.66.8"
	RealWorldValueMappingStorage                                       UID = "1.2.840.10008.5.1.4.1.1.67"       // A.46
	SurfaceScanMeshStorage                                             UID = "1.2.840.10008.5.1.4.1.1.68.1"     // A.68
	SurfaceScanPointCloudStorage                                       UID = "1.2.840.10008.5.1.4.1.1.68.2"     // A.69
	VLEndoscopicImageStorage                                           UID = "1.2.840.10008.5.1.4.1.1.77.1.1"   // A.32.1
	VideoEndoscopicImageStorage                                        UID = "1.2.840.10008.5.1.4.1.1.77.1.1.1" // A.32.5
	VLMicroscopicImageStorage                                          UID = "1.2.840.10008.5.1.4.1.1.77.1.2"   // A.32.2
	VideoMicroscopicImageStorage                                       UID = "1.2.840.10008.5.1.4.1.1.77.1.2.1" // A.32.6
	VLSlideCoordinatesMicroscopicImageStorage                          UID = "1.2.840.10008.5.1.4.1.1.77.1.3"   // A.32.3
	VLPhotographicImageStorage                                         UID = "1.2.840.10008.5.1.4.1.1.77.1.4"   // A.32.4
	VideoPhotographicImageStorage                                      UID = "1.2.840.10008.5.1.4.1.1.77.1.4.1" // A.32.7
	OphthalmicPhotography8BitImageStorage                              UID = "1.2.840.10008.5.1.4.1.1.77.1.5.1" // A.41
	OphthalmicPhotography16BitImageStorage                             UID = "1.2.840.10008.5.1.4.1.1.77.1.5.2" // A.42
	StereometricRelationshipStorage                                    UID = "1.2.840.10008.5.1.4.1.1.77.1.5.3" // A.43
	OphthalmicTomographyImageStorage                                   UID = "1.2.840.10008.5.1.4.1.1.77.1.5.4" // A.52
	WideFieldOphthalmicPhotographyStereographicProjectionImageStorage  UID = "1.2.840.10008.5.1.4.1.1.77.1.5.5" // A.76
	WideFieldOphthalmicPhotography3DCoordinatesImageStorage            UID = "1.2.840.10008.5.1.4.1.1.77.1.5.6" // A.77
	OphthalmicOpticalCoherenceTomographyEnFaceImageStorage             UID = "1.2.840.10008.5.1.4.1.1.77.1.5.7" // A.83
	OphthlamicOpticalCoherenceTomographyBscanVolumeAnalysisStorage     UID = "1.2.840.10008.5.1.4.1.1.77.1.5.8" // A.84
	VLWholeSlideMicroscopyImageStorage                                 UID = "1.2.840.10008.5.1.4.1.1.77.1.6"   // A.32.8
	DermoscopicPhotographyImageStorage                                 UID = "1.2.840.10008.5.1.4.1.1.77.1.7"   // A.32.11
	ConfocalMicroscopyImageStorage                                     UID = "1.2.840.10008.5.1.4.1.1.77.1.8"
	ConfocalMicroscopyTiledPyramidalImageStorage                       UID = "1.2.840.10008.5.1.4.1.1.77.1.9"
	LensometryMeasurementsStorage                                      UID = "1.2.840.10008.5.1.4.1.1.78.1"  // A.60.1
	AutorefractionMeasurementsStorage                                  UID = "1.2.840.10008.5.1.4.1.1.78.2"  // A.60.2
	KeratometryMeasurementsStorage                                     UID = "1.2.840.10008.5.1.4.1.1.78.3"  // A.60.3
	SubjectiveRefractionMeasurementsStorage                            UID = "1.2.840.10008.5.1.4.1.1.78.4"  // A.60.4
	VisualAcuityMeasurementsStorage                                    UID = "1.2.840.10008.5.1.4.1.1.78.5"  // A.60.5
	SpectaclePrescriptionReportStorage                                 UID = "1.2.840.10008.5.1.4.1.1.78.6"  // A.35.9
	OphthalmicAxialMeasurementsStorage                                 UID = "1.2.840.10008.5.1.4.1.1.78.7"  // A.60.6
	IntraocularLensCalculationsStorage                                 UID = "1.2.840.10008.5.1.4.1.1.78.8"  // A.60.7
	MacularGridThicknessAndVolumeReportStorage                         UID = "1.2.840.10008.5.1.4.1.1.79.1"  // A.35.11
	OphthalmicVisualFieldStaticPerimetryMeasurementsStorage            UID = "1.2.840.10008.5.1.4.1.1.80.1"  // A.65
	OphthalmicThicknessMapStorage                                      UID = "1.2.840.10008.5.1.4.1.1.81.1"  // A.67
	CornealTopographyMapStorage                                        UID = "1.2.840.10008.5.1.4.1.1.82.1"  // A.73
	BasicTextSRStorage                                                 UID = "1.2.840.10008.5.1.4.1.1.88.11" // A.35.1
	EnhancedSRStorage                                                  UID = "1.2.840.10008.5.1.4.1.1.88.22" // A.35.2
	ComprehensiveSRStorage                                             UID = "1.2.840.10008.5.1.4.1.1.88.33" // A.35.3
	Comprehensive3DSRStorage                                           UID = "1.2.840.10008.5.1.4.1.1.88.34" // A.35.13
	ExtensibleSRStorage                                                UID = "1.2.840.10008.5.1.4.1.1.88.35" // A.35.15
	ProcedureLogStorage                                                UID = "1.2.840.10008.5.1.4.1.1.88.40" // A.35.7
	MammographyCADSRStorage                                            UID = "1.2.840.10008.5.1.4.1.1.88.50" // A.35.5
	KeyObjectSelectionDocumentStorage                                  UID = "1.2.840.10008.5.1.4.1.1.88.59" // A.35.4
	ChestCADSRStorage                                                  UID = "1.2.840.10008.5.1.4.1.1.88.65" // A.35.6
	XRayRadiationDoseSRStorage                                         UID = "1.2.840.10008.5.1.4.1.1.88.67" // A.35.8
	RadiopharmaceuticalRadiationDoseSRStorage                          UID = "1.2.840.10008.5.1.4.1.1.88.68" // A.35.14
	ColonCADSRStorage                                                  UID = "1.2.840.10008.5.1.4.1.1.88.69" // A.35.10
	ImplantationPlanSRStorage                                          UID = "1.2.840.10008.5.1.4.1.1.88.70" // A.35.12
	AcquisitionContextSRStorage                                        UID = "1.2.840.10008.5.1.4.1.1.88.71" // A.35.16
	SimplifiedAdultEchoSRStorage                                       UID = "1.2.840.10008.5.1.4.1.1.88.72" // A.35.17
	PatientRadiationDoseSRStorage                                      UID = "1.2.840.10008.5.1.4.1.1.88.73" // A.35.18
	PlannedImagingAgentAdministrationSRStorage                         UID = "1.2.840.10008.5.1.4.1.1.88.74" // A.35.19
	PerformedImagingAgentAdministrationSRStorage                       UID = "1.2.840.10008.5.1.4.1.1.88.75" // A.35.20
	EnhancedXRayRadiationDoseSRStorage                                 UID = "1.2.840.10008.5.1.4.1.1.88.76" // A.35.
	WaveformAnnotationSRStorage                                        UID = "1.2.840.10008.5.1.4.1.1.88.77"
	ContentAssessmentResultsStorage                                    UID = "1.2.840.10008.5.1.4.1.1.90.1" // A.81
	MicroscopyBulkSimpleAnnotationsStorage                             UID = "1.2.840.10008.5.1.4.1.1.91.1"
	EncapsulatedPDFStorage                                             UID = "1.2.840.10008.5.1.4.1.1.104.1" // A.45.1
	EncapsulatedCDAStorage                                             UID = "1.2.840.10008.5.1.4.1.1.104.2" // A.45.2
	EncapsulatedSTLStorage                                             UID = "1.2.840.10008.5.1.4.1.1.104.3" // A.85.1
	EncapsulatedOBJStorage                                             UID = "1.2.840.10008.5.1.4.1.1.104.4" // A.85.2
	EncapsulatedMTLStorage                                             UID = "1.2.840.10008.5.1.4.1.1.104.5" // A.85.3
	PositronEmissionTomographyImageStorage                             UID = "1.2.840.10008.5.1.4.1.1.128"   // A.21
	LegacyConvertedEnhancedPETImageStorage                             UID = "1.2.840.10008.5.1.4.1.1.128.1" // A.72
	EnhancedPETImageStorage                                            UID = "1.2.840.10008.5.1.4.1.1.130"   // A.56
	BasicStructuredDisplayStorage                                      UID = "1.2.840.10008.5.1.4.1.1.131"   // A.33.5
	CTPerformedProcedureProtocolStorage                                UID = "1.2.840.10008.5.1.4.1.1.200.2" // A.82.1
	XAPerformedProcedureProtocolStorage                                UID = "1.2.840.10008.5.1.4.1.1.200.8"
	RTImageStorage                                                     UID = "1.2.840.10008.5.1.4.1.1.481.1"  // A.17
	RTDoseStorage                                                      UID = "1.2.840.10008.5.1.4.1.1.481.2"  // A.18
	RTStructureSetStorage                                              UID = "1.2.840.10008.5.1.4.1.1.481.3"  // A.19
	RTBeamsTreatmentRecordStorage                                      UID = "1.2.840.10008.5.1.4.1.1.481.4"  // A.29
	RTPlanStorage                                                      UID = "1.2.840.10008.5.1.4.1.1.481.5"  // A.20
	RTBrachyTreatmentRecordStorage                                     UID = "1.2.840.10008.5.1.4.1.1.481.6"  // A.20
	RTTreatmentSummaryRecordStorage                                    UID = "1.2.840.10008.5.1.4.1.1.481.7"  // A.31
	RTIonPlanStorage                                                   UID = "1.2.840.10008.5.1.4.1.1.481.8"  // A.49
	RTIonBeamsTreatmentRecordStorage                                   UID = "1.2.840.10008.5.1.4.1.1.481.9"  // A.50
	RTPhysicianIntentStorage                                           UID = "1.2.840.10008.5.1.4.1.1.481.10" // A.86.1.2
	RTSegmentAnnotationStorage                                         UID = "1.2.840.10008.5.1.4.1.1.481.11" // A.86.1.3
	RTRadiationSetStorage                                              UID = "1.2.840.10008.5.1.4.1.1.481.12" // A.86.1.4
	CArmPhotonElectronRadiationStorage                                 UID = "1.2.840.10008.5.1.4.1.1.481.13" // A.86.1.5
	TomotherapeuticRadiationStorage                                    UID = "1.2.840.10008.5.1.4.1.1.481.14" // A.86.1.6
	RoboticArmRadiationStorage                                         UID = "1.2.840.10008.5.1.4.1.1.481.15" // A.86.1.7
	RTRadiationRecordSetStorage                                        UID = "1.2.840.10008.5.1.4.1.1.481.16" // A.86.1.8
	RTRadiationSalvageRecordStorage                                    UID = "1.2.840.10008.5.1.4.1.1.481.17" // A.86.1.9
	TomotherapeuticRadiationRecordStorage                              UID = "1.2.840.10008.5.1.4.1.1.481.18" // A.86.1.10
	CArmPhotonElectronRadiationRecordStorage                           UID = "1.2.840.10008.5.1.4.1.1.481.19" // A.86.1.11
	RoboticArmRadiationRecordStorage                                   UID = "1.2.840.10008.5.1.4.1.1.481.20" // A.86.1.12
	RTRadiationSetDeliveryInstructionStorage                           UID = "1.2.840.10008.5.1.4.1.1.481.21"
	RTTreatmentPreparationStorage                                      UID = "1.2.840.10008.5.1.4.1.1.481.22"
	EnhancedRTImageStorage                                             UID = "1.2.840.10008.5.1.4.1.1.481.23"
	EnhancedContinuousRTImageStorage                                   UID = "1.2.840.10008.5.1.4.1.1.481.24"
	RTPatientPositionAcquisitionInstructionStorage                     UID = "1.2.840.10008.5.1.4.1.1.481.25"
	RTBeamsDeliveryInstructionStorage                                  UID = "1.2.840.10008.5.1.4.34.7"  // A.64
	RTBrachyApplicationSetupDeliveryInstructionsStorage                UID = "1.2.840.10008.5.1.4.34.10" // A.79
	StorageCommitmentPushModel                                         UID = "1.2.840.10008.1.20.1"
	InventoryCreation                                                  UID = "1.2.840.10008.5.1.4.1.1.201.5"
	ProductCharacteristicsQuery                                        UID = "1.2.840.10008.5.1.4.41"
	SubstanceApprovalQuery                                             UID = "1.2.840.10008.5.1.4.42"
	UnifiedProcedureStepPush                                           UID = "1.2.840.10008.5.1.4.34.6.1"
	UnifiedProcedureStepWatch                                          UID = "1.2.840.10008.5.1.4.34.6.2"
	UnifiedProcedureStepPull                                           UID = "1.2.840.10008.5.1.4.34.6.3"
	UnifiedProcedureStepEvent                                          UID = "1.2.840.10008.5.1.4.34.6.4"
	UnifiedProcedureStepQuery                                          UID = "1.2.840.10008.5.1.4.34.6.5"
	Verification                                                       UID = "1.2.840.10008.1.1"
)

var (
	AppEventClasses             = []UID{ProceduralEventLogging, SubstanceAdministrationLogging}
	BasicWorklistClasses        = []UID{ModalityWorklistInformationFind}
	ColorPaletteClasses         = []UID{ColorPaletteInformationModelFind, ColorPaletteInformationModelMove, ColorPaletteInformationModelGet}
	DefinedProcedureClasses     = []UID{DefinedProcedureProtocolInformationModelFind, DefinedProcedureProtocolInformationModelMove, DefinedProcedureProtocolInformationModelGet}
	DisplaySystemClasses        = []UID{DisplaySystem}
	InstanceAvailabilityClasses = []UID{InstanceAvailabilityNotification}
	HangingProtocolClasses      = []UID{HangingProtocolInformationModelFind, HangingProtocolInformationModelMove, HangingProtocolInformationModelGet}
	ImplantTemplateClasses      = []UID{GenericImplantTemplateInformationModelFind, GenericImplantTemplateInformationModelMove,
		GenericImplantTemplateInformationModelGet, ImplantAssemblyTemplateInformationModelFind, ImplantAssemblyTemplateInformationModelMove,
		ImplantAssemblyTemplateInformationModelGet, ImplantTemplateGroupInformationModelFind, ImplantTemplateGroupInformationModelMove,
		ImplantTemplateGroupInformationModelGet}
	InventoryClasses        = []UID{InventoryFind, InventoryMove, InventoryGet}
	MediaCreationClasses    = []UID{MediaCreationManagement}
	MediaStorageClasses     = []UID{MediaStorageDirectoryStorage}
	NonPatientObjectClasses = []UID{HangingProtocolStorage, ColorPaletteStorage, GenericImplantTemplateStorage,
		ImplantAssemblyTemplateStorage, ImplantTemplateGroupStorage, CTDefinedProcedureProtocolStorage, ProtocolApprovalStorage,
		XADefinedProcedureProtocolStorage, InventoryStorage}
	PrintManagementClasses = []UID{BasicFilmSession, BasicFilmBox, BasicGrayscaleImageBox, BasicColorImageBox,
		PrintJob, BasicAnnotationBox, Printer, PrinterConfigurationRetrieval, PresentationLUT, BasicGrayscalePrintManagementMeta,
		BasicColorPrintManagementMeta}
	ProcedureStepClasses    = []UID{ModalityPerformedProcedureStep, ModalityPerformedProcedureStepRetrieve, ModalityPerformedProcedureStepNotification}
	ProtocolApprovalClasses = []UID{ProtocolApprovalInformationModelFind, ProtocolApprovalInformationModelMove, ProtocolApprovalInformationModelGet}
	QRClasses               = []UID{
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
	RelevantPatientQueryClasses = []UID{
		GeneralRelevantPatientInformationQuery,
		BreastImagingRelevantPatientInformationQuery,
		CardiacRelevantPatientInformationQuery,
	}
	RTMachineVerificationClasses = []UID{
		RTConventionalMachineVerification,
		RTIonMachineVerification,
	}
	StorageClasses = []UID{
		AmbulatoryECGWaveformStorage,
		ArterialPulseWaveformStorage,
		AutorefractionMeasurementsStorage,
		BasicStructuredDisplayStorage,

		BasicTextSRStorage,
		BasicVoiceAudioWaveformStorage,
		BlendingSoftcopyPresentationStateStorage,
		BreastTomosynthesisImageStorage,
		CardiacElectrophysiologyWaveformStorage,
		ChestCADSRStorage,
		ColonCADSRStorage,
		ColorSoftcopyPresentationStateStorage,
		Comprehensive3DSRStorage,
		ComprehensiveSRStorage,
		ComputedRadiographyImageStorage,
		CTImageStorage,
		DeformableSpatialRegistrationStorage,
		DigitalIntraOralXRayImageStorageForPresentation,
		DigitalIntraOralXRayImageStorageForProcessing,
		DigitalMammographyXRayImageStorageForPresentation,
		DigitalMammographyXRayImageStorageForProcessing,
		DigitalXRayImageStorageForPresentation,
		DigitalXRayImageStorageForProcessing,
		EncapsulatedCDAStorage,
		EncapsulatedPDFStorage,
		EnhancedCTImageStorage,
		EnhancedMRColorImageStorage,
		EnhancedMRImageStorage,
		EnhancedPETImageStorage,
		EnhancedSRStorage,
		EnhancedUSVolumeStorage,
		EnhancedXAImageStorage,
		EnhancedXRFImageStorage,
		GeneralAudioWaveformStorage,
		GeneralECGWaveformStorage,
		GrayscaleSoftcopyPresentationStateStorage,
		HemodynamicWaveformStorage,
		ImplantationPlanSRStorage,
		IntraocularLensCalculationsStorage,
		IntravascularOpticalCoherenceTomographyImageStorageForPresentation,
		IntravascularOpticalCoherenceTomographyImageStorageForProcessing,
		KeratometryMeasurementsStorage,
		KeyObjectSelectionDocumentStorage,
		LegacyConvertedEnhancedCTImageStorage,
		LegacyConvertedEnhancedMRImageStorage,
		LegacyConvertedEnhancedPETImageStorage,
		LensometryMeasurementsStorage,
		MacularGridThicknessAndVolumeReportStorage,
		MammographyCADSRStorage,
		MRImageStorage,
		MRSpectroscopyStorage,
		MultiFrameGrayscaleByteSecondaryCaptureImageStorage,
		MultiFrameGrayscaleWordSecondaryCaptureImageStorage,
		MultiFrameSingleBitSecondaryCaptureImageStorage,
		MultiFrameTrueColorSecondaryCaptureImageStorage,
		NuclearMedicineImageStorage,
		OphthalmicAxialMeasurementsStorage,
		OphthalmicPhotography16BitImageStorage,
		OphthalmicPhotography8BitImageStorage,
		OphthalmicThicknessMapStorage,
		OphthalmicTomographyImageStorage,
		OphthalmicVisualFieldStaticPerimetryMeasurementsStorage,
		PositronEmissionTomographyImageStorage,
		ProcedureLogStorage,
		PseudoColorSoftcopyPresentationStageStorage,
		RawDataStorage,
		RealWorldValueMappingStorage,
		RespiratoryWaveformStorage,
		RTBeamsDeliveryInstructionStorage,
		RTBeamsTreatmentRecordStorage,
		RTBrachyTreatmentRecordStorage,
		RTDoseStorage,
		RTImageStorage,
		RTIonBeamsTreatmentRecordStorage,
		RTIonPlanStorage,
		RTPlanStorage,
		RTStructureSetStorage,
		RTTreatmentSummaryRecordStorage,
		SecondaryCaptureImageStorage,
		SegmentationStorage,
		SpatialFiducialsStorage,
		SpatialRegistrationStorage,
		SpectaclePrescriptionReportStorage,
		StereometricRelationshipStorage,
		SubjectiveRefractionMeasurementsStorage,
		SurfaceScanMeshStorage,
		SurfaceScanPointCloudStorage,
		SurfaceSegmentationStorage,
		TwelveLeadECGWaveformStorage,
		UltrasoundImageStorage,
		UltrasoundMultiFrameImageStorage,
		VideoEndoscopicImageStorage,
		VideoMicroscopicImageStorage,
		VideoPhotographicImageStorage,
		VisualAcuityMeasurementsStorage,
		VLEndoscopicImageStorage,
		VLMicroscopicImageStorage,
		VLPhotographicImageStorage,
		VLSlideCoordinatesMicroscopicImageStorage,
		VLWholeSlideMicroscopyImageStorage,
		XAXRFGrayscaleSoftcopyPresentationStateStorage,
		XRay3DAngiographicImageStorage,
		XRay3DCraniofacialImageStorage,
		XRayAngiographicImageStorage,
		XRayRadiationDoseSRStorage,
		XRayRadiofluoroscopicImageStorage,
	}
	StorageCommitmentClasses       = []UID{StorageCommitmentPushModel}
	StorageManagementClasses       = []UID{InventoryCreation}
	SubstanceAdministrationClasses = []UID{ProductCharacteristicsQuery, SubstanceApprovalQuery}
	UnifiedProcedureStepClasses    = []UID{UnifiedProcedureStepPush, UnifiedProcedureStepWatch, UnifiedProcedureStepPull, UnifiedProcedureStepEvent, UnifiedProcedureStepQuery}
	VerificationClasses            = []UID{Verification}
	QRFindClasses                  = []UID{
		PatientRootQueryRetrieveInformationModelFind,
		StudyRootQueryRetrieveInformationModelFind,
		PatientStudyOnlyQueryRetrieveInformationModelFind,
		ModalityWorklistInformationFind,
	}
	QRMoveClasses = []UID{
		PatientRootQueryRetrieveInformationModelMove,
		StudyRootQueryRetrieveInformationModelMove,
		PatientStudyOnlyQueryRetrieveInformationModelMove,
	}
	QRGetClasses = append([]UID{
		PatientRootQueryRetrieveInformationModelGet,
		StudyRootQueryRetrieveInformationModelGet,
		PatientStudyOnlyQueryRetrieveInformationModelGet,
	}, StorageClasses...)
	AllClasses = append([]UID{Verification}, StorageClasses...)
)
