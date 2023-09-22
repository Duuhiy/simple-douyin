CREATE TABLE comment (
                          id bigint AUTO_INCREMENT,
                          user_id bigint,
                          video_id bigint,
                          contents varchar(255),
                          create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          update_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          delete_at timestamp,
                          PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'comment table';
create index ind_video on comment (video_id);