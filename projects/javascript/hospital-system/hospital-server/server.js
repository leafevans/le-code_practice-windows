const express = require('express');
const cors = require('cors');
const path = require('path');
const { initDatabase, insertInitialData, getDoctorsByDepartment, query, run, db } = require('./database');

const app = express();
const PORT = process.env.PORT || 3000;

const isPkg = typeof process.pkg !== 'undefined';

const distPath = isPkg
    ? path.join(path.dirname(process.execPath), 'hospital-client', 'dist')
    : path.join(__dirname, '../hospital-client/dist');

console.log('前端文件路径:', distPath);

// 简单的请求日志中间件
app.use((req, res, next) => {
    const start = Date.now();
    res.on('finish', () => {
        const duration = Date.now() - start;
        console.log(`[${new Date().toISOString()}] ${req.method} ${req.originalUrl} ${res.statusCode} - ${duration}ms`);
    });
    next();
});

app.use(cors());
app.use(express.json());

initDatabase(() => {
    insertInitialData();
});

app.get('/api/patients', async (req, res) => {
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

    try {
        const rows = await query(sql, params);
        res.json(rows);
    } catch (err) {
        console.error('查询患者失败:', err);
        res.status(500).json({ error: '服务器错误' });
    }
});

app.post('/api/patients', async (req, res) => {
    const { name, gender, id_card, phone, birth_date } = req.body;

    if (!name || !id_card || !phone) {
        return res.status(400).json({ error: '缺少必填字段' });
    }

    const sql = 'INSERT INTO patients (name, gender, id_card, phone, birth_date) VALUES (?, ?, ?, ?, ?)';

    try {
        const result = await run(sql, [name, gender, id_card, phone, birth_date]);
        res.json({ id: result.id, message: '患者档案创建成功' });
    } catch (err) {
        console.error('创建患者失败:', err);
        res.status(500).json({ error: '患者档案创建失败' });
    }
});

app.get('/api/doctors', async (req, res) => {
    const department = req.query.department;

    if (!department) {
        return res.status(400).json({ error: '缺少科室参数 (department)' });
    }

    try {
        const doctors = await getDoctorsByDepartment(department);
        res.json(doctors);
    } catch (err) {
        console.error('数据库查询错误:', err);
        res.status(500).json({ error: '服务器错误' });
    }
});

app.post('/api/registrations', async (req, res) => {
    const { patient_id, doctor_id, department_code, registration_date } = req.body;

    if (!patient_id || !doctor_id || !department_code || !registration_date) {
        return res.status(400).json({ error: '缺少必填字段' });
    }

    // 使用事务处理挂号逻辑
    db.serialize(() => {
        db.run('BEGIN TRANSACTION');

        db.run(
            'UPDATE doctors SET available_slots = available_slots - 1 WHERE id = ? AND available_slots > 0',
            [doctor_id],
            function (err) {
                if (err) {
                    db.run('ROLLBACK');
                    console.error('更新医生号源失败:', err);
                    return res.status(500).json({ error: '系统错误' });
                }

                if (this.changes === 0) {
                    db.run('ROLLBACK');
                    return res.status(400).json({ error: '号源不足' });
                }

                db.run(
                    'INSERT INTO registrations (patient_id, doctor_id, department_code, registration_date) VALUES (?, ?, ?, ?)',
                    [patient_id, doctor_id, department_code, registration_date],
                    function (err) {
                        if (err) {
                            db.run('ROLLBACK');
                            console.error('保存挂号记录失败:', err);
                            return res.status(500).json({ error: '挂号失败' });
                        }

                        db.run('COMMIT');
                        res.json({ id: this.lastID, message: '挂号成功' });
                    }
                );
            }
        );
    });
});

app.get('/api/registrations/patient/:patientId', async (req, res) => {
    const { patientId } = req.params;

    const sql = `
        SELECT 
            r.*,
            d.name as doctor_name,
            dept.name as department_name
        FROM registrations r
        LEFT JOIN doctors d ON r.doctor_id = d.id
        LEFT JOIN departments dept ON r.department_code = dept.code
        WHERE r.patient_id = ?
        ORDER BY r.created_at DESC
    `;

    try {
        const rows = await query(sql, [patientId]);
        res.json(rows);
    } catch (err) {
        console.error('查询挂号记录失败:', err);
        res.status(500).json({ error: '服务器错误' });
    }
});

app.get('/api/medicines', async (req, res) => {
    const { name, category } = req.query;
    let sql = 'SELECT * FROM medicines WHERE 1=1';
    const params = [];

    if (name) {
        sql += ' AND name LIKE ?';
        params.push(`%${name}%`);
    }
    if (category) {
        sql += ' AND category = ?';
        params.push(category);
    }

    try {
        const rows = await query(sql, params);
        res.json(rows);
    } catch (err) {
        console.error('查询药品失败:', err);
        res.status(500).json({ error: '服务器错误' });
    }
});

app.post('/api/medicines', async (req, res) => {
    const { name, specification, unit, price, stock, category } = req.body;

    if (!name || !price) {
        return res.status(400).json({ error: '缺少必填字段' });
    }

    const sql = 'INSERT INTO medicines (name, specification, unit, price, stock, category) VALUES (?, ?, ?, ?, ?, ?)';

    try {
        const result = await run(sql, [name, specification, unit, price, stock || 0, category]);
        res.json({ id: result.id, message: '药品添加成功' });
    } catch (err) {
        console.error('添加药品失败:', err);
        res.status(500).json({ error: '药品添加失败' });
    }
});

app.put('/api/medicines/:id', async (req, res) => {
    const { id } = req.params;
    const { stock } = req.body;

    const sql = 'UPDATE medicines SET stock = ? WHERE id = ?';

    try {
        await run(sql, [stock, id]);
        res.json({ message: '库存更新成功' });
    } catch (err) {
        console.error('更新库存失败:', err);
        res.status(500).json({ error: '库存更新失败' });
    }
});

app.get('/api/statistics/registrations', async (req, res) => {
    const sql = `
        SELECT 
            registration_date as date,
            COUNT(*) as count
        FROM registrations
        GROUP BY registration_date
        ORDER BY registration_date
    `;
    try {
        const rows = await query(sql);
        res.json(rows);
    } catch (err) {
        console.error('查询门诊量统计失败:', err);
        res.status(500).json({ error: '服务器错误' });
    }
});

app.get('/api/statistics/medicines', async (req, res) => {
    const sql = `
        SELECT 
            category,
            SUM(stock) as totalStock,
            COUNT(*) as count
        FROM medicines
        GROUP BY category
    `;
    try {
        const rows = await query(sql);
        res.json(rows);
    } catch (err) {
        console.error('查询药品统计失败:', err);
        res.status(500).json({ error: '服务器错误' });
    }
});

app.get('/api/statistics/departments', async (req, res) => {
    const sql = `
        SELECT 
            d.name as departmentName,
            COUNT(r.id) as count
        FROM departments d
        LEFT JOIN registrations r ON d.code = r.department_code
        GROUP BY d.code, d.name
    `;
    try {
        const rows = await query(sql);
        res.json(rows);
    } catch (err) {
        console.error('查询科室统计失败:', err);
        res.status(500).json({ error: '服务器错误' });
    }
});

app.use(express.static(distPath));

app.use((req, res) => {
    res.sendFile(path.join(distPath, 'index.html'));
});

app.listen(PORT, () => {
    console.log(`========================================`);
    console.log(`医院管理系统已启动！`);
    console.log(`访问地址: http://localhost:${PORT}`);
    console.log(`========================================`);

    if (isPkg) {
        const { exec } = require('child_process');
        setTimeout(() => {
            exec(`start http://localhost:${PORT}`);
        }, 2000);
    }
});
