CREATE TABLE events (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        location VARCHAR(255),
                        date VARCHAR(100),
                        tickets INTEGER DEFAULT 0,
                        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                        deleted_at TIMESTAMP WITH TIME ZONE
);