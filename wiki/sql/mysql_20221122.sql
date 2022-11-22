create table if not exists axg_application
(
    id         bigint unsigned not null comment '主键 ID' primary key,
    created_at datetime        null comment '创建时间',
    updated_at datetime        null comment '更新时间',
    deleted_at datetime        null comment '删除时间',
    name       varchar(30)     null comment '应用名称',
    app_key    varchar(64)     null comment '应用 KEY',
    app_secret varchar(64)     null comment '应用密钥',
    remark     varchar(255)    null comment '应用备注',
    status     int default 1   not null comment '状态(0-禁用、1-正常)'
) comment '应用信息表';

create table if not exists axg_attachment
(
    id           bigint unsigned not null comment '文扩展名' primary key,
    created_at   datetime        not null comment '创建时间',
    updated_at   datetime        null comment '更新时间',
    deleted_at   datetime        null comment '删除时间',
    name         varchar(255)    null comment '文件名称',
    md5          varchar(255)    not null comment 'MD5',
    path         varchar(255)    not null comment '存储路径',
    ext          varchar(255)    null comment '扩展名',
    content_type varchar(255)    null comment 'ContentType',
    etag         varchar(255)    null comment 'HTTP ETag',
    data         varchar(255)    null comment '扩展数据'
) comment '上传附件表';

create table if not exists axg_channel
(
    id         bigint unsigned not null comment '主键 ID' primary key,
    created_at datetime        null comment '创建时间',
    updated_at datetime        null comment '更新时间',
    deleted_at datetime        null comment '删除时间',
    name       varchar(30)     null comment '消息渠道名称',
    type       int             null comment '消息类型(枚举)',
    provider   int             null comment '消息服务提供商',
    weight     int default 0   not null comment '权重',
    config     json            null comment '消息渠道配置（JSON）',
    remark     varchar(255)    null comment '备注',
    status     int default 1   not null comment '状态(0-禁用、1-正常)'
) comment '消息渠道配置表';

create table if not exists axg_channel_template
(
    channel_id  bigint unsigned not null comment '渠道 ID',
    template_id bigint unsigned not null comment '模板 ID',
    type        int             not null comment '消息类型',
    `default`   int             not null comment '是否为默认渠道'
) comment '渠道与消息模板关联表';

create table if not exists axg_mail_log
(
    id          bigint unsigned not null comment '主键 ID' primary key,
    created_at  datetime        null comment '创建时间',
    updated_at  datetime        null comment '更新时间',
    deleted_at  datetime        null comment '删除时间',
    app_id      bigint          not null comment '应用 ID',
    template_id bigint          not null comment '邮件模板 ID',
    request_id  varchar(255)    null comment '唯一请求 ID',
    `to`        json            null comment '发件地址集合',
    cc          json            null comment '抄送地址集合',
    bcc         json            null comment '密送地址集合',
    parameters  json            null comment '邮件参数',
    content     longtext        null comment '邮件具体内容',
    attachments json            not null comment '附件ID集合',
    status      int             null comment '发送状态',
    err_msg     varchar(255)    null comment '错误日志'
) comment '邮件发送日志表';

create table if not exists axg_mail_template
(
    id          bigint unsigned not null comment '主键 ID' primary key,
    created_at  datetime        null comment '创建时间',
    updated_at  datetime        null comment '更新时间',
    deleted_at  datetime        null comment '删除时间',
    app_id      bigint          not null comment '应用 ID',
    name        varchar(30)     null comment '模板名称',
    content     longtext        null comment '模板内容',
    type        int             null comment '模板类型（1-文本、2-HTML）',
    attachments json            not null comment '附件ID集合',
    remark      varchar(255)    null comment '备注',
    status      int default 1   not null comment '状态(0-禁用、1-正常)'
) comment '邮件模板配置表';

create table if not exists axg_menu
(
    id         bigint unsigned           not null comment '主键 ID' primary key,
    created_at datetime                  not null comment '创建时间',
    updated_at datetime                  null comment '更新时间',
    deleted_at datetime                  null comment '删除时间',
    parent_id  bigint     default 1      not null comment '父级 ID',
    path       varchar(255) charset utf8 null comment '菜单路径',
    redirect   varchar(255)              null,
    name       varchar(100) charset utf8 null comment '菜单名称',
    hidden     tinyint(1) default 2      not null comment '是否隐藏',
    component  varchar(255) charset utf8 null comment '对应组件',
    sort       bigint                    null comment '排序号',
    is_frame   tinyint(1) default 2      not null comment '是否为 iframe',
    status     bigint     default 1      not null comment '状态(0-禁用、1-正常)',
    no_cache   tinyint(1) default 1      not null comment '是否禁用缓存',
    title      varchar(100) charset utf8 null comment '标题',
    icon       varchar(50) charset utf8  null comment '图标',
    remark     varchar(255)              null comment '备注'
) comment '菜单表';

create table if not exists axg_role
(
    id         bigint unsigned auto_increment comment '主键 ID' primary key,
    created_at datetime         null comment '创建时间',
    updated_at datetime         null comment '更新时间',
    deleted_at datetime         null comment '删除时间',
    name       varchar(30)      null comment '角色名称',
    alias      varchar(30)      null comment '角色别名',
    remark     varchar(255)     null comment '备注',
    status     bigint default 1 not null comment '状态(0-禁用、1-正常)',
    sort       int              null comment '排序号'
) comment '角色表';

create table if not exists axg_role_menu
(
    role_id bigint not null comment '角色 ID',
    menu_id bigint not null comment '菜单 ID'
) comment '角色菜单关联表';

create table if not exists axg_sms_log
(
    id          bigint unsigned not null comment '主键 ID' primary key,
    created_at  datetime        null comment '创建时间',
    updated_at  datetime        null comment '更新时间',
    deleted_at  datetime        null comment '删除时间',
    app_id      bigint          null comment '应用 ID',
    template_id bigint          null comment '模板 ID',
    request_id  varchar(255)    null comment '唯一请求 ID',
    `to`        json            null comment '接收号码集合',
    parameters  json            null comment '模板参数',
    content     tinytext        null comment '短信内容',
    remark      varchar(255)    null comment '备注',
    status      int             null comment '发送状态',
    err_msg     varchar(255)    null comment '错误日志'
) comment '短信发送日志表';

create table if not exists axg_sms_signature
(
    id         bigint unsigned not null comment '主键 ID' primary key,
    created_at datetime        null comment '创建时间',
    updated_at datetime        null comment '更新时间',
    deleted_at datetime        null comment '删除时间',
    app_id     bigint          null comment '应用 ID',
    name       varchar(30)     null comment '签名名称',
    config     json            null comment '签名所需额外配置',
    remark     varchar(255)    null comment '备注',
    status     int default 1   not null comment '状态(0-禁用、1-正常)'
) comment '短信签名表';

create table if not exists axg_sms_template
(
    id         bigint unsigned not null comment '主键 ID' primary key,
    created_at datetime        null comment '创建时间',
    updated_at datetime        null comment '更新时间',
    deleted_at datetime        null comment '删除时间',
    app_id     bigint          null comment '应用 ID',
    type       int             null comment '消息类型(1-验证码、2-通知短信、3-推广短信)',
    name       varchar(30)     null comment '消息模板名称',
    content    tinytext        not null comment '模板内容',
    sign_id    bigint          null comment '关联签名 ID',
    config     json            null comment '模板所需额外配置',
    remark     varchar(255)    null comment '备注',
    status     int default 1   not null comment '状态(0-禁用、1-正常)'
) comment '短信模板配置表';

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
    status     bigint       default 1          not null comment '状态(0-禁用、1-正常)'
) comment '用户表' charset = utf8;

create index idx_users_deleted_at on axg_user (deleted_at);

create table if not exists axg_user_role
(
    user_id bigint unsigned not null comment '用户 ID',
    role_id bigint unsigned not null comment '角色 ID'
) comment '用户角色关联表';
