function submitForm(){
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

