// Code generated by goctl. DO NOT EDIT.
// Source: evaluation.proto

package eva_info

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateEvaluationRequest = pb.CreateEvaluationRequest
	DeleteEvaluationRequest = pb.DeleteEvaluationRequest
	Evaluation              = pb.Evaluation
	EvaluationResponse      = pb.EvaluationResponse
	GetEvaluationRequest    = pb.GetEvaluationRequest
	GetLikeRequest          = pb.GetLikeRequest
	GetLikeResponse         = pb.GetLikeResponse
	GetReportsRequest       = pb.GetReportsRequest
	GetReportsResponse      = pb.GetReportsResponse
	GetTheReportRequest     = pb.GetTheReportRequest
	GetTheReportResponse    = pb.GetTheReportResponse
	Like                    = pb.Like
	Report                  = pb.Report
	SendReportRequest       = pb.SendReportRequest
	SetLikeRequest          = pb.SetLikeRequest
	SetReportRequest        = pb.SetReportRequest
	StatusResponse          = pb.StatusResponse
	UpdateEvaluationRequest = pb.UpdateEvaluationRequest

	EvaInfo interface {
		CreateEvaluation(ctx context.Context, in *CreateEvaluationRequest, opts ...grpc.CallOption) (*StatusResponse, error)
		DeleteEvaluation(ctx context.Context, in *DeleteEvaluationRequest, opts ...grpc.CallOption) (*StatusResponse, error)
		GetEvaluation(ctx context.Context, in *GetEvaluationRequest, opts ...grpc.CallOption) (*EvaluationResponse, error)
		UpdateEvaluation(ctx context.Context, in *UpdateEvaluationRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	}

	defaultEvaInfo struct {
		cli zrpc.Client
	}
)

func NewEvaInfo(cli zrpc.Client) EvaInfo {
	return &defaultEvaInfo{
		cli: cli,
	}
}

func (m *defaultEvaInfo) CreateEvaluation(ctx context.Context, in *CreateEvaluationRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	client := pb.NewEvaInfoClient(m.cli.Conn())
	return client.CreateEvaluation(ctx, in, opts...)
}

func (m *defaultEvaInfo) DeleteEvaluation(ctx context.Context, in *DeleteEvaluationRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	client := pb.NewEvaInfoClient(m.cli.Conn())
	return client.DeleteEvaluation(ctx, in, opts...)
}

func (m *defaultEvaInfo) GetEvaluation(ctx context.Context, in *GetEvaluationRequest, opts ...grpc.CallOption) (*EvaluationResponse, error) {
	client := pb.NewEvaInfoClient(m.cli.Conn())
	return client.GetEvaluation(ctx, in, opts...)
}

func (m *defaultEvaInfo) UpdateEvaluation(ctx context.Context, in *UpdateEvaluationRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	client := pb.NewEvaInfoClient(m.cli.Conn())
	return client.UpdateEvaluation(ctx, in, opts...)
}
