import Cookie from 'universal-cookie'

function get(key) {
    const cookie = new Cookie();
    return cookie.get(key);
}

function set(key, value){
    const cookie = new Cookie();

    cookie.set(key, value)
}

function remove(key) {
    const cookie = new Cookie();
    cookie.remove(key)
}

export {
    get as cookieGet,
    set as cookieSet,
    remove as cookieRemove
}