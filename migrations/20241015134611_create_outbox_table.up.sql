CREATE TABLE outbox
(
    id             bigserial PRIMARY KEY,
    aggregate_id   BIGINT       NOT NULL,
    aggregate_type VARCHAR(255) NOT NULL,
    payload        JSONB         NOT NULL,
    trace_parent   varchar(255),
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);