-- name: CreateUser :one
INSERT INTO
    Users (
        id,
        username,
        email,
        role_id,
        last_login,
        is_active
    )
VALUES
    (
        -- autoincrement id starting from 1
        -- (
        --     SELECT
        --         COALESCE(MAX(id), 0) + 1
        --     FROM
        --         Users
        -- ),
        @id,
        @username,
        @email,
        (
            SELECT
                id
            FROM
                UserRoles
            WHERE
                role_name = @role
        ),
        CURRENT_TIMESTAMP,
        TRUE
    )
RETURNING
    id;

-- name: SeedRoles :exec
INSERT INTO
    UserRoles (id, role_name, description)
VALUES
    (
        1,
        'admin',
        'System administrator with full access'
    ),
    (
        2,
        'quality_manager',
        'Manages quality processes'
    ),
    (
        3,
        'auditor',
        'Performs internal and external audits'
    ),
    (4, 'staff', 'Standard user with limited access');

-- name: GetAllUsers :many
SELECT
    u.id,
    u.username,
    u.email,
    r.role_name AS role,
    u.last_login,
    u.is_active
FROM
    Users u
    JOIN UserRoles r ON u.role_id = r.id;
