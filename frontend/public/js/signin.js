const loader = document.querySelector('.loader')

const submitBtn = document.querySelector('.submit-btn')
const username = document.querySelector('#username')
const email = document.querySelector('#email')
const password = document.querySelector('#password')

submitBtn.addEventListener('click', async () => {
    const body = {
        "email": email.value,
        "username": username.value,
        "password": password.value,
    }

    await sendRequest('/auth/sign-in', 'post', body)
        .then(data => {
            if (data.status >= 400) {
                showAllert(data.msg)
            } else {
                loader.style.display = 'block'
                
                localStorage.setItem('access_token', data.access_token)
                localStorage.setItem('refresh_token', data.refresh_token)

                setTimeout(() => {
                    window.location.href = '/'
                }, 3000)
            }
        })

})

const showAllert = (msg) => {
    document.querySelector('.alert-msg').innerHTML = msg
    document.querySelector('.alert-box').classList.add('show')
    document.querySelector('.alert-btn').addEventListener('click', () => {
        document.querySelector('.alert-box').classList.remove('show')
    })
}