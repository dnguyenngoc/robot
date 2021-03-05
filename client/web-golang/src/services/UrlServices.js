const  restService = 'http://localhost:8080'
const apiV1 = '/api/v1'


function loginAccessToken() {return restService + apiV1 + '/accounts/login/access-token'}
function signup() {return restService + apiV1 + '/accounts/signup'}


export {loginAccessToken as urlLoginAccessToken, signup as urlSignup};
