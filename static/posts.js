function parseDate(date) {
    const d = new Date(date);
    const year = d.getFullYear();
    const month = d.getMonth() + 1;
    const day = d.getDate();
    const hours = d.getHours();

    let minutes = "0" + d.getMinutes();
    let seconds = "0" + d.getSeconds();

    if (minutes.length === 3) {
        minutes = minutes.substring(1, 4);
    }

    if (seconds.length === 3) {
        seconds = seconds.substring(1, 4);
    }

    const fmt = `${day}/${month}/${year} ${hours}:${minutes}:${seconds}`;
    return fmt;
}

function populatePosts() {
    getPosts().then(resp => {
        posts.innerHTML = "";

        if (resp.data == null) {
            return;
        }

        reqUsername().then(uresp => {
            const myname = uresp.data.username;

            resp.data.forEach(post => {
                const date = parseDate(post.date);
                const id = post.id;
                const uname = post.username;
                const content = post.content;

                let delBtn = "";
                if (myname == uname) {
                    delBtn = `
                    <button id="${id}" class="btn btn-sm btn-delete-post delete">
                        Delete
                    </button>
                    `
                }

                const html = `
                <article class="border">
                    <div class="metadata">
                        <div class="time">
                            ${date}
                        </div>

                        <div class="author">
                            <span>@</span>${uname}
                        </div>

                        ${delBtn}
                    </div>

                    <div class="content">
                        ${content}
                    </div>
                </article>
                `;

                posts.innerHTML += html;

                addEventListeners();
            });
        });
    }).catch(err => console.log(err));
}

function addEventListeners() {
    const btns = document.getElementsByClassName("btn-delete-post");
    for (let i = 0; i < btns.length; i ++) {
        const id = btns[i].id;
        btns[i].addEventListener("click", e => {
            e.preventDefault();
            deletePost(id);
            populatePosts();
        });
    }
}
