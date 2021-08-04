console.log("JS loaded")
    
const url="http://127.0.0.1:5555"

var input_per=document.getElementById("input")

input_per.addEventListener("submit", (e)=>{

    //anti auto submission 

    e.preventDefault()

    const Form_data= new FormData(input_per)

    fetch(url, {

        method:"POST",
        body:Form_data,

    }).then(
        Response => Response.text() // takes a Response stream and reads it to completion

        // => is arrow function 
    ).then(
        (data)=> {console.log(data);document.getElementById(message_box).innerHTML=data}// inner_html: pour examiner la source HTML actuelle de la page
    ).catch(
        error => console.error(error)
    )



})