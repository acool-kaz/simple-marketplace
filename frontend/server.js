const express = require('express')
const path = require('path')

let staticPath = path.join(__dirname, 'public')

const app = express()

app.use(express.static(staticPath))

app.get('/', (req, res) => {
    res.sendFile(path.join(staticPath, 'index.html'))
})

app.get('/product', (req, res) => {
    res.sendFile(path.join(staticPath, 'product.html'))
})

app.get('/signup', (req, res) => {
    res.sendFile(path.join(staticPath, 'signup.html'))
})

app.get('/signin', (req, res) => {
    res.sendFile(path.join(staticPath, 'signin.html'))
})

app.get('/search', (req, res) => {
    res.sendFile(path.join(staticPath, 'search.html'))
})

app.get('/404', (req, res) => {
    res.sendFile(path.join(staticPath, '404.html'))
})

app.use((req, res) => {
    res.redirect('/404')
})

app.listen(3000, () => {
    console.log('Starting frontend on http://localhost:3000');
})