package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var (
	Images     = make(map[string]int)
	ImageSlice = make([]string, 0)
)

func TransferListContainer() {
	if len(Images) != 0 {
		for k, _ := range Images {
			delete(Images, k)
		}
	}

	ImageSlice = ImageSlice[:0]
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if _, exist := Images[container.Image]; exist {
			Images[container.Image]++
		} else {
			Images[container.Image] = 1
		}
	}

	for image, _ := range Images {
		ImageSlice = append(ImageSlice, image)
	}
}

