package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"

	"io"
	"os"
)

func Info()types.Info{
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	info, err := cli.Info(ctx)
	return info
}

func runcon(){
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
		Tty:   true,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	//statusCh, errCh := cli.ContainerWait(ctx, resp.ID)
	//select {
	//case err := <-errCh:
	//	if err != nil {
	//		panic(err)
	//	}
	//case <-statusCh:
	//}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}

func ListContainer() []types.Container{
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	return containers
}

func StopContainer(id string) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	fmt.Print("Stopping container ", id, "... ")
	if err := cli.ContainerStop(ctx, id, nil); err != nil {
			panic(err)
		}
	fmt.Println("Success")
}

func StartContainer(id string) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	fmt.Print("Stopping container ", id, "... ")
	options := types.ContainerStartOptions{}
	if err := cli.ContainerStart(ctx, id, options); err != nil {
		panic(err)
	}
	fmt.Println("Success")
}

func DeleteContainer(id string) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	fmt.Print("Stopping container ", id, "... ")
	options := types.ContainerRemoveOptions{}
	if err := cli.ContainerRemove(ctx, id, options); err != nil {
		panic(err)
	}
	fmt.Println("Success")
}

func PrintLog(id string, dst io.Writer) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	options := types.ContainerLogsOptions{ShowStdout: true}
	out, err := cli.ContainerLogs(ctx, id, options)
	if err != nil {
		panic(err)
	}

	io.Copy(dst, out)
}

func Inspect(id string) types.ContainerJSON{
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	ins, err := cli.ContainerInspect(ctx, id)
	if err != nil {
		panic(err)
	}
	return ins
}

func ListImage() []types.ImageSummary{
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	return images
}

func PullImage(ref string) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	out, err := cli.ImagePull(ctx, ref, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer out.Close()

	io.Copy(os.Stdout, out)
}

func DeleteImage(id string){
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	out, err := cli.ImageRemove(ctx, id, types.ImageRemoveOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Print(out[0].Deleted)
}
