import React from 'react'
import PropTypes from 'prop-types'
import Button from '@material-ui/core/Button'
import { withStyles } from '@material-ui/core/styles'
import { CopyToClipboard } from 'react-copy-to-clipboard'

const styles = theme => ({
  inform: {
    display: 'inline-block',
    opacity: '1'
  }
})

const Copy = ({buttonText, classes, data}) => (
  <CopyToClipboard text={data}>
    <Button variant='outlined' color='primary' className={classes.button}>
      {buttonText}
    </Button>
  </CopyToClipboard>
)

Copy.propTypes = {
  classes: PropTypes.object.isRequired,
  data: PropTypes.string.isRequired,
  buttonText: PropTypes.string
}

Copy.defaultProps = {
  buttonText: 'Copy'
}

export default withStyles(styles)(Copy)
