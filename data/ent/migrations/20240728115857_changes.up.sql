-- create "sys_apis" table
CREATE TABLE `sys_apis` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `path` text NOT NULL, `description` text NOT NULL, `api_group` text NOT NULL, `method` text NOT NULL DEFAULT 'POST');
-- create index "api_path_method" to table: "sys_apis"
CREATE UNIQUE INDEX `api_path_method` ON `sys_apis` (`path`, `method`);
-- create "sys_dictionaries" table
CREATE TABLE `sys_dictionaries` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `status` integer NULL DEFAULT 1, `title` text NOT NULL, `name` text NOT NULL, `description` text NOT NULL);
-- create index "sys_dictionaries_name_key" to table: "sys_dictionaries"
CREATE UNIQUE INDEX `sys_dictionaries_name_key` ON `sys_dictionaries` (`name`);
-- create "sys_dictionary_details" table
CREATE TABLE `sys_dictionary_details` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `status` integer NULL DEFAULT 1, `title` text NOT NULL, `key` text NOT NULL, `value` text NOT NULL, `dictionary_id` integer NULL);
-- create index "dictionarydetail_key_dictionary_id" to table: "sys_dictionary_details"
CREATE UNIQUE INDEX `dictionarydetail_key_dictionary_id` ON `sys_dictionary_details` (`key`, `dictionary_id`);
-- create "sys_logs" table
CREATE TABLE `sys_logs` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `method` text NOT NULL, `api` text NOT NULL, `success` bool NOT NULL, `req_content` text NULL, `resp_content` text NULL, `ip` text NULL, `user_agent` text NULL, `operator` text NULL, `time` integer NULL);
-- create index "logs_api" to table: "sys_logs"
CREATE INDEX `logs_api` ON `sys_logs` (`api`);
-- create "sys_menus" table
CREATE TABLE `sys_menus` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `menu_level` integer NOT NULL, `menu_type` integer NOT NULL, `path` text NULL DEFAULT '', `name` text NOT NULL, `redirect` text NULL DEFAULT '', `component` text NULL DEFAULT '', `order_no` integer NOT NULL DEFAULT 0, `disabled` bool NULL DEFAULT false, `title` text NOT NULL, `icon` text NOT NULL, `hide_menu` bool NULL DEFAULT false, `hide_breadcrumb` bool NULL DEFAULT false, `current_active_menu` text NULL DEFAULT '', `ignore_keep_alive` bool NULL DEFAULT false, `hide_tab` bool NULL DEFAULT false, `frame_src` text NULL DEFAULT '', `carry_param` bool NULL DEFAULT false, `hide_children_in_menu` bool NULL DEFAULT false, `affix` bool NULL DEFAULT false, `dynamic_level` integer NULL DEFAULT 20, `real_path` text NULL DEFAULT '', `parent_id` integer NULL);
-- create "sys_menu_params" table
CREATE TABLE `sys_menu_params` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `type` text NOT NULL, `key` text NOT NULL, `value` text NOT NULL, `menu_params` integer NULL);
-- create "sys_oauth_providers" table
CREATE TABLE `sys_oauth_providers` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `name` text NOT NULL, `app_id` text NULL, `client_id` text NOT NULL, `client_secret` text NULL, `redirect_url` text NOT NULL, `scopes` text NULL, `auth_url` text NOT NULL, `token_url` text NULL, `auth_style` integer NULL, `info_url` text NULL);
-- create index "sys_oauth_providers_name_key" to table: "sys_oauth_providers"
CREATE UNIQUE INDEX `sys_oauth_providers_name_key` ON `sys_oauth_providers` (`name`);
-- create "sys_roles" table
CREATE TABLE `sys_roles` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `status` integer NULL DEFAULT 1, `name` text NOT NULL, `value` text NOT NULL, `default_router` text NOT NULL DEFAULT 'dashboard', `remark` text NOT NULL DEFAULT '', `order_no` integer NOT NULL DEFAULT 0);
-- create index "sys_roles_value_key" to table: "sys_roles"
CREATE UNIQUE INDEX `sys_roles_value_key` ON `sys_roles` (`value`);
-- create "sys_tokens" table
CREATE TABLE `sys_tokens` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `user_id` integer NOT NULL, `token` text NOT NULL, `source` text NOT NULL, `expired_at` datetime NOT NULL, `user_token` integer NULL);
-- create index "sys_tokens_user_id_key" to table: "sys_tokens"
CREATE UNIQUE INDEX `sys_tokens_user_id_key` ON `sys_tokens` (`user_id`);
-- create index "sys_tokens_user_token_key" to table: "sys_tokens"
CREATE UNIQUE INDEX `sys_tokens_user_token_key` ON `sys_tokens` (`user_token`);
-- create index "token_user_id" to table: "sys_tokens"
CREATE INDEX `token_user_id` ON `sys_tokens` (`user_id`);
-- create index "token_expired_at" to table: "sys_tokens"
CREATE INDEX `token_expired_at` ON `sys_tokens` (`expired_at`);
-- create "sys_users" table
CREATE TABLE `sys_users` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `status` integer NULL DEFAULT 1, `username` text NOT NULL, `password` text NOT NULL, `nickname` text NOT NULL, `side_mode` text NULL DEFAULT 'dark', `base_color` text NULL DEFAULT '#fff', `active_color` text NULL DEFAULT '#1890ff', `role_id` integer NULL DEFAULT 2, `mobile` text NOT NULL, `email` text NULL, `wecom` text NULL, `avatar` text NULL DEFAULT '');
-- create index "sys_users_username_key" to table: "sys_users"
CREATE UNIQUE INDEX `sys_users_username_key` ON `sys_users` (`username`);
-- create index "sys_users_nickname_key" to table: "sys_users"
CREATE UNIQUE INDEX `sys_users_nickname_key` ON `sys_users` (`nickname`);
-- create index "sys_users_mobile_key" to table: "sys_users"
CREATE UNIQUE INDEX `sys_users_mobile_key` ON `sys_users` (`mobile`);
-- create index "user_username_email" to table: "sys_users"
CREATE UNIQUE INDEX `user_username_email` ON `sys_users` (`username`, `email`);
-- create "role_menus" table
CREATE TABLE `role_menus` (`role_id` integer NOT NULL, `menu_id` integer NOT NULL, PRIMARY KEY (`role_id`, `menu_id`));
