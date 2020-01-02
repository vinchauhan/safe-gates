import {doPost} from "../http.js";

const template = document.createElement('template')
template.innerHTML = `
    <div class="container">
        <h1>Login</h1>
        <form id="login-form">
            <input type="text" placeholder="Email" autocomplete="email" required>
            <button>Login</button>
        </form>
    </div>
`

export default async function renderLoginPage() {
    const page = /** @type {DocumentFragment}*/template.content.cloneNode(true)
    const loginForm = /** @type {HTMLFormElement} */ page.getElementById('login-form')
    loginForm.addEventListener('submit', onLoginFormSubmit)
    return page
}

/**
 * @param {Event} ev
 */
async function onLoginFormSubmit(ev) {
    ev.preventDefault()
    const form = /** @type {HTMLFormElement} */ ev.currentTarget
    const input = form.querySelector('input')
    const button = form.querySelector('button')
    const email = input.value
    input.disabled = true
    button.disabled = true

    try {
        const out = await http.devLogin(email)
        localStorage.setItem('token', out.token)
        localStorage.setItem('expires_at', typeof out.expiresAt === 'string'
        ? out.expiresAt
        : out.expiresAt.toJSON())
        localStorage.setItem('auth_user', JSON.stringify(out.user))
        location.reload()
    } catch (e) {
        console.error(e)
        alert(e.message)
        setTimeout(input.focus)
    } finally {
        input.disabled = true
        button.disabled = true
    }
}

const http = {
    devLogin: email => doPost('/api/dev_login', {email})
}
