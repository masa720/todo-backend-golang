create database if not exists todo;
use todo;

create table if not exists todos
(
    id          bigint unsigned   not null primary key auto_increment comment 'TODO id',
    user_id     bigint unsigned      comment 'ユーザーid',
    title       varchar(100)      not null comment 'タイトル',
    description varchar(500)      comment '説明',
    deadline    datetime comment '期限',
    is_done     boolean           not null default false comment '完了',
    created_at  datetime          not null default current_timestamp comment '作成日',
    updated_at  datetime          not null default current_timestamp on update current_timestamp comment '更新日'
)