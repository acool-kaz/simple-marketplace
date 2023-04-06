const loader = document.querySelector('.loader')

const submitBtn = document.querySelector('.submit-btn')
const username = document.querySelector('#username')
const email = document.querySelector('#email')
const password = document.querySelector('#password')

submitBtn.addEventListener('click', () => {
    loader.style.display = 'block'

    localStorage.setItem('token', 'value')

    setTimeout(() => {
        window.location.href = '/'
    }, 3000)

    return

    sendData('http://127.0.0.1:8080/auth/sign-in', {
        "email": email.value,
        "username": username.value,
        "password": password.value,
    })
})

const sendData = (path, data) => {
    fetch(path, {
        method: 'post',
        headers: new Headers({
            'Content-Type': 'application/json'
        }),
        body: JSON.stringify(data)
    })
    .then((res) => res.json())
    .then(response => {
        console.log(response);
    })
}

const showAllert = (msg) => {
    document.querySelector('.alert-msg').innerHTML = msg
    document.querySelector('.alert-box').classList.add('show')
    document.querySelector('.alert-btn').addEventListener('click', () => {
        document.querySelector('.alert-box').classList.remove('show')
    })
}