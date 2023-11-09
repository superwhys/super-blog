import {request} from "./requests";

export function getPost(year, month, day, name) {
    return request({
        url: `/post/${year}/${month}/${day}/${name}`,
    })
}