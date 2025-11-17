const express = require('express');
const cors = require('cors');
const path = require('path');
const { initDatabase, insertInitialData, getDoctorsByDepartment } = require('./database');

const app = express();
const PORT = 3000;

app.use(cors());
app.use(express.json());

initDatabase(() => {
    insertInitialData();
});

app.get('/api/patients', (req, res) => {
    const { name, id_card } = req.query;
    let sql = 'SELECT * FROM patients WHERE 1=1';
    const params = [];

    if (name) {
        sql += ' AND name LIKE ?';
        params.push(`%${name}%`);
    }
    if (id_card) {
        sql += ' AND id_card LIKE ?';
        params.push(`%${id_card}%`);
    }

    const { db } = require('./database');
    db.all(sql, params, (err, rows) => {
        if (err) {
            console.error('查询患者失败:', err);
            return res.status(500).json({ error: '服务器错误' });
        }
        res.json(rows);
    });
});

app.post('/api/patients', (req, res) => {
    const { name, gender, id_card, phone, birth_date } = req.body;

    if (!name || !id_card || !phone) {
        return res.status(400).json({ error: '缺少必填字段' });
    }

    const sql = 'INSERT INTO patients (name, gender, id_card, phone, birth_date) VALUES (?, ?, ?, ?, ?)';
    const { db } = require('./database');

    db.run(sql, [name, gender, id_card, phone, birth_date], function (err) {
        if (err) {
            console.error('创建患者失败:', err);
            return res.status(500).json({ error: '患者档案创建失败' });
        }
        res.json({ id: this.lastID, message: '患者档案创建成功' });
    });
});

app.get('/api/doctors', (req, res) => {
    const department = req.query.department;

    if (!department) {
        return res.status(400).json({ error: '缺少科室参数 (department)' });
    }

    getDoctorsByDepartment(department, (err, doctors) => {
        if (err) {
            console.error('数据库查询错误:', err);
            return res.status(500).json({ error: '服务器错误' });
        }
        res.json(doctors);
    });
});

app.use(express.static(path.join(__dirname, '../hospital-client/dist')));

app.use((req, res) => {
    res.sendFile(path.join(__dirname, '../hospital-client/dist/index.html'));
});

app.listen(PORT, () => {
    console.log(`项目已启动,请在浏览器中访问 http://localhost:${PORT}`);
});
