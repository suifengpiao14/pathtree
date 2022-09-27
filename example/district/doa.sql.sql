insert ignore into `source` (`source_id`,`source_type`,`config`) values('rent','SQL','{"logLevel":"debug","dsn":"root:123456@tcp(mysql_address:3306)/rent?charset=utf8&timeout=1s&readTimeout=5s&writeTimeout=5s&parseTime=False&loc=Local&multiStatements=true","timeout":30}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('RentDistrictGetByCode','SQL','rent 地区表(主要数据来源于国家统计局并作清洗) 获取 通过 code','rent 地区表(主要数据来源于国家统计局并作清洗) 获取 通过 code','rent','{{define "GetByCode"}} select * from `district` where `code`=:Code and deleted_at is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('rent-district-GetByCode','rent 地区表(主要数据来源于国家统计局并作清洗) 获取 通过 code','rent 地区表(主要数据来源于国家统计局并作清洗) 获取 通过 code','POST','/api/rent/v1/district/get_by_code','RentDistrictGetByCode','{{define "main"}}
{{execSQLTpl . "GetByCode"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=code,dst=Code,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=code,src=GetByCodeOut.#.code,required
fullname=createdAt,src=GetByCodeOut.#.created_at,required
fullname=deletedAt,src=GetByCodeOut.#.deleted_at,required
fullname=depth,src=GetByCodeOut.#.depth,required
fullname=firstLetter,src=GetByCodeOut.#.first_letter,required
fullname=id,src=GetByCodeOut.#.id,required
fullname=isDeprecated,src=GetByCodeOut.#.is_deprecated,required
fullname=label,src=GetByCodeOut.#.label,required
fullname=parentCode,src=GetByCodeOut.#.parent_code,required
fullname=path,src=GetByCodeOut.#.path,required
fullname=title,src=GetByCodeOut.#.title,required
fullname=updatedAt,src=GetByCodeOut.#.updated_at,required','{{define "input"}}
{{getSetValue . "Code" "input.code"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.code" "GetByCodeOut.#.code"}}
{{getSetValue . "output.createdAt" "GetByCodeOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "GetByCodeOut.#.deleted_at"}}
{{getSetValue . "output.depth" "GetByCodeOut.#.depth"}}
{{getSetValue . "output.firstLetter" "GetByCodeOut.#.first_letter"}}
{{getSetValue . "output.id" "GetByCodeOut.#.id"}}
{{getSetValue . "output.isDeprecated" "GetByCodeOut.#.is_deprecated"}}
{{getSetValue . "output.label" "GetByCodeOut.#.label"}}
{{getSetValue . "output.parentCode" "GetByCodeOut.#.parent_code"}}
{{getSetValue . "output.path" "GetByCodeOut.#.path"}}
{{getSetValue . "output.title" "GetByCodeOut.#.title"}}
{{getSetValue . "output.updatedAt" "GetByCodeOut.#.updated_at"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('RentDistrictGetAllByCode','SQL','rent 地区表(主要数据来源于国家统计局并作清洗) 获取 所有 通过 code','rent 地区表(主要数据来源于国家统计局并作清洗) 获取 所有 通过 code','rent','{{define "GetAllByCode"}} select * from `district` where `code` in ({{in . .CodeList}}) and deleted_at is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('rent-district-GetAllByCode','rent 地区表(主要数据来源于国家统计局并作清洗) 获取 所有 通过 code','rent 地区表(主要数据来源于国家统计局并作清洗) 获取 所有 通过 code','POST','/api/rent/v1/district/get_all_by_code','RentDistrictGetAllByCode','{{define "main"}}
{{execSQLTpl . "GetAllByCode"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=codeList,dst=CodeList,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=code,src=GetAllByCodeOut.#.code,required
fullname=createdAt,src=GetAllByCodeOut.#.created_at,required
fullname=deletedAt,src=GetAllByCodeOut.#.deleted_at,required
fullname=depth,src=GetAllByCodeOut.#.depth,required
fullname=firstLetter,src=GetAllByCodeOut.#.first_letter,required
fullname=id,src=GetAllByCodeOut.#.id,required
fullname=isDeprecated,src=GetAllByCodeOut.#.is_deprecated,required
fullname=label,src=GetAllByCodeOut.#.label,required
fullname=parentCode,src=GetAllByCodeOut.#.parent_code,required
fullname=path,src=GetAllByCodeOut.#.path,required
fullname=title,src=GetAllByCodeOut.#.title,required
fullname=updatedAt,src=GetAllByCodeOut.#.updated_at,required','{{define "input"}}
{{getSetValue . "CodeList" "input.codeList"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.code" "GetAllByCodeOut.#.code"}}
{{getSetValue . "output.createdAt" "GetAllByCodeOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "GetAllByCodeOut.#.deleted_at"}}
{{getSetValue . "output.depth" "GetAllByCodeOut.#.depth"}}
{{getSetValue . "output.firstLetter" "GetAllByCodeOut.#.first_letter"}}
{{getSetValue . "output.id" "GetAllByCodeOut.#.id"}}
{{getSetValue . "output.isDeprecated" "GetAllByCodeOut.#.is_deprecated"}}
{{getSetValue . "output.label" "GetAllByCodeOut.#.label"}}
{{getSetValue . "output.parentCode" "GetAllByCodeOut.#.parent_code"}}
{{getSetValue . "output.path" "GetAllByCodeOut.#.path"}}
{{getSetValue . "output.title" "GetAllByCodeOut.#.title"}}
{{getSetValue . "output.updatedAt" "GetAllByCodeOut.#.updated_at"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('RentDistrictGetByPathPrefixLimitDepth','SQL','rent 地区表(主要数据来源于国家统计局并作清洗) 获取 通过 path prefix limit depth','rent 地区表(主要数据来源于国家统计局并作清洗) 获取 通过 path prefix limit depth','rent','{{define "GetByPathPrefixLimitDepth"}} select * from `district` where `path` like :PathPrefix and `depth`<=:Depth and deleted_at is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('rent-district-GetByPathPrefixLimitDepth','rent 地区表(主要数据来源于国家统计局并作清洗) 获取 通过 path prefix limit depth','rent 地区表(主要数据来源于国家统计局并作清洗) 获取 通过 path prefix limit depth','POST','/api/rent/v1/district/get_by_path_prefix_limit_depth','RentDistrictGetByPathPrefixLimitDepth','{{define "main"}}
{{execSQLTpl . "GetByPathPrefixLimitDepth"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=depth,dst=Depth,format=number,required
fullname=pathPrefix,dst=PathPrefix,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=code,src=GetByPathPrefixLimitDepthOut.#.code,required
fullname=createdAt,src=GetByPathPrefixLimitDepthOut.#.created_at,required
fullname=deletedAt,src=GetByPathPrefixLimitDepthOut.#.deleted_at,required
fullname=depth,src=GetByPathPrefixLimitDepthOut.#.depth,required
fullname=firstLetter,src=GetByPathPrefixLimitDepthOut.#.first_letter,required
fullname=id,src=GetByPathPrefixLimitDepthOut.#.id,required
fullname=isDeprecated,src=GetByPathPrefixLimitDepthOut.#.is_deprecated,required
fullname=label,src=GetByPathPrefixLimitDepthOut.#.label,required
fullname=parentCode,src=GetByPathPrefixLimitDepthOut.#.parent_code,required
fullname=path,src=GetByPathPrefixLimitDepthOut.#.path,required
fullname=title,src=GetByPathPrefixLimitDepthOut.#.title,required
fullname=updatedAt,src=GetByPathPrefixLimitDepthOut.#.updated_at,required','{{define "input"}}
{{getSetNumber . "Depth" "input.depth"}}
{{getSetValue . "PathPrefix" "input.pathPrefix"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.code" "GetByPathPrefixLimitDepthOut.#.code"}}
{{getSetValue . "output.createdAt" "GetByPathPrefixLimitDepthOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "GetByPathPrefixLimitDepthOut.#.deleted_at"}}
{{getSetValue . "output.depth" "GetByPathPrefixLimitDepthOut.#.depth"}}
{{getSetValue . "output.firstLetter" "GetByPathPrefixLimitDepthOut.#.first_letter"}}
{{getSetValue . "output.id" "GetByPathPrefixLimitDepthOut.#.id"}}
{{getSetValue . "output.isDeprecated" "GetByPathPrefixLimitDepthOut.#.is_deprecated"}}
{{getSetValue . "output.label" "GetByPathPrefixLimitDepthOut.#.label"}}
{{getSetValue . "output.parentCode" "GetByPathPrefixLimitDepthOut.#.parent_code"}}
{{getSetValue . "output.path" "GetByPathPrefixLimitDepthOut.#.path"}}
{{getSetValue . "output.title" "GetByPathPrefixLimitDepthOut.#.title"}}
{{getSetValue . "output.updatedAt" "GetByPathPrefixLimitDepthOut.#.updated_at"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('RentDistrictCountByPathPrefix','SQL','rent 地区表(主要数据来源于国家统计局并作清洗) count 通过 path prefix','rent 地区表(主要数据来源于国家统计局并作清洗) count 通过 path prefix','rent','{{define "CountByPathPrefix"}} select count(*) from `district` where `path` like :PathPrefix and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('rent-district-CountByPathPrefix','rent 地区表(主要数据来源于国家统计局并作清洗) count 通过 path prefix','rent 地区表(主要数据来源于国家统计局并作清洗) count 通过 path prefix','POST','/api/rent/v1/district/count_by_path_prefix','RentDistrictCountByPathPrefix','{{define "main"}}
{{execSQLTpl . "CountByPathPrefix"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=pathPrefix,dst=PathPrefix,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "PathPrefix" "input.pathPrefix"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('RentDistrictDeleteByCode','SQL','rent 地区表(主要数据来源于国家统计局并作清洗) delete 通过 code','rent 地区表(主要数据来源于国家统计局并作清洗) delete 通过 code','rent','{{define "DeleteByCode"}} update `district` set `deleted_at`={{currentTime .}} where `code`=:Code; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('rent-district-DeleteByCode','rent 地区表(主要数据来源于国家统计局并作清洗) delete 通过 code','rent 地区表(主要数据来源于国家统计局并作清洗) delete 通过 code','POST','/api/rent/v1/district/delete_by_code','RentDistrictDeleteByCode','{{define "main"}}
{{execSQLTpl . "DeleteByCode"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=code,dst=Code,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "Code" "input.code"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('RentDistrictPaginateWhere','SQL','rent 地区表(主要数据来源于国家统计局并作清洗) 分页列表 where','rent 地区表(主要数据来源于国家统计局并作清洗) 分页列表 where','rent','{{define "PaginateWhere"}} {{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('RentDistrictPaginateTotal','SQL','rent 地区表(主要数据来源于国家统计局并作清洗) 分页列表 总数','rent 地区表(主要数据来源于国家统计局并作清洗) 分页列表 总数','rent','{{define "PaginateTotal"}} select count(*) as `count` from `district` where 1=1 {{template "PaginateWhere" .}} and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('rent-district-PaginateTotal','rent 地区表(主要数据来源于国家统计局并作清洗) 分页列表 总数','rent 地区表(主要数据来源于国家统计局并作清洗) 分页列表 总数','POST','/api/rent/v1/district/paginate_total','RentDistrictPaginateTotal','{{define "main"}}
{{execSQLTpl . "PaginateTotal"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('RentDistrictPaginate','SQL','rent 地区表(主要数据来源于国家统计局并作清洗) 分页列表','rent 地区表(主要数据来源于国家统计局并作清洗) 分页列表','rent','{{define "Paginate"}} select * from `district` where 1=1 {{template "PaginateWhere" .}} and `deleted_at` is null order by `updated_at` desc limit :Offset,:Limit ; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('rent-district-Paginate','rent 地区表(主要数据来源于国家统计局并作清洗) 分页列表','rent 地区表(主要数据来源于国家统计局并作清洗) 分页列表','POST','/api/rent/v1/district/paginate','RentDistrictPaginateWhere,RentDistrictPaginateTotal,RentDistrictPaginate','{{define "main"}}
{{execSQLTpl . "PaginateTotal"}}
{{execSQLTpl . "Paginate"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=limit,dst=Limit,format=number,required
fullname=offset,dst=Offset,format=number,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=code,src=PaginateOut.#.code,required
fullname=createdAt,src=PaginateOut.#.created_at,required
fullname=deletedAt,src=PaginateOut.#.deleted_at,required
fullname=depth,src=PaginateOut.#.depth,required
fullname=firstLetter,src=PaginateOut.#.first_letter,required
fullname=id,src=PaginateOut.#.id,required
fullname=isDeprecated,src=PaginateOut.#.is_deprecated,required
fullname=label,src=PaginateOut.#.label,required
fullname=parentCode,src=PaginateOut.#.parent_code,required
fullname=path,src=PaginateOut.#.path,required
fullname=title,src=PaginateOut.#.title,required
fullname=updatedAt,src=PaginateOut.#.updated_at,required','{{define "input"}}
{{getSetNumber . "Limit" "input.limit"}}
{{getSetNumber . "Offset" "input.offset"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.code" "PaginateOut.#.code"}}
{{getSetValue . "output.createdAt" "PaginateOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "PaginateOut.#.deleted_at"}}
{{getSetValue . "output.depth" "PaginateOut.#.depth"}}
{{getSetValue . "output.firstLetter" "PaginateOut.#.first_letter"}}
{{getSetValue . "output.id" "PaginateOut.#.id"}}
{{getSetValue . "output.isDeprecated" "PaginateOut.#.is_deprecated"}}
{{getSetValue . "output.label" "PaginateOut.#.label"}}
{{getSetValue . "output.parentCode" "PaginateOut.#.parent_code"}}
{{getSetValue . "output.path" "PaginateOut.#.path"}}
{{getSetValue . "output.title" "PaginateOut.#.title"}}
{{getSetValue . "output.updatedAt" "PaginateOut.#.updated_at"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('RentDistrictInsert','SQL','rent 地区表(主要数据来源于国家统计局并作清洗) 新增','rent 地区表(主要数据来源于国家统计局并作清洗) 新增','rent','{{define "Insert"}} insert into `district` (`code`,`title`,`label`,`parent_code`,`path`,`depth`,`first_letter`,`is_deprecated`)values (:Code,:Title,:Label,:ParentCode,:Path,:Depth,:FirstLetter,:IsDeprecated); {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('rent-district-Insert','rent 地区表(主要数据来源于国家统计局并作清洗) 新增','rent 地区表(主要数据来源于国家统计局并作清洗) 新增','POST','/api/rent/v1/district/insert','RentDistrictInsert','{{define "main"}}
{{execSQLTpl . "Insert"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=code,dst=Code,required
fullname=depth,dst=Depth,format=number,required
fullname=firstLetter,dst=FirstLetter,required
fullname=isDeprecated,dst=IsDeprecated,required
fullname=label,dst=Label,required
fullname=parentCode,dst=ParentCode,required
fullname=path,dst=Path,required
fullname=title,dst=Title,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "Code" "input.code"}}
{{getSetNumber . "Depth" "input.depth"}}
{{getSetValue . "FirstLetter" "input.firstLetter"}}
{{getSetValue . "IsDeprecated" "input.isDeprecated"}}
{{getSetValue . "Label" "input.label"}}
{{getSetValue . "ParentCode" "input.parentCode"}}
{{getSetValue . "Path" "input.path"}}
{{getSetValue . "Title" "input.title"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('RentDistrictUpdate','SQL','rent 地区表(主要数据来源于国家统计局并作清洗) 修改','rent 地区表(主要数据来源于国家统计局并作清洗) 修改','rent','{{define "Update"}} {{$preComma:=newPreComma}} update `district` set {{if .Code}} {{$preComma.PreComma}} `code`=:Code {{end}} {{if .Title}} {{$preComma.PreComma}} `title`=:Title {{end}} {{if .Label}} {{$preComma.PreComma}} `label`=:Label {{end}} {{if .ParentCode}} {{$preComma.PreComma}} `parent_code`=:ParentCode {{end}} {{if .Path}} {{$preComma.PreComma}} `path`=:Path {{end}} {{if .Depth}} {{$preComma.PreComma}} `depth`=:Depth {{end}} {{if .FirstLetter}} {{$preComma.PreComma}} `first_letter`=:FirstLetter {{end}} {{if .IsDeprecated}} {{$preComma.PreComma}} `is_deprecated`=:IsDeprecated {{end}} where `id`=:ID; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('rent-district-Update','rent 地区表(主要数据来源于国家统计局并作清洗) 修改','rent 地区表(主要数据来源于国家统计局并作清洗) 修改','POST','/api/rent/v1/district/update','RentDistrictUpdate','{{define "main"}}
{{execSQLTpl . "Update"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=code,dst=Code,required
fullname=depth,dst=Depth,format=number,required
fullname=firstLetter,dst=FirstLetter,required
fullname=id,dst=ID,format=number,required
fullname=isDeprecated,dst=IsDeprecated,required
fullname=label,dst=Label,required
fullname=parentCode,dst=ParentCode,required
fullname=path,dst=Path,required
fullname=title,dst=Title,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "Code" "input.code"}}
{{getSetNumber . "Depth" "input.depth"}}
{{getSetValue . "FirstLetter" "input.firstLetter"}}
{{getSetNumber . "ID" "input.id"}}
{{getSetValue . "IsDeprecated" "input.isDeprecated"}}
{{getSetValue . "Label" "input.label"}}
{{getSetValue . "ParentCode" "input.parentCode"}}
{{getSetValue . "Path" "input.path"}}
{{getSetValue . "Title" "input.title"}}
{{end}}','{{define "output"}}
{{end}}');