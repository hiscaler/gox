package filepathx

import (
	"github.com/hiscaler/gox/inx"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	searchDir = iota
	searchFile
)

type WalkOption struct {
	FilterFunc    func(path string) bool // 自定义函数，返回 true 则会加到列表中，否则忽略。当定义该函数时，将会忽略掉 Except, Only 设置
	Except        []string               // 排除的文件或者目录（仅当 FilterFunc 未设置时起作用）
	Only          []string               // 仅仅符合列表中的文件或者目录才会返回（仅当 FilterFunc 未设置时起作用）
	CaseSensitive bool                   // 是否区分大小写（作用于 Except 和 Only 设置）
	Recursive     bool                   // 是否递归查询下级目录
}

func readDir(root string, recursive bool, searchType int) []string {
	dfs := os.DirFS(root)
	paths := make([]string, 0)
	if recursive {
		fs.WalkDir(dfs, ".", func(path string, d fs.DirEntry, err error) error {
			if err == nil && path != "." && path != ".." &&
				((searchType == searchDir && d.IsDir()) || (searchType == searchFile && !d.IsDir())) {
				paths = append(paths, path)
			}
			return nil
		})
	} else {
		ds, err := fs.ReadDir(dfs, ".")
		if err == nil {
			for _, d := range ds {
				if d.Name() != "." && d.Name() != ".." &&
					((searchType == searchDir && d.IsDir()) || (searchType == searchFile && !d.IsDir())) {
					paths = append(paths, filepath.Join(root, d.Name()))
				}
			}
		}
	}
	return paths
}

func filterPath(path string, opt WalkOption) (ok bool) {
	if (opt.FilterFunc == nil && len(opt.Only) == 0 && len(opt.Except) == 0) ||
		(opt.FilterFunc != nil && opt.FilterFunc(path)) {
		return true
	}

	if len(opt.Except) > 0 || len(opt.Only) > 0 {
		name := filepath.Base(path)
		if len(opt.Except) > 0 {
			if opt.CaseSensitive {
				ok = !inx.StringIn(path, opt.Except...)
			} else {
				ok = true
				for _, s := range opt.Except {
					if s == name {
						ok = false
						break
					}
				}
			}
		}
		if len(opt.Only) > 0 {
			if opt.CaseSensitive {
				ok = inx.StringIn(path, opt.Only...)
			} else {
				for _, s := range opt.Only {
					if s == name {
						ok = true
						break
					}
				}
			}
		}
	}
	return
}

// Dirs 获取指定目录下的所有目录
func Dirs(root string, opt WalkOption) []string {
	dirs := make([]string, 0)
	paths := readDir(root, opt.Recursive, searchDir)
	if len(paths) > 0 {
		for _, path := range paths {
			if filterPath(path, opt) && !strings.EqualFold(path, root) {
				dirs = append(dirs, path)
			}
		}
	}
	return dirs
}

// Files 获取指定目录下的所有文件
func Files(root string, opt WalkOption) []string {
	files := make([]string, 0)
	paths := readDir(root, opt.Recursive, searchFile)
	if len(paths) > 0 {
		for _, path := range paths {
			if filterPath(path, opt) {
				files = append(files, path)
			}
		}
	}
	return files
}
