function render(pageName) {
    fetch(pageName ,{credentials: "include",})
        .then(function (response) {
            if (response.status != 200) {
                alert("token expire")
                location.href = "./index.html"
                return
            } 
            else{

                return response.text();
            }
        })
        .then(function (data) { 
            console.log(data)
            document.getElementById('root').innerHTML = data; // 在root裡塞請求拿到的html
            const newScript = document.createElement('script'); // 建立新的script元素
            newScript.innerHTML = document.getElementById('root').querySelector('script').innerHTML; // 將請求拿到的html的script內容抓出來賦值給新建的script元素
            document.body.appendChild(newScript); // 執行新建的script元素
            document.body.removeChild(newScript);// 執行完就移除新建的script元素
        });
}
