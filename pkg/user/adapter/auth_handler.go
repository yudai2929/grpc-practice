package adapter

import (
	"context"

	"github.com/yudai2929/grpc-practice/pkg/user/usecase"
	pb "github.com/yudai2929/grpc-practice/proto/gen/go/user"
)

func (s *UserServer) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	input := &usecase.SignUpInput{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	output, err := s.authUsecase.SignUp(*input)

	if err != nil {
		return nil, err
	}

	res := &pb.SignUpResponse{
		Token: output.Token,
	}

	return res, nil
}

func (s *UserServer) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	input := &usecase.SignInInput{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	output, err := s.authUsecase.SignIn(*input)

	if err != nil {
		return nil, err
	}

	res := &pb.SignInResponse{
		Token: output.Token,
	}

	return res, nil
}
