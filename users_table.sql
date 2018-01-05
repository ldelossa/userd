-- Create users table
CREATE TABLE users ( u jsonb NOT NULL);

-- Default GIN index over entire document
CREATE INDEX idxginusers ON users USING GIN (u);

-- Force username to be unique
CREATE UNIQUE INDEX users_username ON users( (u->>'username') );

-- Force username to be non-null
ALTER TABLE users ADD CONSTRAINT username_must_exist CHECK (u ? 'username');

-- Force username not to be empty string
ALTER TABLE users ADD CONSTRAINT username_must_not_be_empty_string CHECK ((u->>'username') != '');

--Force password to be non-null
ALTER TABLE users ADD CONSTRAINT username_must_exist CHECK (u ? 'password');

-- Force password not to be empty string
ALTER TABLE users ADD CONSTRAINT password_must_not_be_empty_string CHECK ((u->>'password') != '');
