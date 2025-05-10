-- Таблица персональных данных
CREATE TABLE IF NOT EXISTS personalnye_dannye (
    id SERIAL PRIMARY KEY,
    familiya TEXT NOT NULL,
    imya TEXT NOT NULL,
    otchestvo TEXT,
    nomer_telefona TEXT
);

-- Таблица должностей
CREATE TABLE IF NOT EXISTS dolzhnosti (
    id SERIAL PRIMARY KEY,
    dolzhnost TEXT NOT NULL
);

-- Таблица сотрудников
CREATE TABLE IF NOT EXISTS staff (
    id SERIAL PRIMARY KEY,
    id_pers_dannykh INT REFERENCES personalnye_dannye(id) ON DELETE CASCADE,
    nomer_klyucha TEXT UNIQUE NOT NULL,
    id_dolzhnosti INT REFERENCES dolzhnosti(id) ON DELETE SET NULL
);

-- Таблица зон
CREATE TABLE IF NOT EXISTS zony (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- Связь many-to-many: Сотрудники ↔ Зоны
CREATE TABLE IF NOT EXISTS staff_zones (
    staff_id INT REFERENCES staff(id) ON DELETE CASCADE,
    zona_id INT REFERENCES zony(id) ON DELETE CASCADE,
    PRIMARY KEY (staff_id, zona_id)
);

-- Таблица логов проходов
CREATE TABLE IF NOT EXISTS logs_passes (
    id SERIAL PRIMARY KEY,
    zone_id INT REFERENCES zony(id) ON DELETE CASCADE,
    staff_id INT REFERENCES staff(id) ON DELETE CASCADE,
    vremya TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);








-- Вставим должности
INSERT INTO dolzhnosti (dolzhnost) VALUES
('Системный администратор'),
('Оператор видеонаблюдения'),
('Инженер по безопасности');

-- Вставим персональные данные
INSERT INTO personalnye_dannye (familiya, imya, otchestvo, nomer_telefona) VALUES
('Иванов', 'Иван', 'Иванович', '+79001112233'),
('Петрова', 'Анна', 'Сергеевна', '+79003334455'),
('Сидоров', 'Алексей', 'Михайлович', '+79006667788');

-- Вставим сотрудников
INSERT INTO staff (id_pers_dannykh, nomer_klyucha, id_dolzhnosti) VALUES
(1, 'KEY12345', 1),
(2, 'KEY54321', 2),
(3, 'KEY99999', 3);

-- Вставим зоны
INSERT INTO zony (name) VALUES
('Зона A'),
('Зона B'),
('Зона C'),
('Зона D'),
('Зона E');

-- Назначим сотрудникам доступ в зоны
INSERT INTO staff_zones (staff_id, zona_id) VALUES
(1, 1), -- Иванов -> Зона A
(1, 2), -- Иванов -> Зона B
(2, 2), -- Петрова -> Зона B
(3, 3); -- Сидоров -> Зона C

-- Добавим логи проходов
INSERT INTO logs_passes (zone_id, staff_id) VALUES
(1, 1),
(2, 1),
(2, 2),
(3, 3),
(1, 1);
