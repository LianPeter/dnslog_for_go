window.onload = function () {
    document.getElementById('dnslog-form').addEventListener('submit', function (event) {
        event.preventDefault();
        const domain = document.getElementById('domain_name').value;

        // 调试：查看传递的域名
        console.log(domain);

        fetch(`http://localhost:8080/log_dns?domain=${domain}`)
            .then(response => response.text())
            .then(data => {
                document.getElementById('result').innerText = `查询结果: ${data}`;
            })
            .catch(error => {
                console.error('Error:', error);
            });
    });
};