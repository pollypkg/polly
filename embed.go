package polly

import (
	"embed"
	"io/fs"
	"path/filepath"

	"cuelang.org/go/cue/load"
)

//go:embed schema/*
//go:embed util/*
var CUE embed.FS

func Overlay(moduleRoot string) (map[string]load.Source, error) {
	const mod = "github.com/pollypkg/polly"
	prefix := filepath.Join(moduleRoot, "cue.mod", "pkg", mod)

	overlay := make(map[string]load.Source)
	err := fs.WalkDir(CUE, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		data, err := fs.ReadFile(CUE, path)
		if err != nil {
			return err
		}

		p := filepath.Join(prefix, path)
		overlay[p] = load.FromBytes(data)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return overlay, nil
}
