let pollingInterval = null;

window.onload = function () {
    const form = document.getElementById('dnslog-form');

    form.addEventListener('submit', function (event) {
        event.preventDefault();

        const domain = document.getElementById('domain_name').value;
        console.log("开始轮询 DNS:", domain);

        // 停止已有轮询
        if (pollingInterval) clearInterval(pollingInterval);

        // 立刻查询一次
        fetchDns(domain);

        // 每2秒查询一次
        pollingInterval = setInterval(() => {
            fetchDns(domain);
        }, 2000);
    });
};

function fetchDns(domain) {
    fetch('http://localhost:8080/submit', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ domain_name: domain })
    })
    .then(response => response.json())
    .then(data => {
        const resultDiv = document.getElementById('result');

        if (data.error) {
            resultDiv.innerHTML = `<p style="color:red;">错误: ${data.error}</p>`;
        } else {
            resultDiv.innerHTML = `
                <table border="1" style="border-collapse: collapse; width: 100%; margin-top: 20px;">
                    <thead>
                        <tr style="background-color: #f2f2f2;">
                            <th>域名</th>
                            <th>IP 地址</th>
                            <th>DNS 服务器</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>${data.domain}</td>
                            <td>${data.ip}</td>
                            <td>${data.address}</td>
                        </tr>
                    </tbody>
                </table>
            `;
        }
    })
    .catch(error => {
        console.error('请求失败:', error);
    });
}
