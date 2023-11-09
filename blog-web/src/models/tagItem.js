import {BlogItem} from "@/models/blogItem";

export class TagItem {
    constructor({info, tag, toEndPoint}) {
        this.info = null
        if (info !== undefined) {
            this.info = new BlogItem(info)
        }
        this.tag = tag
        this.toEndPoint = toEndPoint
    }
}

export class TagList {
    constructor({tags}) {
        this.tags = tags.map(item => new TagItem(item));
    }
}

export class TagGroup {
    constructor({tags={}}) {
        this.tagGroup = {};
        for (let tag in tags) {
            this.tagGroup[tag] = tags[tag].map(item => new TagItem(item));
        }
    }
}