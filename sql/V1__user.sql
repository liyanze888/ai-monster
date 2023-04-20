CREATE TABLE `user`
(
    `id`           BIGINT       NOT NULL PRIMARY KEY,
    `type`         int          NOT NULL DEFAULT 0,
    `user_name`    VARCHAR(64)  NOT NULL DEFAULT '',
    `nick_name`    VARCHAR(128) NOT NULL DEFAULT '',
    `created_time` BIGINT       NOT NULL DEFAULT 0,
    `updated_time` BIGINT       NOT NULL DEFAULT 0
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE `prompts_index`
(
    `id`           BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `prompts_id`   BIGINT NOT NULL DEFAULT 0,
    `model`        INT    NOT NULL DEFAULT 0,
    `category`     INT    NOT NULL DEFAULT 0,
    `version`      INT    NOT NULL DEFAULT 0,
    `value`        int    NOT NULL DEFAULT 0,
    `value_num`    int    NOT NULL DEFAULT 0,
    `like_num`     int    NOT NULL DEFAULT 0,
    `created_time` BIGINT NOT NULL DEFAULT 0,
    `updated_time` BIGINT NOT NULL DEFAULT 0,
    UNIQUE `prompts_index_model_category` (`model`, `category`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


CREATE TABLE `prompts`
(
    `id`           BIGINT       NOT NULL PRIMARY KEY,
    `user_id`      BIGINT(64)   NOT NULL DEFAULT 0,
    `p_type`       INT(64)      NOT NULL DEFAULT 0,
    `title`        TEXT         NOT NULL,
    `author`       VARCHAR(128) NOT NULL DEFAULT '',
    `model`        int          NOT NULL DEFAULT 0,
    `version`      int          NOT NULL DEFAULT 0,
    `content_json` LONGTEXT     NOT NULL,
    `created_time` BIGINT       NOT NULL DEFAULT 0,
    `updated_time` BIGINT       NOT NULL DEFAULT 0
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


CREATE TABLE `medias`
(
    `id`           BIGINT NOT NULL PRIMARY KEY,
    `target_id`    BIGINT NOT NULL,
    `media_id`     BIGINT NOT NULL DEFAULT 0,
    `created_time` BIGINT NOT NULL DEFAULT 0,
    `updated_time` BIGINT NOT NULL DEFAULT 0,
    key `medias_target_id` (`target_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;



CREATE TABLE `dict`
(
    `id`        BIGINT       NOT NULL PRIMARY KEY,
    `type_id`   int          NOT NULL,
    `sub_type`  int          NOT NULL DEFAULT 0,
    `parent_id` INT          NOT NULL DEFAULT 0,
    `locale`    char(2)      NOT NULL DEFAULT 'en',
    `content`   varchar(128) NOT NULL,
    key `dict_type_id` (`type_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;