package logic

import (
	mysql "bupt/lep_audit/biz/dal"
	"bupt/lep_audit/kitex_gen/lep_audit"
	"context"
	"errors"
	"time"

	"github.com/zhongershashen/lep_lib/dal"
)

type AuditHandler struct {
	ctx context.Context
	req *lep_audit.AuditReq
}

func NewAuditHandler(ctx context.Context, req *lep_audit.AuditReq) *AuditHandler {
	return &AuditHandler{
		ctx: ctx,
		req: req,
	}
}
func (h *AuditHandler) Audit() error {
	req := h.req
	db := dal.GetDB()
	condition := make(map[string]interface{}, 0)
	updates := make(map[string]interface{}, 0)
	condition["id = ?"] = req.AuditId
	switch req.Operation {
	case 1:
		updates["audit_status"] = 1
	case 2:
		updates["audit_status"] = 1
		if req.FailReason != nil {
			updates["reason"] = *req.FailReason
		}
	}
	updates["audit_user_name"] = req.AuditUserName
	updates["audit_user_id"] = req.AuditUserId
	updates["audit_time"] = time.Now()
	err := mysql.UpdateAudit(h.ctx, db, condition, updates)
	if err != nil {
		return err
	}
	return nil
}

func (h *AuditHandler) Check() error {
	req := h.req
	if req.AuditId == 0 {
		return errors.New("invalid audit_id")
	}
	if req.Operation == 0 {
		return errors.New("invalid op_type")
	}
	if req.AuditUserId == 0 || req.AuditUserName == "" {
		return errors.New("invalid audit_user_info")
	}
	return nil
}
