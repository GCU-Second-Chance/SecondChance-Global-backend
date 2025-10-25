package model

type GyeonggiRandomResponse struct {
	AbdmAnimalProtect []AbdmAnimalProtect `json:"AbdmAnimalProtect"`
}

type AbdmAnimalProtect struct {
	Head []AbdmHead `json:"head,omitempty"`
	Row  []AbdmRow  `json:"row,omitempty"`
}

type AbdmHead struct {
	ListTotalCount int         `json:"list_total_count,omitempty"`
	Result         *AbdmResult `json:"RESULT,omitempty"`
	APIVersion     string      `json:"api_version,omitempty"`
}

type AbdmResult struct {
	Code    string `json:"CODE"`
	Message string `json:"MESSAGE"`
}

type AbdmRow struct {
	SigunCD          string  `json:"SIGUN_CD"`           // 시군 코드
	SigunNM          string  `json:"SIGUN_NM"`           // 시군 이름
	AbdmIDntfyNo     string  `json:"ABDM_IDNTFY_NO"`     // 유기동물 식별번호
	ThumbImageCours  string  `json:"THUMB_IMAGE_COURS"`  // 썸네일 이미지 URL
	ReceptDE         string  `json:"RECEPT_DE"`          // 접수일자 (YYYYMMDD)
	DiscvryPlcInfo   string  `json:"DISCVRY_PLC_INFO"`   // 발견 장소
	SpeciesNM        string  `json:"SPECIES_NM"`         // 종 코드
	ColorNM          string  `json:"COLOR_NM"`           // 색상
	AgeInfo          string  `json:"AGE_INFO"`           // 나이 정보
	BdwghInfo        string  `json:"BDWGH_INFO"`         // 체중 정보
	PblancIDntfyNo   string  `json:"PBLANC_IDNTFY_NO"`   // 공고 식별번호
	PblancBeginDE    string  `json:"PBLANC_BEGIN_DE"`    // 공고 시작일
	PblancEndDE      string  `json:"PBLANC_END_DE"`      // 공고 종료일
	ImageCours       string  `json:"IMAGE_COURS"`        // 이미지 URL
	StateNM          string  `json:"STATE_NM"`           // 상태명 (예: 보호중)
	SexNM            string  `json:"SEX_NM"`             // 성별
	NeutYN           string  `json:"NEUT_YN"`            // 중성화 여부
	SfetrInfo        string  `json:"SFETR_INFO"`         // 특이사항
	ShterNM          string  `json:"SHTER_NM"`           // 보호소 이름
	ShterTelno       string  `json:"SHTER_TELNO"`        // 보호소 전화번호
	ProtectPlc       string  `json:"PROTECT_PLC"`        // 보호장소 주소
	JurisdInstNM     string  `json:"JURISD_INST_NM"`     // 관할 기관
	ChrgpsnNM        *string `json:"CHRGPSN_NM"`         // 담당자 이름 (nullable)
	ChrgpsnContctNo  *string `json:"CHRGPSN_CONTCT_NO"`  // 담당자 연락처 (nullable)
	PartclrMatr      *string `json:"PARTCLR_MATR"`       // 특이 사항 (nullable)
	RefineLotnoAddr  string  `json:"REFINE_LOTNO_ADDR"`  // 지번 주소
	RefineRoadnmAddr string  `json:"REFINE_ROADNM_ADDR"` // 도로명 주소
	RefineZipCd      string  `json:"REFINE_ZIP_CD"`      // 우편번호
	RefineWgs84Logt  string  `json:"REFINE_WGS84_LOGT"`  // 경도 (WGS84)
	RefineWgs84Lat   string  `json:"REFINE_WGS84_LAT"`   // 위도 (WGS84)
}
