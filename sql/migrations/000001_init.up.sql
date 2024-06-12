START TRANSACTION;

CREATE TABLE IF NOT EXISTS projects (
    project_id VARCHAR(36) PRIMARY KEY,
    description VARCHAR(255) NOT NULL,
    start_date DATE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS costs (
    cost_id VARCHAR(36) PRIMARY KEY,
    project_id VARCHAR(36) NOT NULL REFERENCES projects (project_id),
    cost_type VARCHAR(10) NOT NULL,
    description VARCHAR(255) NOT NULL,
    comment VARCHAR(255),
    amount FLOAT8 NOT NULL,
    currency VARCHAR(10) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS installments (
    installment_id VARCHAR(36) PRIMARY KEY,
    cost_id VARCHAR(36) NOT NULL REFERENCES costs (cost_id),
    payment_date DATE NOT NULL,
    amount FLOAT8 NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

COMMIT;