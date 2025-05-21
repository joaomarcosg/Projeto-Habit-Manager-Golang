CREATE TABLE habits (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    category VARCHAR(100) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    frequency SET(
        'sunday', 'monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturday'
    ),
    start_date DATE NOT NULL DEFAULT '2025-01-01',
    target_date DATE NOT NULL '2025-01-01',
    priority TINYINT UNSIGNED NOT NULL DEFAULT 1 CHECK ((priority BETWEEN 1 AND 10)),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);