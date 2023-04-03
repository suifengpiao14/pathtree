package repository

import (
	"fmt"
	"testing"
)

func TestDataFile(t *testing.T) {
	data := `[{"code":"1002","createdAt":"2022-09-26 14:01:20","deletedAt":"","depth":"2","firstLetter":"","id":"2","isDeprecated":"0","label":"city","parentCode":"1001","path":"/1001/1002","title":"城市","updatedAt":"2022-09-26 14:20:26"},{"code":"1003","createdAt":"2022-09-26 14:19:17","deletedAt":"","depth":"0","firstLetter":"","id":"3","isDeprecated":"0","label":"area","parentCode":"1002","path":"","title":"区域","updatedAt":"2022-09-26 14:20:32"},{"code":"1002","createdAt":"2022-09-26 16:10:48","deletedAt":"","depth":"2","firstLetter":"","id":"4","isDeprecated":"","label":"","parentCode":"1001","path":"/1001/1002","title":"子节点","updatedAt":"2022-09-26 16:10:48"},{"code":"1002","createdAt":"2022-09-26 16:21:25","deletedAt":"","depth":"2","firstLetter":"","id":"5","isDeprecated":"","label":"","parentCode":"1001","path":"/1001/1002","title":"子节点","updatedAt":"2022-09-26 16:21:25"},{"code":"1003","createdAt":"2022-09-27 10:22:22","deletedAt":"","depth":"3","firstLetter":"","id":"6","isDeprecated":"","label":"area","parentCode":"1002","path":"/1001/1002","title":"区域","updatedAt":"2022-09-27 10:22:22"},{"code":"1003","createdAt":"2022-09-27 10:44:54","deletedAt":"","depth":"2","firstLetter":"","id":"7","isDeprecated":"0","label":"area","parentCode":"1002","path":"/1001/1002/1003","title":"区域","updatedAt":"2022-09-27 10:44:54"},{"code":"1003","createdAt":"2022-09-27 10:49:06","deletedAt":"","depth":"3","firstLetter":"","id":"8","isDeprecated":"0","label":"area","parentCode":"1002","path":"/1001/1002/1003","title":"区域","updatedAt":"2022-09-27 10:49:06"},{"code":"1003","createdAt":"2022-09-27 10:50:43","deletedAt":"","depth":"3","firstLetter":"","id":"9","isDeprecated":"0","label":"area","parentCode":"1002","path":"/1001/1002/1003","title":"区域","updatedAt":"2022-09-27 10:50:43"},{"code":"1003","createdAt":"2022-09-27 13:48:57","deletedAt":"","depth":"1","firstLetter":"","id":"10","isDeprecated":"0","label":"area","parentCode":"1002","path":"/1003","title":"区域","updatedAt":"2022-09-27 13:48:57"},{"code":"1002","createdAt":"2022-09-28 09:41:37","deletedAt":"","depth":"2","firstLetter":"","id":"11","isDeprecated":"0","label":"","parentCode":"1001","path":"/1001/1002","title":"子节点","updatedAt":"2022-09-28 09:41:37"},{"code":"1003","createdAt":"2022-09-28 09:48:31","deletedAt":"","depth":"1","firstLetter":"","id":"12","isDeprecated":"0","label":"area","parentCode":"1002","path":"/1003","title":"区域","updatedAt":"2022-09-28 09:48:31"}]`
	out, err := DataFill(data)
	if err != nil {
		panic(err)
	}
	fmt.Sprintln(out)
}
