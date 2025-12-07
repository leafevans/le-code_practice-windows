let currentVars = 0;
let currentCons = 0;

function generateTable() {
    currentVars = parseInt(document.getElementById('numVars').value);
    currentCons = parseInt(document.getElementById('numCons').value);
    const objDiv = document.getElementById('obj-inputs');
    objDiv.innerHTML = '';
    for (let i = 0; i < currentVars; i++) {
        objDiv.innerHTML += `
            <div class="join items-center">
                <input type="number" class="input input-bordered input-sm w-20 join-item obj-coeff" placeholder="0" value="0">
                <span class="btn btn-sm btn-ghost join-item no-animation cursor-default normal-case">x<sub>${i + 1}</sub></span>
            </div>
            ${i < currentVars - 1 ? '<span class="text-gray-400">+</span>' : ''}
        `;
    }
    const consDiv = document.getElementById('cons-inputs');
    consDiv.innerHTML = '';
    for (let i = 0; i < currentCons; i++) {
        let rowHtml = `<div class="flex flex-wrap gap-2 items-center p-2 bg-base-200 rounded-lg">`;
        for (let j = 0; j < currentVars; j++) {
            rowHtml += `
                <input type="number" class="input input-bordered input-sm w-16 cons-coeff-${i}" placeholder="0" value="0">
                <span class="text-sm text-gray-600">x<sub>${j + 1}</sub></span>
                ${j < currentVars - 1 ? '<span class="text-gray-400">+</span>' : ''}
            `;
        }
        rowHtml += `
            <select class="select select-bordered select-sm w-20 mx-2 cons-op-${i}">
                <option value="<=">≤</option>
                <option value=">=">≥</option>
                <option value="=">=</option>
            </select>
            <input type="number" class="input input-bordered input-sm w-20 cons-rhs-${i}" placeholder="RHS" value="0">
        </div>`;
        consDiv.innerHTML += rowHtml;
    }
    document.getElementById('input-area').classList.remove('hidden');
    document.querySelectorAll('input, select').forEach(el => {
        el.addEventListener('input', updatePreview);
    });
    updatePreview();
}

function updatePreview() {
    const type = document.getElementById('objType').value;
    let latex = `\\begin{aligned} & \\text{${type}} \\quad Z = `;
    const objCoeffs = Array.from(document.querySelectorAll('.obj-coeff')).map(i => i.value);
    const objParts = objCoeffs.map((c, i) => c != 0 ? `${c}x_{${i + 1}}` : '').filter(s => s);
    latex += (objParts.length ? objParts.join(' + ') : '0') + ' \\\\ & \\text{s.t.} \\\\';
    for (let i = 0; i < currentCons; i++) {
        const coeffs = Array.from(document.querySelectorAll(`.cons-coeff-${i}`)).map(input => input.value);
        const op = document.querySelector(`.cons-op-${i}`).value;
        const rhs = document.querySelector(`.cons-rhs-${i}`).value;

        const lhsParts = coeffs.map((c, j) => c != 0 ? `${c}x_{${j + 1}}` : '').filter(s => s);
        const lhs = lhsParts.length ? lhsParts.join(' + ') : '0';

        let opTex = op === '<=' ? '\\leq' : (op === '>=' ? '\\geq' : '=');
        latex += ` & \\quad ${lhs} ${opTex} ${rhs} \\\\`;
    }
    latex += ' & \\quad x_i \\geq 0 \\end{aligned}';
    const previewDiv = document.getElementById('math-preview');
    previewDiv.innerHTML = `\\[${latex}\\]`;
    if (window.MathJax) MathJax.typesetPromise([previewDiv]);
}

async function solve() {
    const model = {
        description: "Electron Client Problem",
        objectiveType: document.getElementById('objType').value,
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
            alert('计算失败: ' + response.statusText);
        }
    } catch (error) {
        alert('连接错误: 请确保后端服务已启动。\n' + error.message);
    }
}
function displayResult(result) {
    const content = document.getElementById('result-content');
    if (result.isOptimal) {
        let html = `
            <div class="alert alert-success">
                <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
                <span>找到最优解！</span>
            </div>
            <div class="stats shadow w-full">
                <div class="stat">
                    <div class="stat-title">目标函数值 (Z)</div>
                    <div class="stat-value text-primary">${result.maxValue}</div>
                </div>
            </div>
            <div class="divider">变量取值</div>
            <div class="grid grid-cols-2 gap-2">
        `;
        for (const [key, val] of Object.entries(result.variableValues)) {
            html += `
                <div class="flex justify-between p-3 bg-base-200 rounded">
                    <span class="font-bold">${key}</span>
                    <span class="font-mono">${val}</span>
                </div>`;
        }
        html += `</div>`;
        content.innerHTML = html;
    } else {
        content.innerHTML = `
            <div class="alert alert-error">
                <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
                <span>${result.message}</span>
            </div>`;
    }
    document.getElementById('result_modal').showModal();
}
