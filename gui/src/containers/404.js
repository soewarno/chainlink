import React from 'react'
import Grid from '@material-ui/core/Grid'
import Typography from '@material-ui/core/Typography'
import { withSiteData } from 'react-static'
import { withStyles } from '@material-ui/core/styles'

const styles = theme => ({
  title: {
    marginTop: theme.spacing.unit * 5,
    marginBottom: theme.spacing.unit * 5
  }
})

const NotFound = () => (
  <Grid container spacing={24}>
    <Grid item xs={12}>
      <Typography type='display1' align='center'>
        404
      </Typography>
      <Typography type='body1' align='center'>
        Oh no's! We couldn't find that page :(
      </Typography>
    </Grid>
  </Grid>
)

export default withSiteData(
  withStyles(styles)(NotFound)
)
