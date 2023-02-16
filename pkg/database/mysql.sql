create table users
(
    id       int auto_increment comment '自增主键',
    name     varchar(255) not null comment '用户名',
    password varchar(255) not null comment '密码',
    salt     varchar(255) not null comment '加盐',
    constraint users_pk
        primary key (id)
)
    comment '用户表';

create index users_name_index
    on users (name);

create table follows
(
    id        int auto_increment comment '自增主键',
    user_id   int not null comment '用户',
    follow_id int not null comment '被关注者',
    cancel    int null comment '正常为1，取消为0',
    constraint follows_pk
        primary key (id)
)
    comment '用户关系表';

create index follows_id_user_id_follow_id_index
    on follows (id, user_id, follow_id);

create table likes
(
    id       int auto_increment comment '自增主键',
    user_id  int not null comment '点赞用户id',
    video_id int not null comment '点赞视频id',
    cancel   int not null comment '正常为1，取消为0',
    constraint likes_pk
        primary key (id)
)
    comment '喜爱表';

create index likes_user_id_video_id_index
    on likes (user_id, video_id);

create table videos
(
    id           int auto_increment comment '自增主键',
    author_id    int          not null comment '作者id',
    play_url     varchar(255) not null comment '播放url',
    cover_url    varchar(255) not null comment '封面url',
    publish_time datetime     not null comment '发布时间戳',
    title        varchar(255) null comment '视频标题',
    constraint videos_pk
        primary key (id)
)
    comment '视频表';

create index videos_id_author_id_index
    on videos (id, author_id);