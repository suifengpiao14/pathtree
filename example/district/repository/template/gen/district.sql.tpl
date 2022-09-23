{{define "PaginateWhere"}}
{{end}}

{{define "PaginateTotal"}}
select count(*) as `count` from `district` where 1=1 {{template "PaginateWhere" .}} and `deleted_at` is null;
{{end}}

{{define "Paginate"}}
select * from `district` where 1=1 {{template "PaginateWhere" .}} and `deleted_at` is null order by `updated_at` desc limit :Offset,:Limit ;
{{end}}

{{define "Insert"}}
insert into `district` (`code`,`title`,`label`,`parent_code`,`path`,`depth`,`first_letter`,`is_deprecated`)values
(:Code,:Title,:Label,:ParentCode,:Path,:Depth,:FirstLetter,:IsDeprecated);
{{end}}

{{define "Update"}}
{{$preComma:=newPreComma}}
update `district` set {{if .Code}} {{$preComma.PreComma}} `code`=:Code {{end}}
{{if .Title}} {{$preComma.PreComma}} `title`=:Title {{end}}
{{if .Label}} {{$preComma.PreComma}} `label`=:Label {{end}}
{{if .ParentCode}} {{$preComma.PreComma}} `parent_code`=:ParentCode {{end}}
{{if .Path}} {{$preComma.PreComma}} `path`=:Path {{end}}
{{if .Depth}} {{$preComma.PreComma}} `depth`=:Depth {{end}}
{{if .FirstLetter}} {{$preComma.PreComma}} `first_letter`=:FirstLetter {{end}}
{{if .IsDeprecated}} {{$preComma.PreComma}} `is_deprecated`=:IsDeprecated {{end}} where `id`=:ID;
{{end}}
