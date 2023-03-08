const bar = async () => {
    const html = await fetch('/pages/navbar-no-auth.html').then((data) => data.text());
    document.getElementById("nav-bar").innerHTML = html;
}