import axios from "axios";


const Api = {};
async function getTags() {
  try {
    const response = await axios.get('http://localhost:4000/api/cms/tags');
    console.log(response);
  } catch (error) {
    console.error(error);
  }
}

Api.getTags = getTags;


export default Api;
