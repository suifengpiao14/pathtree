{{define "AddCarAttr"}}
insert into `car_attribute`  (`attr_key`,`title`,`parent_id`,`is_leaf`,`order`,`depth`,`path`,`ext`) values(
        "{{.AttrKey}}",
        "{{.Title}}",
        "{{.ParentID}}",
        "{{.IsLeaf}}",
        "{{.Order}}",
        {{.Depth}},
        "{{.Path}}",
        "{{.Ext}}"
    );
{{end}}

{{define "BatchAddCarAttr"}}
insert into `car_attribute`  (`attr_key`,`title`,`parent_id`,`is_leaf`,`order`,`depth`,`path`,`ext`) values
    {{range $k, $v:=. }}
        {{/* 第一个不加,*/}}
        {{if gt $k 0}}
        ,
        {{end}}
        (
            "{{$v.AttrKey}}",
            "{{$v.Title}}",
            "{{$v.ParentID}}",
            "{{$v.IsLeaf}}",
            "{{$v.Order}}",
            {{$v.Depth}},
            "{{$v.Path}}",
            "{{$v.Ext}}"
        )
    {{end}}
    ;
{{end}}



{{define "GetAttr"}}
select * from `car_attribute` where `attr_key`="{{.AttrKey}}" and `deleted_at` ="{{zeroTime}}";
{{end}}

{{define "GetSubAttrLimitDepth"}}
        select * from `car_attribute` where `path` like "{{.ParentPath}}%" and `deleted_at` ="{{zeroTime}}" 
        {{if .Depth }}
            and `depth`<{{.Depth}}
        {{end}}
        order by `depth` asc,`order` asc;
{{end}}



{{define "GetSubAttrLimitDepth"}}
set @path=(select path from `car_attribute` where `attr_key`="{{.AttrKey}}");
select count(*) from `car_attribute` where `path` like concat(@path,"%") and `deleted_at` ="{{zeroTime}}";
{{end}}


{{define "MoveSubAttr"}}
start transaction;
update `car_attribute` set `parent_id`="{{.NewParentID}}",`path`="{{.NewPath}}",`depth`= `depth` + {{.DiffDepth}} where `attr_key`="{{.AttrKey}}";
update `car_attribute` set `path`=replace(`path`,"{{.OldPath}}","{{.NewPath}}"),`depth`= `depth` + {{.DiffDepth}} where `path` like "{{.OldPath}}%";
commit;
{{end}}

{{/* 删除节点和子节点，删除关联节点的 id 集合变量为 @nodeIds */}}
{{define "DeleteSubAttr"}}
update `car_attribute` set `deleted_at`="{{currentTime}}" where `path` like "{{.NodePathPrefix}}%";
{{end}}