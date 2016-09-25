package fs

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
)

type FileMode uint

const (
	ModeRead FileMode = 1 << iota
	ModeWrite
	ModeTruncate
)

//file path object, to differentiate from common string
//file-path treat all paths separated with "/"
//StringSys return system-related string format
type filePath string

//file-path object
func FilePath(str string) filePath {
	var f filePath
	f.Set(str)
	return f
}

//set value
func (this *filePath) Set(str string) {
	*this = filePath(filepath.ToSlash(filepath.Clean(str)))
}

//show as string
func (this filePath) String() string {
	return string(this)
}

//system-related string format
func (this filePath) StringSys() string {
	return filepath.FromSlash(string(this))
}

func (this filePath) SplitList() []string {
	return filepath.SplitList(string(this))
}

func (this filePath) Split() (dir, file string) {
	return filepath.Split(string(this))
}

func (this filePath) Ext() string {
	return filepath.Ext(string(this))
}

func (this filePath) EvalSymlinks() (string, error) {
	return filepath.EvalSymlinks(string(this))
}

func (this filePath) Abs() (string, error) {
	return filepath.Abs(string(this))
}

func (this filePath) Relate(root string) filePath {
	s, _ := filepath.Rel(FilePath(root).String(), string(this))
	return FilePath(s)
}

func (this filePath) RelateGoPath() filePath {
	s, _ := filepath.Rel(GoPath(), string(this))
	return FilePath(s)
}

func (this filePath) RelateWorkPath() filePath {
	s, _ := filepath.Rel(WorkPath(), string(this))
	return FilePath(s)
}

func (this filePath) Walk(walkFn filepath.WalkFunc) error {
	return filepath.Walk(string(this), walkFn)
}

func (this filePath) Base() string {
	return filepath.Base(string(this))
}

func (this filePath) Dir() string {
	return filepath.Dir(string(this))
}

func (this filePath) VolumeName() string {
	return filepath.VolumeName(string(this))
}

func (this filePath) Match(pattern string) (matched bool, err error) {
	return filepath.Match(pattern, string(this))
}

func (this filePath) HasPrefix(prefix string) bool {
	return filepath.HasPrefix(string(this), FilePath(prefix).String())
}

////////////////////////////////
func (this filePath) SplitAll() []string {
	s := this.String()
	maxn := strings.Count(s, "/") + 1
	b := make([]string, maxn, maxn)
	i := maxn - 1
	for ; i >= 0; i-- {
		p, f := filepath.Split(s)
		s = strings.TrimSuffix(p, "/")
		if f != "" {
			b[i] = f
		} else {
			if p != "" {
				b[i] = p
			} else {
				i++
			}
			break
		}
	}
	return b[i:]
}

func Joins(elem ...string) string {
	s := filepath.Join(elem...)
	return FilePath(s).String()
}

func (this filePath) Joins(elem ...string) string {
	return Joins(elem...)
}

func (this filePath) Join(child string) string {
	return Joins(this.String(), child)
}

/////////////////////////////////////////////////////////////////

//func (this filePath) Tree()  {
//	return this
//}
//func (this filePath) CollectSubs(opt FsOption) (subs []string, err error) {
//	return this
//}

//OS operation
func (this filePath) Statistic() (nDir, nFile int, size uint64, info string) {
	var buf bytes.Buffer
	this.Walk(func(path string, info os.FileInfo, err error) error {
		if err == nil {
			if info.IsDir() {
				nDir++
			} else {
				nFile++
				size += uint64(info.Size())
			}
			//fmt.Println(path, info.Name(), info.Size(), info.IsDir())
		} else {
			buf.WriteString(err.Error())
			buf.WriteByte('\n')
			//fmt.Println(err)
		}
		return nil
	})
	info = buf.String()

	//fmt.Printf("%s\n[%s] %ddir(s) %dfile(s) %s\n", info, this.StringSys(), nDir, nFile, FileSize(size))
	return
}

func (this filePath) Remove() error {
	return os.Remove(string(this))
}

func (this filePath) RemoveAll() error {
	return os.RemoveAll(string(this))
}

func (this filePath) Rename(newname string) (newPath filePath, err error) {
	n := FilePath(newname)
	if n.VolumeName() == "" { //related path, then calculate from this.Dir
		n.Set(Joins(this.Dir(), n.String()))
	}
	return n, os.Rename(string(this), n.String())
}
func (this filePath) Truncate(size int64) error {
	return os.Truncate(string(this), size)
}

func (this filePath) MkdirAll(perm os.FileMode) error {
	return os.MkdirAll(string(this), perm)
}

func (this filePath) Open() {
}
func (this filePath) Create() {
}

func (this filePath) Copy(path string) error {
	return nil
}
func (this filePath) Move(path string) error {
	return nil
}
