import {request} from "./requests";

export function getBlogItemList(page, size) {
    return request({
        url: "/post/",
        params: {
            page: page,
            size: size
        }
    })
}