const createNav = () => {
    let nav = document.querySelector('.navbar')

    nav.innerHTML = `
        <div class="nav">
            <a href="/"><img src="img/dark-logo.png" class="brand-logo" alt=""></a>
            <div class="nav-items">
                <div class="search">
                    <input type="text" class="search-box" placeholder="search brand, product">
                    <button class="search-btn">search</button>
                </div>
                <a href="#" onclick="login()"><img src="img/user.png" alt=""></a>
                <a href="#" onclick="logout()" class="logout"><img src="img/logout.png" alt=""></a>
                <a href="#"><img src="img/cart.png" alt=""></a>
            </div>
        </div>
        <ul class="links-container">
            <li class="link-item"><a href="#" class="link">home</a></li>
            <li class="link-item"><a href="#new" class="link">new</a></li>
            <li class="link-item"><a href="#men" class="link">men</a></li>
            <li class="link-item"><a href="#women" class="link">women</a></li>
        </ul>
    `
}

createNav()

window.onload = () => {
    if (localStorage.getItem('access_token')) {
        document.querySelector('.logout').style.display = 'block'
    } else {
        document.querySelector('.logout').style.display = 'none'
    }
}

const login = () => {
    if (localStorage.getItem('access_token')) {
        console.log('has token');
    } else {
        window.location.href = '/signup'
    }
}

const logout = () => {
    localStorage.removeItem('access_token')
    localStorage.removeItem('refresh_token')
    window.location.href = '/'
}

document.querySelector('.search-btn').addEventListener('click', () => {
    window.location.href = '/search?search_by=' + document.querySelector('.search-box').value
})