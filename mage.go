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

func getProto(basePath string) ([]string, error) {
	var files []string
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
func API() error {
	files, err := getProto("api")
	if err != nil {
		return err
	}
	fmt.Println("API files:", files)
	args := []string{"--proto_path=./api", "--proto_path=./third_party",
		"--go_out=paths=source_relative:./api",
		"--go-http_out=paths=source_relative:./api",
		"--go-grpc_out=paths=source_relative:./api",
		"--openapi_out=fq_schema_naming=true,default_response=false:.",
	}
	args = append(args, files...)
	return sh.Run("protoc", args...)
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

func Config() error {
	internalProto, err := getProto("pkg")
	if err != nil {
		return err
	}
	args := []string{"--proto_path=./pkg", "--proto_path=./third_party", "--go_out=paths=source_relative:./pkg"}
	args = append(args, internalProto...)
	return sh.Run("protoc", args...)
}
