//go:build mage
// +build mage

package main

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/magefile/mage/sh"
)

func Build() error {
	return sh.Run("go", "build", "-o", "bin/app", ".")
}

func Test() error {
	return sh.Run("go", "test", "-v", "./...")
}

func Gen() error {
	return sh.Run("wire", "./...")
}

func API() error {
	return sh.Run("swag", "init")
}

func Init() error {
	install := sh.RunCmd("go", "install")
	var tools = []string{
		`google.golang.org/protobuf/cmd/protoc-gen-go@latest`, `google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`,
		`github.com/go-kratos/kratos/cmd/kratos/v2@latest`, `github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest`,
		`github.com/google/gnostic/cmd/protoc-gen-openapi@latest`, `github.com/google/wire/cmd/wire@latest`,
	}
	var err error
	for _, tool := range tools {
		if err = install(tool); err != nil {
			return err
		}
	}
	return err
}

func getInternalProto() ([]string, error) {
	var files []string
	basePath := "internal"
	err := filepath.Walk(basePath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".proto" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func Config() error {
	internalProto, err := getInternalProto()
	if err != nil {
		return err
	}
	fmt.Println(internalProto)
	args := []string{"--proto_path=./internal", "--proto_path=./third_party", "--go_out=paths=source_relative:./internal"}
	args = append(args, internalProto...)
	fmt.Printf("args: %v\n", args)
	return sh.Run("protoc", args...)
}
