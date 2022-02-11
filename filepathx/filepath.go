package filepathx

import (
	"github.com/hiscaler/gox/inx"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	returnDir = iota
	returnFile
)

type WalkOption struct {
	FilterFunc    func(path string) bool // 自定义函数，返回 true 则会加到列表中，否则忽略。当定义该函数时，将会忽略掉 Except, Only 设置
	Except        []string               // 排除的文件或者目录（仅当 FilterFunc 未设置时起作用）
	Only          []string               // 仅仅符合列表中的文件或者目录才会返回（仅当 FilterFunc 未设置时起作用）
	CaseSensitive bool                   // 是否区分大小写（作用于 Except 和 Only 设置）
	Recursive     bool                   // 是否递归查询下级目录
}

func readDir(root string, recursive bool, returnType int) []fs.DirEntry {
	dirEntries := make([]fs.DirEntry, 0)
	if recursive {
		filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err == nil && path != "." && path != ".." &&
				((returnType == returnDir && d.IsDir()) || (returnType == returnFile && !d.IsDir())) {
				dirEntries = append(dirEntries, d)
			}
			return nil
		})
	} else {
		ds, err := os.ReadDir(root)
		if err == nil {
			for _, d := range ds {
				if d.Name() != "." && d.Name() != ".." &&
					((returnType == returnDir && d.IsDir()) || (returnType == returnFile && !d.IsDir())) {
					dirEntries = append(dirEntries, d)
				}
			}
		}
	}
	return dirEntries
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
				ok = !inx.StringIn(path, opt.Only...)
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
	ds := readDir(root, opt.Recursive, returnDir)
	if len(ds) > 0 {
		dirBase := filepath.Base(root)
		dirPath := filepath.Dir(root)
		for _, d := range ds {
			if d.IsDir() && filterPath(d.Name(), opt) && !strings.EqualFold(d.Name(), dirBase) {
				dirs = append(dirs, filepath.Join(dirPath, d.Name()))
			}
		}
	}
	return dirs
}

// Files 获取指定目录下的所有文件
func Files(root string, opt WalkOption) []string {
	files := make([]string, 0)
	ds := readDir(root, opt.Recursive, returnFile)
	if len(ds) > 0 {
		dirPath := filepath.Dir(root)
		for _, d := range ds {
			if !d.IsDir() && filterPath(d.Name(), opt) {
				files = append(files, filepath.Join(dirPath, d.Name()))
			}
		}
	}

	return files
}
