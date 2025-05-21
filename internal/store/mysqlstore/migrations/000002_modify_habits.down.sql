ALTER TABLE habits
    DROP COLUMN category,
    DROP COLUMN description,
    DROP COLUMN frequency,
    DROP COLUMN start_date,
    DROP COLUMN target_date,
    DROP COLUMN priority,
    DROP COLUMN updated_at,
    MODIFY COLUMN name VARCHAR(255) NOT NULL;