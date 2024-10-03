package plugin

import (
	"fmt"
	"plugin"

	fnv1 "github.com/crossplane/function-sdk-go/proto/v1"
)

type PluginCache struct {
	cache map[string]fnv1.FunctionRunnerServiceServer
}

func New() PluginCache {
	return PluginCache{
		cache: make(map[string]fnv1.FunctionRunnerServiceServer),
	}
}

func (c *PluginCache) Load(name, pluginPath string) (fnv1.FunctionRunnerServiceServer, error) {

	if functionRunner, ok := c.cache[name]; ok {
		return functionRunner, nil
	}

	plug, err := plugin.Open(pluginPath)
	if err != nil {
		return nil, err
	}

	symPlugin, err := plug.Lookup("Run")
	if err != nil {
		return nil, err
	}

	var p fnv1.FunctionRunnerServiceServer
	p, ok := symPlugin.(fnv1.FunctionRunnerServiceServer)
	if !ok {
		return nil, fmt.Errorf("unexpected type from module symbol")
	}
	c.cache[name] = p
	return p, nil
}
