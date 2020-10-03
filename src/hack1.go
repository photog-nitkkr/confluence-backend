const admin = require('firebase-admin');
const functions = require('firebase-functions');
const express = require('express');
const bodyParser = require('body-parser');
const cors = require('cors');

// initialize default app
admin.initializeApp(functions.config().firebase);

// database reference
const db = admin.firestore();
db.settings({timestampsInSnapshots:true});

// express app
const app = express();

// ROUTES
const apiRoutes = require('./routes');

// body-parser
app.use(bodyParser.urlencoded({extended:false}));
// cors
app.use(cors({ origin: true }));


// api routes
app.use('/scxaaaa', (req, res) => {

    res.send("saaasasasasds");
})

app.use('/', apiRoutes);

// EXPORT Functions
exports.api = functions.https.onRequest(app);
