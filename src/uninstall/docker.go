package uninstall

import (
	"context"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/TwinProduction/go-color"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"

	"github.com/deployKubernetesInCHINA/dkc-command/src"
	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

func runUseDocker(cmdStr string) {

	var cli *client.Client
	abs, _ := filepath.Abs(".")
	imagePath := filepath.Join(abs, "kubespray_cache", "images", config.KubesprayImage)
	if !src.CheckYes("Uninstall " + os.Args[2] + " with local docker ") {
		os.Exit(0)
	}

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		log.Log.Fatal(err.Error())
	}
	defer cli.Close()

	if o, err := os.Open(imagePath); err != nil {
		log.Log.Fatal(err.Error())
	} else {
		defer o.Close()
		if _, err := cli.ImageLoad(ctx, o, true); err != nil {
			log.Log.Fatal(err.Error())
		} else {
			log.Log.Println("loading docker image: ", imagePath)
		}
	}

	logfile, err := os.OpenFile("./uninstall.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Log.Fatal(err.Error())
	}
	defer logfile.Close()

	log.Log.Println("CMD:", "docker run -ti --mount type=bind,source="+abs+",dst=/kubespray quay.io/kubespray/kubespray:v2.15.1 sh -c "+"\""+cmdStr+"\"")

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "quay.io/kubespray/kubespray:v2.15.1",
		Cmd:   []string{"sh", "-c", cmdStr},
		Tty:   true,
		//AttachStderr: true,
		//AttachStdout: true,
	}, &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: abs,
				Target: "/kubespray",
			},
		},
	}, nil, nil, "")
	if err != nil {
		log.Log.Fatal(err.Error())
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		log.Log.Fatal(err.Error())
	}

	cmd := exec.Command("docker", "logs", "-f", resp.ID)
	cmd.Stdout = io.MultiWriter(os.Stdout, logfile)
	cmd.Stderr = io.MultiWriter(os.Stdout, logfile)
	log.Log.Println("CMD:", cmd.String())
	err = cmd.Run()
	if err != nil {
		log.Log.Fatal(color.Ize(color.Red, err.Error()))
	}
}
