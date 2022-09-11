const tkn = getCookie();

if (tkn.length === 0) {
    // hide logout and posts from nav bar
    hide();

    // go to login page
    switchToLogin();

    setNavUsername(true);
} else {
    // set username in nav bar
    setNavUsername(false);

    // show logout in nav bar
    show();

    // go to posts page
    switchToPosts();
}
