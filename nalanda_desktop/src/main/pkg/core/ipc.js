// This file contains all the IPC listeners used by the electron application.
import {ipcMain} from "electron";

export function listenersInit() {
    ipcMain.on("ping", (() => {
        console.log("pong");
    }))
}