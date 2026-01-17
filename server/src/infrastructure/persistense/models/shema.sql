CREATE TABLE IF NOT EXISTS users (
    id text NOT NULL PRIMARY KEY,

    registered_at TIMESTAMP not null,
    user_name VARCHAR(32) unique NOT NULL,
    password_hash VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS roles (
    role_name text PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS user_roles (
    user_id text REFERENCES users (id),
    role_id text REFERENCES roles (role_name),
    
    created_at TIMESTAMP not null,

    PRIMARY KEY (user_id, role_id)
);


CREATE TABLE IF NOT EXISTS courses (
    id text PRIMARY KEY,

    created_by_id text REFERENCES users (id),

    name text not NULL,
    is_archived boolean not null
);

create table if not EXISTS topics (
    id text PRIMARY key,

    by_course_id text REFERENCES courses (id),
    created_by_id text REFERENCES users (id),

    name text not null,
    is_archived boolean not null
);

create table if not EXISTS questions (
    id text primary key,

    by_topic_id text references topics (id),
    by_course_id text references courses (id),

    text text not null
);

create table if not exists answers (
    by_question_id text references questions (id),

    text text not null,
    is_correct boolean not null,
    serial_number integer not null

    primary key (by_question_id, serial_number)
);

-- create table if not exists groups (
--     id text primary key,

--     by_course_id text references courses (id),
--     created_by_id text references users (id),
--     name text not null,
--     student_count integer default 0,
--     max_student_count integer
-- );

-- create table if not exists enrollments (
--     id text primary key,

--     group_id text references groups (id),
--     course_id text references courses (id),
--     user_id text references users (id), 
--     progress float default 0
-- );

-- create table if not exists topic_progresses (
--     id bigserial primary key,

--     enrollment_id text references enrollments (id),
--     maxScore float
-- );

-- create table if not exists topic_attempts (
--     id text primary key,

--     topic_progress_id bigserial references topic_progresses (id),
--     score float,
--     attempted_at timestamp default now()
-- );
