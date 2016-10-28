const router = require('./router');

// export main function for server side rendering
global.main = router.renderToString;
// global.navigator = { userAgent: 'all' }; 
// global.navigator = { navigator: 'all' };
// start app if it in the browser
if(typeof window !== 'undefined') {
  // Start main application here
  router.run();
}
