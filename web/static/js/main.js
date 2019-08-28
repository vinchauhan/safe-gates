//create router instance
import { createRouter } from './lib/router.js'

const r = createRouter()

r.route('/', () => {
    return 'Home page'
})

r.route(/\//, () => {
    return '404 Not Found'
})

r.subscribe(renderInto(document.querySelector('main')))
r.install()

function renderInto(target) {
    return result => {
        target.innerHTML = result
    }
}