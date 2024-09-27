CREATE TABLE tasks (
    id UUID PRIMARY KEY NOT NULL,
    description VARCHAR(255) NOT NULL,              
    completed BOOLEAN DEFAULT FALSE                 
);