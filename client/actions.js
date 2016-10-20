/**
 * action types
 */

export const SET_CONFIG = 'SET_CONFIG';
export const SET_VERTEX = 'SET_VERTEX';

/**
 * action creators
 */

export function setConfig(config) {
  return { type: SET_CONFIG, config };
}
export function setVertex(config) {
  return {type: SET_VERTEX, config}
}