import { combineReducers } from 'redux';
import deepReducer from './deepReducer';

export default combineReducers({
	deep: deepReducer
});
