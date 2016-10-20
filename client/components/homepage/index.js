import React, { Component } from 'react';
import Helmet from 'react-helmet';
import { Link } from 'react-router';
import { example, p, link } from './styles';

export default class Homepage extends Component {
  /*eslint-disable */
  static onEnter({store, nextState, replaceState, callback}) {
    // Load here any data.
    callback(); // this call is important, don't forget it
  }
  /*eslint-enable */

  render() {
    return <div>
      <Helmet
        title='Home page'
        meta={[
          {
            property: 'og:title',
            content: 'Golang Isomorphic React/Hot Reloadable/Redux/Css-Modules Starter Kits'
          }
        ]} />
      <h1 className={example}>
        Hot Reloadablerrr <br />
        Golang + React + Redux + Css-Modules
        <br />Universy Starter Kit - edited 20161005</h1>
      <br />
      <p className={p}>
        Please take a look at <Link className={link} to='/docs'>usage</Link> page.
        Please take a look at <Link className={link} to='/articles'>article</Link> page.
      </p>
    </div>;
  }

}
