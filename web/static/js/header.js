import {isAuthenticated, getAuthUser} from "./auth.js";

const authenticated = isAuthenticated()
const authUser = getAuthUser()
const header = document.querySelector('header')

header.innerHTML = `
    <nav>
        <a href="/">Home</a>
        ${authenticated ? `
            <a href="/users/${authUser.username}">Profile</a>
            <a href="/codes/${authUser.username}">Passcodes</a>
            <button id="logout-button">Logout</button>
        ` : ''}
    </nav>
`

/**
 *
 * @param {MouseEvent} ev
 */
function onLogoutButtonClick(ev) {
    const button = /** @type {HTMLButtonElement} */ ev.currentTarget
    button.disabled = true
    localStorage.clear()
    location.reload()
}

if (authenticated) {
    const logoutButton = header.querySelector('#logout-button')
    logoutButton.addEventListener('click', onLogoutButtonClick)
}
