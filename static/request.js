async function login(username, password) {
    const resp = await fetch(urls.login, {
        method: "POST",
        body: `username=${username}&password=${password}`,
        headers: {
            "Content-Type": "application/x-www-form-urlencoded",
        }
    });

    return await resp.json();
}

async function register(username, password) {
    const resp = await fetch(urls.register, {
        method: "POST",
        body: `username=${username}&password=${password}`,
        headers: {
            "Content-Type": "application/x-www-form-urlencoded",
        }
    });

    return await resp.json();
}

async function logout() {
    const resp = await fetch(urls.logout, {
        method: "GET",
        headers: {
            'Content-Type': 'application/json;charset=utf-8',
            'Authorization': `Bearer ${getCookie()}`
        }
    });

    return await resp.json();
}

async function newPost(content) {
    const resp = await fetch(urls.newPost, {
        method: "POST",
        body: JSON.stringify({"content": content}),
        headers: {
            'Content-Type': 'application/json;charset=utf-8',
            'Authorization': `Bearer ${getCookie()}`
        }
    });

    return await resp.json();
}

async function getPosts() {
    const resp = await fetch(urls.getPosts, {
        method: "GET",
        headers: {
            'Content-Type': 'application/json;charset=utf-8',
            'Authorization': `Bearer ${getCookie()}`
        }
    });

    return await resp.json();
}

async function deletePost(id) {
    const resp = await fetch(`${urls.deletePost}/${id}`, {
        method: "DELETE",
        headers: {
            'Authorization': `Bearer ${getCookie()}`
        }
    });

    return await resp.json();
}

async function reqUsername() {
    const resp = await fetch(urls.username, {
        method: "GET",
        headers: {
            'Authorization': `Bearer ${getCookie()}`
        }
    });

    return await resp.json();
}

function setNavUsername(toDefault) {
    const ele = document.getElementById("nav-username");

    if (toDefault) {
        ele.innerText = "XSS";
        return;
    }

    reqUsername().then(resp => {
        uname = resp.data.username;
        document.getElementById("nav-username").innerText = uname;
    });
}
