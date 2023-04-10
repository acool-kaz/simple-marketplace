const productCard = (product) => {
    return `
    <div class="product-card">
        <div class="product-image">
            <img src="${product.product_images[0]}" class="product-thumb" alt="">
            <div class="product-buttons">
                <button class="card-btn">add to cart</button>
                <button class="card-btn" onclick="window.location.href = '/web/product?id=${product.product_id}'">more info</button>
            </div>
        </div>
        <div class="product-info">
            <h2 class="product-brand">${product.product_name}</h2>
            <p class="product-short-des">${product.product_short_description}</p>
            <span class="price">${product.product_price}$</span>
        </div>
    </div>
    `
}

const getMenProducts = async () => {
    const menContainer = document.querySelector('.men')
    menContainer.innerHTML = ''
    await sendRequest('/product/men', 'get')
        .then(data => {
            if (data.status >= 400) {
                alert(data.msg)
            } else if (data.data !== null) {
                data.data.forEach(product => {
                    menContainer.innerHTML += productCard(product)
                })
            }
        })
}

getMenProducts()

const getWomenProducts = async () => {
    const womenContainer = document.querySelector('.women')
    womenContainer.innerHTML = ''
    await sendRequest('/product/women', 'get')
        .then(data => {
            if (data.status >= 400) {
                alert(data.msg)
            } else if (data.data !== null) {
                data.data.forEach(product => {
                    womenContainer.innerHTML += productCard(product)
                })
            }
        })
}

getWomenProducts()

const getCurProduct = async () => {
    let curProduct = {}

    const params = new Proxy(new URLSearchParams(window.location.search), {
        get: (searchParams, prop) => searchParams.get(prop),
    });
    let value = params.id;

    await sendRequest(`/product/${value}`, 'get')
        .then(data => {
            if (data.status >= 400) {
                alert(data.msg)
            } else if (data.data !== null) {
                curProduct = data.data[0]
            }
        })

    document.querySelector('.product-images').style.gridTemplateColumns = `repeat(${curProduct.product_images.length}, auto)`

    curProduct.product_images.forEach(img => {
        document.querySelector('.product-images').innerHTML += `<img src="${img}" alt="">`
    })

    document.querySelector('.product-brand').innerHTML = curProduct.product_name
    document.querySelector('.product-short-des').innerHTML = curProduct.product_short_description
    document.querySelector('.product-price').innerHTML = curProduct.product_price + '$'
    document.querySelector('.detail-des .des').innerHTML = curProduct.product_description
}

getCurProduct()

const imageSlider = async () => {
    await sleep(1000)

    const productImages = document.querySelectorAll('.product-images img')
    const productImageSlider = document.querySelector('.image-slider')

    document.querySelector('.product-images').style.gridTemplateColumns = `repeat(${productImages.length}, 1fr)`

    let activeImageSilder = 0
    productImages[activeImageSilder].classList.add('active')
    productImageSlider.style.backgroundImage = `url(${productImages[activeImageSilder].src})`

    productImages.forEach((item, i) => {
        item.addEventListener('click', () => {
            productImages[activeImageSilder].classList.remove('active')
            item.classList.add('active')
            productImageSlider.style.backgroundImage = `url(${item.src})`
            activeImageSilder = i
        })
    })
}

imageSlider()