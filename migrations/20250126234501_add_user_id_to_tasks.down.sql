ALTER TABLE tasks 
    DROP CONSTRAINT IF EXISTS fk_tasks_user,
    DROP COLUMN IF EXISTS user_id; 