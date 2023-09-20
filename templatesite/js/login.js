function show_hide() {
    var login = document.getElementById("container1");
    var signup = document.getElementById("container2");
    var copyright = document.getElementById("copyright");
  
    if (login.style.display === "none") {
        login.style.display = "block";  //lonin出現
        document.getElementById("username").value="";
        document.getElementById("password").value="";
        signup.style.display = "none";  //signup消失
        copyright.style.margin = "200px 0px 0px 0px";
    } else {
        login.style.display = "none";   //login消失
        signup.style.display = "block"; //signup出現
        signup.style.visibility="visible";
        copyright.style.margin = "200px 0px 0px 0px";
     
        document.getElementById("fullname").value="";
        document.getElementById("username2").value="";
        document.getElementById("password2").value="";
        document.getElementById("comfirm_password").value="#/page1";
    }
}
function login(url) {
  var username = document.getElementById("username").value;
  var pass = document.getElementById("password").value;
  
  fetch(url, {
      method: "POST",
      credentials: "include",
      body: JSON.stringify({
          user_id: username,
          pass: pass
          
      }),
      
  })
  .then((response) => {
      if (response.status == 200) {
          alert("Success");
          location.href="./form.html"
      } else {
          // Handle other status codes if needed
          // You can also check response.statusText for a textual description of the status code.
          console.log("Login failed. Status code: " + response.status);
      }
  })
  .catch((error) => {
      console.error("Error:", error);
  });
}

function deleteCookie(name) {
    // Set the cookie's expiration date to a past date
    document.cookie = name + "=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
}

function logout(url){
    fetch(url, {
        method: "POST",
        credentials: "include",
    }).then((response) => {
        if (response.status == 200) {
            deleteCookie("dev-cookie")
            alert("Logout");
            location.href="./index.html"
            
        } else {
            // Handle other status codes if needed
            // You can also check response.statusText for a textual description of the status code.
            console.log(response.status);
        }
    })
}
function register(){
    var user_id = document.getElementById("user_id").value;
    var pass = document.getElementById("pass").value;
    var user_name = document.getElementById("user_name").value;
    var mobile = document.getElementById("mobile").value;
    var email = document.getElementById("email").value;
    var bureau = document.getElementById("bureau").value;
    fetch('/api/register',{
        method: "POST",
      credentials: "include",
      body: JSON.stringify({
        user_id: user_id,
        pass: pass,
        user_name: user_name  ,
        mobile: mobile,
        email: email,
        bureau: bureau,
          
      }),
    }).then(response=>{
        if(response.OK){
            alert("註冊成功！")
        }else{
            alert("註冊失敗")
        }
    })
}



function insertInfo(key,element){
    var heading = document.createElement("h3");
    heading.textContent = key+" : "+element;
    var body = document.body;
    body.insertBefore(heading, body.firstChild);
}

