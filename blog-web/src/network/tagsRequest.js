import {request} from "./requests";

export function getTagsList() {
    return request({
        url: "/tag/",
    })
}


export function  getTagsInfo(tag) {
    let url = "/tag/info"
    if (tag !== "") {
        url += `/${tag}`
    }
    return request({
        url: url
    })
}
