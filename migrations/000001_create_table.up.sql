CREATE TABLE custom_events (
    id UUID PRIMARY KEY,
    user_id UUID not NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    date DATE NOT NULL,
    category VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at bigint default 0
);


CREATE TABLE milestones (
    id UUID PRIMARY KEY,
    user_id UUID not NULL,
    title VARCHAR(255) NOT NULL,
    date DATE NOT NULL,
    category VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at bigint default 0
);

