package uitls

import (
	"bytes"
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
	cmd := exec.Command("go", "list", "-m")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, _ := string(stdout.Bytes()), string(stderr.Bytes())
	if err != nil {
		return ""
	}
	return path.Join(strings.Replace(outStr, "\n", "", -1), p)
}
