package handler

import (
	"bupt/lep_audit/biz/logic"
	"bupt/lep_audit/kitex_gen/base"
	"bupt/lep_audit/kitex_gen/lep_audit"
	"context"
)

func AuditList(ctx context.Context, req *lep_audit.AuditListReq) (*lep_audit.AuditListResp, error) {
	resp := new(lep_audit.AuditListResp)
	resp.BaseResp = base.NewBaseResp()
	list, total, err := logic.NewAuditListHandler(ctx, req).GetAuditList()
	if err != nil {
		return nil, err
	}
	resp.Total = total
	resp.AuditList = list
	return resp, err
}
func CreateAudit(ctx context.Context, req *lep_audit.CreateAuditReq) (*lep_audit.CreateAuditResp, error) {
	resp := new(lep_audit.CreateAuditResp)
	resp.BaseResp = base.NewBaseResp()
	handler := logic.NewCreateHandler(ctx, req)
	err := handler.Check()
	if err != nil {
		return nil, err
	}
	id, err := handler.Create()
	if err != nil {
		return nil, err
	}
	resp.AuditId = id
	return resp, err
}
func Audit(ctx context.Context, req *lep_audit.AuditReq) (*lep_audit.AuditResp, error) {
	resp := new(lep_audit.AuditResp)
	resp.BaseResp = base.NewBaseResp()
	handler := logic.NewAuditHandler(ctx, req)
	err := handler.Check()
	if err != nil {
		return nil, err
	}
	err = handler.Audit()
	if err != nil {
		return nil, err
	}
	return resp, err
}
