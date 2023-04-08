const productList = [
    {
        id: 1,
        img: 'img/card1.png',
        title: 'dress',
        des: 'a short line about the cloth...',
        price: 20
    },
    {
        id: 2,
        img: 'img/card2.png',
        title: 'dress',
        des: 'a short line about the cloth...',
        price: 20
    },
    {
        id: 3,
        img: 'img/card3.png',
        title: 'dress',
        des: 'a short line about the cloth...',
        price: 20
    },
    {
        id: 4,
        img: 'img/card4.png',
        title: 'dress',
        des: 'a short line about the cloth...',
        price: 20
    },
    {
        id: 5,
        img: 'img/card5.png',
        title: 'dress',
        des: 'a short line about the cloth...',
        price: 20
    },
    {
        id: 6,
        img: 'img/card6.png',
        title: 'dress',
        des: 'a short line about the cloth...',
        price: 20
    },
    {
        id: 7,
        img: 'img/card7.png',
        title: 'dress',
        des: 'a short line about the cloth...',
        price: 20
    },
    {
        id: 8,
        img: 'img/card8.png',
        title: 'dress',
        des: 'a short line about the cloth...',
        price: 20
    },
    {
        id: 1,
        img: 'img/card1.png',
        title: 'dress',
        des: 'a short line about the cloth...',
        price: 20
    },
    {
        id: 2,
        img: 'img/card2.png',
        title: 'dress',
        des: 'a short line about the cloth...',
        price: 20
    },
    {
        id: 3,
        img: 'img/card3.png',
        title: 'dress',
        des: 'a short line about the cloth...',
        price: 20
    },
    {
        id: 4,
        img: 'img/card4.png',
        title: 'dress',
        des: 'a short line about the cloth...',
        price: 20
    },
    {
        id: 5,
        img: 'img/card5.png',
        title: 'dress',
        des: 'a short line about the cloth...',
        price: 20
    },
    {
        id: 6,
        img: 'img/card6.png',
        title: 'dress',
        des: 'a short line about the cloth...',
        price: 20
    },
    {
        id: 7,
        img: 'img/card7.png',
        title: 'dress',
        des: 'a short line about the cloth...',
        price: 20
    },
    {
        id: 8,
        img: 'img/card8.png',
        title: 'dress',
        des: 'a short line about the cloth...',
        price: 20
    }
]

const productCard = (product) => {
    return `
    <div class="product-card">
        <div class="product-image">
            <img src="${product.img}" class="product-thumb" alt="">
            <div class="product-buttons">
                <button class="card-btn">add to cart</button>
                <button class="card-btn" onclick="window.location.href = '/product?id=${product.id}'">more info</button>
            </div>
        </div>
        <div class="product-info">
            <h2 class="product-brand">${product.title}</h2>
            <p class="product-short-des">${product.des}</p>
            <span class="price">${product.price}$</span>
        </div>
    </div>
    `
}

const getSearchedProducts = () => {
    const menContainer = document.querySelector('.searched')
    menContainer.innerHTML = ''
    productList.forEach(product => {
        menContainer.innerHTML += productCard(product)
    })
}

getSearchedProducts()