import axios from 'axios'
import {Message} from "element-ui";
let backendApi = process.env.VUE_APP_BACKEND_API
export function request(config) {
    const instance = axios.create({
        baseURL: backendApi,
        timeout: 20000,
        method: 'get'
    })

    instance.interceptors.request.use(config => {
        return config
    }, err => {
        console.log(err)
    })

    instance.interceptors.response.use(res => {
        return res.data
    }, err => {
        console.log(err);
        Message.error(err)
        return err
    })

    return instance(config)
}
