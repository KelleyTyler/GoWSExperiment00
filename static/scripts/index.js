

window.onload=function(){
    //console.log("ONLOAD")
    const path = window.location.pathname.split("/");
    switch(path[1]){
        case "":{
            loadPage("home")
            break;
        }
        case "index":{
            loadPage("home")
            console.log("LOADED INDEX")
            break;
        }
        case "static":{
            loadPage("home")
            console.log("LOADED INDEX HTML")
            break;
        }
        case "home":{
            loadPage("home")
            break;
        }
        case "about":{
            loadPage("about")
            break;
        }
        case "showcase":{
            loadPage("showcase")
            break;
        }
        case "Post":{
            loadPage("Post")
            break;
        }
        case "Posts":{
            loadPage("Posts");
            // PostsOnLoad();
            break;
        }
        default:{
            console.log("LOADED"+path[1])
            loadPage("404")
            break;
        }
    };
    //console.log("ONLOAD 2");
    document.querySelectorAll(".menu_Item").forEach((item)=>{
        //console.log("ITEMS: "+item)
        item.addEventListener("click", function(){
            const path = item.getAttribute("value");
            loadPage(path);
            console.log("HEY "+path)
            if(path==""){
                window.history.pushState("","","/");
                return;
            }
            window.history.pushState("","",path);
        });
        item.addEventListener("DOMContentLoaded",function() {
            alert("aaah");
        })
    });
    //console.log("ONLOAD 3");
} 
function loadPage($path){
    if($path=="") return;

    
    //---- WHERE FETCH  OR PROMISES OR WHATEVER
    //CURRENTLY USING A SIMPLE HTML 
    const container =document.getElementById("container")
    const request = new XMLHttpRequest();
    request.open("GET","./"+$path+".html");
    request.send();
    request.onload=function()
    {
        if(request.status==200){
            container.innerHTML=request.responseText;
            document.title=$path;
            if($path=="Posts"){
                PostsOnLoad()
            };
        }
    }
}
function OnLoad(){
    nums =0;
    isOpen0=false;
    //cv.innerHTML=BUTTON_OPEN("HELLO THERE HELLO THERE HELLO THERE HELLO THERE!",)
}
const canvasID = document.getElementById("Big_Canvas")
const ctx = canvasID.getContext("2d");
//const dyna =document.getElementById("dyna")
//const cv = document.getElementById("CVee")
var nums;
var isOpen0;

function DrawThing(){
const canvasID = document.getElementById("Big_Canvas")
const ctx = canvasID.getContext("2d");
divisor=5;
    // Set line width
ctx.lineWidth = 10/divisor;

// Wall
width = canvasID.clientWidth
height= canvasID.clientHeight
// ctx.strokeRect(75/divisor, 140/divisor, 150/divisor, 110/divisor);

// // Door
// ctx.fillRect(130/divisor, 190/divisor, 40/divisor, 60/divisor);

// // Roof
// ctx.beginPath();
// ctx.moveTo(50/divisor, 140/divisor);
// ctx.lineTo(150/divisor, 60/divisor);
// ctx.lineTo(250/divisor, 140/divisor);
// ctx.closePath();
// ctx.stroke();
var t= 0;
window.requestAnimationFrame(function loop() {
    ctx.clearRect(0,0,width,height)
    ctx.strokeRect((75/divisor)+t, 140/divisor, (150/divisor), 110/divisor);

    // Door
    ctx.fillRect((130/divisor)+t, 190/divisor, (40/divisor), 60/divisor);
    
    // Roof
    ctx.beginPath();
    ctx.moveTo((50/divisor)+t, 140/divisor);
    ctx.lineTo((150/divisor)+t, 60/divisor);
    ctx.lineTo((250/divisor)+t, 140/divisor);
    ctx.closePath();
    ctx.stroke();
    t++;
    if(t<200){
        window.requestAnimationFrame(loop)
    }
})
nums++;
}

function PostsOnLoad(){
    console.log("POSTSONLOAD");
    var userData = null;
    var test = "TEST TEST TEST";
    // fetch('/Posts')
    // .then(response => response.json())
    // .then(data => {
    //     try{
    //         userData= JSON.parse(data);
    //     }catch (e){
    //         userData=data;
    //     }
    //     userData.forEach((item)=>{
    //         // console.log("ID:"+JSON.stringify(item.ID)+" BODY"+JSON.stringify(item.BODY))

    //     })
    //     console.log("USRDATA:"+userData)
    // })
    // .catch(error => console.error('Error:', error))
    postslists = document.getElementById("POST_LISTS")
    if(postslists===null){
        console.log("no postsLists")
    }else{
        //foreach
        // data.forEach((item)=>{
        //     console.log("ITEM"+item.Body)
        //     // postslists.innerHTML+="<li>"+test+"</li>";
        // })
        // postslists.innerHTML+="<li>"+test+"</li>";
        fetch('/Posts')
        .then(response => response.json())
        .then(data => {
            try{
                userData= JSON.parse(data);
            }catch (e){
                userData=data;
            }
            userData.forEach((item)=>{
                // console.log("ID:"+JSON.stringify(item.ID)+" BODY"+JSON.stringify(item.BODY))
                postslists.innerHTML+="<li>"+JSON.stringify(item.BODY)+"</li>";
            })
            //console.log("USRDATA:"+userData)
        })
        .catch(error => console.error('Error:', error))
    }
    //var postlosts=null;
    // if(document.readyState==="loading"){
    //     document.addEventListener("DOMContentLoaded", console.log("LOADING"));
    // }
    // else {
    //   // `DOMContentLoaded` has already fired
    //   console.log("ELSE "+document.readyState);
    // }
}

function submitMyForm(){
    console.log("SENDING DATA!!!")
    var xhr = new XMLHttpRequest();
    xhr.open(form.method, form.action, true);
    xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    var j = {
        "Body":form.address.value
    };
    //console.log("-----"+form.name+"VALUE:"+form.address.value+"----"+j.Body.value);
    xhr.send(JSON.stringify(j));
}

