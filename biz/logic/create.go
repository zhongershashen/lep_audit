package logic

import (
	mysql "bupt/lep_audit/biz/dal"
	"bupt/lep_audit/biz/model"
	"bupt/lep_audit/kitex_gen/lep_audit"
	"context"
	"errors"
	"time"

	"github.com/zhongershashen/lep_lib/dal"
)

type CreateHandler struct {
	ctx context.Context
	req *lep_audit.CreateAuditReq
}

func NewCreateHandler(ctx context.Context, req *lep_audit.CreateAuditReq) *CreateHandler {
	return &CreateHandler{
		ctx: ctx,
		req: req,
	}
}
func (h *CreateHandler) Create() (int64, error) {
	req := h.req
	db := dal.GetDB()
	modelItem := &model.LepAuditTable{
		AuditType:     req.AuditType,
		ApplyUserName: req.ApplyUserName,
		ApplyUserId:   req.ApplyUserId,
		TargetId:      req.TargetId,
		TargetType:    req.TargetType,
		ApplyTime:     time.Now(),
	}
	if req.ApplyInfo != nil {
		modelItem.ApplyInfo = *req.ApplyInfo
	}
	err := mysql.UpsertAudit(h.ctx, db, modelItem)
	if err != nil {
		return 0, err
	}
	return modelItem.ID, nil
}

func (h *CreateHandler) Check() error {
	req := h.req
	if req.ApplyUserName == "" || req.ApplyUserId == 0 {
		return errors.New("invalid apply user")
	}
	if req.TargetType == 0 || req.TargetId == 0 {
		return errors.New("invalid target info")
	}
	if req.AuditType == 0 {
		return errors.New("invalid audit type")
	}
	return nil
}
