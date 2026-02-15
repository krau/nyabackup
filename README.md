# NyaBackup

> 一个简单、高效的 Go 备份工具 🐱

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go](https://img.shields.io/badge/Go-1.16+-00ADD8?style=flat&logo=go)](https://golang.org)
[![GitHub Repo](https://img.shields.io/badge/Repo-krau%2Fnyabackup-blue?logo=github)](https://github.com/krau/nyabackup)

---

## ✨ 特性

- 📦 **简单易用** - 命令行工具，开箱即用
- 🔄 **定时备份** - 支持自定义备份间隔
- 🗂️ **自动归档** - 将源目录打包为 ZIP 格式
- 🧹 **自动清理** - 支持保留指定数量的历史备份
- 📝 **详细日志** - 实时输出备份进度和状态
- ⚡ **高效可靠** - 使用标准库实现，轻量级

---

## 🚀 快速开始

### 安装

```bash
# 克隆仓库
git clone https://github.com/krau/nyabackup.git
cd nyabackup

# 编译
go build -o nyabackup main.go
```

### 基本使用

```bash
# 备份 /home/user/data 目录到 /backup
./nyabackup -s /home/user/data -b /backup
```

---

## 📖 使用方法

### 命令行参数

| 参数 | 短参数 | 说明 | 默认值 | 必需 |
|------|--------|------|--------|------|
| `--source` | `-s` | 源目录路径 | - | ✅ |
| `--backup` | `-b` | 备份目录路径 | - | ✅ |
| `--max` | `-m` | 保留的最大备份数量 | 0 (不限制) | ❌ |
| `--interval` | `-i` | 备份间隔 | 1h | ❌ |

### 示例

#### 1. 基本备份

每 1 小时备份一次 `/home/user/documents` 到 `/backup/documents`：

```bash
./nyabackup -s /home/user/documents -b /backup/documents
```

#### 2. 限制备份数量

每 30 分钟备份一次，最多保留 5 个历史备份：

```bash
./nyabackup -s /home/user/data -b /backup/data -m 5 -i 30m
```

#### 3. 自定义备份间隔

每 6 小时备份一次：

```bash
./nyabackup -s /var/www -b /backup/www -i 6h
```

#### 4. 每日备份

每 24 小时备份一次：

```bash
./nyabackup -s /home/user/projects -b /backup/projects -i 24h
```

---

## 📂 备份文件格式

备份文件以 ZIP 格式存储，命名规则：

```
backup_YYYYMMDDHHMMSS.zip
```

例如：
```
backup_20240214150430.zip
```

---

## 🔄 备份轮转

当设置了 `--max` 参数时，NyaBackup 会自动清理旧的备份文件：

- 按照时间戳排序（从新到旧）
- 保留最新的 N 个备份
- 删除超出限制的旧备份

示例：设置 `-m 5` 时，会保留最新的 5 个备份，删除第 6 个及更旧的备份。

---

## 📝 日志输出

NyaBackup 使用 `gookit/slog` 提供详细的日志输出：

```
[INFO] Starting backup of /home/user/data to /backup/data
[INFO] Backup completed successfully. Waiting for 1h
[DEBUG] Deleted old backup /backup/data/backup_20240213100000.zip
```

---

## 🛠️ 技术栈

- **语言**: Go
- **命令行参数**: [spf13/pflag](https://github.com/spf13/pflag)
- **日志**: [gookit/slog](https://github.com/gookit/slog)

---

## 📄 许可证

本项目采用 [MIT License](LICENSE) 开源协议。

---

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

---

## 👤 作者

[Krau](https://github.com/krau)

---

## ⭐ Star History

如果这个项目对你有帮助，请给个 Star ⭐

---

## 🐱 Nya!

简单、高效、可靠的备份工具，让你的数据更安全！