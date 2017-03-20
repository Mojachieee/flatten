package osabstraction

import (
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type OSWrapper interface {
	// Copy(src, dst FileInfo) error
	// Move(src, dst FileInfo) error
	Copy(src, dst string) error
	Move(src, dst string) error
	// Stat(name string) (FileInfo, error)
	// IsNotExist(err error) bool
	GetFiles(dir string, includeBaseFiles bool) ([]FileInfo, error)
	IsRegularFile(p string) bool
	IsDirectory(p string) bool
}

type FileInfo interface {
	IsDir() bool
	FullPath() string
	Name() string
	Directory() string
	Ext() string
	BaseName() string
	Level() int
}

type File string

func (f File) IsDir() bool {
	fp := f.FullPath()
	fi, err := os.Stat(fp)
	if err != nil {
		return false
	}
	mode := fi.Mode()
	if mode.IsDir() {
		return true
	}
	return false
}

func (f File) FullPath() string {
	return path.Clean(string(f))
}

func (f File) Name() string {
	return filepath.Base(f.FullPath())
}

func (f File) Directory() string {
	j := strings.LastIndex(f.FullPath(), "/")
	dir := f.FullPath()[:j]
	if dir == "" {
		return "/"
	}
	return dir
}

func (f File) Ext() string {
	return filepath.Ext(f.FullPath())
}

func (f File) BaseName() string {
	filename := f.Name()
	ext := f.Ext()
	j := strings.LastIndex(filename, ext)
	return filename[:j]
}

func (f File) Level() int {
	fp := f.FullPath()
	if fp == "" || fp == "/" {
		return 0
	}
	return strings.Count(fp, "/")
}

type RealOS struct{}

func (ros RealOS) Copy(src, dst string) error {
	// TODO: implement this!
	sourceFile := File(src)
	stat, err := ros.Stat(src)
	if err != nil {
		return errors.New(src + " does not exist in file system")
	}
	if stat.IsDir() {
		return errors.New(src + " is a directory")
	}
	stat, err = ros.Stat(dst)
	if err == nil {
		return errors.New(dst + " already exists in file system")
	}
	in, err := os.Open(sourceFile.FullPath())
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(in, out)
	cerr := out.Close()
	if err != nil {
		return err
	}
	return cerr
}

func (ros RealOS) Move(src, dst string) error {
	// TODO: implement this!
	return nil
}

func (ros RealOS) Stat(name string) (FileInfo, error) {
	// TODO
	return nil, nil
}

func (ros RealOS) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

func (ros RealOS) GetFiles(dir string, includeBaseFiles bool) ([]FileInfo, error) {
	return nil, nil
	// files := []os.FileInfo{}
	// filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
	// 	if !info.IsDir() {
	// 		if !includeBaseFiles {
	// 			if path.Dir(p) != path.Clean(dir) {
	// 				files = append(files, info)
	// 			}
	// 		} else {
	// 			files = append(files, info)
	// 		}
	// 	}
	// 	return nil
	// })
	// return []File{}, nil
}