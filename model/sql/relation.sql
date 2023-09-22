CREATE TABLE relation (
                      id bigint AUTO_INCREMENT,
                      user_id bigint ,
                      to_user_id bigint ,
                      create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                      update_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                      PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'relation table';
create index ind_to_from on relation (user_id, to_user_id);
create index ind_from on relation (to_user_id);