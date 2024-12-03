-- +goose Up
-- +goose StatementBegin
SELECT
    'up SQL query';

-- User Management
CREATE TABLE UserRoles (
    id VARCHAR(20) PRIMARY KEY,
    role_name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT
);

CREATE TABLE Users (
    id VARCHAR(20) PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    role_id VARCHAR(20) NOT NULL,
    last_login TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (role_id) REFERENCES UserRoles(id)
);

CREATE TABLE UserPermissions (
    role_id VARCHAR(20),
    permission_code VARCHAR(100),
    PRIMARY KEY (role_id, permission_code),
    FOREIGN KEY (role_id) REFERENCES UserRoles(id)
);

-- Document Control
CREATE TABLE Documents (
    id VARCHAR(20) PRIMARY KEY,
    document_type VARCHAR(100),
    title VARCHAR(255) NOT NULL,
    version VARCHAR(20) NOT NULL,
    created_by VARCHAR(20),
    created_date TIMESTAMP,
    last_reviewed_date TIMESTAMP,
    next_review_date TIMESTAMP,
    STATUS VARCHAR(50),
    content TEXT,
    FOREIGN KEY (created_by) REFERENCES Users(id)
);

CREATE TABLE DocumentAuditTrail (
    id VARCHAR(20) PRIMARY KEY,
    document_id VARCHAR(20),
    action_type VARCHAR(50),
    action_by VARCHAR(20),
    action_date TIMESTAMP,
    previous_version VARCHAR(20),
    FOREIGN KEY (document_id) REFERENCES Documents(id),
    FOREIGN KEY (action_by) REFERENCES Users(id)
);

-- Risk Management
CREATE TABLE RiskRegister (
    id VARCHAR(20) PRIMARY KEY,
    risk_name VARCHAR(255),
    description TEXT,
    risk_level VARCHAR(50),
    likelihood DECIMAL(5, 2),
    impact DECIMAL(5, 2),
    risk_score DECIMAL(5, 2),
    mitigation_strategy TEXT,
    current_status VARCHAR(100),
    owner_id VARCHAR(20),
    FOREIGN KEY (owner_id) REFERENCES Users(id)
);

-- Design and Development
CREATE TABLE DesignProjects (
    id INT PRIMARY KEY,
    project_name VARCHAR(255),
    start_date DATE,
    expected_completion_date DATE,
    STATUS VARCHAR(50)
);

CREATE TABLE DesignRequirements (
    id INT PRIMARY KEY,
    project_id INT,
    requirement_description TEXT,
    requirement_type VARCHAR(100),
    FOREIGN KEY (project_id) REFERENCES DesignProjects(id)
);

CREATE TABLE DesignVerification (
    id INT PRIMARY KEY,
    requirement_id INT,
    verification_method VARCHAR(100),
    verification_result VARCHAR(50),
    verified_by VARCHAR(20),
    verification_date TIMESTAMP,
    FOREIGN KEY (requirement_id) REFERENCES DesignRequirements(id),
    FOREIGN KEY (verified_by) REFERENCES Users(id)
);

-- Training Management
CREATE TABLE TrainingRecords (
    id VARCHAR(20) PRIMARY KEY,
    user_id VARCHAR(20),
    training_name VARCHAR(255),
    completion_date DATE,
    expiry_date DATE,
    certificate_number VARCHAR(100),
    FOREIGN KEY (user_id) REFERENCES Users(id)
);

CREATE TABLE TrainingMatrix (
    role_id VARCHAR(20),
    training_id VARCHAR(20),
    is_required BOOLEAN,
    PRIMARY KEY (role_id, training_id),
    FOREIGN KEY (role_id) REFERENCES UserRoles(id)
);

-- Supplier Management
CREATE TABLE Suppliers (
    id VARCHAR(20) PRIMARY KEY,
    supplier_name VARCHAR(255),
    contact_person VARCHAR(100),
    email VARCHAR(255),
    phone VARCHAR(50),
    qualification_status VARCHAR(50),
    last_evaluated_date DATE
);

CREATE TABLE SupplierPerformance (
    id VARCHAR(20) PRIMARY KEY,
    supplier_id VARCHAR(20),
    evaluation_date DATE,
    overall_score DECIMAL(5, 2),
    quality_rating DECIMAL(5, 2),
    delivery_rating DECIMAL(5, 2),
    comments TEXT,
    FOREIGN KEY (supplier_id) REFERENCES Suppliers(id)
);

-- Nonconformance and CAPA
CREATE TABLE Nonconformances (
    id INT PRIMARY KEY,
    description TEXT,
    detected_date DATE,
    detected_by VARCHAR(20),
    product_id INT,
    severity VARCHAR(50),
    STATUS VARCHAR(50),
    FOREIGN KEY (detected_by) REFERENCES Users(id)
);

CREATE TABLE CAPA (
    id INT PRIMARY KEY,
    nonconformance_id INT,
    root_cause TEXT,
    corrective_action TEXT,
    preventive_action TEXT,
    proposed_date DATE,
    completion_date DATE,
    effectiveness_check TEXT,
    STATUS VARCHAR(50),
    FOREIGN KEY (nonconformance_id) REFERENCES Nonconformances(id)
);

-- Complaint Management
CREATE TABLE Complaints (
    id INT PRIMARY KEY,
    customer_name VARCHAR(255),
    product_id INT,
    batch_lot_number VARCHAR(100),
    complaint_date DATE,
    description TEXT,
    STATUS VARCHAR(50)
);

CREATE TABLE ComplaintInvestigation (
    id INT PRIMARY KEY,
    complaint_id INT,
    root_cause TEXT,
    corrective_action TEXT,
    investigator_id VARCHAR(20),
    investigation_date DATE,
    FOREIGN KEY (complaint_id) REFERENCES Complaints(id),
    FOREIGN KEY (investigator_id) REFERENCES Users(id)
);

-- Audit Management
CREATE TABLE Audits (
    id INT PRIMARY KEY,
    audit_type VARCHAR(100),
    scheduled_date DATE,
    actual_date DATE,
    auditor_id VARCHAR(20),
    STATUS VARCHAR(50),
    FOREIGN KEY (auditor_id) REFERENCES Users(id)
);

CREATE TABLE AuditFindings (
    id INT PRIMARY KEY,
    audit_id INT,
    description TEXT,
    severity VARCHAR(50),
    corrective_action TEXT,
    due_date DATE,
    STATUS VARCHAR(50),
    FOREIGN KEY (audit_id) REFERENCES Audits(id)
);

-- Post-Market Surveillance
CREATE TABLE AdverseEvents (
    id INT PRIMARY KEY,
    product_id INT,
    event_date DATE,
    description TEXT,
    severity VARCHAR(50),
    regulatory_report_status VARCHAR(50)
);

-- Security and Audit Logging
CREATE TABLE SystemAccessLog (
    id INT PRIMARY KEY,
    user_id VARCHAR(20),
    action_type VARCHAR(100),
    timestamp TIMESTAMP,
    ip_address VARCHAR(50),
    FOREIGN KEY (user_id) REFERENCES Users(id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT
    'down SQL query';

DROP TABLE SystemAccessLog CASCADE;

DROP TABLE AdverseEvents CASCADE;

DROP TABLE AuditFindings CASCADE;

DROP TABLE Audits CASCADE;

DROP TABLE ComplaintInvestigation CASCADE;

DROP TABLE Complaints CASCADE;

DROP TABLE CAPA CASCADE;

DROP TABLE Nonconformances CASCADE;

DROP TABLE SupplierPerformance CASCADE;

DROP TABLE Suppliers CASCADE;

DROP TABLE TrainingMatrix CASCADE;

DROP TABLE TrainingRecords CASCADE;

DROP TABLE DesignVerification CASCADE;

DROP TABLE DesignRequirements CASCADE;

DROP TABLE DesignProjects CASCADE;

DROP TABLE RiskRegister CASCADE;

DROP TABLE DocumentAuditTrail CASCADE;

DROP TABLE Documents CASCADE;

DROP TABLE UserPermissions CASCADE;

DROP TABLE UserRoles CASCADE;

DROP TABLE Users CASCADE;

-- +goose StatementEnd
