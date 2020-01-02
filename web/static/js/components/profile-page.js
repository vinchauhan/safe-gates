const template = document.createElement('template')
template.innerHTML = `
    <div class="container">
        <h1>Profile page</h1>
    </div>
`

export default async function renderProfilePage() {
    const page = /** @type {DocumentFragment}*/template.content.cloneNode(true)
    return page
}
