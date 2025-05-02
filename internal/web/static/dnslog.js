let pollingInterval = null;

function init() {
    bindFormSubmit();
    setupGenerateDomainButton();
    ChangeDNSServer();
}

function bindFormSubmit() {
    const form = document.getElementById('dnslog-form');
    form?.addEventListener('submit', function (event) {
        event.preventDefault();
        const domain = document.getElementById('domain_name').value;

        // 停止之前的轮询
        if (pollingInterval) clearInterval(pollingInterval);

        // 开始新的轮询
        fetchDns(domain);

        pollingInterval = setInterval(() => {
            fetchDns(domain);
        }, 2000);// 每2秒请求一次
    });
}

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

function setupGenerateDomainButton() {
    const generateButton = document.getElementById('generate-domain-btn');
    generateButton?.addEventListener('click', function (event) {
        event.preventDefault();

        fetch('http://localhost:8080/random-domain', {
            method: 'POST'
        })
            .then(response => response.json())
            .then(data => {
                const domainInput = document.getElementById('domain_name');
                domainInput.value = data.domain;
                fetchDns(data.domain);
            })
            .catch(error => {
                console.error('Error fetching domain:', error);
            });
    });
}

function ChangeDNSServer() {
    document.getElementById("dns-select")?.addEventListener("change", function () {
        const selectedValue = this.value;
        fetch("http://localhost:8080/change", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ num: parseInt(selectedValue) })
        })
            .then(res => res.text())
            .then(msg => alert(msg))
            .catch(err => console.error("请求失败:", err));
    })
}

window.onload = init;

