-- Create "gateways" table
CREATE TABLE `gateways` (`id` char(36) NOT NULL, `name` varchar(50) NOT NULL, `ip` varchar(255) NOT NULL, `port` bigint NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "compute_images" table
CREATE TABLE `compute_images` (`id` int NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `image` varchar(255) NOT NULL, `tag` varchar(255) NOT NULL, `port` int NOT NULL, `command` varchar(255) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `computeimage_id` (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "compute_instances" table
CREATE TABLE `compute_instances` (`id` char(36) NOT NULL, `owner` varchar(255) NOT NULL, `name` varchar(255) NOT NULL, `core` varchar(255) NOT NULL, `memory` varchar(255) NOT NULL, `image` varchar(255) NOT NULL, `port` varchar(255) NULL, `expiration_time` timestamp NOT NULL, `status` tinyint NOT NULL, `container_id` varchar(255) NULL, `peer_id` varchar(255) NULL, `command` varchar(255) NULL, PRIMARY KEY (`id`), UNIQUE INDEX `computeinstance_id` (`id`), INDEX `computeinstance_owner` (`owner`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "compute_specs" table
CREATE TABLE `compute_specs` (`id` int NOT NULL AUTO_INCREMENT, `core` varchar(255) NOT NULL, `memory` varchar(255) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `computespec_id` (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "employees" table
CREATE TABLE `employees` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `age` int NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "gateway_ports" table
CREATE TABLE `gateway_ports` (`id` char(36) NOT NULL, `fk_gateway_id` varchar(36) NOT NULL, `port` tinyint NOT NULL, `is_use` bool NOT NULL DEFAULT 0, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "agents" table
CREATE TABLE `agents` (`id` char(36) NOT NULL, `peer_id` varchar(255) NOT NULL, `active` bool NOT NULL DEFAULT 1, `last_update_time` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `agent_id` (`id`), UNIQUE INDEX `peer_id` (`peer_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "network_mappings" table
CREATE TABLE `network_mappings` (`id` char(36) NOT NULL, `name` varchar(50) NOT NULL, `fk_gateway_id` char(36) NOT NULL, `fk_computer_id` char(36) NOT NULL, `gateway_port` bigint NOT NULL, `computer_port` bigint NOT NULL, `status` bigint NOT NULL DEFAULT 0, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "storages" table
CREATE TABLE `storages` (`id` char(36) NOT NULL, `owner` varchar(50) NOT NULL, `type` int NOT NULL DEFAULT 0, `name` varchar(50) NOT NULL, `cid` varchar(80) NOT NULL, `size` int NOT NULL, `last_modify` timestamp NOT NULL, `parent_id` varchar(80) NOT NULL, PRIMARY KEY (`id`), INDEX `storage_owner` (`owner`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "tasks" table
CREATE TABLE `tasks` (`id` char(36) NOT NULL, `agent_id` varchar(50) NOT NULL, `cmd` int NOT NULL DEFAULT 0, `params` varchar(255) NOT NULL, `status` bigint NOT NULL, `create_time` timestamp NOT NULL, PRIMARY KEY (`id`), INDEX `task_agent_id` (`agent_id`), INDEX `task_create_time` (`create_time`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "users" table
CREATE TABLE `users` (`id` char(36) NOT NULL, `country_call_coding` varchar(8) NOT NULL, `telephone_number` varchar(50) NOT NULL, `password` varchar(255) NOT NULL, `create_date` timestamp NOT NULL, `last_login_date` timestamp NOT NULL, `name` varchar(255) NOT NULL, `icon` varchar(255) NOT NULL, `pwd_config` bool NOT NULL DEFAULT 0, PRIMARY KEY (`id`), UNIQUE INDEX `user_country_call_coding_telephone_number` (`country_call_coding`, `telephone_number`), UNIQUE INDEX `user_id` (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "scripts" table
CREATE TABLE `scripts` (`id` int NOT NULL AUTO_INCREMENT, `user_id` varchar(255) NOT NULL, `task_number` int NOT NULL, `script_name` varchar(255) NOT NULL, `file_address` varchar(255) NOT NULL, `script_content` longtext NOT NULL, `create_time` timestamp NOT NULL, `update_time` timestamp NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "script_execution_records" table
CREATE TABLE `script_execution_records` (`id` int NOT NULL AUTO_INCREMENT, `user_id` varchar(255) NOT NULL, `fk_script_id` int NOT NULL, `script_content` longtext NOT NULL, `task_number` int NOT NULL, `script_name` varchar(255) NOT NULL, `file_address` varchar(255) NOT NULL, `execute_state` int NOT NULL, `execute_result` longtext NOT NULL, `create_time` timestamp NOT NULL, `update_time` timestamp NOT NULL, `script_script_execution_records` int NULL, PRIMARY KEY (`id`), INDEX `script_execution_records_scripts_scriptExecutionRecords` (`script_script_execution_records`), CONSTRAINT `script_execution_records_scripts_scriptExecutionRecords` FOREIGN KEY (`script_script_execution_records`) REFERENCES `scripts` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL) CHARSET utf8mb4 COLLATE utf8mb4_bin;
