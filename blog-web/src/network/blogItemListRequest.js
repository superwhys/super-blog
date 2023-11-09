import {request} from "./requests";

export function getBlogItemList() {
    return request({
        url: "/post/",
    })
}