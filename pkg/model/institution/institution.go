package institution

import (
	"time"

	"github.com/jinzhu/gorm"
)

type InstitutionInfo struct {
	InsIDCd         string    `gorm:"column:INS_ID_CD;primary_key"`
	InsCompanyCd    string    `gorm:"column:INS_COMPANY_CD"`
	InsType         string    `gorm:"column:INS_TYPE"`
	InsName         string    `gorm:"column:INS_NAME"`
	InsProvCd       string    `gorm:"column:INS_PROV_CD"`
	InsCityCd       string    `gorm:"column:INS_CITY_CD"`
	InsRegionCd     string    `gorm:"column:INS_REGION_CD"`
	InsSta          string    `gorm:"column:INS_STA"`
	InsStlmTp       string    `gorm:"column:INS_STLM_TP"`
	InsAloStlmCycle string    `gorm:"column:INS_ALO_STLM_CYCLE"`
	InsAloStlmMd    string    `gorm:"column:INS_ALO_STLM_MD"`
	InsStlmCNm      string    `gorm:"column:INS_STLM_C_NM"`
	InsStlmCAcct    string    `gorm:"column:INS_STLM_C_ACCT"`
	InsStlmCBkNo    string    `gorm:"column:INS_STLM_C_BK_NO"`
	InsStlmCBkNm    string    `gorm:"column:INS_STLM_C_BK_NM"`
	InsStlmDNm      string    `gorm:"column:INS_STLM_D_NM"`
	InsStlmDAcct    string    `gorm:"column:INS_STLM_D_ACCT"`
	InsStlmDBkNo    string    `gorm:"column:INS_STLM_D_BK_NO"`
	InsStlmDBkNm    string    `gorm:"column:INS_STLM_D_BK_NM"`
	MsgResvFld1     string    `gorm:"column:MSG_RESV_FLD1"`
	MsgResvFld2     string    `gorm:"column:MSG_RESV_FLD2"`
	MsgResvFld3     string    `gorm:"column:MSG_RESV_FLD3"`
	MsgResvFld4     string    `gorm:"column:MSG_RESV_FLD4"`
	MsgResvFld5     string    `gorm:"column:MSG_RESV_FLD5"`
	MsgResvFld6     string    `gorm:"column:MSG_RESV_FLD6"`
	MsgResvFld7     string    `gorm:"column:MSG_RESV_FLD7"`
	MsgResvFld8     string    `gorm:"column:MSG_RESV_FLD8"`
	MsgResvFld9     string    `gorm:"column:MSG_RESV_FLD9"`
	MsgResvFld10    string    `gorm:"column:MSG_RESV_FLD10"`
	RecOprID        string    `gorm:"column:REC_OPR_ID"`
	RecUpdOpr       string    `gorm:"column:REC_UPD_OPR"`
	CreatedAt       time.Time `gorm:"column:REC_CRT_TS"`
	UpdatedAt       time.Time `gorm:"column:REC_UPD_TS"`
}

func (i InstitutionInfo) TableName() string {
	return "TBL_EDIT_INS_INF"
}

func QueryInstitutionInfo(db *gorm.DB, query *InstitutionInfo, page int32, size int32) ([]*InstitutionInfo, int32, error) {
	out := make([]*InstitutionInfo, 0)
	var count int32
	db.Model(&InstitutionInfo{}).Where(query).Count(&count)
	err := db.Where(query).Offset((page - 1) * size).Limit(size).Find(&out).Error
	if err == gorm.ErrRecordNotFound {
		return out, count, nil
	}
	return out, count, err
}

func FindInstitutionInfoById(db *gorm.DB, id string) (*InstitutionInfo, error) {
	out := new(InstitutionInfo)
	err := db.Where("INS_ID_CD = ?", id).Take(out).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return out, err
}

func FindInstitutionInfosByIdList(db *gorm.DB, ids []string) ([]*InstitutionInfo, error) {
	out := make([]*InstitutionInfo, 0)
	err := db.Where("INS_ID_CD in (?)", ids).Find(&out).Error
	if err == gorm.ErrRecordNotFound {
		return out, nil
	}
	return out, err
}

func SaveInstitution(db *gorm.DB, ins *InstitutionInfo) error {
	return db.Save(ins).Error
}
