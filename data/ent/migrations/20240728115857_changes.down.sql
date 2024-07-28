-- reverse: create "role_menus" table
DROP TABLE `role_menus`;
-- reverse: create index "user_username_email" to table: "sys_users"
DROP INDEX `user_username_email`;
-- reverse: create index "sys_users_mobile_key" to table: "sys_users"
DROP INDEX `sys_users_mobile_key`;
-- reverse: create index "sys_users_nickname_key" to table: "sys_users"
DROP INDEX `sys_users_nickname_key`;
-- reverse: create index "sys_users_username_key" to table: "sys_users"
DROP INDEX `sys_users_username_key`;
-- reverse: create "sys_users" table
DROP TABLE `sys_users`;
-- reverse: create index "token_expired_at" to table: "sys_tokens"
DROP INDEX `token_expired_at`;
-- reverse: create index "token_user_id" to table: "sys_tokens"
DROP INDEX `token_user_id`;
-- reverse: create index "sys_tokens_user_token_key" to table: "sys_tokens"
DROP INDEX `sys_tokens_user_token_key`;
-- reverse: create index "sys_tokens_user_id_key" to table: "sys_tokens"
DROP INDEX `sys_tokens_user_id_key`;
-- reverse: create "sys_tokens" table
DROP TABLE `sys_tokens`;
-- reverse: create index "sys_roles_value_key" to table: "sys_roles"
DROP INDEX `sys_roles_value_key`;
-- reverse: create "sys_roles" table
DROP TABLE `sys_roles`;
-- reverse: create index "sys_oauth_providers_name_key" to table: "sys_oauth_providers"
DROP INDEX `sys_oauth_providers_name_key`;
-- reverse: create "sys_oauth_providers" table
DROP TABLE `sys_oauth_providers`;
-- reverse: create "sys_menu_params" table
DROP TABLE `sys_menu_params`;
-- reverse: create "sys_menus" table
DROP TABLE `sys_menus`;
-- reverse: create index "logs_api" to table: "sys_logs"
DROP INDEX `logs_api`;
-- reverse: create "sys_logs" table
DROP TABLE `sys_logs`;
-- reverse: create index "dictionarydetail_key_dictionary_id" to table: "sys_dictionary_details"
DROP INDEX `dictionarydetail_key_dictionary_id`;
-- reverse: create "sys_dictionary_details" table
DROP TABLE `sys_dictionary_details`;
-- reverse: create index "sys_dictionaries_name_key" to table: "sys_dictionaries"
DROP INDEX `sys_dictionaries_name_key`;
-- reverse: create "sys_dictionaries" table
DROP TABLE `sys_dictionaries`;
-- reverse: create index "api_path_method" to table: "sys_apis"
DROP INDEX `api_path_method`;
-- reverse: create "sys_apis" table
DROP TABLE `sys_apis`;
