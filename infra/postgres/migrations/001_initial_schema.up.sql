
-- 1. Organizations
CREATE TABLE organizations (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 2. Cloud Accounts
CREATE TABLE cloud_accounts (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    organization_id BIGINT NOT NULL,
    provider VARCHAR(50) NOT NULL,
    external_account_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_cloud_account_org
        FOREIGN KEY (organization_id)
        REFERENCES organizations(id),

    CONSTRAINT uq_cloud_account_external
        UNIQUE (organization_id, provider, external_account_id),

    CONSTRAINT uq_cloud_account_org_id
        UNIQUE (id, organization_id)
);

-- 3. Cost Records
CREATE TABLE cost_records (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    organization_id BIGINT NOT NULL,
    cloud_account_id BIGINT NOT NULL,
    resource_id VARCHAR(255) NOT NULL,
    usage_date DATE NOT NULL,
    cost_cents BIGINT NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_cost_record_org
        FOREIGN KEY (organization_id)
        REFERENCES organizations(id),

    CONSTRAINT fk_cost_record_account_org
        FOREIGN KEY (cloud_account_id, organization_id)
        REFERENCES cloud_accounts(id, organization_id),

    CONSTRAINT chk_cost_currency
        CHECK (currency ~ '^[A-Z]{3}$')
);

-- 4. Index for tenant-specific date queries
CREATE INDEX idx_cost_records_org_date
    ON cost_records (organization_id, usage_date);
