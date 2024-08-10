START TRANSACTION;

CREATE TABLE IF NOT EXISTS users (
    user_id VARCHAR(36) PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    user_name VARCHAR(8) NOT NULL,
    name VARCHAR(50) NOT NULL,
    user_type VARCHAR(10) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS baselines (
    baseline_id VARCHAR(36) PRIMARY KEY,
    code VARCHAR(20) NOT NULL,
    review INT NOT NULL,
    title VARCHAR(50) NOT NULL,
    description TEXT NULL,
    start_date DATE NOT NULL,
    duration INT NOT NULL,
    manager_id VARCHAR(36) NOT NULL REFERENCES users (user_id),
    estimator_id VARCHAR(36) NOT NULL REFERENCES users (user_id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    UNIQUE (code, review)
);

CREATE TABLE IF NOT EXISTS costs (
    cost_id VARCHAR(36) PRIMARY KEY,
    baseline_id VARCHAR(36) NOT NULL REFERENCES baselines (baseline_id),
    cost_type VARCHAR(10) NOT NULL,
    description VARCHAR(255) NOT NULL,
    comment VARCHAR(255),
    amount FLOAT8 NOT NULL,
    currency VARCHAR(10) NOT NULL,
    tax FLOAT8 NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    UNIQUE (
        baseline_id,
        cost_type,
        description
    )
);

CREATE TABLE IF NOT EXISTS cost_allocations (
    cost_allocation_id VARCHAR(36) PRIMARY KEY,
    cost_id VARCHAR(36) NOT NULL REFERENCES costs (cost_id),
    allocation_date DATE NOT NULL,
    amount FLOAT8 NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS competences (
    competence_id VARCHAR(36) PRIMARY KEY,
    code VARCHAR(20) NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS efforts (
    effort_id VARCHAR(36) PRIMARY KEY,
    baseline_id VARCHAR(36) NOT NULL REFERENCES baselines (baseline_id),
    competence_id VARCHAR(36) NOT NULL REFERENCES competences (competence_id),
    comment VARCHAR(255),
    hours INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    UNIQUE (baseline_id, competence_id)
);

CREATE TABLE IF NOT EXISTS effort_allocations (
    effort_allocation_id VARCHAR(36) PRIMARY KEY,
    effort_id VARCHAR(36) NOT NULL REFERENCES efforts (effort_id),
    allocation_date DATE NOT NULL,
    hours INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS plans (
    plan_id VARCHAR(36) PRIMARY KEY,
    code VARCHAR(10) NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL,
    assumptions JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS portfolios (
    portfolio_id VARCHAR(36) PRIMARY KEY,
    baseline_id VARCHAR(36) NOT NULL REFERENCES baselines (baseline_id),
    plan_id VARCHAR(36) NOT NULL REFERENCES plans (plan_id),
    start_date DATE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    UNIQUE (baseline_id, plan_id)
);

CREATE TABLE IF NOT EXISTS budgets (
    budget_id VARCHAR(36) PRIMARY KEY,
    portfolio_id VARCHAR(36) NOT NULL REFERENCES portfolios (portfolio_id),
    cost_id VARCHAR(36) NOT NULL REFERENCES costs (cost_id),
    amount FLOAT8 NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS budget_allocations (
    budget_allocation_id VARCHAR(36) PRIMARY KEY,
    budget_id VARCHAR(36) NOT NULL REFERENCES budgets (budget_id),
    allocation_date DATE NOT NULL,
    amount FLOAT8 NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS workloads (
    workload_id VARCHAR(36) PRIMARY KEY,
    portfolio_id VARCHAR(36) NOT NULL REFERENCES portfolios (portfolio_id),
    effort_id VARCHAR(36) NOT NULL REFERENCES efforts (effort_id),
    hours INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS workload_allocations (
    workload_allocation_id VARCHAR(36) PRIMARY KEY,
    workload_id VARCHAR(36) NOT NULL REFERENCES workloads (workload_id),
    allocation_date DATE NOT NULL,
    hours INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

COMMIT;