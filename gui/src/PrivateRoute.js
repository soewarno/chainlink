import React from 'react'
import { Redirect, Route } from 'react-static'
import { connect } from 'react-redux'

const PrivateRoute = ({ component: Component, authenticated, ...rest }) => (
  <Route {...rest} render={(props) => (
    authenticated
      ? <Component {...props} />
      : <Redirect to='/sign_in' />
  )} />
)

const mapStateToProps = state => ({
  authenticated: state.authenticated
})

const ConnectedPrivateRoute = connect(mapStateToProps)(PrivateRoute)

export default ConnectedPrivateRoute
