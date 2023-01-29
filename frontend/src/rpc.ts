import * as app from '@/wailsjs/go/main/App'
import * as FileSystemService from '@/wailsjs/go/services/FileSystemService'
import * as epw from '@/wailsjs/go/services/EPW'
import { WindowSetTitle, EventsOnMultiple } from '@/wailsjs/runtime'

const rpc = { app, FileSystemService, epw, on, setPageTitle }

function on(event: string, callback: (...data: any) => void) {
    EventsOnMultiple(event, callback, -1)
}

async function setPageTitle(title: string) {
    const prefix = await app.Title()

    WindowSetTitle(`${prefix} - ${title}`)
}

export default rpc