import {cookieGet, cookieSet } from './CookieService';
import {getNowUtc} from '../utils/DateHelper';


async function tokenCheckCookie() {
    // const token = cookieGet('token')
    // if (token) {
    //     if (token.expire_time < getNowUtc()){
    //         return false
    //     }
    //     else {
    //         return true
    //     }
    // } else {
    //     return false
    // }
}


function makeTokenCookie(token) {
    cookieSet.cookieSet('token', token);
}


export {tokenCheckCookie, makeTokenCookie};