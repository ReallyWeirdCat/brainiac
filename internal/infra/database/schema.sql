-- ============================================================
-- DOMAIN TYPES
-- ============================================================

-- question_type_enum: 0=short_question, 1=normal_question, 2=long_question
CREATE DOMAIN question_type_enum AS SMALLINT
    CHECK (VALUE IN (0, 1, 2));

-- scoring_method_enum: 0=max_score, 1=average_score, 2=latest_score
CREATE DOMAIN scoring_method_enum AS SMALLINT
    CHECK (VALUE IN (0, 1, 2));

-- completion_condition_enum: 0=manual_completion, 1=attestation_test, 2=approved_practice, 3=all
CREATE DOMAIN completion_condition_enum AS SMALLINT
    CHECK (VALUE IN (0, 1, 2, 3));

-- subject_status_enum: 0=untouched, 1=in_progress, 2=forgotten, 3=done
CREATE DOMAIN subject_status_enum AS SMALLINT
    CHECK (VALUE IN (0, 1, 2, 3));

-- grade_enum: 0=unspecified, 1=very_poor, 2=poor, 3=satisfactory, 4=good, 5=very_good
CREATE DOMAIN grade_enum AS SMALLINT
    CHECK (VALUE IN (0, 1, 2, 3, 4, 5));

-- subject_doc_type_enum: 0=theory, 1=practice, 2=example, 3=literature
CREATE DOMAIN subject_doc_type_enum AS SMALLINT
    CHECK (VALUE IN (0, 1, 2, 3));

-- attempt_status_enum: 0=in_progress, 1=finished, 2=flagged
CREATE DOMAIN attempt_status_enum AS SMALLINT
    CHECK (VALUE IN (0, 1, 2));

-- class_type_enum: 0=practice, 1=lecture, 2=attestation, 3=consult
CREATE DOMAIN class_type_enum AS SMALLINT
    CHECK (VALUE IN (0, 1, 2, 3));

-- chat_role_enum: 0=member, 1=admin, 2=owner
CREATE DOMAIN chat_role_enum AS SMALLINT
    CHECK (VALUE IN (0, 1, 2));

-- urgency_enum: 0=low, 1=normal, 2=important
CREATE DOMAIN urgency_enum AS SMALLINT
    CHECK (VALUE IN (0, 1, 2));

-- rarity_enum: 0=common, 1=uncommon, 2=rare, 3=epic, 4=legendary, 5=mythical
CREATE DOMAIN rarity_enum AS SMALLINT
    CHECK (VALUE IN (0, 1, 2, 3, 4, 5));

-- profile_discovery_enum: 0=admins_only, 1=teachers, 2=group_mates, 3=course_mates, 4=users, 5=guests
CREATE DOMAIN profile_discovery_enum AS SMALLINT
    CHECK (VALUE IN (0, 1, 2, 3, 4, 5));

-- student_role_enum: 0=unspecified, 1=student, 2=basic, 3=intermediate
CREATE DOMAIN student_role_enum AS SMALLINT
    CHECK (VALUE IN (0, 1, 2, 3));

-- ============================================================
-- TABLES (ordered by dependencies)
-- ============================================================

-- 1. app_user
CREATE TABLE app_user (
    guid UUID PRIMARY KEY,
    username VARCHAR(18) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

CREATE UNIQUE INDEX idx_appuser_username_not_deleted ON app_user (username) WHERE deleted_at IS NULL;

-- 2. app_user_credential
CREATE TABLE app_user_credential (
    app_user_guid UUID PRIMARY KEY REFERENCES app_user(guid),
    email VARCHAR(255) NULL,
    password_hash VARCHAR(60) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

CREATE UNIQUE INDEX idx_appusercredential_email_not_deleted ON app_user_credential (email) WHERE deleted_at IS NULL;

-- 3. app_user_profile
CREATE TABLE app_user_profile (
    app_user_guid UUID PRIMARY KEY REFERENCES app_user(guid),
    name VARCHAR(100) NULL,
    surname VARCHAR(100) NULL,
    patronymic VARCHAR(100) NULL,
    nickname VARCHAR(32) NULL,
    bio VARCHAR(175) NULL,
    preferred_language VARCHAR(2) NOT NULL DEFAULT 'en', 
    profile_discovery profile_discovery_enum NOT NULL DEFAULT 2,
    avatar_url TEXT NULL,
    editing_locked_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 4. app_user_suspension
CREATE TABLE app_user_suspension (
    guid UUID PRIMARY KEY,
    app_user_guid UUID NOT NULL REFERENCES app_user(guid),
    reason TEXT NOT NULL,
    expire_at TIMESTAMPTZ NULL,
    expired_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 5. app_user_session
CREATE TABLE app_user_session (
    guid UUID PRIMARY KEY,
    app_user_guid UUID NOT NULL REFERENCES app_user(guid),
    last_ipv4 VARCHAR(32) NULL,
    last_ipv6 VARCHAR(39) NULL,
    last_agent VARCHAR(255) NULL,
    last_seen_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    expire_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 6. totp_secret
CREATE TABLE totp_secret (
    app_user_guid UUID PRIMARY KEY REFERENCES app_user(guid),
    secret BYTEA NOT NULL,
    secret_base32 VARCHAR(64) NOT NULL,
    last_used_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 7. teacher
CREATE TABLE teacher (
    app_user_guid UUID PRIMARY KEY REFERENCES app_user(guid),
    active_since TIMESTAMPTZ NULL,
    manage_students BOOLEAN NOT NULL DEFAULT false,
    manage_attendance BOOLEAN NOT NULL DEFAULT false,
    manage_results BOOLEAN NOT NULL DEFAULT false,
    manage_courses BOOLEAN NOT NULL DEFAULT false,
    is_admin BOOLEAN NOT NULL DEFAULT false,
    contact_info JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 8. student_global_stats
CREATE TABLE student_global_stats (
    app_user_guid UUID PRIMARY KEY REFERENCES app_user(guid),
    level SMALLINT NOT NULL DEFAULT 0,
    other_experience BIGINT NOT NULL DEFAULT 0,
    money BIGINT NOT NULL DEFAULT 0,
    money_earned BIGINT NOT NULL DEFAULT 0,
    subjects_completed INT NOT NULL DEFAULT 0,
    tests_completed INT NOT NULL DEFAULT 0,
    tests_failed INT NOT NULL DEFAULT 0,
    tests_terminated INT NOT NULL DEFAULT 0,
    works_submitted INT NOT NULL DEFAULT 0,
    items_collected INT NOT NULL DEFAULT 0,
    items_used INT NOT NULL DEFAULT 0,
    items_sold INT NOT NULL DEFAULT 0,
    items_exchanged INT NOT NULL DEFAULT 0,
    boxes_opened INT NOT NULL DEFAULT 0,
    max_daily_streak INT NOT NULL DEFAULT 0,
    messages_sent INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

CREATE INDEX idx_studentglobalstats_level_not_deleted ON student_global_stats (level) WHERE deleted_at IS NULL;

-- 9. inventory
CREATE TABLE inventory (
    app_user_guid UUID PRIMARY KEY REFERENCES app_user(guid),
    max_slots SMALLINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 10. language
CREATE TABLE language (
    guid UUID PRIMARY KEY,
    language_code VARCHAR(2) UNIQUE NOT NULL,
    english_title VARCHAR(25) NOT NULL,
    local_title VARCHAR(50) NOT NULL,
    emoji TEXT NULL,
    supported BOOLEAN NOT NULL,
    beta BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

CREATE INDEX idx_language_language_code_not_deleted ON language (language_code) WHERE deleted_at IS NULL;

-- 11. achievement
CREATE TABLE achievement (
    guid UUID PRIMARY KEY,
    title_i18n JSONB NOT NULL,
    description_i18n JSONB NULL,
    resource_url VARCHAR(255) NULL,
    is_hidden BOOLEAN NOT NULL DEFAULT false,
    meta JSONB NULL,
    conditions JSONB NULL,
    reward_experience BIGINT NOT NULL DEFAULT 0,
    reward_money BIGINT NOT NULL DEFAULT 0,
    published_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 12. course
CREATE TABLE course (
    guid UUID PRIMARY KEY,
    title_i18n JSONB NULL,
    description_i18n JSONB NULL,
    style JSONB NULL,
    meta JSONB NULL,
    published_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 13. course_language
CREATE TABLE course_language (
    guid UUID PRIMARY KEY,
    course_guid UUID NOT NULL REFERENCES course(guid),
    language_guid UUID NOT NULL REFERENCES language(guid),
    calculated_support_percentage SMALLINT NULL,
    beta BOOLEAN NOT NULL DEFAULT false,
    new BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL,
    CONSTRAINT check_calculated_support_percentage CHECK (calculated_support_percentage BETWEEN 0 AND 100)
);

-- 14. student_group
CREATE TABLE student_group (
    guid UUID PRIMARY KEY,
    title VARCHAR(30) NOT NULL UNIQUE,
    description VARCHAR(175) NULL,
    enable_gamification BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 15. chat
CREATE TABLE chat (
    guid UUID PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    topic VARCHAR(255) NULL,
    avatar_url VARCHAR(255) NULL,
    students_can_text BOOLEAN NOT NULL DEFAULT true,
    students_see_members BOOLEAN NOT NULL DEFAULT true,
    preserve_messages BOOLEAN NOT NULL DEFAULT false,
    meta JSONB NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 16. teacher_group
CREATE TABLE teacher_group (
    guid UUID PRIMARY KEY,
    teacher_guid UUID NOT NULL REFERENCES teacher(guid),
    student_group_guid UUID NOT NULL REFERENCES student_group(guid),
    active_since TIMESTAMPTZ NULL,
    manage_attendance BOOLEAN NOT NULL DEFAULT true,
    manage_results BOOLEAN NOT NULL DEFAULT true,
    manage_students BOOLEAN NOT NULL DEFAULT false,
    manage_achievements BOOLEAN NOT NULL DEFAULT false,
    teacher_until TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

CREATE UNIQUE INDEX idx_teachergroup_unique_active ON teacher_group (teacher_guid, student_group_guid) WHERE deleted_at IS NULL;

-- 17. student_group_member
CREATE TABLE student_group_member (
    guid UUID PRIMARY KEY,
    app_user_guid UUID NOT NULL REFERENCES app_user(guid),
    student_group_guid UUID NOT NULL REFERENCES student_group(guid),
    student_role student_role_enum NOT NULL DEFAULT 1,
    notes TEXT NULL,
    team VARCHAR(18) NULL,
    team_responsibilites JSONB NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 18. student_group_course
CREATE TABLE student_group_course (
    guid UUID PRIMARY KEY,
    student_group_guid UUID NOT NULL REFERENCES student_group(guid),
    course_guid UUID NOT NULL REFERENCES course(guid),
    full_access BOOLEAN NOT NULL DEFAULT false,
    close_submissions BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 19. student_group_class
CREATE TABLE student_group_class (
    guid UUID PRIMARY KEY,
    student_group_guid UUID NOT NULL REFERENCES student_group(guid),
    title VARCHAR(50) NOT NULL,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    objectives JSONB NULL,
    attestation_tests_enabled BOOLEAN NOT NULL DEFAULT true,
    attendance_closed_at TIMESTAMPTZ NULL,
    marks_approved_at TIMESTAMPTZ NULL,
    totp_secret BYTEA NULL,
    room VARCHAR(15) NULL,
    class_type class_type_enum NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 20. student_group_member_class_attendance
CREATE TABLE student_group_member_class_attendance (
    guid UUID PRIMARY KEY,
    student_group_member_guid UUID NOT NULL REFERENCES student_group_member(guid),
    student_group_class_guid UUID NOT NULL REFERENCES student_group_class(guid),
    first_seen_at TIMESTAMPTZ NULL,
    second_seen_at TIMESTAMPTZ NULL,
    present_at TIMESTAMPTZ NULL,
    approved_at TIMESTAMPTZ NULL,
    comment VARCHAR(1024) NULL,
    grade grade_enum NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 21. course_subject
CREATE TABLE course_subject (
    guid UUID PRIMARY KEY,
    course_guid UUID NOT NULL REFERENCES course(guid),
    parent UUID NULL REFERENCES course_subject(guid),
    title_i18n JSONB NULL,
    description_i18n JSONB NULL,
    reward_experience BIGINT NOT NULL DEFAULT 0,
    completion_condition completion_condition_enum NOT NULL DEFAULT 0,
    hide_children BOOLEAN NOT NULL DEFAULT false,
    style JSONB NULL,
    meta JSONB NULL,
    published_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

CREATE INDEX idx_coursesubject_parent_not_deleted ON course_subject (parent) WHERE deleted_at IS NULL;

-- 22. course_subject_doc
CREATE TABLE course_subject_doc (
    guid UUID PRIMARY KEY,
    course_subject_guid UUID NOT NULL REFERENCES course_subject(guid),
    title_i18n JSONB NULL,
    description_i18n JSONB NULL,
    subject_doc_type subject_doc_type_enum NOT NULL DEFAULT 0,
    example_meta JSONB NULL,
    submisson_meta JSONB NULL,
    literature_meta JSONB NULL,
    doc_index SMALLINT NOT NULL DEFAULT 0,
    published_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 23. course_subject_doc_md
CREATE TABLE course_subject_doc_md (
    guid UUID PRIMARY KEY,
    course_subject_doc_guid UUID NOT NULL REFERENCES course_subject_doc(guid),
    language_code VARCHAR(2) NOT NULL,
    md BYTEA,
    published_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);


-- 24. course_subject_test
CREATE TABLE course_subject_test (
    guid UUID PRIMARY KEY,
    course_subject_guid UUID NOT NULL REFERENCES course_subject(guid),
    title_i18n JSONB NULL,
    description_i18n JSONB NULL,
    is_attestation BOOLEAN NOT NULL DEFAULT false,
    hide_answers BOOLEAN NOT NULL DEFAULT false,
    hide_answer_validity BOOLEAN NOT NULL DEFAULT false,
    hide_score BOOLEAN NOT NULL DEFAULT false,
    use_timed_questions BOOLEAN NOT NULL DEFAULT false,
    max_time SMALLINT NULL,
    max_attempts SMALLINT NULL,
    scoring_method scoring_method_enum NOT NULL DEFAULT 0,
    required_score SMALLINT NOT NULL DEFAULT 60,
    short_questions_count SMALLINT NOT NULL DEFAULT 1,
    normal_questions_count SMALLINT NOT NULL DEFAULT 0,
    long_questions_count SMALLINT NOT NULL DEFAULT 0,
    published_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL,
    CONSTRAINT check_required_score CHECK (required_score BETWEEN 0 AND 100)
);

-- 25. course_subject_test_question
CREATE TABLE course_subject_test_question (
    guid UUID PRIMARY KEY,
    course_subject_test_guid UUID NOT NULL REFERENCES course_subject_test(guid),
    title_i18n JSONB NULL,
    description_i18n JSONB NULL,
    question_type question_type_enum NOT NULL DEFAULT 0,
    attachment_url VARCHAR(255) NULL,
    options JSONB NULL,
    correct_options JSONB NULL,
    is_multiple_choice BOOLEAN NOT NULL DEFAULT false,
    max_options SMALLINT NOT NULL DEFAULT 0,
    use_text_answer BOOLEAN NOT NULL DEFAULT false,
    correct_text_answer TEXT NULL,
    example_url VARCHAR(255) NULL,
    published_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 26. course_subject_test_pass
CREATE TABLE course_subject_test_pass (
    guid UUID PRIMARY KEY,
    course_subject_test_guid UUID NOT NULL REFERENCES course_subject_test(guid),
    app_user_guid UUID NOT NULL REFERENCES app_user(guid),
    is_active BOOLEAN NOT NULL DEFAULT true,
    attempts_left SMALLINT NULL,
    expire_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 27. student_course_stats
CREATE TABLE student_course_stats (
    guid UUID PRIMARY KEY,
    app_user_guid UUID NOT NULL REFERENCES app_user(guid),
    course_guid UUID NOT NULL REFERENCES course(guid),
    experience BIGINT NOT NULL DEFAULT 0,
    level SMALLINT NOT NULL DEFAULT 0,
    meta JSONB NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

CREATE UNIQUE INDEX idx_studentcoursestats_unique ON student_course_stats (app_user_guid, course_guid) WHERE deleted_at IS NULL;
CREATE INDEX idx_studentcoursestats_experience ON student_course_stats (experience) WHERE deleted_at IS NULL;
CREATE INDEX idx_studentcoursestats_level ON student_course_stats (level) WHERE deleted_at IS NULL;

-- 28. student_course_subject
CREATE TABLE student_course_subject (
    guid UUID PRIMARY KEY,
    app_user_guid UUID NOT NULL REFERENCES app_user(guid),
    course_subject_guid UUID NOT NULL REFERENCES course_subject(guid),
    completed_at TIMESTAMPTZ NULL,
    subject_status subject_status_enum NOT NULL DEFAULT 0,
    is_favorite BOOLEAN NOT NULL DEFAULT false,
    meta JSONB NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

CREATE UNIQUE INDEX idx_studentcoursesubject_unique ON student_course_subject (app_user_guid, course_subject_guid) WHERE deleted_at IS NULL;

-- 29. student_course_subject_test_attempt
CREATE TABLE student_course_subject_test_attempt (
    guid UUID PRIMARY KEY,
    app_user_guid UUID NOT NULL REFERENCES app_user(guid),
    course_subject_test_guid UUID NOT NULL REFERENCES course_subject_test(guid),
    attempt_status attempt_status_enum NOT NULL DEFAULT 0,
    score SMALLINT NOT NULL DEFAULT 0,
    is_success BOOLEAN NOT NULL DEFAULT false,
    answers JSONB NOT NULL DEFAULT '[]'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

CREATE UNIQUE INDEX idx_studentcoursetestattempt_unique ON student_course_subject_test_attempt (app_user_guid, course_subject_test_guid) WHERE deleted_at IS NULL;

-- 30. student_course_subject_submission_attempt
CREATE TABLE student_course_subject_submission_attempt (
    guid UUID PRIMARY KEY,
    app_user_guid UUID NOT NULL REFERENCES app_user(guid),
    course_subject_doc_guid UUID NOT NULL REFERENCES course_subject_doc(guid),
    submission_status subject_doc_type_enum NOT NULL DEFAULT 0,
    files JSONB NULL,
    meta JSONB NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

CREATE UNIQUE INDEX idx_studentcoursesubmissionattempt_unique ON student_course_subject_submission_attempt (app_user_guid, course_subject_doc_guid) WHERE deleted_at IS NULL;

-- 31. item
CREATE TABLE item (
    guid UUID PRIMARY KEY,
    title_i18n JSONB NOT NULL,
    description_i18n JSONB NULL,
    resource_url VARCHAR(255) NULL,
    meta JSONB NOT NULL DEFAULT '{}'::jsonb,
    stack_size SMALLINT NOT NULL DEFAULT 1,
    rarity rarity_enum NOT NULL DEFAULT 0,
    allow_exchange BOOLEAN NOT NULL DEFAULT false,
    shop_price BIGINT NULL,
    one_time_purchase BOOLEAN NOT NULL DEFAULT false,
    shop_quantity SMALLINT NULL,
    level_required SMALLINT NULL,
    on_sale_since TIMESTAMPTZ NULL,
    on_sale_until TIMESTAMPTZ NULL,
    published_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 32. inventory_slot
CREATE TABLE inventory_slot (
    guid UUID PRIMARY KEY,
    inventory_guid UUID NOT NULL REFERENCES inventory(guid),
    item_guid UUID NOT NULL REFERENCES item(guid),
    amount SMALLINT NOT NULL DEFAULT 1,
    amount_in_exchange SMALLINT NOT NULL DEFAULT 0,
    expire_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 33. course_subject_item
CREATE TABLE course_subject_item (
    guid UUID PRIMARY KEY,
    course_subject_guid UUID NOT NULL REFERENCES course_subject(guid),
    item_guid UUID NOT NULL REFERENCES item(guid),
    amount SMALLINT NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 34. notification
CREATE TABLE notification (
    guid UUID PRIMARY KEY,
    app_user_guid UUID NOT NULL REFERENCES app_user(guid),
    title_i18n JSONB NOT NULL,
    content_i18n JSONB NULL,
    resource_url VARCHAR(255) NULL,
    urgency urgency_enum NOT NULL DEFAULT 1,
    meta JSONB NULL,
    seen_at TIMESTAMPTZ NULL,
    expire_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 35. student_achievement
CREATE TABLE student_achievement (
    guid UUID PRIMARY KEY,
    app_user_guid UUID NOT NULL REFERENCES app_user(guid),
    achievement_guid UUID NOT NULL REFERENCES achievement(guid),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 36. chat_member
CREATE TABLE chat_member (
    guid UUID PRIMARY KEY,
    app_user_guid UUID NOT NULL REFERENCES app_user(guid),
    chat_guid UUID NOT NULL REFERENCES chat(guid),
    chat_role chat_role_enum NOT NULL DEFAULT 0,
    name VARCHAR(150) NULL,
    custom_role_name VARCHAR(100) NULL,
    left_at TIMESTAMPTZ NULL,
    seen_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

-- 37. message
CREATE TABLE message (
    guid UUID PRIMARY KEY,
    chat_member_guid UUID NOT NULL REFERENCES chat_member(guid),
    chat_guid UUID NOT NULL REFERENCES chat(guid),
    resource_url VARCHAR(255) NULL,
    content VARCHAR(1024) NULL,
    urgency urgency_enum NOT NULL DEFAULT 1,
    meta JSONB NULL,
    seen_at TIMESTAMPTZ NULL,
    edited_at TIMESTAMPTZ NULL,
    pinned_at TIMESTAMPTZ NULL,
    expire_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

CREATE INDEX idx_message_chat_created ON message (chat_guid, created_at) WHERE deleted_at IS NULL;

-- 38. daily_activity
CREATE TABLE daily_activity (
    guid UUID PRIMARY KEY,
    app_user_guid UUID NOT NULL REFERENCES app_user(guid),
    day TIMESTAMPTZ NOT NULL,
    experience_earned BIGINT NOT NULL DEFAULT 0,
    levels_earned SMALLINT NOT NULL DEFAULT 0,
    subjects_completed SMALLINT NOT NULL DEFAULT 0,
    tests_completed SMALLINT NOT NULL DEFAULT 0,
    practices_completed SMALLINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);
