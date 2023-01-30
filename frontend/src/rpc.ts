import * as app from '@/wailsjs/go/main/App'
import * as FileSystemService from '@/wailsjs/go/services/FileSystemService'
import * as EPWService from '@/wailsjs/go/services/EPWService'
import { WindowSetTitle, EventsOnMultiple } from '@/wailsjs/runtime'

const rpc = { app, FileSystemService, EPWService, on, setPageTitle }

function on(event: string, callback: (...data: any) => void) {
    EventsOnMultiple(event, callback, -1)
}

async function setPageTitle(title: string) {
    const prefix = await app.Title()

    WindowSetTitle(`${prefix} - ${title}`)
}

export default rpc