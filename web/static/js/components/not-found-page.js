const template = document.createElement('template')
template.innerHTML = `
    <div class="container">
        <h1>404 Not Found</h1>
    </div>
`

export default async function renderNotFoundPage() {
    const page = /** @type {DocumentFragment}*/template.content.cloneNode(true)
    return page
}
