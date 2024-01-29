package adapter

import (
	"context"

	"github.com/yudai2929/grpc-practice/pkg/user/usecase"
	pb "github.com/yudai2929/grpc-practice/proto/gen/go/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *UserServer) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	input := &usecase.GetUserByIDInput{
		ID: req.GetId(),
	}

	output, err := s.userUsecase.GetUserByID(*input)

	if err != nil {
		return nil, err
	}

	res := &pb.GetUserByIDResponse{
		User: &pb.User{
			Id:        output.User.ID,
			Name:      output.User.Name,
			Email:     output.User.Email,
			CreatedAt: timestamppb.New(output.User.CreatedAt),
			UpdatedAt: timestamppb.New(output.User.UpdatedAt),
		},
	}

	return res, nil
}
