// message box 
const mes=document.getElementById("message")
 var xhr= new XMLHttpRequest(); 

function get_infor() {
  
  const pseudo = document.getElementById("username").value;
  
  // create xml http request 
  
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
        console.log("navigate to ")
        // document.location.href="http://localhost:8080/chatroom";
        // document.getElementById("enter_case").style.visibility="visible"// when the pseudo utilisable,you can see the entry button 
      }
    }
  })
}

function exit(){

  const user = document.getElementById("welcome").value;
  
  console.log(user + "are exiting chat room ")
}


function enter_chat(){

  var xhr= new XMLHttpRequest();
  xhr.open("post","/chatroom")

  // console.log("Sending : ", pseudo)
  // xhr.send(pseudo)

  // receiving response

  xhr.addEventListener("readystatechange",()=>{
    if(xhr.readyState==4 && xhr.status == 200){
      var taken = xhr.responseText;
      var code =xhr.responseURL
      console.log("Receiving from server : ", taken,"\n\n")
      mes.innerHTML=taken
    }
  })

}










async function insertText () {

  document.getElementById('infor').innerHTML=''
  document.getElementById('infor').innerHTML='some text '
  sleep(5000)
  
}

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}