const express = require('express');
const cors = require('cors');
const path = require('path');

const app = express();
const PORT = 3000;

app.use(cors());

const allDoctors = {
    internal: [
        { id: 1, doctorName: '张三丰', title: '主任医师', availableSlots: 10 },
        { id: 2, doctorName: '李时珍', title: '副主任医师', availableSlots: 5 },
    ],
    surgical: [
        { id: 3, doctorName: '华佗', title: '主任医师', availableSlots: 8 },
    ],
    pediatrics: [
        { id: 4, doctorName: '扁鹊', title: '主治医师', availableSlots: 12 },
    ],
};

app.get('/api/doctors', (req, res) => {
    const department = req.query.department;

    if (!department) {
        return res.status(400).json({ error: '缺少科室参数 (department)' });
    }

    const doctors = allDoctors[department] || [];
    res.json(doctors);
});

// 托管静态文件
app.use(express.static(path.join(__dirname, '../hospital-client/dist')));

app.use((req, res) => {
    res.sendFile(path.join(__dirname, '../hospital-client/dist/index.html'));
});

app.listen(PORT, () => {
    console.log(`项目已启动,请在浏览器中访问 http://localhost:${PORT}`);
});
