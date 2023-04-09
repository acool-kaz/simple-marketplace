const loader = document.querySelector('.loader')

const submitBtn = document.querySelector('.submit-btn')
const firstName = document.querySelector('#firstName')
const secondName = document.querySelector('#secondName')
const username = document.querySelector('#username')
const email = document.querySelector('#email')
const password = document.querySelector('#password')
const number = document.querySelector('#number')

submitBtn.addEventListener('click', async () => {
    if (firstName.value.length < 3) {
        showAllert('first name must be 3 letters long')
    } else if (secondName.value.length < 3) {
        showAllert('second name must be 3 letters long')
    } else if (username.value.length < 3) {
        showAllert('username must be 3 letters long')
    } else if (!email.value.length) {
        showAllert('enter your email')
    } else if (password.value.length < 8) {
        showAllert('name must be 8 letters long')
    } else if (!number.value.length) {
        showAllert('enter your phone number')
    } else if (!(Number(number.value) || number.value.length < 10)) {
        showAllert('invalid phone number')
    }

    loader.style.display = 'block'

    const body = {
        "first_name": firstName.value,
        "second_name": secondName.value,
        "email": email.value,
        "phone_number": number.value,
        "username": username.value,
        "password": password.value,
    }

    await sendRequest('/auth/sign-up', 'post', body)
        .then(data => {
            console.log(data);
        })

    setTimeout(() => {
        window.location.href = '/'
    }, 3000)
})

const showAllert = (msg) => {
    document.querySelector('.alert-msg').innerHTML = msg
    document.querySelector('.alert-box').classList.add('show')
    document.querySelector('.alert-btn').addEventListener('click', () => {
        document.querySelector('.alert-box').classList.remove('show')
    })
}