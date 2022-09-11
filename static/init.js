const forms = {
    "login": document.getElementById("login"),
    "register": document.getElementById("register"),
    "newPost": document.getElementById("new"),
};

const postPage = document.querySelector("main");
const posts = document.getElementById("posts");

const nav = {
    "login": document.getElementById("nav-login"),
    "register": document.getElementById("nav-register"),
    "posts": document.getElementById("nav-posts"),
    "logout": document.getElementById("nav-logout"),
};

const baseURL = location.protocol + "//" + location.hostname + ":" + location.port

const urls = {
    "login": baseURL + "/login",
    "logout": baseURL + "/logout",
    "register": baseURL + "/register",
    "username": baseURL + "/username",
    "getPosts": baseURL + "/post/getall",
    "newPost": baseURL + "/post/create",
    "deletePost": baseURL + "/post/delete",
}
