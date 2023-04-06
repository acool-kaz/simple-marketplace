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

const getMenProducts = () => {
    const menContainer = document.querySelector('.men')
    menContainer.innerHTML = ''
    productList.forEach(product => {
        menContainer.innerHTML += productCard(product)
    })
}

getMenProducts()

const getWomenProducts = () => {
    const womenContainer = document.querySelector('.women')
    womenContainer.innerHTML = ''
    productList.forEach(product => {
        womenContainer.innerHTML += productCard(product)
    })
}

getWomenProducts()

const getCurProduct = () => {
    const curProduct = {
        images: [
            'img/product image 1.png',
            'img/product image 2.png',
            'img/product image 3.png',
            'img/product image 4.png',
            'img/product image 3.png',
            'img/product image 2.png',
            'img/product image 1.png',
        ],
        brand: 'Calvin Klein',
        shortDes: 'Lorem Ipsum Dolor Sit Amet, Consectetur Adipisicing Elit. Labore, Maiores Assumenda Voluptatibus Eligendi Nulla Consequuntur Laborum Maxime Quam Voluptatem Autem Iure? Sequi Molestias Quae Ut Tenetur Repudiandae Suscipit Tempora Numquam.',
        des: `Lorem ipsum dolor sit amet, consectetur adipisicing elit. Iusto quas tempore porro ex iste odit ab, temporibus fuga ullam saepe deserunt, harum enim a! Voluptate, saepe officiis? Aliquid, deleniti architecto.
        Exercitationem sunt labore distinctio iure illo repellendus. Aliquam incidunt consectetur iusto ab aliquid sint similique quas eaque accusamus aperiam doloribus, veritatis aspernatur laudantium dignissimos explicabo eligendi recusandae ad illum error!
        Atque ut repudiandae pariatur. Odit alias repellat excepturi unde nihil facere illum animi, illo ex iure perferendis ipsam, doloremque corporis, maiores quisquam. Porro consequuntur accusantium quo recusandae ex, explicabo quas.
        Asperiores nihil culpa ea, quis quidem dicta voluptatem laboriosam. Iure iste veniam sapiente labore, voluptatibus quae unde nisi neque similique saepe excepturi? Ipsa laboriosam quam excepturi repellat tempora recusandae facere!
        Maiores aliquam officia veritatis et corporis repellat tenetur quibusdam repellendus maxime ea quasi labore suscipit placeat cupiditate neque earum quisquam consequuntur cumque obcaecati, qui doloremque iusto eaque dolorum? Perferendis, ex.
        Placeat, nam facere? Laboriosam quo hic quaerat cupiditate distinctio, adipisci eveniet dolores dolorum tenetur? Magni non iusto tempora ab quas. Voluptatibus saepe expedita consectetur rerum aliquid asperiores provident obcaecati maiores.
        Recusandae cupiditate eius alias deserunt, quam, sequi laboriosam sed dignissimos libero facilis amet maxime impedit quis veritatis laborum fugit ab, quisquam obcaecati totam? Ipsum similique cumque, facere nostrum enim sunt.
        Perspiciatis mollitia incidunt facilis tempore quos aliquam ea velit nesciunt, repellat suscipit totam delectus ipsa fugit odio, in blanditiis maiores. Reiciendis voluptatibus qui inventore harum! At maxime dolorum vel eaque?
        Vero hic dolore consectetur veritatis et dignissimos, fuga tempora quasi neque accusantium id necessitatibus, eius laudantium earum modi debitis ipsam similique maxime. Et, optio. Commodi numquam eligendi labore voluptate corrupti.
        Fugit magnam fugiat, libero consequatur dolor necessitatibus quos praesentium tenetur quo officiis vero exercitationem similique amet rem velit nesciunt labore aliquam dignissimos quae, sint consectetur earum cum. Alias, expedita possimus.`,
        price: 20
    }
    
    curProduct.images.forEach(img => {
        document.querySelector('.product-images').innerHTML += `<img src="${img}" alt="">`
    })

    document.querySelector('.product-brand').innerHTML = curProduct.brand
    document.querySelector('.product-short-des').innerHTML = curProduct.shortDes
    document.querySelector('.product-price').innerHTML = curProduct.price + '$'
    document.querySelector('.detail-des .des').innerHTML = curProduct.des
}

getCurProduct()

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