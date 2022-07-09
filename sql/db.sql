create table `user`(
    `user_id` bigint(20) unsigned not null auto_increment,
    `username` varchar(255) not null,
    `password` varchar(255) not null,
    `user_avatar` varchar(255) default 'default_user_avatar',
    `gender` tinyint(1),
    `email` varchar(255),
    primary key(`user_id`)
);

create table `blog`(
    `blog_id` bigint(20) unsigned not null auto_increment,
    `user_id` bigint(20) unsigned not null,
    `text` text,
    `imgs` varchar(255),
    `create_time_stamp` bigint(20) unsigned not null,
    `like` int not null default 0,
    primary key(`blog_id`)
);

create table `comment`(
    `comment_id` bigint(20) unsigned not null auto_increment,
    `blog_id` bigint(20) unsigned not null,
    `user_id` bigint(20) unsigned not null,
    `text` text,
    `create_time_stamp` bigint(20) unsigned not null,
    primary key(`comment_id`)
);

create table `like`(
    `blog_id` bigint(20) unsigned not null,
    `user_id` bigint(20) unsigned not null,
);