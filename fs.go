package pathtree

// 实现fs.FS 操作 在mindoc 二次开发中有应用,使得数据下载到本地硬盘、远程硬盘、db 能无缝对接

import (
	"database/sql"
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/**
CREATE TABLE `file_system` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `filename` varchar(1024) NOT NULL DEFAULT '' COMMENT '文件绝对路径',
  `Content` text NOT NULL DEFAULT '' COMMENT '文件内容',
  `is_leaf` varchar(128) NOT NULL DEFAULT '' COMMENT '是否文档',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='文件系统';


**/

// TreeNode 表示数据库中的树形节点
type TreeNode struct {
	Filename  string `db:"filename"`
	Content   string `db:"content"`
	IsLeaf    bool   `db:"is_leaf"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

// TreeFS 实现了 fs.FS 接口用于 MySQL Tree 结构
type TreeFS struct {
	DB *sql.DB
}

var driverName = "mysql"

func (t *TreeFS) MkdirAll(path string, perm os.FileMode) (err error) {
	root := path
	slashIndex := strings.Index(path, "/")
	if slashIndex > -1 {
		root = root[:slashIndex]
	}
	query := "SELECT filename, is_leaf FROM file_system WHERE filename like ? and deleted_at is null"
	db := sqlx.NewDb(t.DB, driverName)
	nodes := make([]TreeNode, 0)
	absolutePath := fmt.Sprintf("%s%", root)
	err = db.Select(&nodes, query, absolutePath)
	if err != nil {
		return err
	}
	return nil
}

type _Repositoy struct {
	db *sqlx.DB
}

func NewRepository(db *sql.DB) _Repositoy {
	return _Repositoy{db: sqlx.NewDb(db, driverName)}
}

// GetTreeNodeInfo 获取记录,不包含文内容
func (r _Repositoy) GetTreeNodeInfo(filename string) (treeNode *TreeNode, err error) {
	treeNode = new(TreeNode)
	err = r.db.Select(treeNode, "SELECT filename, is_leaf, created_at,updated_at FROM file_system WHERE filename = ?", filename)
	if err != nil {
		return nil, err
	}
	return treeNode, nil
}

// GetTreeNodeWithContent 获取记录,包含文内容
func (r _Repositoy) GetTreeNodeWithContent(filename string) (treeNode *TreeNode, err error) {
	treeNode = new(TreeNode)
	err = r.db.Select(treeNode, "SELECT content,filename, is_leaf, created_at,updated_at FROM file_system WHERE filename = ?", filename)
	if err != nil {
		return nil, err
	}
	return treeNode, nil
}

// ReadDir 实现 fs.FS 接口中的 ReadDir 方法
func (t *TreeFS) ReadDir(fullname string) ([]fs.DirEntry, error) {

	// 从数据库中查询指定目录下的文件和子目录信息
	rows, err := t.DB.Query("SELECT filename, is_leaf, created_at,updated_at FROM file_system WHERE filename like '%?'", fullname)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []fs.DirEntry

	for rows.Next() {
		var node TreeNode
		if err := rows.Scan(&node.Filename, &node.IsLeaf, &node.CreatedAt, node.UpdatedAt); err != nil {
			return nil, err
		}

		// 创建 DirEntry
		entry := &TreeDirEntry{
			name:  node.Filename,
			isDir: node.IsLeaf,
			node:  node,
			fs:    t,
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

// Open 实现 fs.FS 接口中的 Open 方法
func (t *TreeFS) Open(fullname string) (fs.File, error) {
	// 查询文件内容并返回对应的文件对象
	row := t.DB.QueryRow("SELECT content,is_leaf,updated_at FROM file_system WHERE filename =? ", fullname)
	var content string
	var isLeaf bool
	var updatedAt string
	if err := row.Scan(&content, &isLeaf, &updatedAt); err != nil {
		return nil, err
	}
	updatedTime, err := time.ParseInLocation("2006-01-02 15:04:05", updatedAt, time.Local)
	if err != nil {
		updatedTime = time.Now()
	}
	file := &TreeFile{
		Fullname:    fullname,
		IsLeaf:      isLeaf,
		UpdatedTime: updatedTime,
		Content:     strings.NewReader(content),
	}
	return file, nil
}

// TreeDirEntry 实现 fs.DirEntry 接口
type TreeDirEntry struct {
	name  string
	isDir bool
	node  TreeNode
	fs    *TreeFS
}

// Name 实现 fs.DirEntry 接口中的 Name 方法
func (d *TreeDirEntry) Name() string {
	return d.name
}

// IsDir 实现 fs.DirEntry 接口中的 IsDir 方法
func (d *TreeDirEntry) IsDir() bool {
	return d.isDir
}

// Type 实现 fs.DirEntry 接口中的 Type 方法
func (d *TreeDirEntry) Type() fs.FileMode {
	if d.isDir {
		return fs.ModeDir
	}
	return 0
}

// Type 实现 fs.FileInfo 接口中的 Type 方法
func (d *TreeDirEntry) Info() (fs.FileInfo, error) {

	return nil, nil
}

// TreeFile 实现 fs.File 接口
type TreeFile struct {
	Fullname    string
	IsLeaf      bool
	UpdatedTime time.Time
	Content     *strings.Reader
}

func (f *TreeFile) Read(b []byte) (n int, err error) {
	return f.Content.Read(b)
}

// Stat 实现 fs.File 接口中的 Stat 方法
func (f *TreeFile) Stat() (fs.FileInfo, error) {
	return &TreeFileInfo{f}, nil
}

// Stat 实现 fs.File 接口中的 Stat 方法
func (f *TreeFile) Close() error {
	return nil
}

// TreeFileInfo 实现 fs.FileInfo 接口
type TreeFileInfo struct {
	*TreeFile
}

// Size 实现 fs.FileInfo 接口中的 Size 方法
func (fi *TreeFileInfo) Name() string {
	return path.Base(fi.Fullname)
}

// Size 实现 fs.FileInfo 接口中的 Size 方法
func (fi *TreeFileInfo) Size() int64 {
	return int64(fi.Content.Len())
}

// Mode 实现 fs.FileInfo 接口中的 Mode 方法
func (fi *TreeFileInfo) Mode() fs.FileMode {
	return 0
}

// ModTime 实现 fs.FileInfo 接口中的 ModTime 方法
func (fi *TreeFileInfo) ModTime() time.Time {

	return fi.UpdatedTime
}

// IsDir 实现 fs.FileInfo 接口中的 IsDir 方法
func (fi *TreeFileInfo) IsDir() bool {
	return !fi.IsLeaf
}

// Sys 实现 fs.FileInfo 接口中的 Sys 方法
func (fi *TreeFileInfo) Sys() interface{} {
	return nil
}
