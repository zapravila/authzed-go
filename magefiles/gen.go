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
	ProtoPath = "proto/authzed/api"
	// BufRepository = "buf.build/authzed/api"
	BufRepository = "api"
	BufTag        = "0.1"
	// BufTag        = "0e9f9eafa6d1cf54b972d6bb37785bbf00adb099a"
	// BufTag        = "14854970e303a3b8e2f9f53efbab9897b5c8f2b1"
)

// Proto runs proto codegen
func (Gen) Proto() error {
	bufRef := BufRepository //+ ":" + BufTag
	fmt.Println("generating", bufRef)
	runDirV("magefiles", "buf", "generate", bufRef, "--exclude-path", "api/google", "--exclude-path", "api/protoc-gen-openapiv2")
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
			tmpl.Execute(&b, map[string]string{
				"package": entry.Name(),
				"bufRepo": BufRepository,
				"bufTag":  BufTag,
			})

			versionPath := filepath.Join(ProtoPath, entry.Name(), "zz_generated.version.go")
			if err := os.WriteFile(versionPath, b.Bytes(), 0o644); err != nil {
				return err
			}
		}
	}
	return nil
}
