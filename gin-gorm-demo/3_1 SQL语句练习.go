package main

/**
CREATE TABLE students (
    id INT AUTO_INCREMENT PRIMARY KEY,       -- 自增主键
    name VARCHAR(50) NOT NULL,               -- 学生姓名（非空）
    age INT,                                 -- 学生年龄（可为空）
    grade VARCHAR(20)                        -- 学生年级（可为空）
);
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

INSERT INTO students (name, age, grade) VALUES ('张三', 20, '三年级');

SELECT * FROM students WHERE age > 18;

UPDATE students SET grade = '四年级' WHERE name = '张三';

DELETE FROM students WHERE age < 15;

2、事务语句
use go;

CREATE TABLE accounts (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '账户ID，主键',
    balance DECIMAL(10, 2) NOT NULL DEFAULT 0.00 COMMENT '账户余额，精确到小数点后两位',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '账户创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '账户最后更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户账户表';


CREATE TABLE transactions (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '转账记录ID，主键',
    from_account_id INT NOT NULL COMMENT '转出账户ID',
    to_account_id INT NOT NULL COMMENT '转入账户ID',
    amount DECIMAL(10, 2) NOT NULL COMMENT '转账金额，必须大于0',
    transaction_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '转账时间',
    status TINYINT DEFAULT 1 COMMENT '转账状态：1-成功，0-失败（可选）',
    FOREIGN KEY (from_account_id) REFERENCES accounts(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (to_account_id) REFERENCES accounts(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    CHECK (amount > 0)  -- 确保转账金额为正数
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='转账记录表';

INSERT INTO accounts (balance) VALUES
(1000.00),  -- 账户ID=1，余额1000元
(500.00);   -- 账户ID=2，余额500元

DROP PROCEDURE IF EXISTS TransferMoney;

DELIMITER //

CREATE PROCEDURE TransferMoney(
    IN from_id INT,
    IN to_id INT,
    IN amount DECIMAL(10, 2)
)
BEGIN
    DECLARE balance_a DECIMAL(10, 2);
    DECLARE EXIT HANDLER FOR SQLEXCEPTION
    BEGIN
        ROLLBACK;
        SELECT '转账失败：发生异常' AS message;
    END;

    START TRANSACTION;

    -- 检查账户A的余额
    SELECT balance INTO balance_a FROM accounts WHERE id = from_id FOR UPDATE;

    -- 判断余额是否足够
    IF balance_a < amount THEN
        ROLLBACK;
        SELECT CONCAT('转账失败：余额不足（当前余额：', balance_a, '）') AS message;
    ELSE
        -- 扣款
        UPDATE accounts SET balance = balance - amount WHERE id = from_id;
        -- 入账
        UPDATE accounts SET balance = balance + amount WHERE id = to_id;
        -- 记录转账日志
        INSERT INTO transactions (from_account_id, to_account_id, amount)
        VALUES (from_id, to_id, amount);

        COMMIT;
        SELECT '转账成功' AS message;
    END IF;
END //

DELIMITER ;

-- 调用存储过程
CALL TransferMoney(1, 2, 100);
*/
