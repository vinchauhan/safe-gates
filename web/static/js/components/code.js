/**
 * @param {import("../types.js").Code} code
 */
export default function renderCode(code) {
    const li = document.createElement('li')
    li.className = 'code-item'
    li.innerHTML = `
        <article class="code">
            <div class="code-content">
                <span>${code}</span>
                <button class="use-code" onclick="useToken()">Use</button>
            </div>
        </article>
    `
    return li
}


