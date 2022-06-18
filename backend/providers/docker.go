package providers

import (
	"context"
	"danielr1996/bashdoard/sse"
	"danielr1996/bashdoard/types"
	dockertypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"time"
)

const URL_LABEL = "de.danielr1996.bashdoard.url"
const ICON_LABEL = "de.danielr1996.bashdoard.icon"
const NAME_LABEL = "de.danielr1996.bashdoard.name"
const ID_LABEL = "de.danielr1996.bashdoard.id"
const PROVIDER = "docker"

type DockerProvider struct {
	entries map[string]types.DashboardEntry
}

func getEntries() map[string]types.DashboardEntry {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	filters := filters.NewArgs()
	filters.Add("label", URL_LABEL)
	containers, err := cli.ContainerList(ctx, dockertypes.ContainerListOptions{Filters: filters})
	if err != nil {
		panic(err)
	}

	entries := make(map[string]types.DashboardEntry)

	for _, container := range containers {
		entries[container.Labels[ID_LABEL]] = types.DashboardEntry{
			Name:     container.Labels[URL_LABEL],
			Url:      container.Labels[ICON_LABEL],
			Icon:     container.Labels[NAME_LABEL],
			Id:       container.Labels[ID_LABEL],
			Provider: PROVIDER,
		}
	}

	return entries
}

func (p *DockerProvider) Push(s *sse.SSE) {
	for {
		entries := getEntries()
		for _, entry := range entries {
			if !s.Contains(entry.Id) {
				s.Add(entry.Id, entry)
			} else if !entry.Equals(s.Get(entry.Id)) {
				s.Update(entry.Id, entry)
			}
		}
		for _, entry := range s.Entries {
			if _, ok := entries[entry.Id]; !ok {
				s.Delete(entry.Id)
			}
		}
		time.Sleep(2 * time.Second)
	}
}
