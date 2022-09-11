function switchToLogin() {
    forms.login.style.display = "block";
    forms.register.style.display = "none";
    postPage.style.display = "none";
}

function switchToRegister() {
    forms.login.style.display = "none";
    forms.register.style.display = "block";
    postPage.style.display = "none";
}

function switchToPosts() {
    forms.login.style.display = "none";
    forms.register.style.display = "none";
    postPage.style.display = "block";
    populatePosts();
}

function show() {
    nav.logout.style.display = "block";
    nav.posts.style.display = "block";
}

function hide() {
    nav.logout.style.display = "none";
    nav.posts.style.display = "none";
}

nav.login.addEventListener("click", e => {
    e.preventDefault();
    switchToLogin();
});

nav.register.addEventListener("click", e => {
    e.preventDefault();
    switchToRegister();
});

nav.logout.addEventListener("click", e => {
    e.preventDefault();
    hide();
    switchToLogin();
    logout();
    setNavUsername(true);
});

nav.posts.addEventListener("click", e => {
    e.preventDefault();
    switchToPosts();
});
