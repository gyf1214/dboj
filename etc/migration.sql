drop table if exists `evaluation`;
drop table if exists `dataset`;
drop table if exists `submition`;
drop table if exists `judge`;
drop table if exists `post`;
drop table if exists `discussion`;
drop table if exists `problem`;
drop table if exists `user`;
drop table if exists `group`;

create table `group` (
    `id`          integer not null auto_increment,
    `name`        varchar(32) not null unique,
    `description` text not null,
    primary key (`id`)
);

create table `user` (
    `id`       int not null auto_increment,
    `name`     varchar(32) not null unique,
    `passwd`   char(32) not null,
    `group`    int,
    `session`  char(32) unique,
    `activity` datetime not null,
    primary key (`id`),
    foreign key (`group`) references `group`(`id`) on delete set null
);

create table `problem` (
    `id`          int not null auto_increment,
    `owner`       int not null,
    `title`       varchar(32) not null,
    `description` text not null,
    primary key (`id`),
    foreign key (`owner`) references `user`(`id`)
);

create table `discussion` (
    `id`      int not null auto_increment,
    `user`    int not null,
    `problem` int not null,
    `title`   varchar(32) not null,
    `content` text not null,
    primary key (`id`),
    foreign key (`user`)    references `user`(`id`),
    foreign key (`problem`) references `problem`(`id`) on delete cascade
);

create table `post` (
    `id`         int not null auto_increment,
    `discussion` int not null,
    `user`       int not null,
    `content`    text not null,
    primary key (`id`),
    foreign key (`discussion`) references `discussion`(`id`) on delete cascade,
    foreign key (`user`)       references `user`(`id`)
);

create table `judge` (
    `id`       int not null auto_increment,
    `name`     varchar(32) not null,
    `language` varchar(32) not null,
    `address`  varchar(32) not null,
    primary key (`id`),
    key         (`language`)
);

create table `submition` (
    `id`       int not null auto_increment,
    `problem`  int not null,
    `user`     int not null,
    `language` varchar(32) not null,
    `code`     text not null,
    primary key (`id`),
    foreign key (`problem`)  references `problem`(`id`) on delete cascade,
    foreign key (`user`)     references `user`(`id`)
);

create table `dataset` (
    `id`      int not null auto_increment,
    `problem` int not null,
    `score`   int not null,
    `input`   text not null,
    `answer`  text not null,
    primary key (`id`),
    foreign key (`problem`) references `problem`(`id`) on delete cascade
);

create table `evaluation` (
    `id`        int not null auto_increment,
    `submition` int not null,
    `dataset`   int,
    `judge`     int,
    `status`    tinyint not null,
    `message`   text not null,
    primary key (`id`),
    foreign key (`submition`) references `submition`(`id`) on delete cascade,
    foreign key (`dataset`)   references `dataset`(`id`) on delete set null,
    foreign key (`judge`)     references `judge`(`id`) on delete set null,
    unique  key (`submition`, `dataset`)
);

create trigger `submit_evaluation` after insert on `submition` for each row
    insert into `evaluation` (`submition`, `dataset`, `status`)
    select NEW.`id`, `dataset`.`id`, 0 from `dataset`
    where `dataset`.`problem` = NEW.`problem`;
