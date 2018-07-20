import React, { Component } from 'react'
import PropTypes from 'prop-types'
import Grid from '@material-ui/core/Grid'
import Typography from '@material-ui/core/Typography'
import TextField from '@material-ui/core/TextField'
import Button from '@material-ui/core/Button'
import Icon from '@material-ui/core/Icon'
import { Redirect } from 'react-static'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { withStyles } from '@material-ui/core/styles'
import { signIn } from 'actions'

const styles = theme => ({
  title: {
    marginTop: theme.spacing.unit * 5,
    marginBottom: theme.spacing.unit * 5
  },
  icon: {
    marginLeft: theme.spacing.unit
  }
})

export class SignIn extends Component {
  constructor (props) {
    super(props)
    this.signIn = this.signIn.bind(this)
    this.state = {
      email: '',
      password: ''
    }
  }

  signIn (e) {
    e.preventDefault()
    const {email, password} = this.state
    this.props.signIn(email, password)
  }

  render () {
    const {classes, authenticated} = this.props

    if (authenticated) {
      return <Redirect to='/' />
    }

    return (
      <Grid container align='center' spacing={40}>
        <Grid item xs={12}>
          <Typography variant='display2' color='inherit' className={classes.title}>
            Sign In to Chainlink
          </Typography>
        </Grid>

        <Grid item xs={12}>
          <form className={classes.form} onSubmit={this.signIn}>
            <Grid container spacing={40}>
              <Grid item xs={12}>
                <TextField
                  id='email'
                  placeholder='Enter your email'
                  onChange={(e) => { this.setState({email: e.target.value}) }}
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  id='password'
                  type='password'
                  placeholder='Password'
                  onChange={(e) => { this.setState({password: e.target.value}) }}
                />
              </Grid>
              <Grid item xs={12}>
                <Button variant='contained' color='primary' type='submit'>
                  Sign In
                  <Icon className={classes.icon}>arrow_forward</Icon>
                </Button>
              </Grid>
            </Grid>
          </form>
        </Grid>
      </Grid>
    )
  }
}

SignIn.propTypes = {
  classes: PropTypes.object.isRequired
}

const mapStateToProps = state => ({
  authenticated: state.authenticated
})

const mapDispatchToProps = (dispatch) => {
  return bindActionCreators({signIn}, dispatch)
}

export const ConnectedSignIn = connect(mapStateToProps, mapDispatchToProps)(SignIn)

export default withStyles(styles)(ConnectedSignIn)
