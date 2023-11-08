//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Dev() error {
	cmd := exec.Command("./dev-server.sh")
	cmd.Dir = "./"
	connectStd(cmd)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func Wire() error {
	fmt.Println("wiring...")

	cmd := exec.Command("wire")
	cmd.Dir = "pkg/server"

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func Proto() error {
	twirpProtoFiles := []string{
		"twirp/auth.proto",
	}
	grpcProtoFiles := []string{
		"models.proto",
	}

	target := "core"

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	fmt.Println("generating twirp protobuf")

	args := append([]string{
		"--go_out", target,
		"--twirp_out", target,
		"--go_opt=paths=source_relative",
		"--twirp_opt=paths=source_relative",
		"--plugin=go=" + "protoc-gen-go",
		"--plugin=twirp=" + "protoc-gen-twirp",
		"-I=./proto",
	}, twirpProtoFiles...)
	cmd := exec.Command("protoc", args...)
	connectStd(cmd)
	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("generating grpc protobuf")

	args = append([]string{
		"--go_out", target,
		"--go-grpc_out", target,
		"--go_opt=paths=source_relative",
		"--go-grpc_opt=paths=source_relative",
		"--plugin=go=" + "protoc-gen-go",
		"--plugin=go-grpc=" + "protoc-gen-go-grpc",
		"-I=./proto",
	}, grpcProtoFiles...)
	cmd = exec.Command("protoc", args...)
	connectStd(cmd)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func Generate() error {
	cmd := exec.Command("go", "generate", "./...")
	connectStd(cmd)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func connectStd(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}
