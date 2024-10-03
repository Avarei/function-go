package main

import (
	"context"

	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/function-go/input/v1beta1"
	"github.com/crossplane/function-go/internal/plugin"
	fnv1 "github.com/crossplane/function-sdk-go/proto/v1"
	"github.com/crossplane/function-sdk-go/request"
	"github.com/crossplane/function-sdk-go/response"
	"github.com/pkg/errors"
)

// Function returns whatever response you ask it to.
type Function struct {
	fnv1.UnimplementedFunctionRunnerServiceServer

	log logging.Logger
}

// RunFunction runs the Function.
func (f *Function) RunFunction(ctx context.Context, req *fnv1.RunFunctionRequest) (*fnv1.RunFunctionResponse, error) {
	f.log.Info("Running function", "tag", req.GetMeta().GetTag())

	rsp := response.To(req, response.DefaultTTL)

	in := &v1beta1.Input{}
	if err := request.GetInput(req, in); err != nil {
		response.Fatal(rsp, errors.Wrapf(err, "cannot get Function input from %T", req))
	}

	plugin := plugin.New()
	path := "./plugin.so"
	// path := "./plugin.so"
	f.log.Info("Before Loading")
	runner, err := plugin.Load(in.Oci, path)
	if err != nil {
		f.log.Info("Error Loading Plugin", "error", err)
		return nil, err
	}
	f.log.Info("Post Error")

	return runner.RunFunction(ctx, req)
}
