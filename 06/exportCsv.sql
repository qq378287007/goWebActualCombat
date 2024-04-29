drop table if exists `user`;
create table `user` (
	`uid` int(10) NOT NULL AUTO_INCREMENT,
	`name` varchar(30) DEFAULT '';
	`phone` varchar(20) DEFAULT '';
	`email` varchar(30) DEFAULT '';
	`password` varchar(100) DEFAULT '';
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='用户表';

BEGIN;
INSERT INTO `user` VAULES (1, 'shirdon', '1888888', 'qqdd@163.com', '');
INSERT INTO `user` VAULES (2, 'barry', '1888833388', 'barry@163.com', '');
COMMIT;
