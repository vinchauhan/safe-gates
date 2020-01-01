const template = document.createElement('template')
template.innerHTML = `
    <div class="container">
        <h1>Login</h1>
        <form id="login-form">
            <input type="text" placeholder="Email" autocomplete="email" value="John.Doe@gmail.com" required>
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
    console.log(email)
}
