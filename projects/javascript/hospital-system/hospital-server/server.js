const express = require('express');
const cors = require('cors');
const path = require('path');
const { initDatabase, insertInitialData, getDoctorsByDepartment, query, run, get, db } = require('./database');

const app = express();
const PORT = process.env.PORT || 3000;

const isPkg = typeof process.pkg !== 'undefined';

const distPath = isPkg
    ? path.join(path.dirname(process.execPath), 'hospital-client', 'dist')
    : path.join(__dirname, '../hospital-client/dist');

console.log('前端文件路径:', distPath);

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
    const { name, id_card, page = 1, limit = 10 } = req.query;
    const limitNum = parseInt(limit);
    const offset = (parseInt(page) - 1) * limitNum;

    let whereSql = 'WHERE 1=1';
    const params = [];

    if (name) {
        whereSql += ' AND name LIKE ?';
        params.push(`%${name}%`);
    }
    if (id_card) {
        whereSql += ' AND id_card LIKE ?';
        params.push(`%${id_card}%`);
    }

    try {
        const countSql = `SELECT COUNT(*) as total FROM patients ${whereSql}`;
        const countResult = await query(countSql, params);
        const total = countResult[0].total;

        const dataSql = `SELECT * FROM patients ${whereSql} LIMIT ? OFFSET ?`;
        const rows = await query(dataSql, [...params, limitNum, offset]);

        res.json({
            total,
            list: rows,
            page: parseInt(page),
            limit: limitNum
        });
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
    const { department, name, page, limit } = req.query;

    if (department && !page) {
        try {
            const doctors = await getDoctorsByDepartment(department);
            return res.json(doctors);
        } catch (err) {
            console.error('数据库查询错误:', err);
            return res.status(500).json({ error: '服务器错误' });
        }
    }

    const pageNum = parseInt(page) || 1;
    const limitNum = parseInt(limit) || 10;
    const offset = (pageNum - 1) * limitNum;

    let whereSql = 'WHERE 1=1';
    const params = [];

    if (department) {
        whereSql += ' AND department_code = ?';
        params.push(department);
    }
    if (name) {
        whereSql += ' AND name LIKE ?';
        params.push(`%${name}%`);
    }

    try {
        const countSql = `SELECT COUNT(*) as total FROM doctors ${whereSql}`;
        const countResult = await query(countSql, params);
        const total = countResult[0].total;

        const dataSql = `
            SELECT d.*, dept.name as department_name 
            FROM doctors d
            LEFT JOIN departments dept ON d.department_code = dept.code
            ${whereSql} 
            LIMIT ? OFFSET ?
        `;
        const rows = await query(dataSql, [...params, limitNum, offset]);

        res.json({
            total,
            list: rows,
            page: pageNum,
            limit: limitNum
        });
    } catch (err) {
        console.error('查询医生列表失败:', err);
        res.status(500).json({ error: '服务器错误' });
    }
});

app.post('/api/doctors', async (req, res) => {
    const { name, title, department_code, available_slots } = req.body;

    if (!name || !title || !department_code) {
        return res.status(400).json({ error: '缺少必填字段' });
    }

    const sql = 'INSERT INTO doctors (name, title, department_code, available_slots) VALUES (?, ?, ?, ?)';

    try {
        const result = await run(sql, [name, title, department_code, available_slots || 0]);
        res.json({ id: result.id, message: '医生添加成功' });
    } catch (err) {
        console.error('添加医生失败:', err);
        res.status(500).json({ error: '医生添加失败' });
    }
});

app.put('/api/doctors/:id', async (req, res) => {
    const { id } = req.params;
    const { name, title, department_code, available_slots } = req.body;

    const sql = 'UPDATE doctors SET name = ?, title = ?, department_code = ?, available_slots = ? WHERE id = ?';

    try {
        await run(sql, [name, title, department_code, available_slots, id]);
        res.json({ message: '医生信息更新成功' });
    } catch (err) {
        console.error('更新医生失败:', err);
        res.status(500).json({ error: '医生信息更新失败' });
    }
});

app.delete('/api/doctors/:id', async (req, res) => {
    const { id } = req.params;
    const sql = 'DELETE FROM doctors WHERE id = ?';

    try {
        await run(sql, [id]);
        res.json({ message: '医生删除成功' });
    } catch (err) {
        console.error('删除医生失败:', err);
        res.status(500).json({ error: '删除医生失败' });
    }
});

app.get('/api/departments', async (req, res) => {
    try {
        const rows = await query('SELECT * FROM departments');
        res.json(rows);
    } catch (err) {
        console.error('查询科室失败:', err);
        res.status(500).json({ error: '服务器错误' });
    }
});

app.post('/api/departments', async (req, res) => {
    const { name, code } = req.body;

    if (!name || !code) {
        return res.status(400).json({ error: '缺少必填字段' });
    }

    const sql = 'INSERT INTO departments (name, code) VALUES (?, ?)';

    try {
        const result = await run(sql, [name, code]);
        res.json({ id: result.id, message: '科室添加成功' });
    } catch (err) {
        console.error('添加科室失败:', err);
        if (err.message.includes('UNIQUE constraint failed')) {
            return res.status(400).json({ error: '科室代码已存在' });
        }
        res.status(500).json({ error: '科室添加失败' });
    }
});

app.put('/api/departments/:id', async (req, res) => {
    const { id } = req.params;
    const { name, code } = req.body;

    const sql = 'UPDATE departments SET name = ?, code = ? WHERE id = ?';

    try {
        await run(sql, [name, code, id]);
        res.json({ message: '科室更新成功' });
    } catch (err) {
        console.error('更新科室失败:', err);
        res.status(500).json({ error: '科室更新失败' });
    }
});

app.delete('/api/departments/:id', async (req, res) => {
    const { id } = req.params;

    try {
        const dept = await get('SELECT code FROM departments WHERE id = ?', [id]);
        if (!dept) {
            return res.status(404).json({ error: '科室不存在' });
        }

        const doctorCount = await get('SELECT COUNT(*) as count FROM doctors WHERE department_code = ?', [dept.code]);
        if (doctorCount.count > 0) {
            return res.status(400).json({ error: '该科室下还有医生，无法删除' });
        }

        await run('DELETE FROM departments WHERE id = ?', [id]);
        res.json({ message: '科室删除成功' });
    } catch (err) {
        console.error('删除科室失败:', err);
        res.status(500).json({ error: '删除科室失败' });
    }
});

app.post('/api/registrations', async (req, res) => {
    const { patient_id, doctor_id, department_code, registration_date } = req.body;

    if (!patient_id || !doctor_id || !department_code || !registration_date) {
        return res.status(400).json({ error: '缺少必填字段' });
    }

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

app.get('/api/registrations', async (req, res) => {
    const { date, department, status, page = 1, limit = 10 } = req.query;
    const limitNum = parseInt(limit);
    const offset = (parseInt(page) - 1) * limitNum;

    let whereSql = 'WHERE 1=1';
    const params = [];

    if (date) {
        whereSql += ' AND r.registration_date = ?';
        params.push(date);
    }
    if (department) {
        whereSql += ' AND r.department_code = ?';
        params.push(department);
    }
    if (status) {
        whereSql += ' AND r.status = ?';
        params.push(status);
    }

    try {
        const countSql = `SELECT COUNT(*) as total FROM registrations r ${whereSql}`;
        const countResult = await query(countSql, params);
        const total = countResult[0].total;

        const dataSql = `
            SELECT 
                r.*,
                p.name as patient_name,
                p.gender as patient_gender,
                d.name as doctor_name,
                dept.name as department_name
            FROM registrations r
            LEFT JOIN patients p ON r.patient_id = p.id
            LEFT JOIN doctors d ON r.doctor_id = d.id
            LEFT JOIN departments dept ON r.department_code = dept.code
            ${whereSql}
            ORDER BY r.created_at DESC
            LIMIT ? OFFSET ?
        `;
        const rows = await query(dataSql, [...params, limitNum, offset]);

        res.json({
            total,
            list: rows,
            page: parseInt(page),
            limit: limitNum
        });
    } catch (err) {
        console.error('查询挂号列表失败:', err);
        res.status(500).json({ error: '服务器错误' });
    }
});

app.put('/api/registrations/:id/status', async (req, res) => {
    const { id } = req.params;
    const { status } = req.body;

    if (!['pending', 'completed', 'cancelled'].includes(status)) {
        return res.status(400).json({ error: '无效的状态' });
    }

    try {
        if (status === 'cancelled') {
            const reg = await get('SELECT doctor_id, status FROM registrations WHERE id = ?', [id]);
            if (!reg) return res.status(404).json({ error: '挂号记录不存在' });

            if (reg.status !== 'cancelled') {
                await run('UPDATE doctors SET available_slots = available_slots + 1 WHERE id = ?', [reg.doctor_id]);
            }
        }

        await run('UPDATE registrations SET status = ? WHERE id = ?', [status, id]);
        res.json({ message: '状态更新成功' });
    } catch (err) {
        console.error('更新挂号状态失败:', err);
        res.status(500).json({ error: '状态更新失败' });
    }
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

app.get('/api/registrations/:id', async (req, res) => {
    const { id } = req.params;
    const sql = `
        SELECT 
            r.*,
            p.name as patient_name,
            p.gender as patient_gender,
            p.birth_date as patient_birth_date,
            d.name as doctor_name,
            dept.name as department_name
        FROM registrations r
        LEFT JOIN patients p ON r.patient_id = p.id
        LEFT JOIN doctors d ON r.doctor_id = d.id
        LEFT JOIN departments dept ON r.department_code = dept.code
        WHERE r.id = ?
    `;
    try {
        const row = await get(sql, [id]);
        if (!row) return res.status(404).json({ error: '挂号记录不存在' });
        res.json(row);
    } catch (err) {
        console.error('查询挂号详情失败:', err);
        res.status(500).json({ error: '服务器错误' });
    }
});

app.post('/api/consultations', async (req, res) => {
    const { registration_id, symptom, diagnosis, prescriptions } = req.body;

    if (!registration_id || !diagnosis) {
        return res.status(400).json({ error: '缺少必填字段' });
    }

    db.serialize(() => {
        db.run('BEGIN TRANSACTION');

        db.run(
            'UPDATE registrations SET symptom = ?, diagnosis = ?, status = ? WHERE id = ?',
            [symptom, diagnosis, 'completed', registration_id],
            function (err) {
                if (err) {
                    db.run('ROLLBACK');
                    console.error('更新诊断信息失败:', err);
                    return res.status(500).json({ error: '保存诊断失败' });
                }

                if (prescriptions && prescriptions.length > 0) {
                    const stmt = db.prepare(`
                        INSERT INTO prescriptions (registration_id, patient_id, doctor_id, medicine_id, quantity, dosage)
                        VALUES (?, ?, ?, ?, ?, ?)
                    `);

                    db.get('SELECT patient_id, doctor_id FROM registrations WHERE id = ?', [registration_id], (err, reg) => {
                        if (err || !reg) {
                            db.run('ROLLBACK');
                            return res.status(500).json({ error: '查询挂号信息失败' });
                        }

                        let errorOccurred = false;
                        prescriptions.forEach(p => {
                            if (errorOccurred) return;
                            stmt.run(registration_id, reg.patient_id, reg.doctor_id, p.medicine_id, p.quantity, p.dosage, (err) => {
                                if (err) {
                                    errorOccurred = true;
                                    console.error('保存处方失败:', err);
                                }
                            });

                            db.run('UPDATE medicines SET stock = stock - ? WHERE id = ?', [p.quantity, p.medicine_id]);
                        });
                        stmt.finalize();

                        if (errorOccurred) {
                            db.run('ROLLBACK');
                            return res.status(500).json({ error: '保存处方失败' });
                        }

                        db.run('COMMIT');
                        res.json({ message: '接诊完成' });
                    });
                } else {
                    db.run('COMMIT');
                    res.json({ message: '接诊完成' });
                }
            }
        );
    });
});

app.get('/api/registrations/doctor/:doctorId', async (req, res) => {
    const { doctorId } = req.params;
    const { status } = req.query;

    let sql = `
        SELECT 
            r.*,
            p.name as patient_name,
            p.gender as patient_gender,
            p.birth_date as patient_birth_date
        FROM registrations r
        LEFT JOIN patients p ON r.patient_id = p.id
        WHERE r.doctor_id = ?
    `;

    const params = [doctorId];
    if (status) {
        sql += ' AND r.status = ?';
        params.push(status);
    }

    sql += ' ORDER BY r.registration_date, r.created_at';

    try {
        const rows = await query(sql, params);
        res.json(rows);
    } catch (err) {
        console.error('查询医生挂号列表失败:', err);
        res.status(500).json({ error: '服务器错误' });
    }
});

app.get('/api/charges', async (req, res) => {
    const { name, id_card } = req.query;

    let sql = `
        SELECT 
            r.*,
            p.name as patient_name,
            p.id_card,
            d.name as doctor_name,
            dept.name as department_name
        FROM registrations r
        LEFT JOIN patients p ON r.patient_id = p.id
        LEFT JOIN doctors d ON r.doctor_id = d.id
        LEFT JOIN departments dept ON r.department_code = dept.code
        WHERE r.status = 'completed' AND (r.payment_status = 'unpaid' OR r.payment_status IS NULL)
    `;

    const params = [];
    if (name) {
        sql += ' AND p.name LIKE ?';
        params.push(`%${name}%`);
    }
    if (id_card) {
        sql += ' AND p.id_card LIKE ?';
        params.push(`%${id_card}%`);
    }

    sql += ' ORDER BY r.registration_date DESC';

    try {
        const rows = await query(sql, params);
        res.json(rows);
    } catch (err) {
        console.error('查询待缴费列表失败:', err);
        res.status(500).json({ error: '服务器错误' });
    }
});

app.get('/api/charges/:id/detail', async (req, res) => {
    const { id } = req.params;

    try {
        const reg = await get(`
            SELECT r.*, p.name as patient_name 
            FROM registrations r 
            JOIN patients p ON r.patient_id = p.id 
            WHERE r.id = ?
        `, [id]);

        if (!reg) return res.status(404).json({ error: '记录不存在' });

        const prescriptions = await query(`
            SELECT 
                p.*,
                m.name as medicine_name,
                m.price,
                m.unit
            FROM prescriptions p
            JOIN medicines m ON p.medicine_id = m.id
            WHERE p.registration_id = ?
        `, [id]);

        let total = reg.fee || 0;
        let medicineFee = 0;

        prescriptions.forEach(item => {
            medicineFee += item.price * item.quantity;
        });

        total += medicineFee;

        res.json({
            registration: reg,
            prescriptions,
            fees: {
                registration: reg.fee,
                medicine: medicineFee,
                total: parseFloat(total.toFixed(2))
            }
        });
    } catch (err) {
        console.error('查询费用详情失败:', err);
        res.status(500).json({ error: '服务器错误' });
    }
});

app.post('/api/charges/:id/pay', async (req, res) => {
    const { id } = req.params;

    try {
        await run("UPDATE registrations SET payment_status = 'paid' WHERE id = ?", [id]);
        res.json({ message: '缴费成功' });
    } catch (err) {
        console.error('缴费失败:', err);
        res.status(500).json({ error: '缴费失败' });
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

app.post('/api/login', (req, res) => {
    const { username, password } = req.body;
    if (username === 'admin' && password === '123456') {
        res.json({
            token: 'mock-token-123456',
            user: { name: '管理员', role: 'admin' }
        });
    } else if (username === 'doctor' && password === '123456') {
        res.json({
            token: 'mock-token-doctor',
            user: { name: '张医生', role: 'doctor' }
        });
    } else {
        res.status(401).json({ error: '用户名或密码错误' });
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
