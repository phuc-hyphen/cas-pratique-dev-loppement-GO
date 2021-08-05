// message box 
const mes=document.getElementById("message")
var user

function get_infor() {
  
  const pseudo = document.getElementById("username").value;
  user=pseudo
  
  // create xml http request 
  var xhr= new XMLHttpRequest(); 
  xhr.open("post","/api/check")

  // sending request 
  console.log("Sending : ", pseudo)
  xhr.send(pseudo)

  // receiving response

  xhr.addEventListener("readystatechange",()=>{
    if(xhr.readyState==4 && xhr.status == 200){
      var taken = xhr.responseText;
      console.log("Receiving from server : ", taken,"\n\n")
      mes.innerHTML=taken
      if(taken==" Pseudo name can be used "){
        console.log("navigate to chat room as pseudo name " + pseudo)
        mes.innerText="Navigating to chat room ... "

        redirect()
      }
    }
  })
}

function exit(){

  const user = document.getElementById("welcome").value;
  
  console.log(user + "are exiting chat room ")
}

function redirect(){
  var xhr = new XMLHttpRequest();
  xhr.open("POST","/")
  xhr.send(user)
  xhr.onreadystatechange = function () {
    if (xhr.readyState === 4) {
      console.log("redirect successful")
    }};
}










async function insertText () {

  document.getElementById('infor').innerHTML=''
  document.getElementById('infor').innerHTML='some text '
  sleep(5000)
  
}

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}