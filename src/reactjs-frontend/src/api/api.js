const axios = require('axios');

const apiUrl = process.env.NODE_ENV === 'production' ? 'https://kleverio-challenge.herokuapp.com' : process.env.REACT_APP_API_ADDRESS;

const getBalance = async (address) => {
  let response = await axios.get(apiUrl + '/balance/' + address);
  return { confirmed: response.data.confirmed, unconfirmed: response.data.unconfirmed };
}



export { getBalance };