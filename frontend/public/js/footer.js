const createFooter = () => {
    let footer = document.querySelector('footer')

    footer.innerHTML = `
        <div class="footer-content">
            <img src="img/light-logo.png" alt="" class="logo">
            <div class="footer-ul-container">
                <ul class="category">
                    <li class="category-title">men</li>
                    <li><a href="#" class="footer-link">t-shirts</a></li>
                    <li><a href="#" class="footer-link">shirts</a></li>
                    <li><a href="#" class="footer-link">jeans</a></li>
                    <li><a href="#" class="footer-link">shoes</a></li>
                    <li><a href="#" class="footer-link">watch</a></li>
                    <li><a href="#" class="footer-link">sports</a></li>
                </ul>
                <ul class="category">
                    <li class="category-title">women</li>
                    <li><a href="#" class="footer-link">t-shirts</a></li>
                    <li><a href="#" class="footer-link">shirts</a></li>
                    <li><a href="#" class="footer-link">jeans</a></li>
                    <li><a href="#" class="footer-link">shoes</a></li>
                    <li><a href="#" class="footer-link">watch</a></li>
                    <li><a href="#" class="footer-link">sports</a></li>
                </ul>
            </div>
        </div>
        <p class="footer-title">about site</p>
        <p class="info">Lorem ipsum dolor sit amet consectetur adipisicing elit. Porro enim eos veniam sunt aliquid. Minima reprehenderit tempore alias perferendis eius quam voluptatem voluptas accusamus quis accusantium, repudiandae quos natus illo.</p>
        <p class="info">support emails - help@gmail.com, customersupport@gmail.com</p>
        <p class="info">telephone - +7 (777) 777 77 77</p>
        <div class="footer-social-container">
            <div>
                <a href="#" class="social-link">terms & services</a>
                <a href="#" class="social-link">privacy page</a>
            </div>
            <div>
                <a href="#" class="social-link">instagram</a>
                <a href="#" class="social-link">facebook</a>
                <a href="#" class="social-link">twitter</a>
            </div>
        </div>
        <p class="footer-credit">clothing, best apparels online shop</p>
    `
}

createFooter()