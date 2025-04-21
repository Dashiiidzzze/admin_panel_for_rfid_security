-- Таблица персональных данных
CREATE TABLE personalnye_dannye (
    id SERIAL PRIMARY KEY,
    familiya TEXT NOT NULL,
    imya TEXT NOT NULL,
    otchestvo TEXT,
    nomer_telefona TEXT
);

-- Таблица должностей
CREATE TABLE dolzhnosti (
    id SERIAL PRIMARY KEY,
    dolzhnost TEXT NOT NULL
);

-- Таблица сотрудников
CREATE TABLE sotrudniki (
    id SERIAL PRIMARY KEY,
    id_pers_dannykh INT REFERENCES personalnye_dannye(id) ON DELETE CASCADE,
    nomer_klyucha TEXT UNIQUE NOT NULL,
    id_dolzhnosti INT REFERENCES dolzhnosti(id) ON DELETE SET NULL
);

-- Таблица зон
CREATE TABLE zony (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- Связь many-to-many: Сотрудники ↔ Зоны
CREATE TABLE sotrudniki_zony (
    sotrudnik_id INT REFERENCES sotrudniki(id) ON DELETE CASCADE,
    zona_id INT REFERENCES zony(id) ON DELETE CASCADE,
    PRIMARY KEY (sotrudnik_id, zona_id)
);

-- Таблица логов проходов
CREATE TABLE logi_prokhodov (
    id SERIAL PRIMARY KEY,
    zona_id INT REFERENCES zony(id) ON DELETE CASCADE,
    sotrudnik_id INT REFERENCES sotrudniki(id) ON DELETE CASCADE,
    vremya TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
