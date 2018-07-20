import reducer from 'connectors/redux/reducers'
import {
  RECEIVE_SIGN_IN_SUCCESS,
  RECEIVE_SIGN_IN_FAIL
} from 'actions'

describe('authenticated reducer', () => {
  it('should return the initial state', () => {
    const state = reducer(undefined, {})

    expect(state.authenticated).toEqual(false)
  })

  it('RECEIVE_SIGN_IN_SUCCESS sets the state to true', () => {
    const action = {type: RECEIVE_SIGN_IN_SUCCESS}
    const state = reducer(undefined, action)

    expect(state.authenticated).toEqual(true)
  })

  it('RECEIVE_SIGN_IN_FAIL sets the state to false', () => {
    const action = {type: RECEIVE_SIGN_IN_FAIL}
    const state = reducer({authenticated: true}, action)

    expect(state.authenticated).toEqual(false)
  })
})
