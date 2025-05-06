-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `pre_go_acc_user_info_9999` (
    `user_id` INT NOT NULL COMMENT 'User ID', -- Adjusted to match pre_go_acc_user_base_9999
    `user_account` VARCHAR(255) NOT NULL COMMENT 'User account',
    `user_nickname` VARCHAR(255) COMMENT 'User nickname',
    `user_avatar` VARCHAR(255) COMMENT 'User avatar',
    `user_state` TINYINT UNSIGNED NOT NULL COMMENT 'User state: 0-Locked, 1-Activated, 2-Not Activated',
    `user_mobile` VARCHAR(20) COMMENT 'Mobile phone number',
    `user_gender` TINYINT UNSIGNED COMMENT 'User gender: 0-Secret, 1-Male, 2-Female',
    `user_birthday` DATE COMMENT 'User birthday',
    `user_email` VARCHAR(255) COMMENT 'User email address',
    `user_is_authentication` TINYINT UNSIGNED NOT NULL COMMENT 'Authentication status: 0-Not Authenticated, 1-Pending, 2-Authenticated',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Record creation time',
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update time',
    CONSTRAINT `fk_user_base_to_info` FOREIGN KEY (`user_id`) REFERENCES `pre_go_acc_user_base_9999` (`user_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='account_user_info';



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_acc_user_info_9999`;
-- +goose StatementEnd
