import { combineReducers } from 'redux'
import authenticated from './authenticated'
import errors from './errors'
import accountBalance from './accountBalance'
import bridges from './bridges'
import jobs from './jobs'
import jobRuns from './jobRuns'
import configuration from './configuration'

const reducer = combineReducers({
  authenticated,
  errors,
  accountBalance,
  bridges,
  jobs,
  jobRuns,
  configuration
})

export default reducer
