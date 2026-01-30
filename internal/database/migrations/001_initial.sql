-- +goose Up
CREATE TABLE settings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    key TEXT UNIQUE NOT NULL,
    value TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Seed initial settings (can be used for dynamic content later)
INSERT INTO settings (key, value) VALUES
    ('site_name', 'Kids First Childcare'),
    ('phone', '715-313-0578'),
    ('address', '400 Woodside Dr, Cornell, WI 54732'),
    ('hours_time', '5:30am - 6:00pm'),
    ('hours_days', 'Monday through Friday');

-- +goose Down
DROP TABLE IF EXISTS settings;
