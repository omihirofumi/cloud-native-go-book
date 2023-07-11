package main

import (
	"fmt"
	"github.com/hashicorp/go-plugin"
	"github.com/omihirofumi/cloud-native-go-book/chap08/hashicorp-plugin/commons"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: run ./chap08/hashicorp-plugin/main/main.go animal")
	}

	name := os.Args[1]
	module := fmt.Sprintf("./chap08/hashicorp-plugin/%s/%s", name, name)

	_, err := os.Stat(module)
	if os.IsNotExist(err) {
		log.Fatalf("can't find an animal named", name)
	}

	var pluginMap = map[string]plugin.Plugin{
		"sayer": &commons.SayerPlugin{},
	}

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: commons.HandshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command(module),
	})
	defer client.Kill()

	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	raw, err := rpcClient.Dispense("sayer")
	if err != nil {
		log.Fatal(err)
	}

	sayer := raw.(commons.Sayer)

	fmt.Printf("A %s says: %q\n", name, sayer.Says())
}
