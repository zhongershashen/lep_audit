package dal

import (
	"bupt/lep_audit/biz/model"
	"context"

	"github.com/zhongershashen/lep_lib/dal"
	"gorm.io/gorm"
)

func QueryAudit(ctx context.Context, db *gorm.DB, conditions map[string]interface{}, offset, limit int) (modelList []*model.LepAuditTable, err error) {
	if db == nil {
		db = dal.GetDB()
	}
	modelList = make([]*model.LepAuditTable, 0)
	query := db.Table(model.LepAuditTableName)
	for key, value := range conditions {
		query = query.Where(key, value)
	}
	if offset != 0 {
		query = query.Offset(offset)
	}
	if limit != 0 {
		query = query.Limit(limit)
	}
	err = query.WithContext(ctx).Debug().Order("id desc").Find(&modelList).Error
	if err != nil {
		return nil, err
	}
	return modelList, nil
}
func UpsertAudit(ctx context.Context, db *gorm.DB, modelItem *model.LepAuditTable) error {
	if db == nil {
		db = dal.GetDB()
	}
	query := db.Table(model.LepAuditTableName)
	return query.WithContext(ctx).Debug().Omit("created_time").Save(&modelItem).Error
}
func CountAudit(ctx context.Context, db *gorm.DB, conditions map[string]interface{}) (total int64, err error) {
	if db == nil {
		db = dal.GetDB()
	}
	query := db.Table(model.LepAuditTableName)
	for key, value := range conditions {
		query = query.Where(key, value)
	}
	err = query.WithContext(ctx).Debug().Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
func UpdateAudit(ctx context.Context, db *gorm.DB, conditions, updates map[string]interface{}) error {
	if db == nil {
		db = dal.GetDB()
	}
	query := db.Table(model.LepAuditTableName)
	for key, value := range conditions {
		query = query.Where(key, value)
	}
	return query.WithContext(ctx).Debug().Updates(updates).Error
}
