package pathtree_test

import (
	"database/sql"
	"log"
	"net/http"
	"testing"

	"github.com/suifengpiao14/pathtree"
)

func TestFs(t *testing.T) {
	// 创建一个数据库连接
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建 TreeFS 实例
	treeFS := &pathtree.TreeFS{DB: db}

	// 作为示例，将 TreeFS 用于一个简单的 HTTP 服务器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.FS(treeFS)).ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
