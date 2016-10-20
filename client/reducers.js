import { combineReducers } from 'redux';
import { SET_CONFIG, SET_VERTEX } from './actions';

function config(state = {}, action) {
  switch (action.type) {
    case SET_CONFIG:
      return action.config;
    case SET_VERTEX:
      return {...state, vertex: action.config};
    default:
      return state;
  }
}

export default combineReducers({config});
