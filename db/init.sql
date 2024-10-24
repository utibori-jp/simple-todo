-- テーブルの作成
CREATE TABLE todos (
    id SERIAL PRIMARY KEY,         -- Todo項目の一意のID
    task_name VARCHAR(255) NOT NULL, -- タスク名
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 作成日時
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- 更新日時
);

-- サンプルデータの挿入
INSERT INTO todos (task_name) VALUES
('Sample Todo 1'),
('Sample Todo 2'),
('Sample Todo 3');

-- 追加のインデックスを作成することで検索性能を向上
CREATE INDEX idx_task_name ON todos (task_name);
