const getCurProduct = async () => {
    let curProduct = {
        product_images: ['https://www.hallensteins.com/content/products/hb-organic-oversized-tee-white-front-10001791.jpg?width=2058', 'https://img.sonofatailor.com/images/customizer/product/White_O_Crew_Regular_NoPocket.jpg'],
        product_name: 'white t shirt',
        product_short_description: `Amazon Essentials Men's Slim-Fit Short-Sleeve Crewneck T-Shirt, Pack of 2`,
        product_price: '12',
        product_description: `
        Solids: 100% Cotton; Heathers: 60% Cotton, 40% Polyester
        Imported
        Machine Wash
        Fits close to the body for fitted, slim silhouette
        Smooth and comfortable lightweight jersey fabric
        Crew neckline
        Tag-free
        An Amazon brand
        `,
    }

    document.querySelector('.product-images').style.gridTemplateColumns = `repeat(${curProduct.product_images.length}, auto)`

    curProduct.product_images.forEach(img => {
        document.querySelector('.product-images').innerHTML += `<img src="${img}" alt="">`
    })

    document.querySelector('.product-brand').innerHTML = curProduct.product_name
    document.querySelector('.product-short-des').innerHTML = curProduct.product_short_description
    document.querySelector('.product-price').innerHTML = curProduct.product_price + '$'
    document.querySelector('.detail-des .des').innerHTML = curProduct.product_description
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

getCurProduct()

const imageSlider = async () => {
    await sleep(100)

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