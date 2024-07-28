-- disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- create "new_sys_apis" table
CREATE TABLE `new_sys_apis` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `path` text NOT NULL, `description` text NOT NULL, `api_group` text NOT NULL, `method` text NOT NULL DEFAULT 'POST');
-- copy rows from old table "sys_apis" to new temporary table "new_sys_apis"
INSERT INTO `new_sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `method`) SELECT `id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `method` FROM `sys_apis`;
-- drop "sys_apis" table after copying rows
DROP TABLE `sys_apis`;
-- rename temporary table "new_sys_apis" to "sys_apis"
ALTER TABLE `new_sys_apis` RENAME TO `sys_apis`;
-- create index "api_path_method" to table: "sys_apis"
CREATE UNIQUE INDEX `api_path_method` ON `sys_apis` (`path`, `method`);
-- create "new_sys_dictionaries" table
CREATE TABLE `new_sys_dictionaries` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `status` integer NULL DEFAULT 1, `title` text NOT NULL, `name` text NOT NULL, `description` text NOT NULL);
-- copy rows from old table "sys_dictionaries" to new temporary table "new_sys_dictionaries"
INSERT INTO `new_sys_dictionaries` (`id`, `created_at`, `updated_at`, `status`, `title`, `name`, `description`) SELECT `id`, `created_at`, `updated_at`, `status`, `title`, `name`, `description` FROM `sys_dictionaries`;
-- drop "sys_dictionaries" table after copying rows
DROP TABLE `sys_dictionaries`;
-- rename temporary table "new_sys_dictionaries" to "sys_dictionaries"
ALTER TABLE `new_sys_dictionaries` RENAME TO `sys_dictionaries`;
-- create index "sys_dictionaries_name_key" to table: "sys_dictionaries"
CREATE UNIQUE INDEX `sys_dictionaries_name_key` ON `sys_dictionaries` (`name`);
-- create "new_sys_dictionary_details" table
CREATE TABLE `new_sys_dictionary_details` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `status` integer NULL DEFAULT 1, `title` text NOT NULL, `key` text NOT NULL, `value` text NOT NULL, `dictionary_id` integer NULL);
-- copy rows from old table "sys_dictionary_details" to new temporary table "new_sys_dictionary_details"
INSERT INTO `new_sys_dictionary_details` (`id`, `created_at`, `updated_at`, `status`, `title`, `key`, `value`, `dictionary_id`) SELECT `id`, `created_at`, `updated_at`, `status`, `title`, `key`, `value`, `dictionary_id` FROM `sys_dictionary_details`;
-- drop "sys_dictionary_details" table after copying rows
DROP TABLE `sys_dictionary_details`;
-- rename temporary table "new_sys_dictionary_details" to "sys_dictionary_details"
ALTER TABLE `new_sys_dictionary_details` RENAME TO `sys_dictionary_details`;
-- create index "dictionarydetail_key_dictionary_id" to table: "sys_dictionary_details"
CREATE UNIQUE INDEX `dictionarydetail_key_dictionary_id` ON `sys_dictionary_details` (`key`, `dictionary_id`);
-- create "new_sys_logs" table
CREATE TABLE `new_sys_logs` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `method` text NOT NULL, `api` text NOT NULL, `success` bool NOT NULL, `req_content` text NULL, `resp_content` text NULL, `ip` text NULL, `user_agent` text NULL, `operator` text NULL, `time` integer NULL);
-- copy rows from old table "sys_logs" to new temporary table "new_sys_logs"
INSERT INTO `new_sys_logs` (`id`, `created_at`, `updated_at`, `type`, `method`, `api`, `success`, `req_content`, `resp_content`, `ip`, `user_agent`, `operator`, `time`) SELECT `id`, `created_at`, `updated_at`, `type`, `method`, `api`, `success`, `req_content`, `resp_content`, `ip`, `user_agent`, `operator`, `time` FROM `sys_logs`;
-- drop "sys_logs" table after copying rows
DROP TABLE `sys_logs`;
-- rename temporary table "new_sys_logs" to "sys_logs"
ALTER TABLE `new_sys_logs` RENAME TO `sys_logs`;
-- create index "logs_api" to table: "sys_logs"
CREATE INDEX `logs_api` ON `sys_logs` (`api`);
-- create "new_sys_menus" table
CREATE TABLE `new_sys_menus` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `menu_level` integer NOT NULL, `menu_type` integer NOT NULL, `path` text NULL DEFAULT '', `name` text NOT NULL, `redirect` text NULL DEFAULT '', `component` text NULL DEFAULT '', `order_no` integer NOT NULL DEFAULT 0, `disabled` bool NULL DEFAULT false, `title` text NOT NULL, `icon` text NOT NULL, `hide_menu` bool NULL DEFAULT false, `hide_breadcrumb` bool NULL DEFAULT false, `current_active_menu` text NULL DEFAULT '', `ignore_keep_alive` bool NULL DEFAULT false, `hide_tab` bool NULL DEFAULT false, `frame_src` text NULL DEFAULT '', `carry_param` bool NULL DEFAULT false, `hide_children_in_menu` bool NULL DEFAULT false, `affix` bool NULL DEFAULT false, `dynamic_level` integer NULL DEFAULT 20, `real_path` text NULL DEFAULT '', `parent_id` integer NULL);
-- copy rows from old table "sys_menus" to new temporary table "new_sys_menus"
INSERT INTO `new_sys_menus` (`id`, `created_at`, `updated_at`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `order_no`, `disabled`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `current_active_menu`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id`) SELECT `id`, `created_at`, `updated_at`, `menu_level`, `menu_type`, `path`, `name`, `redirect`, `component`, `order_no`, `disabled`, `title`, `icon`, `hide_menu`, `hide_breadcrumb`, `current_active_menu`, `ignore_keep_alive`, `hide_tab`, `frame_src`, `carry_param`, `hide_children_in_menu`, `affix`, `dynamic_level`, `real_path`, `parent_id` FROM `sys_menus`;
-- drop "sys_menus" table after copying rows
DROP TABLE `sys_menus`;
-- rename temporary table "new_sys_menus" to "sys_menus"
ALTER TABLE `new_sys_menus` RENAME TO `sys_menus`;
-- create "new_sys_menu_params" table
CREATE TABLE `new_sys_menu_params` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `key` text NOT NULL, `value` text NOT NULL, `menu_params` integer NULL);
-- copy rows from old table "sys_menu_params" to new temporary table "new_sys_menu_params"
INSERT INTO `new_sys_menu_params` (`id`, `created_at`, `updated_at`, `type`, `key`, `value`, `menu_params`) SELECT `id`, `created_at`, `updated_at`, `type`, `key`, `value`, `menu_params` FROM `sys_menu_params`;
-- drop "sys_menu_params" table after copying rows
DROP TABLE `sys_menu_params`;
-- rename temporary table "new_sys_menu_params" to "sys_menu_params"
ALTER TABLE `new_sys_menu_params` RENAME TO `sys_menu_params`;
-- create "new_sys_oauth_providers" table
CREATE TABLE `new_sys_oauth_providers` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `name` text NOT NULL, `app_id` text NULL, `client_id` text NOT NULL, `client_secret` text NULL, `redirect_url` text NOT NULL, `scopes` text NULL, `auth_url` text NOT NULL, `token_url` text NULL, `auth_style` integer NULL, `info_url` text NULL);
-- copy rows from old table "sys_oauth_providers" to new temporary table "new_sys_oauth_providers"
INSERT INTO `new_sys_oauth_providers` (`id`, `created_at`, `updated_at`, `name`, `app_id`, `client_id`, `client_secret`, `redirect_url`, `scopes`, `auth_url`, `token_url`, `auth_style`, `info_url`) SELECT `id`, `created_at`, `updated_at`, `name`, `app_id`, `client_id`, `client_secret`, `redirect_url`, `scopes`, `auth_url`, `token_url`, `auth_style`, `info_url` FROM `sys_oauth_providers`;
-- drop "sys_oauth_providers" table after copying rows
DROP TABLE `sys_oauth_providers`;
-- rename temporary table "new_sys_oauth_providers" to "sys_oauth_providers"
ALTER TABLE `new_sys_oauth_providers` RENAME TO `sys_oauth_providers`;
-- create index "sys_oauth_providers_name_key" to table: "sys_oauth_providers"
CREATE UNIQUE INDEX `sys_oauth_providers_name_key` ON `sys_oauth_providers` (`name`);
-- create "new_sys_roles" table
CREATE TABLE `new_sys_roles` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `status` integer NULL DEFAULT 1, `name` text NOT NULL, `value` text NOT NULL, `default_router` text NOT NULL DEFAULT 'dashboard', `remark` text NOT NULL DEFAULT '', `order_no` integer NOT NULL DEFAULT 0);
-- copy rows from old table "sys_roles" to new temporary table "new_sys_roles"
INSERT INTO `new_sys_roles` (`id`, `created_at`, `updated_at`, `status`, `name`, `value`, `default_router`, `remark`, `order_no`) SELECT `id`, `created_at`, `updated_at`, `status`, `name`, `value`, `default_router`, `remark`, `order_no` FROM `sys_roles`;
-- drop "sys_roles" table after copying rows
DROP TABLE `sys_roles`;
-- rename temporary table "new_sys_roles" to "sys_roles"
ALTER TABLE `new_sys_roles` RENAME TO `sys_roles`;
-- create index "sys_roles_value_key" to table: "sys_roles"
CREATE UNIQUE INDEX `sys_roles_value_key` ON `sys_roles` (`value`);
-- create "new_sys_tokens" table
CREATE TABLE `new_sys_tokens` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `user_id` integer NOT NULL, `token` text NOT NULL, `source` text NOT NULL, `expired_at` datetime NOT NULL, `user_token` integer NULL);
-- copy rows from old table "sys_tokens" to new temporary table "new_sys_tokens"
INSERT INTO `new_sys_tokens` (`id`, `created_at`, `updated_at`, `user_id`, `token`, `source`, `expired_at`, `user_token`) SELECT `id`, `created_at`, `updated_at`, `user_id`, `token`, `source`, `expired_at`, `user_token` FROM `sys_tokens`;
-- drop "sys_tokens" table after copying rows
DROP TABLE `sys_tokens`;
-- rename temporary table "new_sys_tokens" to "sys_tokens"
ALTER TABLE `new_sys_tokens` RENAME TO `sys_tokens`;
-- create index "sys_tokens_user_id_key" to table: "sys_tokens"
CREATE UNIQUE INDEX `sys_tokens_user_id_key` ON `sys_tokens` (`user_id`);
-- create index "sys_tokens_user_token_key" to table: "sys_tokens"
CREATE UNIQUE INDEX `sys_tokens_user_token_key` ON `sys_tokens` (`user_token`);
-- create index "token_user_id" to table: "sys_tokens"
CREATE INDEX `token_user_id` ON `sys_tokens` (`user_id`);
-- create index "token_expired_at" to table: "sys_tokens"
CREATE INDEX `token_expired_at` ON `sys_tokens` (`expired_at`);
-- create "new_sys_users" table
CREATE TABLE `new_sys_users` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `status` integer NULL DEFAULT 1, `username` text NOT NULL, `password` text NOT NULL, `nickname` text NOT NULL, `side_mode` text NULL DEFAULT 'dark', `base_color` text NULL DEFAULT '#fff', `active_color` text NULL DEFAULT '#1890ff', `role_id` integer NULL DEFAULT 2, `mobile` text NOT NULL, `email` text NULL, `wecom` text NULL, `avatar` text NULL DEFAULT '');
-- copy rows from old table "sys_users" to new temporary table "new_sys_users"
INSERT INTO `new_sys_users` (`id`, `created_at`, `updated_at`, `status`, `username`, `password`, `nickname`, `side_mode`, `base_color`, `active_color`, `role_id`, `mobile`, `email`, `wecom`, `avatar`) SELECT `id`, `created_at`, `updated_at`, `status`, `username`, `password`, `nickname`, `side_mode`, `base_color`, `active_color`, `role_id`, `mobile`, `email`, `wecom`, `avatar` FROM `sys_users`;
-- drop "sys_users" table after copying rows
DROP TABLE `sys_users`;
-- rename temporary table "new_sys_users" to "sys_users"
ALTER TABLE `new_sys_users` RENAME TO `sys_users`;
-- create index "sys_users_username_key" to table: "sys_users"
CREATE UNIQUE INDEX `sys_users_username_key` ON `sys_users` (`username`);
-- create index "sys_users_nickname_key" to table: "sys_users"
CREATE UNIQUE INDEX `sys_users_nickname_key` ON `sys_users` (`nickname`);
-- create index "sys_users_mobile_key" to table: "sys_users"
CREATE UNIQUE INDEX `sys_users_mobile_key` ON `sys_users` (`mobile`);
-- create index "user_username_email" to table: "sys_users"
CREATE UNIQUE INDEX `user_username_email` ON `sys_users` (`username`, `email`);
-- enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
