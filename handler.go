package main

import (
	"bupt/lep_audit/biz/handler"
	lep_audit "bupt/lep_audit/kitex_gen/lep_audit"
	"context"
)

// LepAuditImpl implements the last service interface defined in the IDL.
type LepAuditImpl struct{}

// Audit implements the LepAuditImpl interface.
func (s *LepAuditImpl) Audit(ctx context.Context, req *lep_audit.AuditReq) (resp *lep_audit.AuditResp, err error) {
	resp, err = handler.Audit(ctx, req)
	return
}

// CreateAudit implements the LepAuditImpl interface.
func (s *LepAuditImpl) CreateAudit(ctx context.Context, req *lep_audit.CreateAuditReq) (resp *lep_audit.CreateAuditResp, err error) {
	resp, err = handler.CreateAudit(ctx, req)
	return
}

// AuditList implements the LepAuditImpl interface.
func (s *LepAuditImpl) AuditList(ctx context.Context, req *lep_audit.AuditListReq) (resp *lep_audit.AuditListResp, err error) {
	resp, err = handler.AuditList(ctx, req)
	return
}
