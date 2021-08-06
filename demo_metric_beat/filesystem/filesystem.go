package filesystem

import (
	"demo_metric-beat/util"
	"fmt"
	"github.com/prometheus/common/log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

const mountTableFileName = "/etc/mtab"

func PrintMetric() {
	fss, err := getFileSystemList()
	if err != nil {
		panic(err)
	}
	log.Info(fmt.Sprintf("lenth is %d", len(fss)))
	for _, fs := range fss {
		stat, err := getFileSystemStat(fs)
		if err != nil {
			panic(err)
		}
		fmt.Println(stat)
	}
}

func getFileSystemList() ([]FileSystem, error) {
	fsList := make([]FileSystem, 0)
	var count = 0
	err := util.ReadFile(mountTableFileName, func(line string) bool {
		fields := strings.Fields(line)
		fs := FileSystem{}
		fs.DevName = fields[0]
		fs.DirName = fields[1]
		fs.SysTypeName = fields[2]
		fs.Options = fields[3]
		fsList = append(fsList, fs)
		count = count + 1
		return true
	})
	fmt.Println("count is ", count)
	fmt.Println("fsList is ", len(fsList))
	return filterFileSystemList(fsList), err
}

func filterFileSystemList(fsList []FileSystem) []FileSystem {
	var filtered []FileSystem
	devices := make(map[string]FileSystem)
	fmt.Println("begin to filter ", len(fsList))
	for _, fs := range fsList {
		// Ignore relative mount points, which are present for example
		// in /proc/mounts on Linux with network namespaces.
		if !filepath.IsAbs(fs.DirName) {
			fmt.Println("ignore abs dir ", fs.DirName)
			continue
		}

		//// Don't do further checks in special devices
		if !filepath.IsAbs(fs.DevName) {
			fmt.Println("ignore abs dev ", fs.DirName)
			filtered = append(filtered, fs)
			continue
		}

		// If the device name is a directory, this is a bind mount or nullfs,
		// don't count it as it'd be counting again its parent filesystem.
		devFileInfo, _ := os.Stat(fs.DevName)
		if devFileInfo != nil && devFileInfo.IsDir() {
			continue
		}

		// If a block device is mounted multiple times (e.g. with some bind mounts),
		// store it only once, and use the shorter mount point path.
		if seen, found := devices[fs.DevName]; found {
			if len(fs.DirName) < len(seen.DirName) {
				devices[fs.DevName] = fs
			}
			continue
		}

		devices[fs.DevName] = fs
	}

	for _, fs := range devices {
		filtered = append(filtered, fs)
	}

	return filtered
}

func getFileSystemStat(fs FileSystem) (*FSStat, error) {
	usage, err := getFileSystemUsage(fs.DirName)
	if err != nil {
		return nil, err
	}

	filesystem := FSStat{
		FileSystemUsage: usage,
		DevName:         fs.DevName,
		Mount:           fs.DirName,
		SysTypeName:     fs.SysTypeName,
	}

	return &filesystem, nil
}

func getFileSystemUsage(path string) (FileSystemUsage, error) {
	result := FileSystemUsage{}
	stat := syscall.Statfs_t{}
	err := syscall.Statfs(path, &stat)
	if err != nil {
		return result, err
	}

	result.Total = uint64(stat.Blocks) * uint64(stat.Bsize)
	result.Free = uint64(stat.Bfree) * uint64(stat.Bsize)
	result.Avail = uint64(stat.Bavail) * uint64(stat.Bsize)
	result.Used = result.Total - result.Free
	result.Files = stat.Files
	result.FreeFiles = uint64(stat.Ffree)
	return result, nil
}

// FSStat contains filesystem metrics
type FSStat struct {
	FileSystemUsage
	DevName     string  `json:"device_name"`
	Mount       string  `json:"mount_point"`
	UsedPercent float64 `json:"used_p"`
	SysTypeName string  `json:"type"`
	ctime       time.Time
}

// FileSystemUsage contains basic stats for the specified filesystem
type FileSystemUsage struct {
	Total     uint64
	Used      uint64
	Free      uint64
	Avail     uint64
	Files     uint64
	FreeFiles uint64
}

type FileSystem struct {
	DirName     string
	DevName     string
	TypeName    string
	SysTypeName string
	Options     string
	Flags       uint32
}
