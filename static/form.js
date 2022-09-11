function reset() {
    for (let k in forms) {
        forms[k].reset();
    }
}

forms.login.addEventListener("submit", e => {
    e.preventDefault();

    const username = document.getElementById("login-username").value;
    const password = document.getElementById("login-password").value;

    if (username.length === 0 || password === 0) {
        alert("Fields cannot be blank", "danger");
        reset();
        return
    }

    login(username, password).then(resp => {
        if (!resp.success) {
            alert(resp.err, "danger");
            reset();
            return;
        }

        show();
        switchToPosts();
        setNavUsername(false);
    });

    reset();
});

forms.register.addEventListener("submit", e => {
    e.preventDefault();

    const username = document.getElementById("register-username").value;
    const password = document.getElementById("register-password").value;
    const confirm = document.getElementById("register-confirm-password").value;

    if (username.length === 0 || password === 0 || confirm === 0) {
        alert("Fields cannot be blank", "danger");
        reset();
        return;
    }

    if (password != confirm) {
        alert("passwords donot match", "danger");
        reset();
        return;
    }

    register(username, password).then(resp => {
        if (!resp.success) {
            alert(resp.err, "danger");
            reset();
            return;
        }

        alert(resp.data, "success");
        switchToLogin();
    });

    reset();
});

forms.newPost.addEventListener("submit", e => {
    e.preventDefault();

    const content = document.querySelector("#new textarea").value;

    if (content.length === 0) {
        alert("Post cannot be blank", "danger");
        reset();
        return;
    }

    newPost(content).then(resp => {
        if (!resp.success) {
            alert(resp.err, "danger");
            reset();
            return;
        }

        populatePosts();
    });

    reset();
});
