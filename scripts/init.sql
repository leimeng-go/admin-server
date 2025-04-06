CREATE DATABASE IF NOT EXISTS admin DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE admin;

CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    nickname VARCHAR(50) NOT NULL,
    avatar VARCHAR(255) DEFAULT '',
    role VARCHAR(20) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_username (username)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 插入默认管理员账号，密码为 admin123
INSERT INTO users (username, password, nickname, role) VALUES 
('admin', '$2a$10$RD6c0YT0TiXMgf/2qAFXpehwuTVoQbLLBa.DCUvIKqFj.Ld7SL9Hy', '管理员', 'admin')
ON DUPLICATE KEY UPDATE id=id; 

-- 菜单表
CREATE TABLE menu (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '主键ID',
    menu_id BIGINT NOT NULL COMMENT '业务菜单ID',
    parent_menu_id BIGINT NOT NULL COMMENT '父级业务菜单ID',
    name VARCHAR(50) NOT NULL COMMENT '菜单名称',
    path VARCHAR(100) NOT NULL COMMENT '路由路径',
    component VARCHAR(100) COMMENT '组件路径',
    title VARCHAR(50) NOT NULL COMMENT '菜单标题',
    icon VARCHAR(50) COMMENT '图标',
    keep_alive TINYINT(1) DEFAULT 0 COMMENT '是否缓存',
    is_hide TINYINT(1) DEFAULT 0 COMMENT '是否隐藏',
    is_hide_tab TINYINT(1) DEFAULT 0 COMMENT '是否隐藏标签页',
    is_iframe TINYINT(1) DEFAULT 0 COMMENT '是否是iframe',
    link VARCHAR(255) COMMENT '外部链接',
    show_badge TINYINT(1) DEFAULT 0 COMMENT '是否显示徽章',
    show_text_badge VARCHAR(50) COMMENT '徽章文本',
    is_in_main_container TINYINT(1) DEFAULT 0 COMMENT '是否在主容器中',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY uk_menu_id (menu_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单表';

-- 权限表
CREATE TABLE auth (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '主键ID',
    auth_id BIGINT NOT NULL COMMENT '业务权限ID',
    menu_id BIGINT NOT NULL COMMENT '关联的菜单业务ID',
    title VARCHAR(50) NOT NULL COMMENT '权限标题',
    auth_mark VARCHAR(100) NOT NULL COMMENT '权限标识',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY uk_auth_id (auth_id),
    KEY idx_menu_id (menu_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限表';

 -- 菜单表插入语句
INSERT INTO menu (menu_id, parent_menu_id, name, path, component, title, icon, keep_alive, is_hide, is_hide_tab, is_iframe, link, show_badge, show_text_badge, is_in_main_container) VALUES
-- Dashboard
(1, 0, 'Dashboard', '/dashboard', 'RoutesAlias.Home', 'menus.dashboard.title', '&#xe721;', 0, 0, 0, 0, NULL, 0, NULL, 0),
(101, 1, 'Console', 'console', 'RoutesAlias.Dashboard', 'menus.dashboard.console', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(102, 1, 'Analysis', 'analysis', 'RoutesAlias.Analysis', 'menus.dashboard.analysis', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(103, 1, 'Ecommerce', 'ecommerce', 'RoutesAlias.Ecommerce', 'menus.dashboard.ecommerce', NULL, 1, 0, 0, 0, NULL, 0, 'new', 0),

-- Widgets
(5, 0, 'Widgets', '/widgets', 'RoutesAlias.Home', 'menus.widgets.title', '&#xe81a;', 0, 0, 0, 0, NULL, 0, NULL, 0),
(503, 5, 'IconList', 'icon-list', 'RoutesAlias.IconList', 'menus.widgets.iconList', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(504, 5, 'IconSelector', 'icon-selector', 'RoutesAlias.IconSelector', 'menus.widgets.iconSelector', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(505, 5, 'ImageCrop', 'image-crop', 'RoutesAlias.ImageCrop', 'menus.widgets.imageCrop', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(506, 5, 'Excel', 'excel', 'RoutesAlias.Excel', 'menus.widgets.excel', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(507, 5, 'Video', 'video', 'RoutesAlias.Video', 'menus.widgets.video', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(508, 5, 'CountTo', 'count-to', 'RoutesAlias.CountTo', 'menus.widgets.countTo', NULL, 0, 0, 0, 0, NULL, 0, NULL, 0),
(509, 5, 'WangEditor', 'wang-editor', 'RoutesAlias.WangEditor', 'menus.widgets.wangEditor', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(510, 5, 'Watermark', 'watermark', 'RoutesAlias.Watermark', 'menus.widgets.watermark', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(511, 5, 'ContextMenu', 'context-menu', 'RoutesAlias.ContextMenu', 'menus.widgets.contextMenu', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(512, 5, 'Qrcode', 'qrcode', 'RoutesAlias.Qrcode', 'menus.widgets.qrcode', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(513, 5, 'Drag', 'drag', 'RoutesAlias.Drag', 'menus.widgets.drag', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(514, 5, 'TextScroll', 'text-scroll', 'RoutesAlias.TextScroll', 'menus.widgets.textScroll', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(515, 5, 'Fireworks', 'fireworks', 'RoutesAlias.Fireworks', 'menus.widgets.fireworks', NULL, 1, 0, 0, 0, NULL, 0, 'Hot', 0),
(516, 5, 'ElementUI', '/outside/iframe/elementui', '', 'menus.widgets.elementUI', NULL, 0, 0, 0, 1, 'https://element-plus.org/zh-CN/component/overview.html', 1, NULL, 0),

-- Template
(126, 0, 'Template', '/template', 'RoutesAlias.Home', 'menus.template.title', '&#xe860;', 0, 0, 0, 0, NULL, 0, NULL, 0),
(12602, 126, 'Cards', 'cards', 'RoutesAlias.Cards', 'menus.template.cards', NULL, 0, 0, 0, 0, NULL, 0, NULL, 0),
(12603, 126, 'Banners', 'banners', 'RoutesAlias.Banners', 'menus.template.banners', NULL, 0, 0, 0, 0, NULL, 0, NULL, 0),
(12604, 126, 'Charts', 'charts', 'RoutesAlias.Charts', 'menus.template.charts', NULL, 0, 0, 0, 0, NULL, 0, NULL, 0),
(12609, 126, 'Map', 'map', 'RoutesAlias.Map', 'menus.template.map', NULL, 1, 0, 0, 0, NULL, 0, 'new', 0),
(12601, 126, 'Chat', 'chat', 'RoutesAlias.Chat', 'menus.template.chat', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(12605, 126, 'Calendar', 'calendar', 'RoutesAlias.Calendar', 'menus.template.calendar', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(12622, 126, 'Pricing', 'pricing', 'RoutesAlias.Pricing', 'menus.template.pricing', NULL, 1, 0, 1, 0, NULL, 0, NULL, 0),

-- Article
(4, 0, 'Article', '/article', 'RoutesAlias.Home', 'menus.article.title', '&#xe7ae;', 1, 0, 0, 0, NULL, 0, NULL, 0),
(202, 4, 'ArticleList', 'article-list', 'RoutesAlias.ArticleList', 'menus.article.articleList', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(204, 4, 'ArticleDetail', 'detail', 'RoutesAlias.ArticleDetail', 'menus.article.articleDetail', NULL, 1, 1, 0, 0, NULL, 0, NULL, 0),
(205, 4, 'Comment', 'comment', 'RoutesAlias.Comment', 'menus.article.comment', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(201, 4, 'ArticlePublish', 'article-publish', 'RoutesAlias.ArticlePublish', 'menus.article.articlePublish', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),

-- User
(2, 0, 'User', '/user', 'RoutesAlias.Home', 'menus.user.title', '&#xe86e;', 0, 0, 0, 0, NULL, 0, NULL, 0),
(301, 2, 'Account', 'account', 'RoutesAlias.Account', 'menus.user.account', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(302, 2, 'Department', 'department', 'RoutesAlias.Department', 'menus.user.department', NULL, 0, 0, 0, 0, NULL, 0, NULL, 0),
(303, 2, 'Role', 'role', 'RoutesAlias.Role', 'menus.user.role', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(304, 2, 'UserCenter', 'user', 'RoutesAlias.UserCenter', 'menus.user.userCenter', NULL, 1, 1, 1, 0, NULL, 0, NULL, 0),

-- Menu
(3, 0, 'Menu', '/menu', 'RoutesAlias.Home', 'menus.menu.title', '&#xe8a4;', 0, 0, 0, 0, NULL, 0, NULL, 0),
(401, 3, 'Menus', 'menu', 'RoutesAlias.Menu', 'menus.menu.menu', '&#xe8a4;', 1, 0, 0, 0, NULL, 0, NULL, 0),
(411, 3, 'Permission', 'permission', 'RoutesAlias.Permission', 'menus.menu.permission', '&#xe831;', 1, 0, 0, 0, NULL, 0, 'new', 0),
(402, 3, 'Nested', 'nested', '', 'menus.menu.nested', '&#xe676;', 1, 0, 0, 0, NULL, 0, NULL, 0),
(40201, 402, 'NestedMenu1', 'menu1', 'RoutesAlias.NestedMenu1', 'menus.menu.menu1', '&#xe676;', 1, 0, 0, 0, NULL, 0, NULL, 0),
(40202, 402, 'NestedMenu2', 'menu2', '', 'menus.menu.menu2', '&#xe676;', 1, 0, 0, 0, NULL, 0, NULL, 0),
(4020201, 40202, 'NestedMenu2-1', 'menu2-1', 'RoutesAlias.NestedMenu21', 'menus.menu.menu21', '&#xe676;', 1, 0, 0, 0, NULL, 0, NULL, 0),
(40203, 402, 'NestedMenu3', 'menu3', '', 'menus.menu.menu3', '&#xe676;', 1, 0, 0, 0, NULL, 0, NULL, 0),
(4020301, 40203, 'NestedMenu3-1', 'menu3-1', 'RoutesAlias.NestedMenu31', 'menus.menu.menu31', '&#xe676;', 1, 0, 0, 0, NULL, 0, NULL, 0),
(4020302, 40203, 'NestedMenu3-2', 'menu3-2', '', 'menus.menu.menu32', '&#xe676;', 1, 0, 0, 0, NULL, 0, NULL, 0),
(402030201, 4020302, 'NestedMenu3-2-1', 'menu3-2-1', 'RoutesAlias.NestedMenu321', 'menus.menu.menu321', '&#xe676;', 1, 0, 0, 0, NULL, 0, NULL, 0),

-- Result
(18, 0, 'Result', '/result', 'RoutesAlias.Home', 'menus.result.title', '&#xe715;', 0, 0, 0, 0, NULL, 0, NULL, 0),
(1801, 18, 'Success', 'success', 'RoutesAlias.Success', 'menus.result.success', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(1802, 18, 'Fail', 'fail', 'RoutesAlias.Fail', 'menus.result.fail', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),

-- Exception
(8, 0, 'Exception', '/exception', 'RoutesAlias.Home', 'menus.exception.title', '&#xe820;', 0, 0, 0, 0, NULL, 0, NULL, 0),
(801, 8, '403', '403', 'RoutesAlias.Exception403', 'menus.exception.notFound', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(802, 8, '404', '404', 'RoutesAlias.Exception404', 'menus.exception.notFoundEn', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(803, 8, '500', '500', 'RoutesAlias.Exception500', 'menus.exception.serverError', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),

-- System
(9, 0, 'System', '/system', 'RoutesAlias.Home', 'menus.system.title', '&#xe7b9;', 0, 0, 0, 0, NULL, 0, NULL, 0),
(901, 9, 'Setting', 'setting', 'RoutesAlias.Setting', 'menus.system.setting', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(902, 9, 'Api', 'api', 'RoutesAlias.Api', 'menus.system.api', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),
(903, 9, 'Log', 'log', 'RoutesAlias.Log', 'menus.system.log', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),

-- Safeguard
(10, 0, 'Safeguard', '/safeguard', 'RoutesAlias.Home', 'menus.safeguard.title', '&#xe816;', 0, 0, 0, 0, NULL, 0, NULL, 0),
(1010, 10, 'Server', 'server', 'RoutesAlias.Server', 'menus.safeguard.server', NULL, 1, 0, 0, 0, NULL, 0, NULL, 0),

-- Help
(12, 0, '', '', 'RoutesAlias.Home', 'menus.help.title', '&#xe719;', 0, 0, 0, 0, NULL, 0, NULL, 0),
(1101, 12, 'Document', '', '', 'menus.help.document', NULL, 0, 0, 0, 0, 'https://www.lingchen.kim/art-design-pro/docs/', 0, NULL, 0),

-- ChangeLog
(11912, 0, 'ChangeLog', '/log/changeLog', '/log/ChangeLog', 'menus.plan.log', '&#xe712;', 0, 0, 0, 0, NULL, 0, '${upgradeLogList.value[0].version}', 1);

-- 权限表插入语句
INSERT INTO auth (auth_id, menu_id, title, auth_mark) VALUES
-- ArticleList权限
(2021, 202, '新增', 'add'),
(2022, 202, '编辑', 'edit'),

-- ArticlePublish权限
(2010, 201, '发布', 'article/article-publish/add'),

-- Menu权限
(4011, 401, '新增', 'add'),
(4012, 401, '编辑', 'edit'),
(4013, 401, '删除', 'delete'),

-- Permission权限
(4111, 411, '新增', 'add'),
(4112, 411, '编辑', 'edit'),
(4113, 411, '删除', 'delete');