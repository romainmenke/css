package runtime

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
)

func (r *Runtime) CompileGo(ctx context.Context, name string, src io.Reader) (string, error) {
	fileName := fmt.Sprintf("%s.go", path.Join(r.tmpDir, name))
	pluginName := fmt.Sprintf("%s.so", path.Join(r.tmpDir, name))

	if fileExists(pluginName) {
		return pluginName, nil
	}

	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(file, src)
	if err != nil {
		return "", err
	}

	cmd := exec.CommandContext(
		ctx,
		"go",
		"build",
		"-buildmode",
		"plugin",
		"-o",
		pluginName,
		fileName,
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		os.Remove(fileName)
		os.Remove(pluginName)
		return "", err
	}

	return pluginName, err
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
