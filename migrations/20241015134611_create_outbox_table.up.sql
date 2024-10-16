CREATE TABLE outbox
(
    id             BIGINT PRIMARY KEY AUTO_INCREMENT,
    aggregate_id   BIGINT       NOT NULL,
    aggregate_type VARCHAR(255) NOT NULL,
    payload        JSON         NOT NULL,
    trace_parent   varchar(255),
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);