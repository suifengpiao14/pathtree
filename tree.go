package tree

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func TreeSQL() (sql string) {
	sql = fmt.Sprintf(`
	create table  if not exists %s(
        node_id varchar(64) not null default "" comment "外部节点标识",
        title varchar(64) not null default "" comment "标题",
        parent_id varchar(64) not null default "" comment "父节点ID",
        %s tinyint(1) not null  default 0 comment "序列号(兄弟节点排序)",
        is_leaf enum(1,2)  default 1 comment "是否为叶子节点1-是,2-否",
        depth tinyint(4) not null  default 0 comment "节点深度",
        path varchar(2048) not null default "/" comment "路径",
        ext varchar(124) not null default "" comment "存储字段",
        created_at datetime  not null default CURRENT_TIMESTAMP  comment "创建时间",
        updated_at datetime  not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP comment "更新时间",
        deleted_at datetime  not null default "0000-00-00 00:00:00"  comment "删除时间",
        primary key (node_id),
        key idx_path(%s(768))
    )ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 comment "树关系模型";
	`, "`tree_relation`", "`order`", "`path`")
	return sql
}

//NodeModel 树结构模型
type NodeModel struct {
	NodeID    string `json:"nodeId"`
	Title     string `json:"title"`
	ParentID  string `json:"parentId"`
	IsLeaf    int    `json:"isLeaf"`
	Depth     int    `json:"depth"`
	Order     int    `json:"order"`
	Path      string `json:"path"`
	Ext       string `json:"ext"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt string `json:"deletedAt"`
}

var zeroTime = "0000-00-00 00:00:00"

func CurrentTime() (datetime string) {
	datetime = time.Now().Format("2006-01-02 13:04:05")
	return datetime
}

func AddNode(node NodeModel) (sql string) {
	sql = fmt.Sprintf("insert into `tree_relation`  (`node_id`,`title`,`parent_id`,`is_leaf`,`order`,`depth`,`path`,`ext`) values('%s','%s','%s',%d,%d,%d,'%s','%s')", node.NodeID, node.Title, node.ParentID, node.IsLeaf, node.Order, node.Depth, node.Path, node.Ext)
	return sql
}

func BatchAddNode(nodes []*NodeModel) (sql string) {
	var w bytes.Buffer
	w.WriteString("insert into `tree_relation`  (`node_id`,`title`,`parent_id`,`is_leaf`,`order`,`depth`,`path`,`ext`) values ")
	for i, node := range nodes {
		if i != 0 {
			w.WriteString(",")
		}
		w.WriteString(fmt.Sprintf("('%s','%s','%s',%d,%d,%d,'%s','%s')", node.NodeID, node.Title, node.ParentID, node.IsLeaf, node.Order, node.Depth, node.Path, node.Ext))
	}
	w.WriteString(";")

	return w.String()
}

func GetNode(nodeId string) (sql string) {
	sql = fmt.Sprintf("select * from `tree_relation` where `node_id`='%s' and `deleted_at` ='%s'", nodeId, zeroTime)
	return sql
}

func UpdatePath(nodeId string, path string) (sql string) {
	sql = fmt.Sprintf("update `tree_relation` set `path`='%s' where `node_id`='%s'", path, nodeId)
	return sql
}

func GetSubTreeLimitDepth(parentPath string, depth int) (sql string) {
	var w bytes.Buffer
	w.WriteString(fmt.Sprintf("select * from `tree_relation` where `path` like '%s%%' and `deleted_at` ='%s'", parentPath, zeroTime))
	if depth > 0 {
		w.WriteString(fmt.Sprintf(" and `depth`=%d", depth))
	}
	w.WriteString("order by `depth` asc,`order` asc;")
	sql = w.String()
	return sql
}

func GetSubTreeNodeCount(nodeId string) (sql string) {
	var w bytes.Buffer
	w.WriteString(fmt.Sprintf("set @path=(select path from `tree_relation` where `node_id`='%s';)", nodeId))
	w.WriteString(fmt.Sprintf("select count(*) from `tree_relation` where `path` like concat(@path,'%%') and `deleted_at` ='%s';", zeroTime))
	sql = w.String()
	return sql
}

func MoveSubTree(newPath string, oldPath string) (sql string) {
	lastIndex := strings.LastIndex(newPath, "/")
	if lastIndex < 0 {
		err := errors.Errorf("newPath required contains char '/',got:%s", newPath)
		panic(err)
	}
	nodeId := newPath[lastIndex+1:]
	newParentPath := newPath[:lastIndex]
	newParentId := ""
	if lastIndex := strings.LastIndex(newParentPath, "/"); lastIndex > -1 {
		newParentId = newParentPath[lastIndex+1:]
	}
	diffDepth := strings.Count(newPath, "/") - strings.Count(oldPath, "/") //计算深度变化量
	var w bytes.Buffer
	w.WriteString("start transaction;")
	w.WriteString(fmt.Sprintf("update `tree_relation` set `parent_id`='%s',`path`='%s',`depth`= `depth` + %d where `node_id`='%s';", newParentId, newPath, diffDepth, nodeId))
	w.WriteString(fmt.Sprintf("update `tree_relation` set `path`=replace(`path`,'%s','%s'),`depth`= `depth` + %d where `path` like '%s%%';", oldPath, newPath, diffDepth, oldPath))
	w.WriteString("commit;")
	return sql
}

// 删除节点和子节点，删除关联节点的 id 集合变量为 @nodeIds
func DeleteSubTree(nodePathPrefix string) (sql string) {
	sql = fmt.Sprintf("update `tree_relation` set `deleted_at`='%s' where `path` like '%s%%';", CurrentTime(), nodePathPrefix)
	return sql
}
