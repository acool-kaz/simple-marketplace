const productCard = (product) => {
    return `
    <div class="product-card">
        <div class="product-image">
            <img src="${product.product_images[0]}" class="product-thumb" alt="">
            <div class="product-buttons">
                <button class="card-btn">add to cart</button>
                <button class="card-btn" onclick="window.location.href = '/product?id=${product.product_id}'">more info</button>
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

const getNewProducts = async () => {
    const newContainer = document.querySelector('.new')
    newContainer.innerHTML = ''
    await sendRequest('/product/new', 'get')
        .then(data => {
            if (data.status >= 400) {
                alert(data.msg)
            } else if (data.data !== null) {
                data.data.forEach(product => {
                    newContainer.innerHTML += productCard(product)
                })
            }
        })
}

getNewProducts()

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
