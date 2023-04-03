{{define "GetByAreaID"}}
    select * from `t_city_info_tmp` where `Farea_id`=:AreaID and `Fcity_status`=1;
{{end}}

{{define "GetAllByAreaIDList"}}
    select * from `t_city_info_tmp` where `Farea_id` in ({{in . .AreaIDList}})   and `Fcity_status`=1;
{{end}}

{{define "GetByCityPathPrefix"}}
    select * from `t_city_info_tmp` where `Fcity_path` like :PathPrefix and `Fcity_level`<=:CityLevel  and `Fcity_status`=1;
{{end}}

{{define "GetAllByPathPrefix"}}
    select * from `t_city_info_tmp` where 1=1 {{if .PathPrefix}} `Fcity_path`  like :PathPrefix {{end}}  and `Fcity_status`=1;
{{end}}

{{define "GetByCityLevel"}}
    select * from `t_city_info_tmp` where `Fcity_level`=:CityLevel and  `Fcity_status`=1;
{{end}}

{{define "ListByKeyword"}}
    select * from `t_city_info_tmp` where `Farea_name` like :AreaName and  `Fcity_status`=1;
{{end}}
