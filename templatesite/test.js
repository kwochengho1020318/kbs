
const form = document.getElementById('test-form')
form.addEventListener('submit', function (e) {
    e.preventDefault();
    const formData = utils.getFormValue(this);
    alert(JSON.stringify(formData));
    fetch("./api/common/DevUser",{
      method: "POST",
      credentials: "include",
      body: JSON.stringify(formData),
      headers: new Headers({
        'Content-Type': 'application/json'
    })
  }).then((response) => {
    if (response.status == 200) {
        alert("Insert Success");
    } else {
        // Handle other status codes if needed
        // You can also check response.statusText for a textual description of the status code.
        console.log("Insert failed. Status code: " + response.status);
    }
})
.catch((error) => {
    console.error("Error:", error);
})
    // 發請求儲存資料**********
    return false;
  });
  
  
const table = document.getElementById('data-table');
fetch("./api/common/DevUser",{
      method: "GET",
      credentials: "include",
    }).then(response=>response.json())
  .then(data=>{
        utils.setTableValue(table, data);
	new DataTable(`#${table.id}`, {
        pageResize: true
});
});
