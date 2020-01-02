import {createRouter} from "./lib/router.js";
import {gaurd} from "./auth.js";

const modulesCache = new Map()
const viewsCache = new Map()
const disconnectEvent = new CustomEvent("disconnect")
const r = createRouter()
r.route("/", gaurd(view('home'), view('login')))
r.route("/login", view('login'))
r.route(/^\/codes\/(?<username>[a-zA-Z][a-zA-Z0-9_-]{0,17})$/, gaurd(view('passcode'), view('login')))
r.route(/^\/users\/(?<username>[a-zA-Z][a-zA-Z0-9_-]{0,17})$/, view('profile'))
r.route(/\//, view('not-found'))

function renderInto(target) {
    let currentPage = /** @type {Node=}*/ (null)
    return async result => {
        if (currentPage instanceof Node) {
            console.log("disconnect event sent")
            currentPage.dispatchEvent(disconnectEvent)
            target.innerHTML=''
        }
        try {
            currentPage = await result
        } catch (e) {
            console.error(e)
            currentPage = renderErrorPage(e)
        }
        target.appendChild(currentPage)
    }
}

r.subscribe(renderInto(document.querySelector('main')))

r.install()

function view(name) {
    return (...args) => {
        if (viewsCache.has(name)){
            const renderPage = viewsCache.get(name)
            return renderPage(...args)
        }
        return import(`/js/components/${name}-page.js`).then(m => {
            const renderPage = m.default
            viewsCache.set(name, renderPage)
            return renderPage(...args)
        })
    }
}
