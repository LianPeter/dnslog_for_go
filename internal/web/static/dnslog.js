let pollingInterval = null;

function init() {
    bindFormSubmit();
    setupGenerateDomainButton();
    ChangeDNSServer();
    ChangePact();
}

/**
 * 绑定表单提交事件
 * @author william
 * @returns {void}
 * @throws {Error} - 如果表单元素未找到，则抛出错误
 */
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

/**
 * 从服务器获取DNS日志数据并更新UI
 * @author william
 * @param {string} domain - 要查询的域名
 * @returns {void}
 * @throws {Error} - 如果请求失败或返回错误，则抛出错误
 */
function fetchDns(domain) {
    fetch('http://localhost:8080/submit', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ domain_name: domain })
    })
    .then(response => response.json())  // 解析 JSON 响应
    .then(data => {
        const resultDiv = document.getElementById('result');

        if (data.error) {
            resultDiv.innerHTML = `<p style="color:red;">错误: ${data.error}</p>`;
        } else {
            // 初始化表格结构
            let tableHtml = `
                <table border="1" style="border-collapse: collapse; width: 100%; margin-top: 20px;">
                    <thead>
                        <tr style="background-color: #f2f2f2;">
                            <th>域名</th>
                            <th>IP 地址</th>
                            <th>DNS 服务器</th>
                        </tr>
                    </thead>
                    <tbody>
            `;
            
            // 循环遍历每个结果并生成表格行
            data.results.forEach(result => {
                // 访问正确的属性：result.ip 和 result.address
                tableHtml += `
                    <tr>
                        <td>${data.domain}</td>  <!-- 显示域名 -->
                        <td>${result.ip}</td>    <!-- 显示 IP 地址 -->
                        <td>${result.address}</td> <!-- 显示 DNS 服务器地址 -->
                    </tr>
                `;
            });

            // 结束表格标签
            tableHtml += `</tbody></table>`;

            // 将生成的 HTML 插入到 resultDiv 中
            resultDiv.innerHTML = tableHtml;
        }
    })
    .catch(error => {
        console.error('请求失败:', error);
    });
}




/**
 * 设置生成域名按钮的点击事件
 * @author william
 */
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

/**
 * 更改DNS服务器
 * @author william
 * @returns {void}
 * @throws {Error} - 如果请求失败，则抛出错误
 */
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


/**
 * 更改协议
 * @author william
 * @returns {void}
 * @throws {Error} - 如果请求失败，则抛出错误
 */
function ChangePact() {
    document.getElementById("pact")?.addEventListener("change", function () {
        const selectPact = this.value.toLowerCase();  // 将选中的协议转为小写（"udp" 或 "tcp"）
        
        fetch("http://localhost:8080/change-pact", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ pact: selectPact })
        })
        .then(res => res.json())
        .then(msg => {
            if (msg.message) {
                alert(msg.message);
            } else {
                alert("发生错误：" + msg.error);
            }
        })
        .catch(err => console.error("请求失败:", err));
    });
}


window.onload = init;

