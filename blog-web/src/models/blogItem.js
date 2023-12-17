export class BlogMetaData {
    constructor({
                    layout = "",
                    title = "",
                    subTitle = "",
                    date = "",
                    auth = "",
                    tags = []
                }) {
        this.layout = layout
        this.title = title
        this.subTitle = subTitle
        this.date = date
        this.auth = auth
        this.tags = tags
    }
}

export class BlogItem {
    constructor({
                    metaData = {},
                    fileName = "",
                    title = "",
                    subTitle = "",
                    postContent = "",
                    postedTime = "",
                    toEndPoint = ""
                }) {
        this.metaData = new BlogMetaData(metaData)
        this.fileName = fileName
        this.title = title
        this.subTitle = subTitle
        this.postContent = postContent
        this.postedTime = postedTime
        this.toEndPoint = toEndPoint
    }
}

export class BlogList {
    constructor({items = [], total = 0}) {
        console.debug("items: ", items)
        this.items = items.map(item => new BlogItem(item));
        this.total = total;
    }
}