function CheckPassword(){
    let Confirmation_Password = document.getElementById("Confirmation_Password").value
    let Password = document.getElementById("password").value
    let min = false;
    let maj = false;
    let num = false;
    let spe = false;
    let len = false;
    let pw = false;
    let empty = false;

    let clickable = document.getElementById("clickable");
    let error = document.getElementById("error");

    const LowerExp = /[a-z]/g;
    let LowerMatch = Password.match(LowerExp);
    const UpperExp = /[A-Z]/g;
    let UpperMatch = Password.match(UpperExp);
    const NumExp = /[0-9]/g;
    let NumMatch = Password.match(NumExp);
    const SpeExp = /[!@#$%^&*]/g;
    let SpeMatch = Password.match(SpeExp);

    if (LowerMatch !== null){
        min = true;
    } else {
        min = false;
    }
    if (UpperMatch !== null) {
        maj = true;
    } else {
        maj = false;
    }
    if (NumMatch !== null) {
        num = true;
    } else {
        num = false;
    }
    if (SpeMatch !== null) {
        spe = true;
    } else {
        spe = false;
    }
    if (Password.length >= 8) {
        len = true;
    } else {
        len = false;
    }
    if (Password !== Confirmation_Password) {
        pw = false;
    } else {
        pw = true;
    }
    if ((Password === null && Confirmation_Password === null) || (Password === "" && Confirmation_Password === "")) {
        empty = true;
    }

    if (min === true && maj === true && num === true && spe === true && len === true && pw === true || empty === true) {
        clickable.removeAttribute("disabled");
        error.classList.add("valid");
        error.classList.remove("invalid");
        error.innerHTML = "Valid Password";
    } else if (min === true && maj === true && num === true && spe === true && len === true && pw !== true) {
        let str = "Confirmation Password different from Password";
        clickable.setAttribute("disabled","");
        error.classList.add("invalid");
        error.classList.remove("valid");
        error.innerHTML = str;
    } else {
        let str = "Password must contain";
        clickable.setAttribute("disabled","");
        error.classList.add("invalid");
        error.classList.remove("valid");
        if (min === true && maj === true && num === true && spe === true && len !== true) {
            error.innerHTML = "Password is too short";
        } else {
            if (min !== true) {
                str += " lowercase letter";
            }
            if (maj !== true) {
                str += " uppercase letter";
            }
            if (num !== true) {
                str += " number";
            }
            if (spe !== true) {
                str += " special character";
            }
            error.innerHTML = str;
        }
    }
}


function ppUpload()
{
    var upl = document.getElementById("profilepicture");
    var max = 2000000

    if(upl.files[0].size > max)
    {
       alert("File too big!");
       upl.value = "";
    }
};