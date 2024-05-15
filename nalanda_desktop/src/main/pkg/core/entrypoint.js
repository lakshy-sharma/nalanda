/*
This file is our main entrypoint for the electron backend.
Once the init is completed the electron application calls the entrypoint and starts the backend logic.

All folders are briefly described here.

1. Bridge folder is used for exposing the API for all OS specific commands.
2. The controllers folder is the main folder that contains the features required, the functions described here call the bridge functions.
3. Windows Folder contais the backend functions for windows operations.
4. Linux folder contains the backend functions for linux operations.
5. API folder is used for performing any API related operations.
6. Globals File contains all required variables and must be initialized first.
7. Utils folder contains common utilities like logger and file writer etc.
*/

import { schedulerInit } from "./schedules";
import { listenersInit } from "./ipc";
import "./globals";

export function entrypoint() {
    console.log(STARTUP_MESSAGE);

    // Initialize all the IPC listeners.
    listenersInit();

    // Setup all schedulers.
    schedulerInit();
}