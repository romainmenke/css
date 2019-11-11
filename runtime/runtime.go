package runtime

import (
	"io/ioutil"
	"os"
)

type Runtime struct {
	tmpDir string
}

func New() *Runtime {
	tmpDir, err := ioutil.TempDir("", "css-runtime")
	if err != nil {
		panic(err)
	}

	return &Runtime{
		tmpDir: tmpDir,
	}
}

func (r *Runtime) Close() error {
	return os.RemoveAll(r.tmpDir)
}
