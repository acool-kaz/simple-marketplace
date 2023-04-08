const productContainers = [...document.querySelectorAll('.product-container')]
const nxtBtn = [...document.querySelectorAll('.nxt-btn')]
const preBtn = [...document.querySelectorAll('.pre-btn')]

productContainers.forEach((item, i) => {
    let conatinerDim = item.getBoundingClientRect()
    let conatinerWidth = conatinerDim.width

    nxtBtn[i].addEventListener('click', ()=>{
        item.scrollLeft += conatinerWidth/2
        if (item.scrollLeft >= conatinerWidth/2) {
            item.scrollLeft = 0
        }
    })

    preBtn[i].addEventListener('click', ()=>{
        item.scrollLeft -= conatinerWidth/2
        if (item.scrollLeft <= 0) {
            item.scrollLeft = conatinerWidth*5
        }
    })
})