function AddComment() {

    const form = document.getElementById('form');
    if (form !== null) {
        form.remove()
    } else {
        var f = document.createElement("form");
        f.setAttribute('method',"post");
        f.setAttribute('action',"#");
        f.setAttribute('id',"form");

        var comment = document.createElement("input"); //input element, text
        comment.setAttribute('type',"text");
        comment.setAttribute('name',"content");

        var s = document.createElement("input"); //input element, Submit button
        s.setAttribute('type',"submit");
        s.setAttribute('value',"Submit");

        f.appendChild(comment);
        f.appendChild(s);

        document.getElementById("cc").appendChild(f);
    }
}
