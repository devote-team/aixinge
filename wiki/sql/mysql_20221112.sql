create table if not exists axg_application
(
    id         bigint unsigned auto_increment comment '主键 ID' primary key,
    created_at datetime         null comment '创建时间',
    updated_at datetime         null comment '更新时间',
    deleted_at datetime         null comment '删除时间',
    name       varchar(30)      null comment '应用名称',
    app_key    varchar(64)      null comment '应用 ID',
    app_secret varchar(64)      null comment '应用密钥',
    remark     varchar(255)     null comment '应用备注',
    status     bigint default 1 not null comment '状态(0-删除、1-正常)'
) comment '应用信息表';

create table if not exists axg_menu
(
    id         bigint unsigned           not null comment '主键 ID' primary key,
    created_at datetime                  not null comment '创建时间',
    updated_at datetime                  null comment '更新时间',
    deleted_at datetime                  null comment '删除时间',
    parent_id  bigint     default 1      not null comment '父级 ID',
    path       varchar(255) charset utf8 null comment '菜单路径',
    name       varchar(100) charset utf8 null comment '菜单名称',
    hidden     tinyint(1) default 2      not null comment '是否隐藏',
    component  varchar(255) charset utf8 null comment '对应组件',
    sort       bigint                    null comment '排序号',
    is_frame   tinyint(1) default 2      not null comment '是否为 iframe',
    status     bigint     default 1      not null comment '状态(0-删除、1-正常)',
    no_cache   tinyint(1) default 1      not null comment '是否禁用缓存',
    title      varchar(100) charset utf8 null comment '标题',
    icon       varchar(50) charset utf8  null comment '图标',
    remark     varchar(255)              null comment '备注'
) comment '菜单表';

create table if not exists axg_msg_channel
(
    id         bigint unsigned auto_increment comment '主键 ID' primary key,
    created_at datetime         null comment '创建时间',
    updated_at datetime         null comment '更新时间',
    deleted_at datetime         null comment '删除时间',
    app_id     bigint           null comment '应用 ID',
    name       varchar(30)      null comment '消息渠道名称',
    type       int              null comment '消息类型(枚举)',
    provider   int              null comment '消息服务提供商',
    config     json             null comment '消息渠道配置（JSON）',
    remark     varchar(255)     null comment '备注',
    status     bigint default 1 not null comment '状态(0-删除、1-正常)'
) comment '消息渠道配置表';

create table if not exists axg_msg_log
(
    id          bigint unsigned auto_increment comment '主键 ID' primary key,
    created_at  datetime         null comment '创建时间',
    updated_at  datetime         null comment '更新时间',
    deleted_at  datetime         null comment '删除时间',
    template_id bigint           null comment '模板 ID',
    parameter   json             null comment '消息参数（JSON）',
    receiver    json             null comment '接收人（JSON）',
    content     longtext         null comment '消息内容',
    extend      json             null comment '扩展字段（JSON）',
    remark      varchar(255)     null comment '备注',
    status      bigint default 1 not null comment '状态(0-删除、1-正常)'
) comment '消息发送记录表';

create table if not exists axg_msg_signature
(
    id         bigint unsigned auto_increment comment '主键 ID' primary key,
    created_at datetime         null comment '创建时间',
    updated_at datetime         null comment '更新时间',
    deleted_at datetime         null comment '删除时间',
    name       varchar(30)      null comment '签名名称',
    remark     varchar(255)     null comment '备注',
    status     bigint default 1 not null comment '状态(0-删除、1-正常)'
) comment '消息主体签名表';

create table if not exists axg_msg_template
(
    id         bigint unsigned auto_increment comment '主键 ID' primary key,
    created_at datetime         null comment '创建时间',
    updated_at datetime         null comment '更新时间',
    deleted_at datetime         null comment '删除时间',
    app_id     bigint           null comment '应用 ID',
    name       varchar(30)      null comment '模板名称',
    type       int              null comment '消息类型',
    content    longtext         null comment '模板内容',
    extend     json             null comment '扩展字段（JSON）',
    remark     varchar(255)     null comment '备注',
    status     bigint default 1 not null comment '状态(0-删除、1-正常)'
) comment '消息模板配置表';

create table if not exists axg_role
(
    id         bigint unsigned auto_increment comment '主键 ID' primary key,
    created_at datetime         null comment '创建时间',
    updated_at datetime         null comment '更新时间',
    deleted_at datetime         null comment '删除时间',
    name       varchar(30)      null comment '角色名称',
    alias      varchar(30)      null comment '角色别名',
    remark     varchar(255)     null comment '备注',
    status     bigint default 1 not null comment '状态(0-删除、1-正常)',
    sort       int              null comment '排序号'
) comment '角色表';

create table if not exists axg_role_menu
(
    role_id bigint not null comment '角色 ID',
    menu_id bigint not null comment '菜单 ID'
) comment '角色菜单关联表';

create table if not exists axg_user
(
    id         bigint unsigned                 not null comment '主键 ID' primary key,
    created_at datetime                        null comment '创建时间',
    updated_at datetime                        null comment '更新时间',
    deleted_at datetime                        null comment '删除时间',
    uuid       varchar(255)                    null comment '唯一标识',
    username   varchar(255)                    null comment '用户名',
    password   varchar(255)                    null comment '密码',
    nick_name  varchar(255) default '系统用户' null comment '昵称',
    avatar     varchar(255)                    null comment '头像',
    status     bigint       default 1          not null comment '状态(0-删除、1-正常)'
) comment '用户表' charset = utf8;

create index idx_users_deleted_at on axg_user (deleted_at);

create table if not exists axg_user_role
(
    user_id bigint unsigned not null comment '用户 ID',
    role_id bigint unsigned not null comment '角色 ID'
) comment '用户角色关联表';

insert into aixinge.axg_user (id, created_at, updated_at, deleted_at, uuid, username, password, nick_name, avatar, status)
values (1, '2022-02-15 23:25:20', '2022-11-06 20:26:03', null, 'f5f10f02-d4fe-472a-b6ed-c84698a3c007', 'admin', '1751ad04bed81acd2d1bc5720b15f0f5',
        '超级管理员', '1589232642040004608', 1);
