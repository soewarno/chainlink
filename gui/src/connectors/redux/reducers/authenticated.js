import {
  RECEIVE_SIGN_IN_SUCCESS,
  RECEIVE_SIGN_IN_FAIL
} from 'actions'

const initialState = false

export default (state = initialState, action = {}) => {
  switch (action.type) {
    case RECEIVE_SIGN_IN_SUCCESS:
      return true
    case RECEIVE_SIGN_IN_FAIL:
      return false
    default:
      return state
  }
}
