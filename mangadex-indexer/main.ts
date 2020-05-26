import { connect, Db } from 'mongodb';
import * as express from 'express';
import Agenda from 'agenda';
import Parser from "rss-parser";

const agendash = require('agendash');

const url = 'mongodb://localhost:27017';
const dbName = 'mangafox';

const app = express.default();
const agenda = new Agenda({
    db: {
        address: url, options: {
            useUnifiedTopology: true,
        }
    }
});

const parser = new Parser()

var database: Db;

enum AgendaJob {
    mangadex_sync_latest = "mangadex_sync_latest"
}

agenda.define("mangadex_sync_latest", { concurrency: 1 }, async (job) => {
    let feed = await parser.parseURL('https://mangadex.org/rss/2ZevhabKgkstB6DPzQpMcdSRnxwf78uC');
    console.log(feed.title);

});

async function main() {
    const client = await connect(url, {
        useUnifiedTopology: true,
    });

    database = client.db(dbName);

    await agenda.start();
    await agenda.every('30 seconds', "mangadex_sync_latest");


    app.use('/agenda', agendash(agenda))

    process.on('SIGTERM', graceful);
    process.on('SIGINT', graceful);

    app.listen(8080);
}

async function graceful() {
    await agenda.stop();
    process.exit(0);
}

main();