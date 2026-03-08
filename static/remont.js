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
        }
    })
    .catch(function() {
        msg.textContent = 'Ошибка при отправке';
        msg.style.color = '#ff4444';
    });
}