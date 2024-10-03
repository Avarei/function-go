// Package main implements a Composition Function.
package main

import (
	"context"

	"github.com/alecthomas/kong"

	"github.com/crossplane/function-sdk-go"
	v1 "github.com/crossplane/function-sdk-go/proto/v1"
)

// CLI of this Function.
type CLI struct {
	Debug bool `short:"d" help:"Emit debug logs in addition to info logs."`

	Network     string `help:"Network on which to listen for gRPC connections." default:"tcp"`
	Address     string `help:"Address at which to listen for gRPC connections." default:":9443"`
	TLSCertsDir string `help:"Directory containing server certs (tls.key, tls.crt) and the CA used to verify client certificates (ca.crt)" env:"TLS_SERVER_CERTS_DIR"`
	Insecure    bool   `help:"Run without mTLS credentials. If you supply this flag --tls-server-certs-dir will be ignored."`

	ImagePullSecretName []string `help:"ImagePullSecret to use to pull plugins"`
}

// Run this Function.
func (c *CLI) Run() error {
	log, err := function.NewLogger(c.Debug)
	if err != nil {
		return err
	}

	log.Debug("test")
	f := Function{
		log: log,
	}
	f.RunFunction(context.Background(), &v1.RunFunctionRequest{})
	return nil

	// return function.Serve(&Function{log: log},
	// 	function.Listen(c.Network, c.Address),
	// 	function.MTLSCertificates(c.TLSCertsDir),
	// 	function.Insecure(c.Insecure))
}

func main() {
	ctx := kong.Parse(&CLI{}, kong.Description("A Crossplane Composition Function."))
	ctx.FatalIfErrorf(ctx.Run())
}
