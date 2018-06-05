package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
	"log"
)

//CreateOverlayNetwork creates an overlay network in docker swarm
func CreateOverlayNetwork(name string) (success bool, err error) {
	cli, err := client.NewEnvClient()
	ctx := context.Background()
	if err != nil {
		log.Panicln(fmt.Sprintf("Error getting docker client environment: %s", err))
		return false, err
	}

	res, err := cli.NetworkCreate(ctx, name, types.NetworkCreate{Driver: "overlay", Attachable: true})

	logrus.Infoln(fmt.Sprintf(res.ID))
	logrus.Infoln(fmt.Sprintf(res.Warning))

	if err != nil {
		logrus.Infoln(fmt.Sprintf(err.Error()))
		return false, err
	}
	success = true
	return

}

//CheckNetworkExists tells us if a network name exists
func CheckNetworkExists(name string) (exists bool, err error) {
	cli, err := client.NewEnvClient()
	ctx := context.Background()
	if err != nil {
		log.Panicln(fmt.Sprintf("Error getting docker client environment: %s", err))
		return false, err
	}
	list, err := cli.NetworkList(ctx, types.NetworkListOptions{})
	if len(list) == 0 {
		exists = false
		return
	}

	for i := 0; i < len(list); i++ {
		if list[i].Name == name {
			exists = true
			return
		}
	}

	exists = false
	return

}

//DeleteNetwork removes a network by name
func DeleteNetwork(name string) (success bool, err error) {
	cli, err := client.NewEnvClient()
	ctx := context.Background()
	if err != nil {
		log.Panicln(fmt.Sprintf("Error getting docker client environment: %s", err))
		return false, err
	}

	list, err := cli.NetworkList(ctx, types.NetworkListOptions{})
	if len(list) == 0 {
		success = false
		return
	}

	for i := 0; i < len(list); i++ {
		if list[i].Name == name {
			cli.NetworkRemove(ctx, list[i].ID)
			success = true
			return
		}
	}

	success = false
	return

}
