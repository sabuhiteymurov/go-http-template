ALTER TABLE auth.users
    add column soft_deleted BOOLEAN DEFAULT FALSE;