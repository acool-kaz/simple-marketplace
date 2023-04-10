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

const getSearchedProducts = async () => {
    const params = new Proxy(new URLSearchParams(window.location.search), {
        get: (searchParams, prop) => searchParams.get(prop),
    });
    let value = params.search_by;

    document.querySelector('.heading span').innerHTML = value

    let path = ''

    if (value === null) {
        path = `/product`
    } else {
        path = `/product?search_by=${value}`
    }

    const menContainer = document.querySelector('.searched')
    await sendRequest(path, 'get')
        .then(data => {
            if (data.status >= 400) {
                alert(data.msg)
            }
        })
}

getSearchedProducts()