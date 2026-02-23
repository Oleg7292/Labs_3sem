-- Создание таблицы Характеры
CREATE TABLE personalities (
    id SERIAL PRIMARY KEY,
    description TEXT NOT NULL,
    main_trait VARCHAR(100) NOT NULL
);

-- Создание таблицы Локации
CREATE TABLE locations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT
);

-- Создание таблицы Персонажи
CREATE TABLE characters (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    species VARCHAR(50) NOT NULL,
    gender VARCHAR(10) CHECK (gender IN ('Мужской', 'Женский')),
    personality_id INT REFERENCES personalities(id) ON DELETE SET NULL
);

-- Создание таблицы Эпизоды
CREATE TABLE episodes (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    release_date DATE,
    season INT
);

-- Создание таблицы Персонажи_Эпизоды (Многие-ко-Многим)
CREATE TABLE character_episodes (
    character_id INT REFERENCES characters(id) ON DELETE CASCADE,
    episode_id INT REFERENCES episodes(id) ON DELETE CASCADE,
    PRIMARY KEY (character_id, episode_id)
);

-- Создание таблицы Предметы
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    owner_id INT REFERENCES characters(id) ON DELETE SET NULL,
    location_id INT REFERENCES locations(id) ON DELETE SET NULL
);

-- Создание таблицы Отношения
CREATE TABLE relationships (
    id SERIAL PRIMARY KEY,
    char1_id INT REFERENCES characters(id) ON DELETE CASCADE,
    char2_id INT REFERENCES characters(id) ON DELETE CASCADE,
    relationship_type VARCHAR(50) NOT NULL,
    CHECK (char1_id < char2_id) -- Предотвращение дубликатов пар (1,2) и (2,1)
);

-- Вставка Характеров
INSERT INTO personalities (description, main_trait) VALUES
('Энергичный и оптимистичный, любит приключения.', 'Оптимизм'),
('Спокойный, интеллектуальный и философский.', 'Мудрость'),
('Романтичный, поэтичный и чувствительный.', 'Чувствительность'),
('Активная, спортивная, но иногда неуклюжая.', 'Активность'),
('Изобретательный, научный и прагматичный.', 'Изобретательность'),
('Заботливая, любит готовить, материнская фигура.', 'Забота'),
('Артистичная, драматичная и любит внимание.', 'Артистичность'),
('Любит огород, практичный и приземленный.', 'Практичность'),
('Таинственный, магический и наблюдательный.', 'Загадочность');

-- Вставка Локаций
INSERT INTO locations (name, description) VALUES
('Дом Кроша', 'Разноцветный домик со скамейкой снаружи.'),
('Дом Ёжика', 'Уютный домик, наполненный коллекциями фантиков и кактусами.'),
('Дом Бараша', 'Домик на холме, идеален для написания стихов.'),
('Дом Лосяша', 'Дом с библиотекой и компьютером.'),
('Огород Копатыча', 'Большое поле с ульями и овощами.'),
('Дом Совуньи', 'Дом в большом дереве, содержит кухню и медицинские принадлежности.'),
('Ангар Пина', 'Большой металлический ангар, наполненный изобретениями.'),
('Дом Нюши', 'Розовый домик с цветами.'),
('Море', 'Большой водоем, окружающий долину.');

-- Вставка Персонажей
INSERT INTO characters (name, species, gender, personality_id) VALUES
('Крош', 'Кролик', 'Мужской', 1),
('Ёжик', 'Ёж', 'Мужской', 2),
('Бараш', 'Баран', 'Мужской', 3),
('Нюша', 'Свинья', 'Женский', 7),
('Лосяш', 'Лось', 'Мужской', 5),
('Копатыч', 'Медведь', 'Мужской', 8),
('Совунья', 'Сова', 'Женский', 6),
('Пин', 'Пингвин', 'Мужской', 5),
('Кар-Карыч', 'Ворон', 'Мужской', 9);

-- Вставка Эпизодов
INSERT INTO episodes (title, release_date, season) VALUES
('Скамейка', '2004-05-17', 1),
('Принц для Нюши', '2004-05-18', 1),
('Фанерное солнце', '2004-05-19', 1),
('Куда уходит старый год', '2004-05-20', 1),
('Энергия храпа', '2004-05-21', 1),
('Железная няня', '2004-05-22', 1);

-- Вставка Персонажи_Эпизоды
INSERT INTO character_episodes (character_id, episode_id) VALUES
(1, 1), (2, 1), -- Крош и Ёжик в Скамейке
(4, 2), (3, 2), -- Нюша и Бараш в Принце для Нюши
(1, 3), (2, 3), (3, 3), (4, 3), -- Многие в Фанерном солнце
(9, 4), (6, 4), -- Кар-Карыч и Копатыч в Куда уходит старый год
(8, 6), (1, 6), (2, 6); -- Пин, Крош, Ёжик в Железной няне

-- Вставка Предметов
INSERT INTO items (name, owner_id, location_id) VALUES
('Скамейка', 1, 1),
('Коллекция фантиков', 2, 2),
('Зонтик', 3, 3),
('Компьютер', 5, 4),
('Горшок меда', 6, 5),
('Аптечка', 7, 6),
('Гаечный ключ', 8, 7),
('Зеркало', 4, 8),
('Рояль', 9, 9);

-- Вставка Отношений
INSERT INTO relationships (char1_id, char2_id, relationship_type) VALUES
(1, 2, 'Лучшие друзья'), -- Крош и Ёжик
(3, 4, 'Романтический интерес'), -- Бараш и Нюша
(6, 7, 'Старые друзья'), -- Копатыч и Совунья
(5, 8, 'Коллеги'), -- Лосяш и Пин
(1, 4, 'Друзья');
