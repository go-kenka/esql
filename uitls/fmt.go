package uitls

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func GoFmt(p string) error {
	root, _ := os.Getwd()
	p = filepath.Join(root, p)

	cmd := exec.Command("gofmt", "-s", "-w", ".")
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func PkgPath(p string) string {
	local, _ := os.Getwd()
	cmd := exec.Command("go", "list", "-m", "-json")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	err := cmd.Run()
	outStr := string(stdout.Bytes())
	if err != nil {
		return ""
	}

	var mod struct {
		Path      string
		Main      bool
		Dir       string
		GoMod     string
		GoVersion string
	}

	err = json.Unmarshal([]byte(outStr), &mod)
	if err != nil {
		return ""
	}

	relPath := strings.Replace(local, mod.Dir, mod.Path, -1)
	return path.Join(filepath.ToSlash(relPath), p)
}
