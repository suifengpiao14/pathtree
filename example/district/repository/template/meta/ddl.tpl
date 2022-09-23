CREATE TABLE `district` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `code` char(10)  NOT NULL  COMMENT '城市编码',
  `title` char(20) NOT NULL COMMENT '城市名称',
  `label` enum("country","province","city","area","street","county","town","village") NOT NULL COMMENT '分级标签(country-国家,province-省,city-市,area-区,street-街道,county-县,town-镇,village-村)',
  `parent_code` char(10)  NOT NULL default "" COMMENT '上级城市编码',
  `path` char(255)  NOT NULL default "/" COMMENT '城市层级路径',
  `depth` int(11) unsigned NOT NULL default 0 COMMENT '城市层级',
  `first_letter` char(1) NOT NULL DEFAULT "" COMMENT '名称首字母',
  `is_deprecated` enum("0","1")  NOT NULL DEFAULT '0' COMMENT '是否废弃(0-弃用,1-废弃)',
  `created_at` timestamp  NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT  '更新时间',
  `deleted_at` timestamp NOT NULL DEFAULT "0000-00-00 00:00:00"  COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `path` (`path`),
  KEY `code` (`code`,`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='地区表(主要数据来源于国家统计局并作清洗)';