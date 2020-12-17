CREATE TABLE events
(
    id         BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    sequence   BIGINT UNSIGNED NULL,
    data       MEDIUMBLOB      NOT NULL,
    created_at TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX idx_sequence ON events (sequence);
