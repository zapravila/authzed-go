//go:build mage

package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/magefile/mage/mg"
)

type Gen mg.Namespace

// All runs all generators in parallel
func (g Gen) All() error {
	mg.Deps(g.Proto)
	return nil
}

const (
	ProtoPath     = "proto/authzed/api"
	BufRepository = "api"
	BufTag        = "8a6892e674e78add466b64606dc3ccf5a75f8a97"
)

// Proto runs proto codegen
func (Gen) Proto() error {
	bufRef := BufRepository
	// + ":" + BufTag
	fmt.Println("generating", bufRef)
	runDirV("magefiles", "go", "run", "github.com/bufbuild/buf/cmd/buf", "generate", bufRef, "--exclude-path", "api/google", "--exclude-path", "api/protoc-gen-openapiv2")
	return generateVersionFiles()
	// return nil
}

func generateVersionFiles() error {
	tmpl, err := template.ParseFiles("magefiles/version.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse version template: %w", err)
	}

	entries, err := os.ReadDir(ProtoPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() && !strings.HasSuffix(entry.Name(), "_test") {
			var b bytes.Buffer
			if err := tmpl.Execute(&b, map[string]string{
				"package": entry.Name(),
				"bufRepo": BufRepository,
				"bufTag":  BufTag,
			}); err != nil {
				return fmt.Errorf("failed to execute version template: %w", err)
			}

			versionPath := filepath.Join(ProtoPath, entry.Name(), "zz_generated.version.go")
			if err := os.WriteFile(versionPath, b.Bytes(), 0o644); err != nil {
				return err
			}
		}
	}
	return nil
}
