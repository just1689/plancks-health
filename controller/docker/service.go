package docker

import (
	"context"
	"fmt"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
)

func createService(service *model.Service, contract *model.Contract) {
	log.Debugln(fmt.Sprintf("createService method!"))

	cli, err := client.NewEnvClient()
	ctx := context.Background()
	if err != nil {
		log.Panicln(fmt.Sprintf("Error getting docker client environment: %s", err))
	}

	replicas := uint64(contract.Replicas)

	spec := swarm.ServiceSpec{
		Annotations: swarm.Annotations{
			Name: contract.ServiceName,
		},
		Mode: swarm.ServiceMode{
			Replicated: &swarm.ReplicatedService{
				Replicas: &replicas,
			},
		},
		TaskTemplate: swarm.TaskSpec{
			ContainerSpec: swarm.ContainerSpec{
				Image: service.Image,
			},
			Resources: &swarm.ResourceRequirements{
				Limits: &swarm.Resources{
					MemoryBytes: int64(contract.RequiredMBMemory * 1024 * 1024),
				},
			},
		},
	}

	_, err = cli.ServiceCreate(
		ctx,
		spec,
		types.ServiceCreateOptions{},
	)

	if err != nil {
		log.Errorln(fmt.Sprintf("Error creating docker service: %s", err))
	}
}
