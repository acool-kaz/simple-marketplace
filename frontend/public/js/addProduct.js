let loader = document.querySelector('.loader')

window.onload = () => {
    if (!localStorage.getItem('access_token')) {
        window.location.href = '/signup'
    }
}

const productName = document.querySelector('#product-name')
const shortDes = document.querySelector('#short-des')
const des = document.querySelector('#des')
const price = document.querySelector('#price')
const productTag = document.querySelector('#product-tag')
const uploadImages = document.querySelectorAll('.fileupload')

uploadImages.forEach((fileUpload, index) => {
    fileUpload.addEventListener('change', () => {
        const file = fileUpload.files[0]

        var reader = new FileReader();

        reader.onload = () => {
            var imgPath = "url('" + reader.result + "')";

            document.querySelector(`label[for=${fileUpload.id}]`).style.backgroundImage = imgPath
            document.querySelector('.product-image').style.backgroundImage = imgPath
        }

        reader.readAsDataURL(file);
    })
})

const addBtn = document.querySelector('#add-btn')

addBtn.addEventListener('click', async () => {
    let formData = new FormData();

    formData.append('name', productName.value)
    formData.append('description', shortDes.value)
    formData.append('price', price.value)

    uploadImages.forEach(file => {
        formData.append('images', file.files[0])
    })

    await sendRequest('/product/api', 'post', formData, true, true)
    .then(data => {
        console.log(data);
    })
})

const saveBtn = document.querySelectorAll('#save-btn')

const showAllert = (msg) => {
    document.querySelector('.alert-msg').innerHTML = msg
    document.querySelector('.alert-box').classList.add('show')
    document.querySelector('.alert-btn').addEventListener('click', () => {
        document.querySelector('.alert-box').classList.remove('show')
    })
}