function submitRepair() {
    const car = document.getElementById('repairCar').value;
    const problem = document.getElementById('repairProblem').value.trim();
    const priority = document.getElementById('repairPriority').value;
    const mechanic = document.getElementById('repairMechanic').value;
    const msg = document.getElementById('formMessage');

    if (!problem) {
        msg.textContent = 'Опишите неисправность.';
        msg.style.color = '#ffaa00';
        return;
    }

    const data = { car, problem, priority, mechanic };

    fetch('/api/repair', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
    })
    .then(function(res) { return res.json(); })
    .then(function(res) {
        if (res.status === 'ok') {
            msg.textContent = 'Заявка успешно отправлена!';
            msg.style.color = '#4caf50';
            document.getElementById('repairProblem').value = '';

            //localStorage
            let locCount = parseInt(localStorage.getItem('repair_count') || '0') + 1;
            localStorage.setItem('repair_count', locCount);
            localStorage.setItem('car_' + locCount, car);
            localStorage.setItem('problem_' + locCount, problem);
            localStorage.setItem('priority_' + locCount, priority);
            localStorage.setItem('mechanic_' + locCount, mechanic);

            //sessionStorage
            let sesCount = parseInt(sessionStorage.getItem('repair_count') || '0') + 1;
            sessionStorage.setItem('repair_count', sesCount);
            sessionStorage.setItem('car_' + sesCount, car);
            sessionStorage.setItem('problem_' + sesCount, problem);
            sessionStorage.setItem('priority_' + sesCount, priority);
            sessionStorage.setItem('mechanic_' + sesCount, mechanic);
        }
    })
    .catch(function() {
        msg.textContent = 'Ошибка при отправке';
        msg.style.color = '#ff4444';
    });
}

function showRequests() {
    const block = document.getElementById('requestsBlock');
    const isVisible = block.style.display === 'block';

    if (isVisible) {
        block.style.display = 'none';
        return;
    }

    const priorityLabel = { low: 'Низкий', medium: 'Средний', high: 'Высокий' };

    function buildTable(storage, title) {
        const count = parseInt(storage.getItem('repair_count') || '0');
        if (count === 0) {
            return '<h3 style="color: var(--accent); margin-top: 20px;">' + title + '</h3>'
                 + '<p style="color: var(--text-muted);">Нет заявок.</p>';
        }

        let html = '<h3 style="color: var(--accent); margin-top: 20px;">' + title + '</h3>';
        html += '<table class="main-table"><tr class="table-header">';
        html += '<th>№</th><th>Автомобиль</th><th>Неисправность</th><th>Приоритет</th><th>Механик</th></tr>';

        for (let i = 1; i <= count; i++) {
            const car      = storage.getItem('car_' + i);
            const problem  = storage.getItem('problem_' + i);
            const priority = storage.getItem('priority_' + i);
            const mechanic = storage.getItem('mechanic_' + i);
            html += '<tr>';
            html += '<td>' + i + '</td>';
            html += '<td>' + car + '</td>';
            html += '<td>' + problem + '</td>';
            html += '<td>' + (priorityLabel[priority] || priority) + '</td>';
            html += '<td>' + mechanic + '</td>';
            html += '</tr>';
        }
        html += '</table>';
        return html;
    }

    block.innerHTML =
        buildTable(localStorage,   'Все заявки') +
        buildTable(sessionStorage, 'Заявки текущей сессии');

    block.style.display = 'block';
}