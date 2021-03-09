import Axios from 'axios'


export const userServiceLoginAccessToken = (username, password)=>{
    loginAccessToken(username, password);
}

async function loginAccessToken(username, password) {
    const requestOptions = {
        method: 'POST',
        headers: {'Content-Type': 'application/json',}
    };
    try {
        const response = await Axios.post(
            "http://localhost:8080/api/v1/accounts/login/access-token", 
            requestOptions, 
            {"email": username, "password": password }
        );
        return response.data;
    }catch (error) {
        console.error('Error', error.response);
        return false
    }
}


// async function signup(username, password, verifyPassword) {
//     const requestOptions = {
//         method: 'POST',
//         headers: {'Content-Type': 'application/json'}
//     };
//     try {
//         const response = await Axios.post(
//             urlServiceSignup(),
//             requestOptions, 
//             { email: username, password: password, verifyPassword: verifyPassword }
//         );
//         return response.data;
//     }catch (error) {
//         console.error('Error', error.response);
//         return false
//     }
// }



// function logout() {
// }

// function getById(id) {
    
// }

// async function update(id, params) {
   
// }

// async function _delete(id) {
   
// }

