import {urlService} from './UrlServices';
import Axios from 'axios';

export const fileService = {
    uploadPnl,
};

async function uploadPnl(file) {
    const requestOptions = {
        method: 'POST',
        headers: {'Content-Type': 'application/json'}
    };
    try{
        const response = await Axios.post(
            urlService.uploadPnl(), 
            requestOptions, 
            { file: file }
        );
        console.log(response)
    }catch{
        console.error('Error', error.response);
    }
}