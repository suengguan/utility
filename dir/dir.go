package dir

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
)

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func ParentDir(dir string) string {
	return substr(dir, 0, strings.LastIndex(dir, "/"))
}

func CurrentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		beego.Debug(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
