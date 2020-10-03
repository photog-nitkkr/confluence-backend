const express = require('express');
const app = express.Router();

const mediaRoutes = require('./mediaRoutes');
const eventRoutes = require('./eventsRoutes');


app.use('/gallery', mediaRoutes);
app.use('/events', eventRoutes);

app.use('/', (req, res) => {

	return res.status(400).json({
		success: false,
		message: 'wrong route'
	});
})

module.exports = app;
