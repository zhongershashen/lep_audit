package model

import "time"

const (
	LepAuditTableName = "lep_audit_table"
)

// LepAuditTable 审核信息表
type LepAuditTable struct {
	ID              int64      `json:"id" gorm:"id"`                               // 主键id
	CreatedTime     *time.Time `json:"created_time" gorm:"created_time"`           // 创建时间
	UpdatedTime     *time.Time `json:"updated_time" gorm:"updated_time"`           // 修改时间
	IsDeleted       int32      `json:"is_deleted" gorm:"is_deleted"`               // 0-未删除，1-已删除
	AuditType       int32      `json:"audit_type" gorm:"audit_type"`               // 审核类型，业务自定义
	ApplyUserName   string     `json:"apply_user_name" gorm:"apply_user_name"`     // 申请人
	ApplyUserId     int64      `json:"apply_user_id" gorm:"apply_user_id"`         // 申请人id
	TargetId        int32      `json:"target_id" gorm:"target_id"`                 // 审核目标ID
	TargetType      int32      `json:"target_type" gorm:"target_type"`             // 审核目标类型
	ApplyInfo       string     `json:"apply_info" gorm:"apply_info"`               // 附加申请信息
	AuditStatus     int32      `json:"audit_status" gorm:"audit_status"`           // 审核状态：1-待审核，2-审核通过，3-审核驳回
	BizStatus       int32      `json:"biz_status" gorm:"biz_status"`               // 业务自定义status
	AuditorUserId   int64      `json:"auditor_user_id" gorm:"auditor_user_id"`     // 审核人id
	AuditorUserName string     `json:"auditor_user_name" gorm:"auditor_user_name"` // 审核人
	Reason          string     `json:"reason" gorm:"reason"`                       // 驳回原因
	ApplyTime       time.Time  `json:"apply_time" gorm:"apply_time"`               // 申请时间
	AuditTime       time.Time  `json:"audit_time" gorm:"audit_time"`               // 审核时间
	Extra           string     `json:"extra" gorm:"extra"`                         // 扩展信息
	IsTest          int32      `json:"is_test" gorm:"is_test"`                     // 是否测试数据，1是，0不是
}
