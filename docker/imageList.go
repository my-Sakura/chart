package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var (
	LocalImages = make(map[string]int64)
	ImageSlice  = make([]string, 0)
)

func ListImage() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{All: true})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		id := image.ID[:10]
		LocalImages[id] = image.Size
		ImageSlice = append(ImageSlice, id)
	}
}
