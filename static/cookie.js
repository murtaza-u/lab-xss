function getCookie() {
    return document.cookie.split("; ").reduce((r, v) => {
        const parts = v.split("=");
        return parts[0] === "token" ? decodeURIComponent(parts[1]) : r;
      }, "");
}

function setCookie(token) {
    document.cookie = `token=${token};`;
}
