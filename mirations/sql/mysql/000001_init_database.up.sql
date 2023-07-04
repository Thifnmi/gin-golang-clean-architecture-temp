CREATE TABLE examples (
    id INT NOT NULL AUTO_INCREMENT,
    uuid varchar(36) not null,
    name varchar(255),
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    INDEX logs_id (uuid),
    PRIMARY KEY(id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;