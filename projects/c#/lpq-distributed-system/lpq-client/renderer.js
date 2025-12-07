let currentVars = 2;
let currentCons = 2;

window.onload = () => {
    generateTable();
};

function generateTable() {
    currentVars = parseInt(document.getElementById('numVars').value);
    currentCons = parseInt(document.getElementById('numCons').value);
    const thead = document.getElementById('table-header');
    const tbody = document.getElementById('table-body');
    let headerHtml = `
        <tr>
            <th rowspan="2" class="w-24 bg-gray-100 border border-gray-400">资源</th>
            <th colspan="${currentVars}" class="bg-white text-center border border-gray-400 py-1">单位竞争性活动</th>
            <th rowspan="2" class="w-32 bg-gray-100 border border-gray-400">资源<br>限制</th>
        </tr>
        <tr>
    `;
    for (let i = 0; i < currentVars; i++) {
        headerHtml += `
            <th class="p-0 border border-gray-400 min-w-[80px] align-top">
                <div class="flex flex-col h-full">
                    <!-- 上半部分：活动名称输入 -->
                    <input type="text" 
                           class="w-full text-center text-xs py-2 bg-white outline-none border-b border-gray-300 text-gray-800 placeholder-gray-300" 
                           value="活动${i + 1}" 
                           placeholder="名称">
                    <!-- 下半部分：变量名 -->
                    <div class="bg-gray-50 py-1 text-gray-600 font-normal">
                        x<sub>${i + 1}</sub>
                    </div>
                </div>
            </th>`;
    }
    headerHtml += `</tr>`;
    const table = document.querySelector('.table');
    table.className = 'eng-table w-full text-sm border-collapse';
    table.innerHTML = `<thead>${headerHtml}</thead><tbody id="table-body"></tbody>`;
    const newTbody = document.getElementById('table-body');
    let bodyHtml = '';
    for (let i = 0; i < currentCons; i++) {
        bodyHtml += `<tr>`;
        bodyHtml += `<td class="border border-gray-400"><input type="text" value="资源 ${String.fromCharCode(65 + i)}" class="font-bold text-gray-700 w-full text-center bg-transparent outline-none py-1"></td>`;
        for (let j = 0; j < currentVars; j++) {
            bodyHtml += `<td class="border border-gray-400"><input type="number" class="cons-coeff-${i} w-full text-center outline-none h-full py-1" data-row="${i}" data-col="${j}" value="" placeholder="0"></td>`;
        }
        bodyHtml += `
            <td class="border border-gray-400 p-0">
                <div class="flex h-full w-full items-center">
                    <select class="bg-gray-50 text-xs px-1 border-r border-gray-300 h-full cons-op-${i} outline-none cursor-pointer appearance-none text-center" style="width: 40px;">
                        <option value="<=">≤</option>
                        <option value=">=">≥</option>
                        <option value="=">=</option>
                    </select>
                    <input type="number" class="cons-rhs-${i} flex-1 h-full text-center outline-none" value="" placeholder="0">
                </div>
            </td>
        </tr>`;
    }
    bodyHtml += `<tr>`;
    bodyHtml += `<td class="bg-gray-50 font-bold border border-gray-400 text-center">单位<br>利润</td>`;
    for (let j = 0; j < currentVars; j++) {
        bodyHtml += `<td class="border border-gray-400"><input type="number" class="obj-coeff font-bold text-blue-600 w-full text-center outline-none h-full py-1" value="" placeholder="0"></td>`;
    }
    bodyHtml += `<td class="bg-gray-50 text-xs text-center text-gray-500 border border-gray-400">「参数均可编辑」</td>`;
    bodyHtml += `</tr>`;
    newTbody.innerHTML = bodyHtml;
    document.querySelectorAll('input, select').forEach(el => {
        el.addEventListener('input', updatePreview);
    });
    updatePreview();
}

function updatePreview() {
    const type = document.querySelector('input[name="objType"]:checked').value;
    const formatTerm = (val, idx) => {
        const c = parseFloat(val);
        if (c === 0 || isNaN(c)) return null;
        if (c === 1) return `x_{${idx + 1}}`;
        if (c === -1) return `-x_{${idx + 1}}`;
        return `${c}x_{${idx + 1}}`;
    };
    let latex = `\\begin{aligned} & \\text{${type.toLowerCase()}} \\, Z = `;
    const objCoeffs = Array.from(document.querySelectorAll('.obj-coeff')).map(i => i.value);
    const objParts = objCoeffs.map((c, i) => formatTerm(c, i)).filter(s => s !== null);
    latex += (objParts.length ? objParts.join(' + ') : '0') + ' \\\\';
    let constraintsLatex = '';
    let hasConstraints = false;
    for (let i = 0; i < currentCons; i++) {
        const coeffs = Array.from(document.querySelectorAll(`.cons-coeff-${i}`)).map(input => input.value);
        const op = document.querySelector(`.cons-op-${i}`).value;
        let rhs = document.querySelector(`.cons-rhs-${i}`).value;
        if (rhs === '') rhs = '0';
        const lhsParts = coeffs.map((c, j) => formatTerm(c, j)).filter(s => s !== null);
        if (lhsParts.length === 0 && parseFloat(rhs) === 0) {
            continue;
        }
        const lhs = lhsParts.length ? lhsParts.join(' + ') : '0';
        let opTex = op === '<=' ? '\\leq' : (op === '>=' ? '\\geq' : '=');
        constraintsLatex += ` & ${lhs} ${opTex} ${rhs} \\\\`;
        hasConstraints = true;
    }

    if (hasConstraints) {
        latex += ' & \\text{s.t.} \\quad \\left\\{ \\begin{aligned}';
        latex += constraintsLatex;
        latex += ' & x_i \\geq 0 \\end{aligned} \\right.';
    } else {
        latex += ' & \\text{s.t.} \\quad x_i \\geq 0';
    }
    latex += ' \\end{aligned}';
    const previewDiv = document.getElementById('math-preview');
    previewDiv.innerHTML = `\\[${latex}\\]`;
    if (window.MathJax) MathJax.typesetPromise([previewDiv]);
}

async function solve() {
    const resultPanel = document.getElementById('result-panel');
    resultPanel.innerHTML = '<div class="loading loading-spinner loading-lg text-primary mx-auto block"></div>';
    const model = {
        description: "Electron Table Client",
        objectiveType: document.querySelector('input[name="objType"]:checked').value,
        variables: [],
        constraints: []
    };
    const objInputs = document.querySelectorAll('.obj-coeff');
    objInputs.forEach((input, index) => {
        model.variables.push({
            name: `x${index + 1}`,
            coefficient: parseFloat(input.value) || 0
        });
    });
    for (let i = 0; i < currentCons; i++) {
        const coeffs = {};
        const coeffInputs = document.querySelectorAll(`.cons-coeff-${i}`);
        coeffInputs.forEach((input, index) => {
            const val = parseFloat(input.value);
            if (val !== 0) coeffs[`x${index + 1}`] = val;
        });
        model.constraints.push({
            coefficients: coeffs,
            operator: document.querySelector(`.cons-op-${i}`).value,
            rightHandSide: parseFloat(document.querySelector(`.cons-rhs-${i}`).value) || 0
        });
    }
    try {
        const response = await fetch('http://localhost:5278/api/Solver/solve', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(model)
        });
        if (response.ok) {
            const result = await response.json();
            displayResult(result);
        } else {
            resultPanel.innerHTML = `<div class="alert alert-error">计算失败: ${response.statusText}</div>`;
        }
    } catch (error) {
        resultPanel.innerHTML = `<div class="alert alert-error text-xs">连接错误: 请检查后端服务<br>${error.message}</div>`;
    }
}

function displayResult(result) {
    const panel = document.getElementById('result-panel');
    panel.classList.remove('items-center', 'justify-center');

    if (result.isOptimal) {
        let html = `
            <div class="bg-green-50 border border-green-200 p-4 rounded-md mb-4 shadow-sm">
                <div class="flex items-center gap-2 mb-2">
                    <!-- 使用 RemixIcon 成功图标 -->
                    <i class="ri-checkbox-circle-fill text-green-600 text-lg"></i>
                    <div class="font-bold text-green-800 text-sm">求解成功</div>
                </div>
                <div id="res-z" class="text-xl font-mono text-center py-1"></div>
            </div>
            <div class="space-y-3">
                <div class="text-xs font-bold text-gray-500 uppercase tracking-wider border-b pb-1">决策变量最优取值</div>
        `;
        let varIndex = 0;
        for (const [key, val] of Object.entries(result.variableValues)) {
            const texKey = key.replace(/(\d+)/g, '_{$1}');
            html += `
                <div class="flex justify-between items-center text-sm border-b border-gray-100 py-2">
                    <div id="res-var-${varIndex}" data-tex="${texKey} = ${val}"></div>
                </div>`;
            varIndex++;
        }
        html += `</div>`;
        panel.innerHTML = html;

        const zDiv = document.getElementById('res-z');
        zDiv.innerHTML = `\\[ Z = ${result.maxValue} \\]`;
        const promises = [MathJax.typesetPromise([zDiv])];
        varIndex = 0;
        for (const [key, val] of Object.entries(result.variableValues)) {
            const vDiv = document.getElementById(`res-var-${varIndex}`);
            const tex = vDiv.getAttribute('data-tex');
            vDiv.innerHTML = `\\( ${tex} \\)`;
            promises.push(MathJax.typesetPromise([vDiv]));
            varIndex++;
        }
        Promise.all(promises);

    } else {
        panel.innerHTML = `
            <div class="bg-red-50 border border-red-200 p-4 rounded-md shadow-sm">
                <div class="flex items-center gap-2 mb-2">
                    <!-- 使用 RemixIcon 失败图标 -->
                    <i class="ri-close-circle-fill text-red-600 text-lg"></i>
                    <div class="font-bold text-red-800 text-sm">求解失败</div>
                </div>
                <div class="text-xs text-red-600 leading-relaxed">${result.message}</div>
            </div>`;
    }
}
