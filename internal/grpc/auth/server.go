package auth

import ssov1 "github.com/N1gaz/protos/gen/go/sso"

type serverApi struct {
	ssov1.UnimplementedAuthServer
}
