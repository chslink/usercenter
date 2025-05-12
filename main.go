package main

import (
	"flag"
	"os"

	"usercenter/bc"
	"usercenter/config"
	"usercenter/pkg/server"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "config/config.yaml", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, s *server.Server, loader *bc.Loader) *kratos.App {
	err := loader.Migrate()
	if err != nil {
		panic(err)
	}
	loader.StartService()

	app := kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			s.Grpc,
			s.Http,
		),
	)

	return app
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	conf, err := config.LoadConfig(flagconf)
	if err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(conf, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
