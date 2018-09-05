import axios from 'axios';
import { FETCH_DEEP_API } from './types';

//async dispatch is a function
export const fetchDeepAPI = () => async dispatch => {
	const res = await axios.get('http://localhost:5001');
	dispatch({ type: FETCH_DEEP_API, payload: res.data });
};