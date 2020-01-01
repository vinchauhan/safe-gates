import {createRouter} from "./lib/router.js";


const r = createRouter()
r.route("/", view('home'))
r.route("/login", view('login'))
r.route(/\//, view('not-found'))

function renderInto(target) {
    return async result => {
        target.innerHTML = ''
        target.appendChild(await result)
    }
}

r.subscribe(renderInto(document.querySelector('main')))

r.install()

function view(name) {
    return (...args) => import(`/js/components/${name}-page.js`).then(m => m.default(...args))
}
