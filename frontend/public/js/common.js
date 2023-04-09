const host = `http://127.0.0.1:8080`

const sendRequest = async (path, method, body = {}, auth = false, form = false) => {
    let headers = new Headers()

    let sendBody

    if (!form) {
        headers.append('Content-Type', 'application/json')
        sendBody = JSON.stringify(body)
    } else {
        sendBody = body
    }

    if (auth) {
        headers.append('Authorization', `Bearer ${localStorage.getItem('access_token')}`)
    }

    var params = {}

    if (method === 'get' || method === 'head') {
        params = {
            method: method,
            headers: headers,
        }
    } else {
        params = {
            method: method,
            headers: headers,
            body: sendBody,
        }
    }

    try {
        const data = await fetch(host + path, params)

        return data.json()
    } catch (err) {
        alert(err)
    }
}