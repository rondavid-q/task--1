const express = require('express');
const axios = require('axios');

const app = express();
const port = 8000;

app.use(express.json());

app.post('/process-data', async (req, res) => {
  try {

    const pythonResponse = await axios.post('http://python-service:8001/process-data', req.body);


    const goResponse = await axios.post('http://go-service:8002/log-request', req.body);


    res.json({
      pythonResponse: pythonResponse.data,
      goResponse: goResponse.data,
    });
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

app.listen(port, () => {
  console.log(`Node.js service listening at http://0.0.0.0:${port}`);
});

