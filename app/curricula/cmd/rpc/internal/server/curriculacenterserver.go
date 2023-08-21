// Code generated by goctl. DO NOT EDIT.
// Source: curriculacenter.proto

package server

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/internal/logic"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"
)

type CurriculacenterServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedCurriculacenterServer
}

func NewCurriculacenterServer(svcCtx *svc.ServiceContext) *CurriculacenterServer {
	return &CurriculacenterServer{
		svcCtx: svcCtx,
	}
}

func (s *CurriculacenterServer) AddCurricula(ctx context.Context, in *pb.AddCurriculaRequest) (*pb.AddCurriculaResponse, error) {
	l := logic.NewAddCurriculaLogic(ctx, s.svcCtx)
	return l.AddCurricula(in)
}

func (s *CurriculacenterServer) DeleteCurricula(ctx context.Context, in *pb.DeleteCurriculaRequest) (*pb.DeleteCurriculaResponse, error) {
	l := logic.NewDeleteCurriculaLogic(ctx, s.svcCtx)
	return l.DeleteCurricula(in)
}

func (s *CurriculacenterServer) UpdateCurricula(ctx context.Context, in *pb.UpdateCurriculaRequest) (*pb.UpdateCurriculaResponse, error) {
	l := logic.NewUpdateCurriculaLogic(ctx, s.svcCtx)
	return l.UpdateCurricula(in)
}

func (s *CurriculacenterServer) SearchCurricula(ctx context.Context, in *pb.SearchCurriculaRequest) (*pb.SearchCurriculaResponse, error) {
	l := logic.NewSearchCurriculaLogic(ctx, s.svcCtx)
	return l.SearchCurricula(in)
}

func (s *CurriculacenterServer) CurriculumDetail(ctx context.Context, in *pb.CurriculumDetailRequest) (*pb.CurriculumDetailResponse, error) {
	l := logic.NewCurriculumDetailLogic(ctx, s.svcCtx)
	return l.CurriculumDetail(in)
}

func (s *CurriculacenterServer) CheckCharacteristics(ctx context.Context, in *pb.CheckCharacteristicsRequest) (*pb.CheckCharacteristicsResponse, error) {
	l := logic.NewCheckCharacteristicsLogic(ctx, s.svcCtx)
	return l.CheckCharacteristics(in)
}

func (s *CurriculacenterServer) CollectCurriculum(ctx context.Context, in *pb.CollectCurriculumRequest) (*pb.CollectCurriculumResponse, error) {
	l := logic.NewCollectCurriculumLogic(ctx, s.svcCtx)
	return l.CollectCurriculum(in)
}
