-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    data TEXT NOT NULL,
    createdDate DATETIME NOT NULL
);

INSERT INTO tasks (name, data, createdDate)
VALUES ('Task 1', 'Complete the financial report by EOD.', '2024-12-24 10:00:00');

INSERT INTO tasks (name, data, createdDate)
VALUES ('Task 2', 'Prepare slides for the team meeting.', '2024-12-23 15:30:00');

INSERT INTO tasks (name, data, createdDate)
VALUES ('Task 3', 'Send invitations for the annual event.', '2024-12-22 09:45:00');

INSERT INTO tasks (name, data, createdDate)
VALUES ('Task 4', 'Update the project roadmap document.', '2024-12-24 11:15:00');

INSERT INTO tasks (name, data, createdDate)
VALUES ('Task 5', 'Review the new design prototypes.', '2024-12-23 08:00:00');


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd
