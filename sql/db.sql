create table `user`(
    `id` bigint(20) unsigned not null auto_increment,
    `username` varchar(255) not null,
    `password` varchar(255) not null,
    `icon` varchar(1024),
    `gender` tinyint(1),
    `email` varchar(255),
    primary key(`id`)
);

create table `blog`(
    `id` bigint(20) unsigned not null auto_increment,
    `user_id` bigint(20) unsigned not null,
    `title` varchar(255) not null,
    `text` text,
    `like` int not null default 0,
    `comment` text,
    primary key(`id`)
);