import {doGet} from "../http.js";
import renderCode from "./code.js";

let passCodes = /** @type {import("../types.js").Code[]}*/ (null)
const template = document.createElement('template')

template.innerHTML = `
    <div class="container">
        <h1>Passcodes</h1>
        <ol id="passcode-list"></ol>
        <button id="regen-passcode-btn">Regenerate</button>
    </div>
`

export default async function renderPasscodePage(params) {
    passCodes = await http.codes(params.username)
    const page = /** @type {DocumentFragment}*/ template.content.cloneNode(true)
    // page.addEventListener("disconnect", function() {
    //     alert("Ready!");
    // }, false);
    const passcodeList = /** @type {HTMLOListElement} */ page.getElementById('passcode-list')
    const regenerateButton = /** @type {HTMLButtonElement} */ page.getElementById('regen-passcode-btn')
    for(const passCode of passCodes) {
        passcodeList.appendChild(renderCode(passCode.code))
    }
    return page
}

const http = {
    /**
     *
     * @returns {Promise<any>}
     */
    codes: username => doGet(`/api/passcodes/${username}`)
}



