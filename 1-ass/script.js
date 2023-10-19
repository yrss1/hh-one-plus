async function getFetch() {
    const res = await fetch('https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1');
    const data = await res.json();
    return data;
}

async function main() {
    const list = await getFetch();
    list.forEach(i => {
        const table = document.getElementById('main');
        const row = document.createElement('tr');
        row.innerHTML = `
        <td>${i.id}</td>
        <td>${i.symbol}</td>
        <td>${i.name}</td>
        `;

        if (i.symbol === 'usdt') {
            row.classList.add('green-bg');
        } else if (table.rows.length <= 6) {
            row.classList.add('blue-bg');
        }
        table.appendChild(row);
    });
}

main();
