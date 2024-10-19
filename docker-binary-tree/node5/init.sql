-- This script initializes the database with a table.
CREATE TABLE IF NOT EXISTS test_data (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100)
);

-- Insert sample data
INSERT INTO test_data (name) VALUES ('Node 5 Data 1'), ('Node 5 Data 2');