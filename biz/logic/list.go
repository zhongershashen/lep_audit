package logic

import (
	mysql "bupt/lep_audit/biz/dal"
	"bupt/lep_audit/kitex_gen/lep_audit"
	"context"

	"github.com/zhongershashen/lep_lib/dal"
)

type ListHandler struct {
	ctx context.Context
	req *lep_audit.AuditListReq
}

func NewAuditListHandler(ctx context.Context, req *lep_audit.AuditListReq) *ListHandler {
	return &ListHandler{
		ctx: ctx,
		req: req,
	}
}
func (h *ListHandler) GetAuditList() ([]*lep_audit.AuditInfo, int64, error) {
	list := make([]*lep_audit.AuditInfo, 0)
	db := dal.GetDB()
	condition := h.buildCondition()
	total, err := mysql.CountAudit(h.ctx, db, condition)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return list, 0, nil
	}
	modelList, err := mysql.QueryAudit(h.ctx, db, condition, int(h.req.Offset), int(h.req.Limit))
	for _, item := range modelList {
		tempPermission := &lep_audit.AuditInfo{
			Id:              item.ID,
			AuditType:       item.AuditType,
			TargetId:        item.TargetId,
			TargetType:      item.TargetType,
			ApplyInfo:       item.ApplyInfo,
			AuditStatus:     item.AuditStatus,
			ApplyUserId:     item.ApplyUserId,
			ApplyUserName:   item.ApplyUserName,
			AuditorUserId:   item.AuditorUserId,
			FailReason:      item.Reason,
			ApplyTime:       item.ApplyTime.Unix(),
			AuditTime:       item.AuditTime.Unix(),
			AuditorUserName: item.AuditorUserName,
		}
		list = append(list, tempPermission)
	}
	return list, total, nil
}

func (h *ListHandler) buildCondition() map[string]interface{} {
	condition := make(map[string]interface{}, 0)
	req := h.req
	if req.UserId != nil {
		condition["apply_usd_id = ?"] = *req.UserId
	}
	if req.AuditType != nil {
		condition["audit_type = ?"] = *req.AuditType
	}
	if req.AuditStatus != nil {
		condition["audit_status = ?"] = *req.AuditStatus
	}
	return condition
}
