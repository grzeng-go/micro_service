package institution

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Control struct {
	InsIdCd      string    `gorm:"column:INS_ID_CD;primary_key"`
	InsCompanyCd string    `gorm:"column:INS_COMPANY_CD"`
	ProdCd       string    `gorm:"column:PROD_CD;primary_key"`
	BizCd        string    `gorm:"column:BIZ_CD;primary_key"`
	CtrlSta      string    `gorm:"column:CTRL_STA"`
	InsBegTm     string    `gorm:"column:INS_BEG_TM"`
	InsEndTm     string    `gorm:"column:INS_END_TM"`
	MsgResvFld1  string    `gorm:"column:MSG_RESV_FLD1"`
	MsgResvFld2  string    `gorm:"column:MSG_RESV_FLD2"`
	MsgResvFld3  string    `gorm:"column:MSG_RESV_FLD3"`
	MsgResvFld4  string    `gorm:"column:MSG_RESV_FLD4"`
	MsgResvFld5  string    `gorm:"column:MSG_RESV_FLD5"`
	MsgResvFld6  string    `gorm:"column:MSG_RESV_FLD6"`
	MsgResvFld7  string    `gorm:"column:MSG_RESV_FLD7"`
	MsgResvFld8  string    `gorm:"column:MSG_RESV_FLD8"`
	MsgResvFld9  string    `gorm:"column:MSG_RESV_FLD9"`
	MsgResvFld10 string    `gorm:"column:MSG_RESV_FLD10"`
	RecOprId     string    `gorm:"column:REC_OPR_ID"`
	RecUpdOpr    string    `gorm:"column:REC_UPD_OPR"`
	CreatedAt    time.Time `gorm:"column:REC_CRT_TS"`
	UpdatedAt    time.Time `gorm:"column:REC_UPD_TS"`
}

func (c *Control) TableName() string {
	return "TBL_EDIT_INS_CTRL_INF"
}

func SaveInstitutionControl(db *gorm.DB, control *Control) error {
	return db.Save(control).Error
}

func FindInstitutionControlByPrimaryKey(
	db *gorm.DB,
	insId string,
	proCd string,
	bizCd string,
) (*Control, error) {
	out := new(Control)
	query := &Control{
		InsIdCd: insId,
		ProdCd:  proCd,
		BizCd:   bizCd,
	}
	err := db.Where(query).Take(out).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return out, err
}