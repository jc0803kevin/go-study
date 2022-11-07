package main

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"net"
	"net/http"
)

// NewServeMux builds a ServeMux that will route requests
// to the given EchoHandler.
//func NewServeMux(echo *EchoHandler) *http.ServeMux {
//	mux := http.NewServeMux()
//	mux.Handle("/echo", echo)
//	return mux
//}

// 显示定义，如果我们添加更多的处理程序，这将很快变得不方便。
// NewServeMux builds a ServeMux that will route requests
// to the given Route.
//func NewServeMux(route Route, route2 Route) *http.ServeMux {
//	mux := http.NewServeMux()
//	mux.Handle(route.Pattern(), route)
//	mux.Handle(route2.Pattern(), route2)
//	return mux
//}

func NewServeMux(routes []Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}
	return mux
}

// Route is an http.Handler that knows the mux pattern
// under which it will be registered.
type Route interface {
	http.Handler

	// Pattern reports the path at which this is registered.
	Pattern() string
}

// AsRoute annotates the given constructor to state that
// it provides a route to the "routes" group.
func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}

func NewHttpServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
	srv := &http.Server{Addr: ":18080", Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			//fmt.Println("Starting HTTP server at", srv.Addr)
			log.Info("Starting HTTP server at", zap.String("addr", srv.Addr))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

func RunServer() {
	fx.New(

		//使用Provide将具体反射的类型添加到container中
		//可以按需添加任意多个构造函数

		//fx.Provide(NewHttpServer,
		//	NewServeMux,
		//	NewEchoHandler,
		//	zap.NewExample,
		//	),

		//fx.Provide(NewHttpServer,
		//	//NewServeMux,
		//	fx.Annotate(NewServeMux,
		//		fx.ParamTags(`name:"echo"`, `name:"hello"`),
		//	),
		//	fx.Annotate(NewEchoHandler,
		//		fx.As(new(Route)),
		//		fx.ResultTags(`name:"echo"`),
		//	),
		//	fx.Annotate(
		//		NewHelloHandler,
		//		fx.As(new(Route)),
		//		fx.ResultTags(`name:"hello"`),
		//	),
		//	zap.NewExample,
		//),

		fx.Provide(
			NewHttpServer,

			fx.Annotate(
				NewServeMux,
				fx.ParamTags(`group:"routes"`),
			),

			AsRoute(NewEchoHandler),
			AsRoute(NewHelloHandler),
			zap.NewExample,
		),

		fx.Invoke(func(server *http.Server) {}),

		// 配置替换fx默认的日志框架
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	).Run()
}
