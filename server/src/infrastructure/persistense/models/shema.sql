CREATE TABLE IF NOT EXISTS users (
    id uuid NOT NULL PRIMARY KEY,

    registered_at TIMESTAMP DEFAULT NOW (),
    user_name VARCHAR(32) unique NOT NULL CHECK (LENGTH (user_name) >= 5),
    password_hash VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS roles (
    role_name text PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS user_roles (
    created_at TIMESTAMP not null,
    user_id text REFERENCES users (id),
    role_id text REFERENCES roles (role_name),
    PRIMARY KEY (user_id, role_id)
);


CREATE TABLE IF NOT EXISTS courses (
    id uuid PRIMARY KEY,

    created_by_id text REFERENCES users (id),
    name text not NULL,
    is_archived boolean not null
);

create table if not EXISTS topics (
    id uuid PRIMARY key,

    by_course_id text REFERENCES courses (id),
    created_by_id text REFERENCES users (id),
    name text not null,
    is_archived boolean default false,
    question_count integer default 0
    question_ids text[]
);

create table if not EXISTS questions (
    id uuid primary key,

    by_topic_id text references topics (id),
    text text not null
);

create table if not exists answers (
    id bigserial primary key, 

    by_question_id text references questions (id),
    text text not null,
    is_correct boolean default false,
    serial_number integer,

    unique(text, by_question_id)
);

create table if not exists groups (
    id text primary key,

    by_course_id text references courses (id),
    created_by_id text references users (id),
    name text not null,
    student_count integer default 0,
    max_student_count integer
);

create table if not exists enrollments (
    id UUID primary key,

    group_id text references groups (id),
    course_id text references courses (id),
    user_id text references users (id), 
    progress float default 0
);

create table if not exists topic_progresses (
    id bigserial primary key,

    enrollment_id text references enrollments (id),
    maxScore float
);

create table if not exists topic_attempts (
    id UUID primary key,

    topic_progress_id bigserial references topic_progresses (id),
    score float,
    attempted_at timestamp default now()
);
