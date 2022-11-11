import express from 'express';
import AWS from 'aws-sdk';
import morgan from 'morgan';
import { faker } from '@faker-js/faker';
import cors from 'cors';

const s3 = new AWS.S3({
  accessKeyId: process.env.AWS_S3_ACCESS_KEY_ID,
  secretAccessKey: process.env.AWS_S3_SECRET_ACCESS_KEY
});

const app = express();

app.use(morgan('tiny'));
app.use(express.json());
app.use(cors());

app.get('/', (req, res) => {
  res.setHeader('Content-Type', 'application/json');
  res.send({ status: 'OK' });
});

app.get('/person', (req, res) => {
  res.setHeader('Content-Type', 'application/json');
  res.send({
    name: faker.name.firstName(),
    age: faker.random.numeric(2),
    favoriteColor: faker.color.human()
  });
});

app.post('/file', async (req, res) => {
  const content = req.body.content;

  await s3
    .upload({
      Bucket: 'nodejs-sample',
      Key: new Date().toISOString() + '.txt',
      Body: content
    })
    .promise();

  res.sendStatus(201);
});

app.use('/web', express.static('static'));

const port = process.env.PORT ?? 3000;

app
  .listen(port, () => {
    console.log(`Server running on port ${port}`);
  })
  .on('error', (err) => {
    console.error(err);
  });
