package runtime

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Runtime struct {
	tmpDir string
	wDir   string
}

func New() *Runtime {
	tmpDir, err := ioutil.TempDir("", "css-runtime")
	if err != nil {
		panic(err)
	}

	wDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return &Runtime{
		tmpDir: tmpDir,
		wDir:   wDir,
	}
}

func (r *Runtime) Close() error {
	return os.RemoveAll(r.tmpDir)
}

func (runtime *Runtime) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*30)
	defer cancel()

	r = r.Clone(ctx)

	path := strings.TrimPrefix(r.URL.Path, "/")
	path, err := filepath.Abs(path)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if !strings.HasPrefix(path, runtime.wDir) {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	plugin, err := runtime.CompileGo(ctx, "serve-http-test", file)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	handler, err := runtime.LoadHTTP(ctx, plugin)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	handler.ServeHTTP(w, r)
}
