package filepathx

import (
	"github.com/hiscaler/gox/inx"
	"io/fs"
	"path/filepath"
)

type WalkOption struct {
	Filter        func(path string) bool // 自定义函数，返回 true 则会加到列表中，否则忽略。当定义该函数时，将会忽略掉 Except, Only 设置
	Except        []string               // 排除的文件或者目录（仅当 Filter 未设置时起作用）
	Only          []string               // 仅仅符合列表中的文件或者目录才会返回（仅当 Filter 未设置时起作用）
	CaseSensitive bool                   // 是否区分大小写（作用于 Except 和 Only 设置）
}

func filterPath(path string, opt WalkOption) (ok bool) {
	if opt.Filter == nil && len(opt.Only) == 0 && len(opt.Except) == 0 {
		return true
	}

	if opt.Filter != nil {
		ok = opt.Filter(path)
	}
	if !ok {
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

func Dirs(root string, opt WalkOption) []string {
	dirs := make([]string, 0)
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err == nil && path != "." && path != ".." && info.IsDir() && filterPath(path, opt) {
			dirs = append(dirs, path)
		}
		return nil
	})
	return dirs
}

func Files(root string, opt WalkOption) []string {
	files := make([]string, 0)
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err == nil && path != "." && path != ".." && !info.IsDir() && filterPath(path, opt) {
			files = append(files, path)
		}
		return nil
	})
	return files
}
