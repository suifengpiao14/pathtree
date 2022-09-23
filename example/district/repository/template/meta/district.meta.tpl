{{define "GetByCode"}}
    select * from `district` where `code`=:Code and deleted_at="0000-00-00 00:00:00";
{{end}}

{{define "GetAllByCode"}}
    select * from `district` where `code` in ({{in . .CodeList}})   and deleted_at="0000-00-00 00:00:00";
{{end}}

{{define "GetByPathPrefixLimitDepth"}}
    select * from `district` where `path` like .PathPrefix and `depth`<=:Depth  and deleted_at="0000-00-00 00:00:00";
{{end}}

{{define "CountByPathPrefix"}}
    select count(*) from `district` where `path` like :PathPrefix and `deleted_at`="0000-00-00 00:00:00";
{{end}}

{{define "DeleteByCode"}}
    update `district` set `deleted_at`={{currentTime .}} where `code`=:Code;
{{end}}