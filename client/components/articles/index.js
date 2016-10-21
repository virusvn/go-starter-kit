import React, { Component } from 'react';
import { connect } from 'react-redux';
import Helmet from 'react-helmet';
import { IndexLink } from 'react-router';
import { usage, todo } from './styles';
import { example, p, link } from '../homepage/styles';
import { setConfig, setVertex } from '../../actions';

class UsageV2 extends Component {
  /*eslint-enable */
  componentWillMount() {
   fetch('/api/v1/articles').then((r) => {
      return r.json();
   }).then((conf) => {
      store.dispatch(setVertex(conf));
      console.warn('Faked connection latency! Please, take a look ---> `server/api.go:22`');
    }); 
  }
  

  render() {
    return <div className={usage}>
      <Helmet title='Usage' />
      <h2 className={example}>Articles:</h2>
      <div className={p}>
        <span className={todo}>// TODO: write an article</span>
        <pre className={todo}>vertex:
          {JSON.stringify(this.props.vertex, null, 2)}</pre>
      </div>
      <br />
      go <IndexLink to='/' className={link}>home</IndexLink>
    </div>;
  }

}
class Usage extends Component {

  /*eslint-disable */
  static onEnter({store, nextState, replaceState, callback}) {
    fetch('/api/v1/conf').then((r) => {
      return r.json();
    }).then((conf) => {
      store.dispatch(setConfig(conf));
      console.warn('Faked connection latency! Please, take a look ---> `server/api.go:22`');
      callback();
    });
  }
  /*eslint-enable */

  render() {
    return <div className={usage}>
      <Helmet title='Usage' />
      <h2 className={example}>Articles:</h2>
      <div className={p}>
        <span className={todo}>// TODO: write an article</span>
        <pre className={todo}>config:
          {JSON.stringify(this.props.config, null, 2)}</pre>
      </div>
      <br />
      go <IndexLink to='/' className={link}>home</IndexLink>
      <UsageV2 vertex={(this.props.config.vertex)}/>
    </div>;
  }

}

export default connect(store => ({ config: store.config }))(Usage);
