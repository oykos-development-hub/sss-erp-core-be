ALTER TABLE users
    DROP COLUMN IF EXISTS secondary_email,
    DROP COLUMN IF EXISTS phone,
    DROP COLUMN IF EXISTS pin,
    DROP COLUMN IF EXISTS active,
    DROP COLUMN IF EXISTS verified_email,
    DROP COLUMN IF EXISTS verified_phone,
    DROP COLUMN IF EXISTS folder_id;