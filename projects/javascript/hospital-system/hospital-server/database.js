const sqlite3 = require('sqlite3').verbose();
const path = require('path');

const isPkg = typeof process.pkg !== 'undefined';

const dbPath = isPkg
    ? path.join(path.dirname(process.execPath), 'hospital.db')
    : path.join(__dirname, 'hospital.db');

console.log('数据库文件路径:', dbPath);

const db = new sqlite3.Database(dbPath, (err) => {
    if (err) {
        console.error('数据库连接失败:', err.message);
    } else {
        console.log('数据库连接成功');
    }
});

function query(sql, params = []) {
    return new Promise((resolve, reject) => {
        db.all(sql, params, (err, rows) => {
            if (err) reject(err);
            else resolve(rows);
        });
    });
}

function run(sql, params = []) {
    return new Promise((resolve, reject) => {
        db.run(sql, params, function (err) {
            if (err) reject(err);
            else resolve({ id: this.lastID, changes: this.changes });
        });
    });
}

function get(sql, params = []) {
    return new Promise((resolve, reject) => {
        db.get(sql, params, (err, row) => {
            if (err) reject(err);
            else resolve(row);
        });
    });
}

function initDatabase(callback) {
    db.serialize(() => {
        db.run(`
            CREATE TABLE IF NOT EXISTS departments (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                name TEXT NOT NULL,
                code TEXT UNIQUE NOT NULL
            )
        `);

        db.run(`
            CREATE TABLE IF NOT EXISTS doctors (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                name TEXT NOT NULL,
                title TEXT NOT NULL,
                department_code TEXT NOT NULL,
                available_slots INTEGER DEFAULT 0,
                FOREIGN KEY (department_code) REFERENCES departments(code)
            )
        `);

        db.run(`
            CREATE TABLE IF NOT EXISTS patients (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                name TEXT NOT NULL,
                id_card TEXT UNIQUE NOT NULL,
                phone TEXT NOT NULL,
                gender TEXT,
                birth_date TEXT,
                created_at DATETIME DEFAULT CURRENT_TIMESTAMP
            )
        `);

        // 创建挂号表
        db.run(`CREATE TABLE IF NOT EXISTS registrations (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            patient_id INTEGER,
            doctor_id INTEGER,
            department_code TEXT,
            registration_date TEXT,
            status TEXT DEFAULT 'pending',
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY(patient_id) REFERENCES patients(id),
            FOREIGN KEY(doctor_id) REFERENCES doctors(id)
        )`);

        // 尝试添加新字段（如果不存在）
        const columnsToAdd = [
            'ALTER TABLE registrations ADD COLUMN symptom TEXT',
            'ALTER TABLE registrations ADD COLUMN diagnosis TEXT',
            'ALTER TABLE registrations ADD COLUMN fee REAL DEFAULT 0',
            'ALTER TABLE registrations ADD COLUMN payment_status TEXT DEFAULT "unpaid"'
        ];

        columnsToAdd.forEach(sql => {
            db.run(sql, (err) => {
                // 忽略 "duplicate column name" 错误
                if (err && !err.message.includes('duplicate column name')) {
                    console.error('添加字段失败:', err.message);
                }
            });
        });

        db.run(`
            CREATE TABLE IF NOT EXISTS medicines (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                name TEXT NOT NULL,
                specification TEXT,
                unit TEXT DEFAULT '盒',
                price REAL NOT NULL,
                stock INTEGER DEFAULT 0,
                category TEXT,
                created_at DATETIME DEFAULT CURRENT_TIMESTAMP
            )
        `);

        db.run(`
            CREATE TABLE IF NOT EXISTS prescriptions (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                registration_id INTEGER NOT NULL,
                patient_id INTEGER NOT NULL,
                doctor_id INTEGER NOT NULL,
                medicine_id INTEGER NOT NULL,
                quantity INTEGER NOT NULL,
                dosage TEXT,
                created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                FOREIGN KEY (registration_id) REFERENCES registrations(id),
                FOREIGN KEY (patient_id) REFERENCES patients(id),
                FOREIGN KEY (doctor_id) REFERENCES doctors(id),
                FOREIGN KEY (medicine_id) REFERENCES medicines(id)
            )
        `);

        db.run(`CREATE INDEX IF NOT EXISTS idx_patients_id_card ON patients(id_card)`);
        db.run(`CREATE INDEX IF NOT EXISTS idx_patients_name ON patients(name)`);
        db.run(`CREATE INDEX IF NOT EXISTS idx_registrations_patient_id ON registrations(patient_id)`);
        db.run(`CREATE INDEX IF NOT EXISTS idx_registrations_date ON registrations(registration_date)`);
        db.run(`CREATE INDEX IF NOT EXISTS idx_medicines_name ON medicines(name)`);
        db.run(`CREATE INDEX IF NOT EXISTS idx_medicines_category ON medicines(category)`);

        console.log('数据库表结构及索引初始化完成');
        if (callback) callback();
    });
}

function insertInitialData() {
    const departments = [
        { name: '内科', code: 'internal' },
        { name: '外科', code: 'surgical' },
        { name: '儿科', code: 'pediatrics' }
    ];

    const doctors = [
        { name: '张三丰', title: '主任医师', department_code: 'internal', available_slots: 10 },
        { name: '李时珍', title: '副主任医师', department_code: 'internal', available_slots: 5 },
        { name: '华佗', title: '主任医师', department_code: 'surgical', available_slots: 8 },
        { name: '扁鹊', title: '主治医师', department_code: 'pediatrics', available_slots: 12 }
    ];

    const medicines = [
        { name: '阿莫西林胶囊', specification: '0.25g*24粒', unit: '盒', price: 15.50, stock: 200, category: '抗生素' },
        { name: '布洛芬片', specification: '0.1g*20片', unit: '盒', price: 8.00, stock: 300, category: '解热镇痛' },
        { name: '感冒灵颗粒', specification: '10g*12袋', unit: '盒', price: 12.00, stock: 150, category: '感冒药' },
        { name: '维生素C片', specification: '0.1g*100片', unit: '瓶', price: 6.50, stock: 500, category: '维生素' },
        { name: '板蓝根颗粒', specification: '10g*20袋', unit: '盒', price: 18.00, stock: 180, category: '中成药' }
    ];

    db.serialize(() => {
        db.run("BEGIN TRANSACTION");

        const insertDept = db.prepare('INSERT OR IGNORE INTO departments (name, code) VALUES (?, ?)');
        departments.forEach(dept => {
            insertDept.run(dept.name, dept.code);
        });
        insertDept.finalize();

        const insertDoc = db.prepare('INSERT OR IGNORE INTO doctors (name, title, department_code, available_slots) VALUES (?, ?, ?, ?)');
        doctors.forEach(doc => {
            insertDoc.run(doc.name, doc.title, doc.department_code, doc.available_slots);
        });
        insertDoc.finalize();

        const insertMed = db.prepare('INSERT OR IGNORE INTO medicines (name, specification, unit, price, stock, category) VALUES (?, ?, ?, ?, ?, ?)');
        medicines.forEach(med => {
            insertMed.run(med.name, med.specification, med.unit, med.price, med.stock, med.category);
        });
        insertMed.finalize();

        db.run("COMMIT", (err) => {
            if (err) console.error("初始化数据提交失败", err);
            else console.log("初始化数据插入完成");
        });
    });
}

async function getDoctorsByDepartment(departmentCode) {
    const sql = `
        SELECT id, name as doctorName, title, available_slots as availableSlots 
        FROM doctors 
        WHERE department_code = ?
    `;
    return query(sql, [departmentCode]);
}

module.exports = {
    db,
    query,
    run,
    get,
    initDatabase,
    insertInitialData,
    getDoctorsByDepartment
};
