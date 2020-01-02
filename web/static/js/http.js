import {isAuthenticated} from "./auth.js";
import {isObject} from "./utl.js";

/**
 * @param {Response} res
 * @returns {Promise<any>}
 */
async function parseResponse(res) {
    let body = await res.clone().json().catch(() => res.text())
    if (!res.ok) {
        const msg = String(body).trim().toLowerCase()
        const err = new Error(msg)
        err.name = msg
                .split(" ")
                .map(word => word.charAt(0).toUpperCase() + word.slice(1))
                .join("")
            + "Error"
        err["statusCode"] = res.status
        err["statusText"] = res.statusText
        err["url"] = res.url
        throw err
    }
    return body
}

/**
 * @param {string} url
 * @param {{[key:string]:string}=} headers
 * @return Promise<Response>
 */
export function doGet(url, headers) {
    const request = new Request(url, {
        method: 'GET',
        mode: "no-cors",
        headers: Object.assign(defaultHeaders(), headers)
    })
    return fetch(request).then(parseResponse)
    // return fetch(url, {
    //     method: 'GET',
    //     headers: Object.assign(defaultHeaders(), headers),
    // }).then(parseResponse)
}

export function doPost(url, body, headers) {
    const init = {
        method: 'POST',
        headers: defaultHeaders()
    }
    if(isObject(body)) {
        init['body'] = JSON.stringify(body)
        init.headers['content-type'] = 'application/json; charset=utf-8'
    }
    Object.assign(init.headers, headers)
    return fetch(url, init).then(parseResponse)
}

function defaultHeaders() {
    return isAuthenticated() ? {
        Authorization: `Bearer `+ localStorage.getItem('token'),
    } : {}
}

export default {
    get: doGet,
    post: doPost,
}
