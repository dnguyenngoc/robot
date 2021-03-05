import {urlLoginAccessToken, urlSignup} from './UrlServices';


export const userService = {
    loginAccessToken,
    logout,
    signup,
    getById,
    update,
    delete: _delete 
};


async function loginAccessToken(username, password) {
    const requestOptions = {
        method: 'POST',
        headers: {'Content-Type': 'application/json'}
    };
    try {
        const response = await Axios.post(
            urlLoginAccessToken(), 
            requestOptions, 
            { email: username, password: password, verifyPassword: verifyPassword }
        );
        return response.data;
    }catch (error) {
        console.error('Error', error.response);
        return false
    }
}


async function signup(username, password, verifyPassword) {
    const requestOptions = {
        method: 'POST',
        headers: {'Content-Type': 'application/json'}
    };
    try {
        const response = await Axios.post(
            urlSignup(),
            requestOptions, 
            { email: username, password: password, verifyPassword: verifyPassword }
        );
        return response.data;
    }catch (error) {
        console.error('Error', error.response);
        return false
    }
}



function logout() {
}

function getById(id) {
    
}

async function update(id, params) {
   
}

async function _delete(id) {
   
}

