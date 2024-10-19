import express from 'express';
import axios from 'axios';

const app = express();
const port = 3000;

app.get('/', async (req, res) => {
  try {
    // Query left child (Node 2 - Go Service)
    const node2Response = await axios.get('http://node2:8080');
    
    // Query right child (Node 3 - Python Service)
    const node3Response = await axios.get('http://node3:5000');
    
    res.send(`Node 2 says: ${node2Response.data}, Node 3 says: ${node3Response.data}`);
  } catch (error) {
    res.status(500).send('Error communicating with child nodes');
  }
});

app.listen(port, () => {
  console.log(`Node 1 (TypeScript) running at http://localhost:${port}`);
});
