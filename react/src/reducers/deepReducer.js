import { FETCH_DEEP_API } from '../actions/types';

const INITIAL_STATE = {
	deep_results: null
}

export default function(state = INITIAL_STATE, action) {
	// console.log(action);
	switch(action.type) {
		case FETCH_DEEP_API:
			return {...state, deep_results: action.payload};
		default: 
			return state;
	}
}