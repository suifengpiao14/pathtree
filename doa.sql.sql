insert ignore into `source` (`source_id`,`source_type`,`config`) values('doa','SQL','{"logLevel":"debug","dsn":"root:123456@tcp(mysql_address:3306)/doa?charset=utf8&timeout=1s&readTimeout=5s&writeTimeout=5s&parseTime=False&loc=Local&multiStatements=true","timeout":30}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DoaTreeRelationAddNode','SQL','doa tree relation add node','doa tree relation add node','doa','{{define "AddNode"}} insert into `tree_relation` (`node_id`,`title`,`parent_id`,`is_leaf`,`order`,`depth`,`path`,`ext`) values( '{{.NodeID}}', '{{.Title}}', '{{.ParentID}}', '{{.IsLeaf}}', '{{.Order}}', {{.Depth}}, '{{.Path}}', '{{.Ext}}' ); {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`) values('doa-tree_relation-AddNode','doa tree relation add node','doa tree relation add node','POST','/api/doa/v1/tree_relation/add_node','DoaTreeRelationAddNode','{{define "main"}}
{{execSQLTpl . "AddNode"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=depth,dst=Depth,format=number,required
fullname=ext,dst=Ext,required
fullname=isLeaf,dst=IsLeaf,required
fullname=nodeId,dst=NodeID,required
fullname=order,dst=Order,format=number,required
fullname=parentId,dst=ParentID,required
fullname=path,dst=Path,required
fullname=title,dst=Title,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DoaTreeRelationBatchAddNode','SQL','doa tree relation batch add node','doa tree relation batch add node','doa','{{define "BatchAddNode"}} insert into `tree_relation` (`node_id`,`title`,`parent_id`,`is_leaf`,`order`,`depth`,`path`,`ext`) values {{range $k, $v:=. }} {{/* 第一个不加,*/}} {{if gt $k 0}} , {{end}} ( '{{$v.NodeID}}', '{{$v.Title}}', '{{$v.ParentID}}', '{{$v.IsLeaf}}', '{{$v.Order}}', {{$v.Depth}}, '{{$v.Path}}', '{{$v.Ext}}' ) {{end}} ; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`) values('doa-tree_relation-BatchAddNode','doa tree relation batch add node','doa tree relation batch add node','POST','/api/doa/v1/tree_relation/batch_add_node','DoaTreeRelationBatchAddNode','{{define "main"}}
{{execSQLTpl . "BatchAddNode"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in','version=http://json-schema.org/draft-07/schema,id=output,direction=out');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DoaTreeRelationGetNode','SQL','doa tree relation 获取 node','doa tree relation 获取 node','doa','{{define "GetNode"}} select * from `tree_relation` where `node_id`='{{.NodeID}}' and `deleted_at` ='{{zeroTime}}'; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`) values('doa-tree_relation-GetNode','doa tree relation 获取 node','doa tree relation 获取 node','POST','/api/doa/v1/tree_relation/get_node','DoaTreeRelationGetNode','{{define "main"}}
{{execSQLTpl . "GetNode"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=nodeId,dst=NodeID,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=createdAt,src=GetNodeOut.#.created_at,required
fullname=deletedAt,src=GetNodeOut.#.deleted_at,required
fullname=depth,src=GetNodeOut.#.depth,required
fullname=ext,src=GetNodeOut.#.ext,required
fullname=isLeaf,src=GetNodeOut.#.is_leaf,required
fullname=nodeId,src=GetNodeOut.#.node_id,required
fullname=order,src=GetNodeOut.#.order,required
fullname=parentId,src=GetNodeOut.#.parent_id,required
fullname=path,src=GetNodeOut.#.path,required
fullname=title,src=GetNodeOut.#.title,required
fullname=updatedAt,src=GetNodeOut.#.updated_at,required');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DoaTreeRelationGetSubTreeLimitDepth','SQL','doa tree relation 获取 sub limit depth','doa tree relation 获取 sub limit depth','doa','{{define "GetSubTreeLimitDepth"}} select * from `tree_relation` where `path` like '{{.ParentPath}}%' and `deleted_at` ='{{zeroTime}}' {{if .Depth }} and `depth`<{{.Depth}} {{end}} order by `depth` asc,`order` asc; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`) values('doa-tree_relation-GetSubTreeLimitDepth','doa tree relation 获取 sub limit depth','doa tree relation 获取 sub limit depth','POST','/api/doa/v1/tree_relation/get_sub_tree_limit_depth','DoaTreeRelationGetSubTreeLimitDepth','{{define "main"}}
{{execSQLTpl . "GetSubTreeLimitDepth"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=depth,dst=Depth,format=number,required
fullname=parentPath,dst=ParentPath,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=createdAt,src=GetSubTreeLimitDepthOut.#.created_at,required
fullname=deletedAt,src=GetSubTreeLimitDepthOut.#.deleted_at,required
fullname=depth,src=GetSubTreeLimitDepthOut.#.depth,required
fullname=ext,src=GetSubTreeLimitDepthOut.#.ext,required
fullname=isLeaf,src=GetSubTreeLimitDepthOut.#.is_leaf,required
fullname=nodeId,src=GetSubTreeLimitDepthOut.#.node_id,required
fullname=order,src=GetSubTreeLimitDepthOut.#.order,required
fullname=parentId,src=GetSubTreeLimitDepthOut.#.parent_id,required
fullname=path,src=GetSubTreeLimitDepthOut.#.path,required
fullname=title,src=GetSubTreeLimitDepthOut.#.title,required
fullname=updatedAt,src=GetSubTreeLimitDepthOut.#.updated_at,required');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DoaTreeRelationGetSubTreeNodeCount','SQL','doa tree relation 获取 sub node count','doa tree relation 获取 sub node count','doa','{{define "GetSubTreeNodeCount"}} set @path=(select path from `tree_relation` where `node_id`='{{.NodeID}}'); select count(*) from `tree_relation` where `path` like concat(@path,'%') and `deleted_at` ='{{zeroTime}}'; {{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DoaTreeRelationMoveSubTree','SQL','doa tree relation move sub','doa tree relation move sub','doa','{{define "MoveSubTree"}} start transaction; update `tree_relation` set `parent_id`='{{.NewParentID}}',`path`='{{.NewPath}}',`depth`= `depth` + {{.DiffDepth}} where `node_id`='{{.NodeID}}'; update `tree_relation` set `path`=replace(`path`,'{{.OldPath}}','{{.NewPath}}'),`depth`= `depth` + {{.DiffDepth}} where `path` like '{{.OldPath}}%'; commit; {{end}} {{/* 删除节点和子节点，删除关联节点的 id 集合变量为 @nodeIds */}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DoaTreeRelationDeleteSubTree','SQL','doa tree relation delete sub','doa tree relation delete sub','doa','{{define "DeleteSubTree"}} update `tree_relation` set `deleted_at`='{{currentTime}}' where `path` like '{{.NodePathPrefix}}%'; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`) values('doa-tree_relation-DeleteSubTree','doa tree relation delete sub','doa tree relation delete sub','POST','/api/doa/v1/tree_relation/delete_sub_tree','DoaTreeRelationDeleteSubTree','{{define "main"}}
{{execSQLTpl . "DeleteSubTree"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=nodePathPrefix,dst=NodePathPrefix,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out');

