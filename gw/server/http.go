package server

import (
	"context"
	"crypto/x509"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kk-no/expapp/gw/config"
	"github.com/kk-no/expapp/gw/gcp"
	sample "github.com/kk-no/proto-terminal/sample/v1"
	idtoken "github.com/salrashid123/oauth2/idtoken"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type HTTPServer struct {
	mux *runtime.ServeMux
}

var _ Server = (*HTTPServer)(nil)

func NewHTTPServer(ctx context.Context) (*HTTPServer, error) {
	s := &HTTPServer{
		mux: runtime.NewServeMux(),
	}

	conf := config.Conf
	if err := registerEndpoint(ctx, s.mux, conf.ExpAppDomain, conf.ExpAppPort, sample.RegisterSampleServiceHandlerFromEndpoint); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *HTTPServer) Serve(port string) error {
	return http.ListenAndServe(":"+port, s.mux)
}

type registerFunc func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error

func registerEndpoint(ctx context.Context, mux *runtime.ServeMux, domain, port string, register registerFunc) error {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	if gcp.OnGCP() {
		credOpts, err := makeCredentialOptions(ctx, domain)
		if err != nil {
			return err
		}
		opts = credOpts
	}
	endpoint := fmt.Sprintf("%s:%s", domain, port)
	return register(ctx, mux, endpoint, opts)
}

func makeCredentialOptions(ctx context.Context, domain string) ([]grpc.DialOption, error) {
	pool, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}
	cred, err := rpcCredential(ctx, fmt.Sprintf("https://%s", domain))
	if err != nil {
		return nil, err
	}
	return []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(pool, "")),
		grpc.WithPerRPCCredentials(cred),
	}, nil
}

func rpcCredential(ctx context.Context, audience ...string) (credentials.PerRPCCredentials, error) {
	scope := "https://www.googleapis.com/auth/userinfo.email"

	cred, err := google.FindDefaultCredentials(ctx, scope)
	if err != nil {
		return nil, err
	}
	idTokenSource, err := idtoken.IdTokenSource(&idtoken.IdTokenConfig{
		Credentials: cred,
		Audiences:   audience,
	})
	if err != nil {
		return nil, err
	}
	return idtoken.NewIDTokenRPCCredential(ctx, idTokenSource)
}
