package transfersyntax

import "encoding/binary"

type UID string

const (
	ImplicitVRLittleEndian                          UID = "1.2.840.10008.1.2"
	ExplicitVRLittleEndian                          UID = "1.2.840.10008.1.2.1"
	ExplicitVRBigEndian                             UID = "1.2.840.10008.1.2.2"
	DeflatedExplicitVRLittleEndian                  UID = "1.2.840.10008.1.2.1.99"
	JPEGBaseline8Bit                                UID = "1.2.840.10008.1.2.4.50"
	JPEGExtended12Bit                               UID = "1.2.840.10008.1.2.4.51"
	JPEGLossless                                    UID = "1.2.840.10008.1.2.4.57"
	JPEGLosslessSV1                                 UID = "1.2.840.10008.1.2.4.70"
	JPEGLSLossless                                  UID = "1.2.840.10008.1.2.4.80"
	JPEGLSLossy                                     UID = "1.2.840.10008.1.2.4.81"
	JPEG2000Lossless                                UID = "1.2.840.10008.1.2.4.90"
	JPEG2000                                        UID = "1.2.840.10008.1.2.4.91"
	JPEG2000MultiComponentLossless                  UID = "1.2.840.10008.1.2.4.92"
	JPEG2000MultiComponent                          UID = "1.2.840.10008.1.2.4.93"
	JPIPReferenced                                  UID = "1.2.840.10008.1.2.4.94"
	JPIPReferencedDeflate                           UID = "1.2.840.10008.1.2.4.95"
	MPEG2MainProfile                                UID = "1.2.840.10008.1.2.4.100"
	FragmentableMPEG2MainProfile                    UID = "1.2.840.10008.1.2.4.100.1"
	MPEG2MainProfileHighLevel                       UID = "1.2.840.10008.1.2.4.101"
	FragmentableMPEG2MainProfileHighLevel           UID = "1.2.840.10008.1.2.4.101.1"
	MPEG4AVCH264HighProfileLevel                    UID = "1.2.840.10008.1.2.4.102"
	FragmentableMPEG4AVCH264HighProfileLevel        UID = "1.2.840.10008.1.2.4.102.1"
	MPEG4AVCH264BDcompatibleHighProfile             UID = "1.2.840.10008.1.2.4.103"
	FragmentableMPEG4AVCH264BDcompatibleHighProfile UID = "1.2.840.10008.1.2.4.103.1"
	MPEG4AVCH264HighProfileFor2DVideo               UID = "1.2.840.10008.1.2.4.104"
	FragmentableMPEG4AVCH264HighProfileFor2DVideo   UID = "1.2.840.10008.1.2.4.104.1"
	MPEG4AVCH264HighProfileFor3DVideo               UID = "1.2.840.10008.1.2.4.105"
	FragmentableMPEG4AVCH264HighProfileFor3DVideo   UID = "1.2.840.10008.1.2.4.105.1"
	MPEG4AVCH264StereoHighProfile                   UID = "1.2.840.10008.1.2.4.106"
	FragmentableMPEG4AVCH264StereoHighProfile       UID = "1.2.840.10008.1.2.4.106.1"
	HEVCH265MainProfileLevel5                       UID = "1.2.840.10008.1.2.4.107"
	HEVCH265Main10ProfileLevel5                     UID = "1.2.840.10008.1.2.4.108"
	JPEGXLLossless                                  UID = "1.2.840.10008.1.2.4.110"
	JPEGXLJPEGRecompression                         UID = "1.2.840.10008.1.2.4.111"
	JPEGXL                                          UID = "1.2.840.10008.1.2.4.112"
	HighThroughputJPEG2000Lossless                  UID = "1.2.840.10008.1.2.4.201"
	HighThroughputJPEG2000RPCL                      UID = "1.2.840.10008.1.2.4.202"
	HighThroughputJPEG2000                          UID = "1.2.840.10008.1.2.4.203"
	JPIPHT2KReferenced                              UID = "1.2.840.10008.1.2.4.204"
	JPIPHTJ2kReferencedDeflate                      UID = "1.2.840.10008.1.2.4.205"
	RLELossless                                     UID = "1.2.840.10008.1.2.5"
	SMPTEST211020UncompressedProgressiveActiveVideo UID = "1.2.840.10008.1.2.7.1"
	SMPTEST211020UncompressedInterlacedActiveVideo  UID = "1.2.840.10008.1.2.7.2"
	SMPTEST211030PCMDigitalAudio                    UID = "1.2.840.10008.1.2.7.3"
	DeflatedImageFrameCompression                   UID = "1.2.840.10008.1.2.8.1"
)

var StandardSyntaxes = []UID{
	ImplicitVRLittleEndian,
	ExplicitVRLittleEndian,
	ExplicitVRBigEndian,
	DeflatedExplicitVRLittleEndian,
}

func Info(uid UID) (bo binary.ByteOrder, implicit bool) {
	switch uid {
	case ImplicitVRLittleEndian:
		return binary.LittleEndian, true
	case DeflatedExplicitVRLittleEndian, ExplicitVRLittleEndian:
		return binary.LittleEndian, false
	case ExplicitVRBigEndian:
		fallthrough
	default:
		return binary.BigEndian, false
	}
}
