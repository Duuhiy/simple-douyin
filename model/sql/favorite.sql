CREATE TABLE favorite (
                       id bigint AUTO_INCREMENT,
                       user_id bigint,
                       video_id bigint,
                       create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       update_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                       delete_at timestamp,
                       PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'favorite table';
create index ind_user_video on favorite (user_id, video_id, create_at);