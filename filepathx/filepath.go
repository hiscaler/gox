package filepathx

import (
	"github.com/hiscaler/gox/filex"
	"github.com/hiscaler/gox/inx"
	"io/fs"
	"mime"
	"net/http"
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
	CaseSensitive bool                   // 区分大小写（作用于 Except 和 Only 设置）
	Recursive     bool                   // 是否递归查询下级目录
}

func read(root string, recursive bool, searchType int) []string {
	dfs := os.DirFS(root)
	paths := make([]string, 0)
	if recursive {
		fs.WalkDir(dfs, ".", func(path string, d fs.DirEntry, err error) error {
			if err == nil && path != "." && path != ".." &&
				((searchType == searchDir && d.IsDir()) || (searchType == searchFile && !d.IsDir())) {
				paths = append(paths, filepath.Join(root, path))
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

	pathPrefix := ""
	if strings.HasPrefix(root, "..") {
		pathPrefix = ".."
	} else if strings.HasPrefix(root, ".") {
		pathPrefix = "."
	}
	if pathPrefix != "" {
		pathPrefix += string(filepath.Separator)
		for i, path := range paths {
			paths[i] = pathPrefix + path
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
				ok = true
				for _, s := range opt.Except {
					if s == name {
						ok = false
						break
					}
				}
			} else {
				ok = !inx.StringIn(name, opt.Except...)
			}
		}
		if len(opt.Only) > 0 {
			if opt.CaseSensitive {
				for _, s := range opt.Only {
					if s == name {
						ok = true
						break
					}
				}
			} else {
				ok = inx.StringIn(name, opt.Only...)
			}
		}
	}
	return
}

// Dirs 获取指定目录下的所有目录
func Dirs(root string, opt WalkOption) []string {
	dirs := make([]string, 0)
	paths := read(root, opt.Recursive, searchDir)
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
	paths := read(root, opt.Recursive, searchFile)
	if len(paths) > 0 {
		for _, path := range paths {
			if filterPath(path, opt) {
				files = append(files, path)
			}
		}
	}
	return files
}

// GenerateDirNames 生成目录名
func GenerateDirNames(s string, n, level int, caseSensitive bool) []string {
	if s == "" {
		return []string{}
	}

	isValidCharFunc := func(r rune) bool {
		return 'A' <= r && r <= 'Z' || 'a' <= r && r <= 'z' || '0' <= r && r <= '9'
	}
	var b strings.Builder
	for _, r := range s {
		if isValidCharFunc(r) {
			b.WriteRune(r)
		}
	}
	if b.Len() == 0 {
		return []string{}
	}

	s = b.String() // Clean s string
	if !caseSensitive {
		s = strings.ToLower(s)
	}
	if n <= 0 {
		return []string{s}
	}

	if level <= 0 {
		level = 1
	}
	names := make([]string, 0)
	sLen := len(s)
	for i := 0; i < sLen; i += n {
		if len(names) == level {
			break
		}

		lastIndex := i + n
		if lastIndex >= sLen {
			lastIndex = sLen
		}
		names = append(names, s[i:lastIndex])
	}
	return names
}

// Ext 获取资源扩展名
func Ext(path string, b []byte) string {
	if path == "" && b == nil {
		return ""
	}

	if b == nil && filex.Exists(path) {
		if b1, err := os.ReadFile(path); err == nil {
			b = b1[:512]
		}
	}
	ext := ""
	if b != nil {
		contentType := http.DetectContentType(b)
		// https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types
		extTypes := map[string][]string{
			".aac": {"audio/aac"},
			".abw": {"application/x-abiword"},
			".arc": {"application/x-freearc"},
			".avi": {"video/x-msvideo"},
			".azw": {"application/vnd.amazon.ebook"},
			// ".bin": {"application/octet-stream"},
			".bmp":    {"image/bmp"},
			".bz":     {"application/x-bzip"},
			".bz2":    {"application/x-bzip2"},
			".csh":    {"application/x-csh"},
			".css":    {"text/css"},
			".csv":    {"text/csv"},
			".doc":    {"application/msword"},
			".docx":   {"application/vnd.openxmlformats-officedocument.wordprocessingml.document"},
			".eot":    {"application/vnd.ms-fontobject"},
			".epub":   {"application/epub+zip"},
			".gif":    {"image/gif"},
			".htm":    {"text/html"},
			".html":   {"text/html"},
			".ico":    {"image/vnd.microsoft.icon"},
			".ics":    {"text/calendar"},
			".jar":    {"application/java-archive"},
			".jpg":    {"image/jpeg"},
			".jpeg":   {"image/jpeg"},
			".js":     {"text/javascript"},
			".json":   {"application/json"},
			".jsonld": {"application/ld+json"},
			".mid":    {"audio/midi", "audio/x-midi"},
			".midi":   {"audio/midi", "audio/x-midi"},
			".mjs":    {"text/javascript"},
			".mp3":    {"audio/mpeg"},
			".mpeg":   {"video/mpeg"},
			".mpkg":   {"application/vnd.apple.installer+xml"},
			".odp":    {"application/vnd.oasis.opendocument.presentation"},
			".ods":    {"application/vnd.oasis.opendocument.spreadsheet"},
			".odt":    {"application/vnd.oasis.opendocument.text"},
			".oga":    {"audio/ogg"},
			".ogv":    {"video/ogg"},
			".ogx":    {"application/ogg"},
			".otf":    {"font/otf"},
			".png":    {"image/png"},
			".pdf":    {"application/pdf"},
			".ppt":    {"application/vnd.ms-powerpoint"},
			".pptx":   {"application/vnd.openxmlformats-officedocument.presentationml.presentation"},
			".rar":    {"application/x-rar-compressed"},
			".rtf":    {"application/rtf"},
			".sh":     {"application/x-sh"},
			".svg":    {"image/svg+xml"},
			".swf":    {"application/x-shockwave-flash"},
			".tar":    {"application/x-tar"},
			".tif":    {"image/tiff"},
			".tiff":   {"image/tiff"},
			".ttf":    {"font/ttf"},
			".txt":    {"text/plain"},
			".vsd":    {"application/vnd.visio"},
			".wav":    {"audio/wav"},
			".weba":   {"audio/webm"},
			".webm":   {"video/webm"},
			".webp":   {"image/webp"},
			".woff":   {"font/woff"},
			".woff2":  {"font/woff2"},
			".xhtml":  {"application/xhtml+xml"},
			".xls":    {"application/vnd.ms-excel"},
			".xlsx":   {"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"},
			".xml":    {"application/xml", "text/xml"},
			".xul":    {"application/vnd.mozilla.xul+xml"},
			".zip":    {"application/zip"},
			".3gp":    {"video/3gpp", "audio/3gpp"},
			".3g2":    {"video/3gpp2", "audio/3gpp2"},
			".7z":     {"application/x-7z-compressed"},
		}
		for k, types := range extTypes {
			for _, v := range types {
				mime.AddExtensionType(k, v)
			}
		}
		if extensions, err := mime.ExtensionsByType(contentType); err == nil && extensions != nil {
			n := len(extensions)
			if n == 1 {
				ext = extensions[0]
			} else {
				typeExt := map[string]string{
					"text/plain; charset=utf-8": ".txt",
					"image/jpeg":                ".jpg",
				}
				if v, exists := typeExt[contentType]; exists {
					ext = v
				} else {
					ext = extensions[0]
				}
			}
		}
	}
	if ext == "" {
		ext = filepath.Ext(path)
	}
	return ext
}
