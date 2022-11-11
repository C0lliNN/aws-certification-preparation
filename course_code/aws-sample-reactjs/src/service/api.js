import axios from 'axios';

const baseUrl = process.env.REACT_APP_BASE_URL ?? 'http://localhost:3001';

export async function getPerson() {
  const response = await axios.get(`${baseUrl}/person`);
  return response.data;
}
