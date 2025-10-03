package auth

import (
	"context"

	ssov1 "github.com/N1gaz/protos/gen/go/sso"
	"google.golang.org/grpc"
)

type serverApi struct {
	ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &serverApi{})
}

func (s *serverApi) SignIn(
	ctx context.Context,
	req *ssov1.SignInRequest,
) (*ssov1.SignInResponse, error) {
	panic("Implement me")
}

func (s *serverApi) SignOut(
	ctx context.Context,
	req *ssov1.SignOutRequest,
) (*ssov1.SignOutResponse, error) {
	panic("Implement me")
}

func (s *serverApi) SignUp(
	ctx context.Context,
	req *ssov1.SignUpRequest,
) (*ssov1.SignUpResponse, error) {
	panic("Implement me")
}

func (s *serverApi) RefreshToken(
	ctx context.Context,
	req *ssov1.RefreshTokenRequest,
) (*ssov1.RefreshTokenResponse, error) {
	panic("Implement me")
}
