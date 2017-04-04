package resources

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/docker/swarmkit/api"
)

//--resources=gpus:2;disk:409600;ports:[21000-24000];nvidiaGpus:{/dev/nvidia0,/dev/nvida1}
func ParseCustomizedResources(resContent string) (resources []*api.CustomizedResource, err error) {
	resources = []*api.CustomizedResource{}
	strs := strings.Split(resContent, ";")

	for _, str := range strs {
		var resource *api.CustomizedResource
		res := strings.Split(str, ":")
		if len(res) != 2 {
			return resources, fmt.Errorf("Illegal format %s", str)
		}
		key, value := res

		if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {

		} else if strings.HasPrefix(value, "{") && strings.HasSuffix(value, "}") {

		} else {
			scalar, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return resources, fmt.Errorf("Illegal format %s", str)
			}
			resource = &api.CustomizedResource{
				Name:  key,
				Value: scalar,
			}
		}

	}

}

func parseRange(key string, value string) (resource *api.CustomizedResource, err error) {

}

func parseSet(key string, value string) (resource *api.CustomizedResource, err error) {

}
