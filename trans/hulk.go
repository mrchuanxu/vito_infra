package trans

import (
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"github.com/VitoChueng/vito_infra/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// HulkGo 执行grpc服务与http结构体
type HulkGo struct {
	RegisterGRPC func(grpcServer *grpc.Server) error
	//RegisterGW              func(ctx context.Context, gateway *runtime.ServeMux, endPoint string, dopts []grpc.DialOption) error
	UnaryServerInterceptors []grpc.UnaryServerInterceptor
}

func HulkGoRun(hulkGo *HulkGo) {
	hulkGo.RunGrpc()
}

const (
	EndingPoint = ":52242"
)

//var gatewayEndingPoint = flag.String("share_endpoint", "localhost:52242", "endpoint of Gateway")
//
// 实现链接
//func (hg *HulkGo) Run() {
//	// 监听本地端口
//	lis, err := net.Listen("tcp", EndingPoint)
//	if err != nil {
//		logger.TransLogger.Sugar().Panicf("port listen fail:[%v]", err)
//		return
//	}
//	srv := hg.newServer(context.Background())
//	logger.TransLogger.Sugar().Infof("gRPC and https listen on %s\\n", EndingPoint)
//	if err := srv.Serve(lis); err != nil {
//		logger.TransLogger.Sugar().Panicf("Run grpc and gateway err:[%v]", err)
//	}
//}

func (hg *HulkGo) RunGrpc() { // work
	// 监听本地端口
	lis, err := net.Listen("tcp", EndingPoint)
	if err != nil {
		logger.TransLogger.Sugar().Panicf("port listen fail:[%v]", err)
		return
	}
	srv := hg.newGrpc()
	logger.TransLogger.Sugar().Infof("gRPC listen on %s\\n", EndingPoint)
	if err := srv.Serve(lis); err != nil {
		logger.TransLogger.Sugar().Panicf("Run grpc Serve err:[%v]", err)
	}
}

func (hg *HulkGo) newGrpc() *grpc.Server {
	// 拦截器
	var serverInterceptors []grpc.UnaryServerInterceptor
	if len(hg.UnaryServerInterceptors) > 0 {
		serverInterceptors = append(serverInterceptors, hg.UnaryServerInterceptors...)
	}
	opts := []grpc.ServerOption{grpc_middleware.WithUnaryServerChain(serverInterceptors...)}
	server := grpc.NewServer(opts...) // grpc请求
	if err := hg.RegisterGRPC(server); err != nil {
		logger.TransLogger.Sugar().Panicf("newGrpc failure err:[%v]", err)
		return nil
	}
	reflection.Register(server)
	return server
}

//func (hg *HulkGo) newGateway() http.Handler {
//	ctx := context.Background()
//	dopts := []grpc.DialOption{
//		grpc.WithInsecure(), // http协议请求
//	}
//	gwmux := runtime.NewServeMux()
//	if err := hg.RegisterGW(ctx, gwmux, *gatewayEndingPoint, dopts); err != nil {
//		logger.TransLogger.Sugar().Panicf("newGateway failure err:[%v]", err)
//		return nil
//	}
//	return gwmux
//}

//func (hg *HulkGo) newServer(ctx context.Context) *http.Server {
//	grpcServer := hg.newGrpc(ctx)
//	gwmux := hg.newGateway()
//
//	mux := http.NewServeMux()
//	mux.Handle("/", gwmux)
//	return &http.Server{
//		Addr:    EndingPoint,
//		Handler: util.GrpcHandlerFunc(grpcServer, mux),
//		// No tls
//	}
//}
