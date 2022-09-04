{{define "AddNode"}}
insert into `tree_relation`  (`node_id`,`title`,`parent_id`,`is_leaf`,`order`,`depth`,`path`,`ext`) values(
        '{{.NodeID}}',
        '{{.Title}}',
        '{{.ParentID}}',
        '{{.IsLeaf}}',
        '{{.Order}}',
        {{.Depth}},
        '{{.Path}}',
        '{{.Ext}}'
    );
{{end}}

{{define "BatchAddNode"}}
insert into `tree_relation`  (`node_id`,`title`,`parent_id`,`is_leaf`,`order`,`depth`,`path`,`ext`) values
    {{range $k, $v:=. }}
        {{/* 第一个不加,*/}}
        {{if gt $k 0}}
        ,
        {{end}}
        (
            '{{$v.NodeID}}',
            '{{$v.Title}}',
            '{{$v.ParentID}}',
            '{{$v.IsLeaf}}',
            '{{$v.Order}}',
            {{$v.Depth}},
            '{{$v.Path}}',
            '{{$v.Ext}}'
        )
    {{end}}
    ;
{{end}}



{{define "GetNode"}}
select * from `tree_relation` where `node_id`='{{.NodeID}}' and `deleted_at` ='{{zeroTime}}';
{{end}}

{{define "GetSubTreeLimitDepth"}}
        select * from `tree_relation` where `path` like '{{.ParentPath}}%' and `deleted_at` ='{{zeroTime}}' 
        {{if .Depth }}
            and `depth`<{{.Depth}}
        {{end}}
        order by `depth` asc,`order` asc;
{{end}}



{{define "GetSubTreeNodeCount"}}
set @path=(select path from `tree_relation` where `node_id`='{{.NodeID}}');
select count(*) from `tree_relation` where `path` like concat(@path,'%') and `deleted_at` ='{{zeroTime}}';
{{end}}


{{define "MoveSubTree"}}
start transaction;
update `tree_relation` set `parent_id`='{{.NewParentID}}',`path`='{{.NewPath}}',`depth`= `depth` + {{.DiffDepth}} where `node_id`='{{.NodeID}}';
update `tree_relation` set `path`=replace(`path`,'{{.OldPath}}','{{.NewPath}}'),`depth`= `depth` + {{.DiffDepth}} where `path` like '{{.OldPath}}%';
commit;
{{end}}

{{/* 删除节点和子节点，删除关联节点的 id 集合变量为 @nodeIds */}}
{{define "DeleteSubTree"}}
update `tree_relation` set `deleted_at`='{{currentTime}}' where `path` like '{{.NodePathPrefix}}%';
{{end}}