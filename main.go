package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/gookit/slog"

	flag "github.com/spf13/pflag"
)

func backup(sourceDir, backupDir string, maxBackups int) error {
	// 创建备份目录
	if err := os.MkdirAll(backupDir, os.ModePerm); err != nil {
		return err
	}

	backupFileName := fmt.Sprintf("backup_%s.zip", time.Now().Format("20060102150405"))

	backupPath := filepath.Join(backupDir, backupFileName)
	backupFile, err := os.Create(backupPath)
	if err != nil {
		return err
	}
	defer backupFile.Close()

	zipWriter := zip.NewWriter(backupFile)
	defer zipWriter.Close()

	err = filepath.Walk(sourceDir, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(sourceDir, filePath)
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = relPath

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	// 删除多余的备份文件
	if maxBackups > 0 {
		backups, err := filepath.Glob(filepath.Join(backupDir, "backup_*.zip"))
		if err != nil {
			return err
		}

		sortBackups(backups)

		for i := 0; i < len(backups)-maxBackups; i++ {
			err := os.Remove(backups[i])
			if err != nil {
				return err
			}
			slog.Debugf("Deleted old backup %s\n", backups[i])
		}
	}
	return nil
}

// 将备份文件按照时间戳从新到旧排序
func sortBackups(backups []string) {
	sort.Slice(backups, func(i, j int) bool {
		return getTimeFromBackupName(backups[i]).After(getTimeFromBackupName(backups[j]))
	})
}

// 解析备份文件名中的时间戳
func getTimeFromBackupName(backupName string) time.Time {

	t, err := time.Parse("backup_20060102150405.zip", backupName)
	if err != nil {
		return time.Time{}
	}
	return t
}

func main() {
	var sourceDir string
	var backupDir string
	var maxBackups int
	var interval time.Duration

	flag.StringVarP(&sourceDir, "source", "s", "", "Source directory to backup")
	flag.StringVarP(&backupDir, "backup", "b", "", "Backup directory")
	flag.IntVarP(&maxBackups, "max", "m", 0, "Maximum number of backups to keep")
	flag.DurationVarP(&interval, "interval", "i", time.Hour, "Backup interval")
	flag.Parse()

	if sourceDir == "" || backupDir == "" {
		flag.Usage()
		return
	}

	for {
		slog.Infof("Starting backup of %s to %s\n", sourceDir, backupDir)
		err := backup(sourceDir, backupDir, maxBackups)
		if err != nil {
			slog.Errorf("Backup failed: %s, retrying in %s\n", err, interval)
		} else {
			slog.Infof("Backup completed successfully. Waiting for %s\n", interval)
		}
		time.Sleep(interval)
	}
}
