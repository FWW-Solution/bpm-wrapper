CREATE TABLE IF NOT EXISTS workflow (
    id serial PRIMARY KEY,
    case_id VARCHAR(255) NOT NULL,
    task_name VARCHAR(255) NOT NULL,
    process_name VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);