package model

import (
	"encoding/binary"
	"fmt"
)

type HWPVersion struct {
	Major    int
	Minor    int
	Build    int
	Revision int
}

func NewHWPVersion(data []byte) (HWPVersion, error) {
	expectedLength := 4

	if len(data) != expectedLength {
		return HWPVersion{}, &ByteLengthError{
			ExpectedLength: expectedLength,
			ActualLength:   len(data),
		}
	}

	return HWPVersion{
		Major:    int(data[3]),
		Minor:    int(data[2]),
		Build:    int(data[1]),
		Revision: int(data[0]),
	}, nil
}

func (v HWPVersion) IsCompatible(target HWPVersion) bool {
	return v.Major == target.Major && v.Minor <= target.Minor
}

// Gte returns true if v is greater than or equal to target
func (v HWPVersion) Gte(target HWPVersion) bool {
	if v.Major != target.Major {
		return v.Major > target.Major
	}
	if v.Minor != target.Minor {
		return v.Minor > target.Minor
	}
	if v.Build != target.Build {
		return v.Build > target.Build
	}
	return v.Revision >= target.Revision
}

func (v HWPVersion) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", v.Major, v.Minor, v.Build, v.Revision)
}

type Attributes1 struct {
	Compressed                        bool // 압축 여부
	Encrypted                         bool // 암호 설정 여부
	Distribution                      bool // 배포용 문서 여부
	Script                            bool // 스크립트 저장 여부
	Drm                               bool // DRM 보안 문서 여부
	HasXmlTemplateStorage             bool // XMLTemplate 스토리지 존재 여부
	Vcs                               bool // 문서 이력 관리 존재 여부
	HasElectronicSignatureInformation bool // 전자 서명 정보 존재 여부
	CertificateEncryption             bool // 공인 인증서 암호화 여부
	PrepareSignature                  bool // 전자 서명 예비 저장 여부
	CertificateDRM                    bool // 공인 인증서 DRM 보안 문서 여부
	Ccl                               bool // CCL 문서 여부
	Mobile                            bool // 모바일 최적화 여부
	IsPrivacySecurityDocument         bool // 개인 정보 보안 문서 여부
	TrackChanges                      bool // 변경 추적 문서 여부
	KOGL                              bool // 공공누리(KOGL) 저작권 문서
	HasVideoControl                   bool // 비디오 컨트롤 포함 여부
	HasOrderFieldControl              bool // 차례 필드 컨트롤 포함 여부
}

func NewAttributes1(data []byte) (*Attributes1, error) {
	if len(data) != 4 {
		return nil, &ByteLengthError{
			ExpectedLength: 4,
			ActualLength:   len(data),
		}
	}

	val := binary.LittleEndian.Uint32(data)

	// Check if rest(after 17) is invalid
	if val&(0xFFFC0000) != 0 {
		return nil, fmt.Errorf("bit 18~31 are reserved but has value")
	}

	return &Attributes1{
		Compressed:                        val&0x01 != 0,
		Encrypted:                         val&(1<<1) != 0,
		Distribution:                      val&(1<<2) != 0,
		Script:                            val&(1<<3) != 0,
		Drm:                               val&(1<<4) != 0,
		HasXmlTemplateStorage:             val&(1<<5) != 0,
		Vcs:                               val&(1<<6) != 0,
		HasElectronicSignatureInformation: val&(1<<7) != 0,
		CertificateEncryption:             val&(1<<8) != 0,
		PrepareSignature:                  val&(1<<9) != 0,
		CertificateDRM:                    val&(1<<10) != 0,
		Ccl:                               val&(1<<11) != 0,
		Mobile:                            val&(1<<12) != 0,
		IsPrivacySecurityDocument:         val&(1<<13) != 0,
		TrackChanges:                      val&(1<<14) != 0,
		KOGL:                              val&(1<<15) != 0,
		HasVideoControl:                   val&(1<<16) != 0,
		HasOrderFieldControl:              val&(1<<17) != 0,
	}, nil
}
