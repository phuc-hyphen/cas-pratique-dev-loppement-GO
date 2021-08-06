// message box
const mes = document.getElementById("message");
var user

function get_infor() {
  const pseudo = document.getElementById("username").value;
  user=pseudo

  // create xml http request
  var xhr = new XMLHttpRequest();
  xhr.open("post", "/api/check");

  // sending request
  console.log("Sending : ", pseudo);
  xhr.send(pseudo);

  // receiving response
  xhr.addEventListener("readystatechange", () => {
    if (xhr.readyState == 4 && xhr.status == 200) {
      var taken = xhr.responseText;
      console.log("Receiving from server : ", taken, "\n");
      mes.innerHTML = taken;

      if (taken == " Pseudo name can be used ") {
        console.log("navigate to chat room as pseudo name " + pseudo);
        mes.innerText = "Navigating to chat room ... ";
        redirect();
      }
    }
  });
}

function redirect() {
  var xhr = new XMLHttpRequest();
  xhr.open("POST", "/");
  xhr.send(user)
}

async function insertText() {
  var text=document.getElementById("usermsg").innerHTML = "";
  sleep(5000);
}

function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}
