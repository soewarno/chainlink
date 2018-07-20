import {
  RECEIVE_SIGN_IN_FAIL,
  RECEIVE_SIGN_IN_SUCCESS
} from 'actions'
const initialState = []

const FAIL_MSG = 'Your email or password are incorrect. Please try again'

export default (state = initialState, action = {}) => {
  switch (action.type) {
    case RECEIVE_SIGN_IN_FAIL: {
      return [...state, FAIL_MSG]
    }
    case RECEIVE_SIGN_IN_SUCCESS: {
      return state.filter(m => m !== FAIL_MSG)
    }
    default:
      return state
  }
}
