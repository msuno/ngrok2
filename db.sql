CREATE TABLE `ngrok_user`(  
    id int NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    union_id varchar(64) NOT NULL DEFAULT '' COMMENT '微信unionId',
    domain varchar(64) NOT NULL DEFAULT '' COMMENT '代理子域名',
    sk varchar(64) NOT NULL DEFAULT '' COMMENT '代理SK',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间'
) DEFAULT CHARSET UTF8 COMMENT 'Ngrok代理用户';

CREATE TABLE `users` (
  `id` int(11) AUTO_INCREMENT NOT NULL COMMENT '主键',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `avatar` varchar(1023) NOT NULL DEFAULT '',
  `description` varchar(1023) NOT NULL DEFAULT '',
  `email` varchar(127) NOT NULL DEFAULT '',
  `expire_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `mfa_key` varchar(64) NOT NULL DEFAULT '',
  `mfa_type` int(11) NOT NULL DEFAULT '0',
  `nickname` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `username` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
);

CREATE TABLE `system_info` (
  `id` int(11) AUTO_INCREMENT NOT NULL COMMENT '主键',
  `cpu` decimal(18,2) NOT NULL DEFAULT '0.00',
  `mem` decimal(18,2) NOT NULL DEFAULT '0.00',
  `disk` decimal(18,2) NOT NULL DEFAULT '0.00',
  `net_i` decimal(18,2) NOT NULL DEFAULT '0.00',
  `net_o` decimal(18,2) NOT NULL DEFAULT '0.00',
  `load` decimal(18,2) NOT NULL DEFAULT '0.00',
  `pid` decimal(18,2) NOT NULL DEFAULT '0.00',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);