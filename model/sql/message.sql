CREATE TABLE message (
                       id bigint AUTO_INCREMENT,
                       to_user_id bigint,
                       from_user_id bigint,
                       content varchar(255),
                       create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'message table';