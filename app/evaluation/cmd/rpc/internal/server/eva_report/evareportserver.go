// Code generated by goctl. DO NOT EDIT.
// Source: evaluation.proto

package server

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/internal/logic/eva_report"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"
)

type EvaReportServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedEvaReportServer
}

func NewEvaReportServer(svcCtx *svc.ServiceContext) *EvaReportServer {
	return &EvaReportServer{
		svcCtx: svcCtx,
	}
}

func (s *EvaReportServer) SendReport(ctx context.Context, in *pb.SendReportRequest) (*pb.StatusResponse, error) {
	l := evareportlogic.NewSendReportLogic(ctx, s.svcCtx)
	return l.SendReport(in)
}

func (s *EvaReportServer) GetReports(ctx context.Context, in *pb.GetReportsRequest) (*pb.GetReportsResponse, error) {
	l := evareportlogic.NewGetReportsLogic(ctx, s.svcCtx)
	return l.GetReports(in)
}

func (s *EvaReportServer) GetTheReport(ctx context.Context, in *pb.GetTheReportRequest) (*pb.GetTheReportResponse, error) {
	l := evareportlogic.NewGetTheReportLogic(ctx, s.svcCtx)
	return l.GetTheReport(in)
}

func (s *EvaReportServer) SetReport(ctx context.Context, in *pb.SetReportRequest) (*pb.StatusResponse, error) {
	l := evareportlogic.NewSetReportLogic(ctx, s.svcCtx)
	return l.SetReport(in)
}
