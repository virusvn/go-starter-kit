import React, { Component } from 'react';
import Helmet from 'react-helmet';
import injectTapEventPlugin from 'react-tap-event-plugin';
import { AppBar,DatePicker } from 'material-ui';

// Needed for onTouchTap
// http://stackoverflow.com/a/34015469/988941
injectTapEventPlugin();
export default class App extends Component {

  render() {
    return <div id="parent">
        <AppBar
    title="Title"
    iconClassNameRight="muidocs-icon-navigation-expand-more"
  />
      <Helmet title='Go + React + Redux = rocks!' />
      {this.props.children}
        <div>
          <DatePicker hintText="Portrait Dialog" />
          <DatePicker hintText="Landscape Dialog" mode="landscape" />
          <DatePicker hintText="Dialog Disabled" disabled={true} />
        </div>
    </div>;
  }

}
