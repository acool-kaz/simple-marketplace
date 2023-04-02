const productCard = (product) => {
    return `
    <a href="/product/${product.id}">
        <div class="product">
            <img src="${product.image}" alt="" class="product-image">
            <div class="product-text">
                <p>${product.name}</p>
                <p>${product.price}</p>
            </div>
        </div>
    </a>
    `
}

const randomProductCard = (product) => {
    const tags = []

    product.tags.forEach(element => {
        tags.push(`<a href="#">${element}</a>`)
    });

    const carouselItems = []
    const liItems = []

    product.images.forEach((element, index) => {
        if (index === 0) {
            liItems.push(`<li data-target="#carouselExampleIndicators" data-slide-to="0" class="active"></li>`)

            carouselItems.push(`
            <div class="carousel-item active">
                <img class="d-block w-100" src="${element}" alt="">
            </div>
            `)
        } else {
            liItems.push(`<li data-target="#carouselExampleIndicators" data-slide-to="${index}"></li>`)

            carouselItems.push(`
            <div class="carousel-item">
                <img class="d-block w-100" src="${element}" alt="">
            </div>
            `)
        }
    });

    return `
    <div class="left-side">
        <div id="carouselExampleIndicators" class="carousel slide" data-ride="carousel">
            <ol class="carousel-indicators">
            ${liItems.join('\n')}            
            </ol>
            <div class="carousel-inner" style="width: 70%; margin: auto;">
            ${carouselItems.join('\n')}
            </div>
            <a class="carousel-control-prev" href="#carouselExampleIndicators" role="button" data-slide="prev">
            <span class="carousel-control-prev-icon" aria-hidden="true"></span>
            <span class="sr-only">Previous</span>
            </a>
            <a class="carousel-control-next" href="#carouselExampleIndicators" role="button" data-slide="next">
            <span class="carousel-control-next-icon" aria-hidden="true"></span>
            <span class="sr-only">Next</span>
            </a>
        </div>
    </div>
    <div class="right-side">
        <a href="/user/${product.user_id}"><p>${product.username}</p></a>
        <h4>${product.name}</h4>
        <h5>${product.price}</h5>
        <h5>${tags.join(' ')}</h5>
        <div class="button"><h3>Buy</h3></div>
        <h4>Description</h4>
        <h5>${product.description}</h5>
    </div>
    `
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

async function setRandomProduct() {
    product = {
        id: 1,
        user_id: 1,
        username: 'acool',
        tags: ['Tech', 'Computer'],
        images: ['../assets/images/test-product.jpg','../assets/images/test-product.jpg','../assets/images/test-product.jpg','../assets/images/test-product.jpg','../assets/images/test-product.jpg'],
        name: 'Test Product',
        price: '120$',
        description: 'Lorem ipsum dolor sit, amet consectetur adipisicing elit. Sint odio doloremque nulla officiis veritatis iste, repellendus dolorum quia quas unde accusamus provident esse tempore maiores quibusdam eius voluptates, ut labore.'
    }

    await sleep(100);

    document.querySelector('.random-product').innerHTML = randomProductCard(product)
}

setRandomProduct()

async function setNewProductList() {
    products = [
        {
            id: 1,
            image: '../assets/images/test-product.jpg',
            name: 'Test Product',
            price: '120$'
        },
        {
            id: 1,
            image: '../assets/images/test-product.jpg',
            name: 'Test Product',
            price: '120$'
        },
        {
            id: 1,
            image: '../assets/images/test-product.jpg',
            name: 'Test Product',
            price: '120$'
        },
        {
            id: 1,
            image: '../assets/images/test-product.jpg',
            name: 'Test Product',
            price: '120$'
        },
        {
            id: 1,
            image: '../assets/images/test-product.jpg',
            name: 'Test Product',
            price: '120$'
        },
    ]
    
    var productList = ''

    products.forEach(product => {
        productList += productCard(product)
    });

    await sleep(100);

    document.querySelector('.products').innerHTML = productList
}

setNewProductList();

