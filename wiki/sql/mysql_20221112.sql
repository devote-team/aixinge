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
    redirect   varchar(255) charset utf8 null comment '重定向路径',
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

INSERT INTO aixinge.axg_role 
(id, created_at, updated_at, deleted_at, name, alias, remark, status, sort)
VALUES 
(1, '2022-05-12 11:27:36', '2022-05-12 11:27:36', NULL, '产品经理', 'ttt', 'string', 1, 0);

INSERT INTO aixinge.axg_role 
(id, created_at, updated_at, deleted_at, name, alias, remark, status, sort)
VALUES 
(2, '2022-05-13 08:27:16', '2022-10-30 18:28:33', NULL, '管理员', 'admin', '系统管理员身份', 1, 0);

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


INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(1, CURRENT_TIMESTAMP, 0, '', '/dashboard/workplace', 'dashboard', 0, 'RouteView', 1, 'menu.dashboard', 'dashboard', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(7, CURRENT_TIMESTAMP, 1, '', null, 'Workplace', 0, 'Workplace', 1, 'menu.dashboard.workplace', '', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(2, CURRENT_TIMESTAMP, 1, '/dashboard/analysis', null, 'Analysis', 0, 'Analysis', 1, 'menu.dashboard.analysis', '', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10, CURRENT_TIMESTAMP, 0, null, '/form/base-form', 'form', 0, 'RouteView', 2, 'menu.form', 'form', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(6, CURRENT_TIMESTAMP, 10, null, null, 'basic-form', 0, 'BasicForm', 2, 'menu.form.basic-form', '', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(5, CURRENT_TIMESTAMP, 10, null, null, 'step-form', 0, 'StepForm', 2, 'menu.form.step-form', '', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(4, CURRENT_TIMESTAMP, 10, null, null, 'advanced-form', 0, 'AdvanceForm', 2, 'menu.form.advanced-form', '', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10010, CURRENT_TIMESTAMP, 0, null, '/list/table-list', 'list', 0, 'RouteView', 3, 'menu.list', 'table', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10011, CURRENT_TIMESTAMP, 10010, '/list/table-list/:pageNo([1-9]\\d*)?', null, 'table-list', 0, 'TableList', 3, 'menu.list.table-list', '', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10012, CURRENT_TIMESTAMP, 10010, null, null, 'basic-list', 0, 'StandardList', 3, 'menu.list.basic-list', '', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10013, CURRENT_TIMESTAMP, 10010, null, null, 'card', 0, 'CardList', 3, 'menu.list.card-list', '', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10014, CURRENT_TIMESTAMP, 10010, null, '/list/search/article', 'search', 0, 'SearchLayout', 3, 'menu.list.search-list', '', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10015, CURRENT_TIMESTAMP, 10014, null, null, 'article', 0, 'SearchArticles', 3, 'menu.list.search-list.articles', '', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10016, CURRENT_TIMESTAMP, 10014, null, null, 'project', 0, 'SearchProjects', 3, 'menu.list.search-list.projects', '', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10017, CURRENT_TIMESTAMP, 10014, null, null, 'application', 0, 'SearchApplications', 3, 'menu.list.search-list.applications', '', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10018, CURRENT_TIMESTAMP, 0, null, '/profile/basic', 'profile', 0, 'RouteView', 4, 'menu.profile', 'profile', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10019, CURRENT_TIMESTAMP, 10018, null, null, 'basic', 0, 'ProfileBasic', 4, 'menu.profile.basic', '', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10020, CURRENT_TIMESTAMP, 10018, null, null, 'advanced', 0, 'ProfileAdvanced', 4, 'menu.profile.advanced', '', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10021, CURRENT_TIMESTAMP, 0, null, '/result/success', 'result', 0, 'PageView', 5, 'menu.result', 'check-circle-o', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10022, CURRENT_TIMESTAMP, 10021, null, null, 'success', 0, 'ResultSuccess', 5, 'menu.result.success', '', 'hiddenHeaderContent');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10023, CURRENT_TIMESTAMP, 10021, null, null, 'fail', 0, 'ResultFail', 5, 'menu.result.fail', '', 'hiddenHeaderContent');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10024, CURRENT_TIMESTAMP, 10021, '/exception/icon-selector', null, 'fail', 0, 'other/IconSelectorView', 5, 'IconSelector', 'tool', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10086, CURRENT_TIMESTAMP, 0, '/management', '/management/permission/list', 'management', 0, 'RouteView', 6, 'menu.management', 'align-center', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10087, CURRENT_TIMESTAMP, 10086, '/management/permission/list', null, 'PermissionList', 0, 'other/PermissionList', 6, '权限列表', 'safety-certificate', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10088, CURRENT_TIMESTAMP, 10086, '/management/user/list', null, 'UserList', 0, 'other/UserList', 6, '用户列表', 'team', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10089, CURRENT_TIMESTAMP, 10086, '/management/role/list', null, 'RoleList', 0, 'other/RoleList', 6, '角色列表', 'gateway', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10028, CURRENT_TIMESTAMP, 0, null, '/account/center', 'account', 0, 'RouteView', 7, 'menu.account', 'user', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10029, CURRENT_TIMESTAMP, 10028, 'path', 'redirect', 'center', 0, 'AccountCenter', 7, 'menu.account.center', '', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10030, CURRENT_TIMESTAMP, 10028, null, '/account/settings/basic', 'settings', 0, 'AccountSettings', 8, 'menu.account.settings', '', 'hideHeader,hideChildren');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10031, CURRENT_TIMESTAMP, 10030, '/account/settings/basic', null, 'BasicSettings', 0, 'BasicSetting', 8, 'account.settings.menuMap.basic', '', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10032, CURRENT_TIMESTAMP, 10030, '/account/settings/security', null, 'SecuritySettings', 0, 'SecuritySettings', 8, 'account.settings.menuMap.security', '', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10033, CURRENT_TIMESTAMP, 10030, '/account/settings/custom', null, 'CustomSettings', 0, 'component', 8, 'account.settings.menuMap.custom', '', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10034, CURRENT_TIMESTAMP, 10030, '/account/settings/binding', null, 'BindingSettings', 0, 'BindingSettings', 8, 'account.settings.menuMap.binding', '', '');

INSERT INTO aixinge.axg_menu 
(id, created_at, parent_id, path, redirect, name, hidden, component, sort, title, icon, remark)
VALUES
(10035, CURRENT_TIMESTAMP, 10030, '/account/settings/notification', null, 'NotificationSettings', 0, 'NotificationSettings', 8, 'account.settings.menuMap.notification', '', '');

INSERT INTO aixinge.axg_user_role (user_id, role_id) VALUES (1, 2);
-- 根节点
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 0);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 1);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 2);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 4);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 5);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 6);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 7);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10010);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10011);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10012);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10013);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10014);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10015);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10016);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10017);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10018);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10019);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10020);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10021);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10022);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10023);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10024);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10025);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10026);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10027);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10028);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10029);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10030);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10031);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10032);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10033);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10034);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10035);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10086);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10087);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10088);
INSERT INTO aixinge.axg_role_menu (role_id, menu_id) VALUES (2, 10089);

