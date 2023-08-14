CREATE TABLE leaves IF NOT  (
    id SERIAL PRIMARY KEY, 
    name VARCHAR(200) NOT NULL,
    leave_type TEXT NOT NULL,
    from_date DATE NOT NULL,
    to_date DATE NOT NULL,
    team_name VARCHAR(200) NOT NULL,
    sick_leave_file BYTEA,
    reporter TEXT NOT NULL
);
