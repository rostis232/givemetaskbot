CREATE TABLE users
(
    user_id      SERIAL PRIMARY KEY            NOT NULL,
    chat_id      BIGINT                         NOT NULL,
    user_name    VARCHAR(50) DEFAULT 'no_name' NOT NULL,
    language     VARCHAR(5)  DEFAULT 'EN'      NOT NULL,
    status       INT                           NOT NULL DEFAULT 0 NOT NULL,
    active_group INT         DEFAULT 0         NOT NULL,
    active_task  INT         DEFAULT 0         NOT NULL
);

CREATE TABLE groups
(
    group_id      SERIAL PRIMARY KEY            NOT NULL,
    chief_user_id INT                           NOT NULL,
    group_name    VARCHAR(50) DEFAULT 'no_name' NOT NULL,
    FOREIGN KEY (chief_user_id) REFERENCES users (user_id) ON DELETE CASCADE
);

CREATE TABLE tasks
(
    task_id          SERIAL PRIMARY KEY                   NOT NULL,
    task_name        VARCHAR(50) DEFAULT 'no_name'        NOT NULL,
    task_description TEXT        DEFAULT 'no_description' NOT NULL,
    group_id         INT                                  NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups (group_id) ON DELETE CASCADE
);

CREATE TABLE task_employee
(
    task_employee_id SERIAL PRIMARY KEY,
    task_id          INT NOT NULL,
    employee_user_id INT NOT NULL,
    FOREIGN KEY (task_id) REFERENCES tasks (task_id) ON DELETE CASCADE,
    FOREIGN KEY (employee_user_id) REFERENCES users (user_id) ON DELETE CASCADE
);

CREATE TABLE group_employee
(
    group_employee_id SERIAL PRIMARY KEY,
    group_id          INT NOT NULL,
    employee_user_id  INT NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups (group_id) ON DELETE CASCADE,
    FOREIGN KEY (employee_user_id) REFERENCES users (user_id) ON DELETE CASCADE
);