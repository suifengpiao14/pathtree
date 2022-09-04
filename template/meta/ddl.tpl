
{{define "ddlTreeRelation"}}
    create table  if not exists `tree_relation`(
        node_id varchar(64) not null default "" comment "外部节点标识",
        title varchar(64) not null default "" comment "标题",
        parent_id varchar(64) not null default "" comment "父节点ID",
        `order` tinyint(1) not null  default 0 comment "序列号(兄弟节点排序)",
        is_leaf enum("YES","NO")  default 1 comment "是否为叶子节点(YES-是,NO-否)",
        depth tinyint(4) not null  default 0 comment "节点深度",
        path varchar(2048) not null default "/" comment "路径",
        ext varchar(124) not null default "" comment "存储字段",
        created_at datetime  not null default CURRENT_TIMESTAMP  comment "创建时间",
        updated_at datetime  not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP comment "更新时间",
        deleted_at datetime  not null default "0000-00-00 00:00:00"  comment "删除时间",
        primary key (`node_id`),
        key `idx_path`(`path`(768))
    )ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 comment "树关系模型";
{{end}}