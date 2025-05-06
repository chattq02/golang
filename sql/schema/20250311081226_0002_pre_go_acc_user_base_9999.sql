-- +goose Up
-- +goose StatementBegin
CREATE TABLE `pre_go_acc_user_base_9999` (
  `user_id` INT NOT NULL AUTO_INCREMENT,           -- User ID
  `user_account` VARCHAR(255) NOT NULL,            -- User account name
  `user_password` VARCHAR(255) NOT NULL,           -- User password
  `user_salt` VARCHAR(255) NOT NULL,               -- Salt for password hashing
  `user_login_time` TIMESTAMP NULL DEFAULT NULL,   -- Last login time
  `user_logout_time` TIMESTAMP NULL DEFAULT NULL,  -- Last logout time
  `user_login_ip` VARCHAR(45) DEFAULT NULL,        -- Last login IP address
  `user_created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP, -- Account creation time
  `user_updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Last update time
  PRIMARY KEY (`user_id`),                         -- Primary key for the table
  UNIQUE KEY `unique_user_account` (`user_account`) -- Ensure user_account is unique
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
COMMENT='pre_go_acc_user_base_9999';


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_acc_user_base_9999`;
-- +goose StatementEnd
