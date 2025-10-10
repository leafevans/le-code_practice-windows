CREATE TABLE student(
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    name VARCHAR(10) COMMENT '姓名',
    no VARCHAR(10) COMMENT '学号'
) COMMENT = '学生表';

INSERT INTO student
VALUES(NULL, '黛绮丝', '2000100101'),
    (NULL, '谢逊', '2000100102'),
    (NULL, '殷天正', '2000100103'),
    (NULL, '韦一笑', '2000100104');
