CREATE TABLE users (
                       user_id SERIAL PRIMARY KEY,
                       chat_id INT NOT NULL,
                       user_name TEXT NOT NULL
);

CREATE TABLE groups (
                        group_id SERIAL PRIMARY KEY,
                        chief_user_id INT NOT NULL,
                        group_name TEXT NOT NULL,
                        FOREIGN KEY (chief_user_id) REFERENCES users (user_id) ON DELETE CASCADE
);

CREATE TABLE tasks (
                       task_id SERIAL PRIMARY KEY,
                       task_name TEXT NOT NULL,
                       task_description TEXT,
                       group_id INT NOT NULL,
                       FOREIGN KEY (group_id) REFERENCES groups (group_id) ON DELETE CASCADE
);

CREATE TABLE task_employee (
                               task_employee_id SERIAL PRIMARY KEY ,
                               task_id INT NOT NULL,
                               employee_user_id INT NOT NULL,
                               FOREIGN KEY (task_id) REFERENCES tasks (task_id) ON DELETE CASCADE,
                               FOREIGN KEY (employee_user_id) REFERENCES users (user_id) ON DELETE CASCADE
);

CREATE TABLE group_employee (
                                group_employee_id SERIAL PRIMARY KEY,
                                group_id INT NOT NULL,
                                employee_user_id INT NOT NULL,
                                FOREIGN KEY (group_id) REFERENCES groups (group_id) ON DELETE CASCADE,
                                FOREIGN KEY (employee_user_id) REFERENCES users (user_id) ON DELETE CASCADE
);